package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	productv1 "github.com/ydssx/morphix/api/product/v1"
	"github.com/ydssx/morphix/app/product/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.ProductService) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.Product.Server)

	productv1.RegisterProductServiceServer(srv, svc)

	return srv
}
