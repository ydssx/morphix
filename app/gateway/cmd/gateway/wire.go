//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/gateway/internal/server"
	"github.com/ydssx/morphix/common/conf"
)

func wireApp(ctx context.Context, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
