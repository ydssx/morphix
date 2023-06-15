package server

import (

	// "github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	user "github.com/ydssx/morphix/app/user/api"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/app/user/internal/server/middleware"
	"github.com/ydssx/morphix/app/user/internal/service"
	"go.uber.org/zap"
)

func NewGRPCServer(c *conf.Server, userSvc *service.UserService) *grpc.Server {
	logopts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	var opts = []grpc.ServerOption{
		grpc.Middleware(recovery.Recovery()),
		grpc.UnaryInterceptor(logging.UnaryServerInterceptor(middleware.InterceptorLogger(zap.NewExample()), logopts...)),
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
