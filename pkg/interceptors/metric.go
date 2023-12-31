package interceptors

import (
	"context"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/ydssx/morphix/pkg/logger"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	reg = prometheus.DefaultRegisterer

	exemplarFromContext = func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}

	panicsTotal = promauto.With(reg).NewCounter(prometheus.CounterOpts{
		Name: "grpc_req_panics_recovered_total",
		Help: "Total number of gRPC requests recovered from internal panic.",
	})

	grpcPanicRecoveryHandler = func(ctx context.Context, p any) (err error) {
		panicsTotal.Inc()
		reportPanic(ctx, p)
		return status.Errorf(codes.Internal, "%s", p)
	}
)

func reportPanic(ctx context.Context, p any) {
	logger.Error(ctx, p)
}

func MetricServer() grpc.UnaryServerInterceptor {
	clMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg.MustRegister(clMetrics)
	return clMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
}

func MetricClient() grpc.UnaryClientInterceptor {
	clMetrics := grpcprom.NewClientMetrics(
		grpcprom.WithClientHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(clMetrics)
	return clMetrics.UnaryClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
}
