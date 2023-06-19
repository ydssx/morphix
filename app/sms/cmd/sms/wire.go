//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/sms/internal/conf"
	"github.com/ydssx/morphix/app/sms/internal/server"
	"github.com/ydssx/morphix/app/sms/internal/service"
)

func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet,service.ProviderSet,newApp))
}
