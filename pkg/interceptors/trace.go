package interceptors

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func TraceServer() grpc.UnaryServerInterceptor {
	return otelgrpc.UnaryServerInterceptor()
}

func TraceClient() grpc.UnaryClientInterceptor {
	return otelgrpc.UnaryClientInterceptor()
}
