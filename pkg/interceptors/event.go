package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

type EventSource struct{}
type EventType struct{}

func EventInterceptors() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, EventType{}, info.FullMethod)
		ctx = context.WithValue(ctx, EventSource{}, nil)
		return handler(ctx, req)
	}
}
