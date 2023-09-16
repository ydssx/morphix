package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	paymentv1 "github.com/ydssx/morphix/api/payment/v1"
	"github.com/ydssx/morphix/app/payment/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.PaymentService, logger log.Logger) *grpc.Server {

	srv := common.NewGRPCServer(c.ServiceSet.Payment.Server)
	
	paymentv1.RegisterPaymentServiceServer(srv, svc)

	return srv
}
