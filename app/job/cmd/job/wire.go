//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/job/internal/server"
	"github.com/ydssx/morphix/app/job/internal/service"
	"github.com/ydssx/morphix/common/conf"
)

func wireApp(*conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
