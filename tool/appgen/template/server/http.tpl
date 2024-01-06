package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	{{.serviceInfo.PkgName}} "{{.serviceInfo.PkgPath}}"
	"{{.module}}/app/{{.appName}}/internal/service"
	"{{.module}}/common"
	"{{.module}}/common/conf"
)

func NewHTTPServer(c *conf.Bootstrap, svc *service.{{.serviceInfo.Name}}) *http.Server {
	srv := common.NewHTTPServer(c.ServiceSet.{{.appName | Title}}.Server)

	{{.serviceInfo.PkgName}}.Register{{.serviceInfo.Name}}HTTPServer(srv, svc)

	return srv
}
