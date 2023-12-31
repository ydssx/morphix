package provider

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	resource          *sdkresource.Resource
	initResourcesOnce sync.Once
)

func initResource() *sdkresource.Resource {
	initResourcesOnce.Do(func() {
		extraResources, _ := sdkresource.New(
			context.Background(),
			sdkresource.WithOS(),
			sdkresource.WithProcess(),
			sdkresource.WithContainer(),
			sdkresource.WithHost(),
		)
		resource, _ = sdkresource.Merge(sdkresource.Default(), extraResources)
	})
	return resource
}

// InitTraceProvider initializes and returns a TracerProvider for exporting
// traces to the specified endpoint. It creates an OTLP exporter, configures
// a resource with service name and attributes, sets the global tracer
// provider, and returns the provider.
func InitTraceProvider(url string, tracename string) (*sdktrace.TracerProvider, error) {
	// 创建 Jaeger exporter
	// exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	ctx := context.Background()
	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(url), otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	sche := sdkresource.NewSchemaless(
		semconv.ServiceNameKey.String(tracename),
		attribute.String("exporter", "jaeger"),
		attribute.Float64("float", 312.23),
	)
	source, _ := sdkresource.Merge(initResource(), sche)

	tp := sdktrace.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1.0))),
		// 始终确保在生产中批量处理
		sdktrace.WithBatcher(exp),
		// 在资源中记录有关此应用程序的信息
		sdktrace.WithResource(source),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

// InitMeterProvider initializes and returns a MeterProvider for exporting
// metrics to the specified endpoint. It creates an OTLP exporter, a
// PeriodicReader, and a MeterProvider with the exporter and default
// resource. It also sets the global MeterProvider.
func InitMeterProvider(endpoint string) *sdkmetric.MeterProvider {
	exporter, err := otlpmetricgrpc.New(context.Background(), otlpmetricgrpc.WithEndpoint(endpoint), otlpmetricgrpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	reader := sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(time.Second*30))
	resource := initResource()

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(reader),
		sdkmetric.WithResource(resource),
	)
	otel.SetMeterProvider(provider)

	return provider
}
