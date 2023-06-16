package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	user "github.com/ydssx/morphix/app/user/api"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/app/user/internal/server/interceptors"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/pkg/trace"
	"go.uber.org/zap"
)

func NewGRPCServer(c *conf.Server, userSvc *service.UserService, logger log.Logger, zaplog *zap.Logger) *grpc.Server {
	err := trace.InitTracer("http://localhost:14268/api/traces", "user-rpc")
	if err != nil {
		panic(err)
	}

	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			// interceptors.TraceInterceptor(),
			interceptors.LoggingInterceptor(zaplog),
		),
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
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
