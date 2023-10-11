package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	{{.appName}}v1 "github.com/ydssx/morphix/api/{{.appName}}/v1"
	"github.com/ydssx/morphix/app/{{.appName}}/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

func NewGRPCServer(c *conf.Bootstrap, svc *service.{{.serviceName}}) *grpc.Server {

	srv := common.NewGRPCServer(c.ServiceSet.{{.appName | Title}}.Server)

	{{.appName}}v1.Register{{.serviceName}}Server(srv, svc)

	return srv
}
