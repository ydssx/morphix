package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/app/sms/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, smsSvc *service.SMSService) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.Sms.Server)

	smsv1.RegisterSMSServiceServer(srv, smsSvc)

	return srv
}
