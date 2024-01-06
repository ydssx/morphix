package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewHTTPServer(c *conf.Bootstrap, svc *service.UserService) *http.Server {
	srv := common.NewHTTPServer(c.ServiceSet.User.Server)

	userv1.RegisterUserServiceHTTPServer(srv, svc)

	return srv
}
