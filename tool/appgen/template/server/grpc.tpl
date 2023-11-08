package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	{{.serviceInfo.PkgName}} "{{.serviceInfo.PkgPath}}"
	"{{.module}}/app/{{.appName}}/internal/service"
	"{{.module}}/common"
	"{{.module}}/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.{{.serviceInfo.Name}}) *grpc.Server {
	srv := common.NewGRPCServer(c.ServiceSet.{{.appName | Title}}.Server)

	{{.serviceInfo.PkgName}}.Register{{.serviceInfo.Name}}Server(srv, svc)

	return srv
}
