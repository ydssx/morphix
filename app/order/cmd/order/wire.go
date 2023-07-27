//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/order/internal/listener"
	"github.com/ydssx/morphix/app/order/internal/server"
	"github.com/ydssx/morphix/app/order/internal/service"
	"github.com/ydssx/morphix/common"
)

func wireApp(*common.Config, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, listener.ProviderSet, newApp))
}