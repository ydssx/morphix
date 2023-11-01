package interceptors

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/ydssx/morphix/pkg/logger"
	"google.golang.org/grpc"
)

func interceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		logger.Log(ctx, log.Level(lvl), msg, fields...)
	})
}

func LoggingServer() grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(interceptorLogger(), initOpt()...)
}

func LoggingStreamServer() grpc.StreamServerInterceptor {
	return logging.StreamServerInterceptor(interceptorLogger(), initOpt()...)
}

func LoggingClient() grpc.UnaryClientInterceptor {
	return logging.UnaryClientInterceptor(interceptorLogger(), initOpt()...)
}

func LoggingStreamClient() grpc.StreamClientInterceptor {
	return logging.StreamClientInterceptor(interceptorLogger(), initOpt()...)
}

func initOpt() []logging.Option {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.PayloadReceived, logging.PayloadSent),
		logging.WithTimestampFormat("2006-01-02 15:04:05"),
		// Add any other option (check functions starting with logging.With).
	}
	return opts
}
