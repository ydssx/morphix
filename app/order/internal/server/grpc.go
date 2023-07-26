package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	"github.com/ydssx/morphix/app/order/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/middleware/kratos"
)

func NewGRPCServer(c *common.Config, svc *service.OrderService, logger log.Logger) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(
			interceptors.TraceServerInterceptor(),
			interceptors.LoggingServerInterceptor(logger),
			interceptors.ValidatorServerInterceptor(),
			interceptors.EventInterceptors(),
		),
		grpc.Middleware(
			kratos.MetricServer(),
			// tracing.Server(),
			// logging.Server(logger),
			recovery.Recovery(),
		),
	}
	server := c.Order.Server
	if server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(server.Grpc.Addr))
	}
	if server.Grpc.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(server.Grpc.Timeout)*time.Second))
	}
	srv := grpc.NewServer(opts...)

	orderv1.RegisterOrderServiceServer(srv, svc)

	return srv
}
