package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	orderv1 "github.com/ydssx/morphix/api/order/v1"
	"github.com/ydssx/morphix/app/order/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewHTTPServer(c *conf.Bootstrap, svc *service.OrderService) *http.Server {
	srv := common.NewHTTPServer(c.ServiceSet.Order.Server)

	orderv1.RegisterOrderServiceHTTPServer(srv, svc)

	return srv
}
