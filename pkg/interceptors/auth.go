package interceptors

import (
	"context"

	"github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func AuthServer() grpc.UnaryServerInterceptor {
	authFn := func(ctx context.Context) (context.Context, error) {
		md := metadata.ExtractIncoming(ctx)
		ctx = md.ToOutgoing(ctx)
		if !isExternalRequest(md) {
			newCtx, err := parseToken(ctx)
			if err == nil {
				return newCtx, nil
			}
			return ctx, nil
		}
		return parseToken(ctx)
	}

	// Setup auth matcher.
	authMatcher := func(_ context.Context, callMeta interceptors.CallMeta) bool {
		srvNames := []string{
			healthpb.Health_ServiceDesc.ServiceName,
			runtime.AppCallbackAlpha_ServiceDesc.ServiceName,
		}
		methNames := []string{
			userv1.UserService_Login_FullMethodName,
			// userv1.UserService_Register_FullMethodName,
			smsv1.SMSService_SendSMS_FullMethodName,
		}

		result := !(util.SliceContain(srvNames, callMeta.Service) || util.SliceContain(methNames, callMeta.FullMethod()))

		return result
	}

	return selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(authMatcher))
}

func parseToken(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	claims, err := jwt.VerifyToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}

	return newContext(ctx, claims), nil
}

func isExternalRequest(md metadata.MD) bool {
	// Determine if the request is coming from an external client based on the presence of custom headers.
	// For example, you can define a custom "external-request" header.
	return md.Get("external-request") == "true"
}

type authKey struct{}

type AuthInfo struct {
	Claims *jwt.Claims
}

func newContext(ctx context.Context, c *jwt.Claims) context.Context {
	info := &AuthInfo{
		Claims: c,
	}
	return context.WithValue(ctx, authKey{}, info)
}

func AuthFromContext(ctx context.Context) (*jwt.Claims, bool) {
	info, ok := ctx.Value(authKey{}).(*AuthInfo)
	if !ok {
		return nil, false
	}
	return info.Claims, true
}
