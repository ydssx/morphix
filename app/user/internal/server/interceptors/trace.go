package interceptors

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func TraceInterceptor() grpc.UnaryServerInterceptor {
	// exporter, err := stdout.New(stdout.WithPrettyPrint())
	// if err != nil {
	// 	panic(err)
	// }
	// tp := sdktrace.NewTracerProvider(
	// 	sdktrace.WithSampler(sdktrace.AlwaysSample()),
	// 	sdktrace.WithBatcher(exporter),
	// )
	// otel.SetTracerProvider(tp)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return otelgrpc.UnaryServerInterceptor()
}
