package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	productv1 "github.com/ydssx/morphix/api/product/v1"
	"github.com/ydssx/morphix/app/product/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewHTTPServer(c *conf.Bootstrap, svc *service.ProductService) *http.Server {
	srv := common.NewHTTPServer(c.ServiceSet.Product.Server)

	productv1.RegisterProductServiceHTTPServer(srv, svc)

	return srv
}
