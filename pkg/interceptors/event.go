package interceptors

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"google.golang.org/grpc"
)

type EventSource struct{}
type EventType struct{}

func EventServerInterceptors() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		app, ok := kratos.FromContext(ctx)
		if ok {
			ctx = context.WithValue(ctx, EventSource{}, app.Name())
		}
		ctx = context.WithValue(ctx, EventType{}, info.FullMethod)
		return handler(ctx, req)
	}
}
