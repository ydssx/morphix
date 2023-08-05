package interceptors

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/filters"
	"google.golang.org/grpc"
)

func TraceServer() grpc.UnaryServerInterceptor {
	return otelgrpc.UnaryServerInterceptor(otelgrpc.WithInterceptorFilter(filters.Not(filters.HealthCheck())))
}

func TraceClient() grpc.UnaryClientInterceptor {
	return otelgrpc.UnaryClientInterceptor()
}
