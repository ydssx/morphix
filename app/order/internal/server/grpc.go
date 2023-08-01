package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	"github.com/ydssx/morphix/app/order/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.OrderService) *grpc.Server {

	srv := common.NewGRPCServer(c.Order.Server)

	orderv1.RegisterOrderServiceServer(srv, svc)

	return srv
}
