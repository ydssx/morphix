package common

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
)

func NewGRPCServer(server *conf.Server) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			interceptors.TraceServer(),
			interceptors.LoggingServer(logger.DefaultLogger),
			interceptors.ValidatorServer(),
			interceptors.EventServer(),
			interceptors.AuthServer(),
			interceptors.MetricServer(),
			interceptors.RecoveryServer(),
		),
		grpc.StreamInterceptor(
			interceptors.TraceStreamServer(),
			interceptors.LoggingStreamServer(logger.DefaultLogger),
			interceptors.ValidatorStreamServer(),
			interceptors.AuthStreamServer(),
			interceptors.RecoveryStreamServer(),
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
