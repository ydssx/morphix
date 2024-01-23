//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/product/internal/biz"
	"github.com/ydssx/morphix/app/product/internal/data"
	"github.com/ydssx/morphix/app/product/internal/server"
	"github.com/ydssx/morphix/app/product/internal/service"
	"github.com/ydssx/morphix/common/conf"
)

func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
