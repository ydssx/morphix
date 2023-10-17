//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"{{.module}}/app/{{.appName}}/internal/server"
	"{{.module}}/app/{{.appName}}/internal/service"
	"{{.module}}/app/{{.appName}}/internal/biz"
	"{{.module}}/common/conf"
)

func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet,service.ProviderSet,biz.ProviderSet,newApp))
}