package metric

import (
	"context"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	eptprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func InitMeterProvider() error {
	exporter, err := eptprom.New()
	if err != nil {
		return err
	}
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	_ = provider.Meter("github.com/open-telemetry/opentelemetry-go/example/prometheus")

	otel.SetMeterProvider(provider)
	return nil
}

func init() {
	_metricSeconds := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "server",
		Subsystem: "requests",
		Name:      "duration_sec",
		Help:      "server requests duratio(sec).",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	_metricRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "client",
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})

	prometheus.MustRegister(_metricSeconds, _metricRequests)
}

func InitMetric() grpc.UnaryClientInterceptor {
	reg := prometheus.NewRegistry()
	clMetrics := grpcprom.NewClientMetrics(
		grpcprom.WithClientHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg.MustRegister(clMetrics)

	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}
	return clMetrics.UnaryClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
}
