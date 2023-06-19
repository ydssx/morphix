package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/middleware/kratos"
)

func NewGRPCServer(c *conf.Server, userSvc *service.UserService, logger log.Logger) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			// interceptors.TraceServerInterceptor(),
			interceptors.LoggingServerInterceptor(logger),
		),
		grpc.Middleware(
			kratos.MetricServer(),
			tracing.Server(),
			// logging.Server(logger),
			recovery.Recovery(),
		),
	}

	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	userv1.RegisterUserServiceServer(srv, userSvc)

	return srv
}
