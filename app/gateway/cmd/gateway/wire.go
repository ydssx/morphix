//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/gateway/internal/server"
	"github.com/ydssx/morphix/common"
)

func wireApp(*common.Config) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet,newApp))
}
