package server

import (
	"time"

	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/middleware/kratos"
)

func NewGRPCServer(c *common.Config, userSvc *service.UserService, logger log.Logger) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			interceptors.TraceServerInterceptor(),
			interceptors.LoggingServerInterceptor(logger),
		),
		grpc.Middleware(
			kratos.MetricServer(),
			// tracing.Server(),
			// logging.Server(logger),
			recovery.Recovery(),
		),
	}
	server := c.User.Server
	if server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(server.Grpc.Addr))
	}
	if server.Grpc.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(server.Grpc.Timeout)*time.Second))
	}
	srv := grpc.NewServer(opts...)

	userv1.RegisterUserServiceServer(srv, userSvc)

	daprd.NewServiceWithGrpcServer(nil, srv.Server)

	return srv
}
