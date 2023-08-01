package common

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/middleware/kratos"
)

func NewGRPCServer(server *conf.Server) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			interceptors.TraceServerInterceptor(),
			interceptors.LoggingServerInterceptor(logger.DefaultLogger),
			interceptors.ValidatorServerInterceptor(),
			interceptors.EventServerInterceptors(),
		),
		grpc.Middleware(
			kratos.MetricServer(),
			recovery.Recovery(),
		),
	}
	if server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(server.Grpc.Addr))
	}
	if server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	return srv
}
