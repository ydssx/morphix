package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	{{.PkgName}} "{{.PkgPath}}"
	"github.com/ydssx/morphix/app/{{.appName}}/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.{{.serviceName}}) *grpc.Server {

	srv := common.NewGRPCServer(c.ServiceSet.{{.appName | Title}}.Server)

	{{.PkgName}}.Register{{.serviceName}}Server(srv, svc)

	return srv
}
