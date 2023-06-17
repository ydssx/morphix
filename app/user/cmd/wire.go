//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/app/user/internal/data"
	"github.com/ydssx/morphix/app/user/internal/server"
	"github.com/ydssx/morphix/app/user/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger,*zap.Logger, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
