package interceptors

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/ydssx/morphix/pkg/logger"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func interceptorLogger(lg log.Logger) logging.Logger {
	l := lg.(*logger.Logger)
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.Zlog.WithOptions(zap.AddCallerSkip(2)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}

	})
}

func LoggingServer(l log.Logger) grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(interceptorLogger(l), initOpt()...)
}

func LoggingStreamServer(l log.Logger) grpc.StreamServerInterceptor {
	return logging.StreamServerInterceptor(interceptorLogger(l), initOpt()...)
}


func LoggingClient(l log.Logger) grpc.UnaryClientInterceptor {
	return logging.UnaryClientInterceptor(interceptorLogger(l), initOpt()...)
}

func LoggingStreamClient(l log.Logger) grpc.StreamClientInterceptor {
	return logging.StreamClientInterceptor(interceptorLogger(l), initOpt()...)
}

func initOpt() []logging.Option {
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.PayloadReceived, logging.PayloadSent),
		logging.WithFieldsFromContext(logTraceID),
		logging.WithTimestampFormat("2006-01-02 15:04:05"),
		// Add any other option (check functions starting with logging.With).
	}
	return opts
}
