package interceptors

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"google.golang.org/grpc"
)

type eventSource struct{}
type eventType struct{}

func EventServer() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		app, ok := kratos.FromContext(ctx)
		if ok {
			ctx = context.WithValue(ctx, eventSource{}, app.Name())
		}
		ctx = context.WithValue(ctx, eventType{}, info.FullMethod)
		return handler(ctx, req)
	}
}

func EventSourceFromCtx(ctx context.Context) string {
	t, ok := ctx.Value(eventSource{}).(string)
	if ok {
		return t
	}
	return "api"
}

func EventTypeFromCtx(ctx context.Context) string {
	t, ok := ctx.Value(eventType{}).(string)
	if ok {
		return t
	}
	return "null"
}
