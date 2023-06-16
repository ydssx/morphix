package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	user "github.com/ydssx/morphix/app/user/api"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/middleware/kratos"
	"go.uber.org/zap"
)

func NewGRPCServer(c *conf.Server, userSvc *service.UserService, logger log.Logger, zaplog *zap.Logger) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			interceptors.TraceServerInterceptor(),
			interceptors.LoggingServerInterceptor(zaplog),
		),
		grpc.Middleware(
			kratos.MetricServer(),
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

	user.RegisterUserServiceServer(srv, userSvc)

	return srv
}
