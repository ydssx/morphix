package interceptors

import (
	"context"

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

// 需要跳过鉴权的服务(api集合)和api
var (
	srvNames = []string{
		healthpb.Health_ServiceDesc.ServiceName,
	}
	methNames = []string{
		userv1.UserService_Login_FullMethodName,
		userv1.UserService_Register_FullMethodName,
		smsv1.SMSService_SendSMS_FullMethodName,
	}
)

func match(_ context.Context, callMeta interceptors.CallMeta) bool {
	result := !(util.SliceContain(srvNames, callMeta.Service) || util.SliceContain(methNames, callMeta.FullMethod()))
	return result
}

func AuthServer() grpc.UnaryServerInterceptor {
	return selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authCtx), selector.MatchFunc(match))
}

func AuthStreamServer() grpc.StreamServerInterceptor {
	return selector.StreamServerInterceptor(auth.StreamServerInterceptor(authCtx), selector.MatchFunc(match))
}

// authCtx authenticates the incoming request context.
// It checks for a valid authentication token and populates the
// context with the parsed claims. If no token is present, it
// returns the existing context unmodified.
func authCtx(ctx context.Context) (context.Context, error) {
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

// parseToken 解析传入的 context 中的 token,验证其有效性。
// 如果 token 验证失败,返回错误。
// 如果验证成功,使用提取的 claims 生成新的 context 并返回。
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

// isExternalRequest checks if the request is from an external client.
// It looks for a custom "external-request" header and returns true if it's set to "true".
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
