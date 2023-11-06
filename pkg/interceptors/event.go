package interceptors

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			err := grpc.SendHeader(ctx, metadata.Pairs("trace-id", span.TraceID().String()))
			if err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}

func EventStreamServer() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		app, ok := kratos.FromContext(ctx)
		if ok {
			ctx = context.WithValue(ctx, eventSource{}, app.Name())
		}

		ctx = context.WithValue(ctx, eventType{}, info.FullMethod)

		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			err := ss.SendHeader(metadata.Pairs("trace-id", span.TraceID().String()))
			if err != nil {
				return err
			}
		}

		wrapped := middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
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
