package interceptors

import (
	"context"

	"github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
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
		token, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid auth token")
		}
		md := metadata.ExtractIncoming(ctx)
		return newContext(md.ToOutgoing(ctx), claims), nil
	}

	// Setup auth matcher.
	authMatcher := func(ctx context.Context, callMeta interceptors.CallMeta) bool {
		srvNames := []string{
			healthpb.Health_ServiceDesc.ServiceName,
			runtime.AppCallbackAlpha_ServiceDesc.ServiceName,
		}
		methNames := []string{
			userv1.UserService_Login_FullMethodName,
			userv1.UserService_Register_FullMethodName,
		}

		result := !(util.SliceContain(srvNames, callMeta.Service) || util.SliceContain(methNames, callMeta.FullMethod()))

		return result
	}

	return selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(authMatcher))
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
