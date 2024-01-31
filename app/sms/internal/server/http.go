package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/app/sms/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewHTTPServer(c *conf.Bootstrap, svc *service.SMSService) *http.Server {
	srv := common.NewHTTPServer(c.ServiceSet.Sms.Server)

	smsv1.RegisterSMSServiceHTTPServer(srv, svc)

	return srv
}
