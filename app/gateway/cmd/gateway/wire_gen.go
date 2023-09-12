// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/ydssx/morphix/app/gateway/internal/server"
	"github.com/ydssx/morphix/common/conf"
)

// Injectors from wire.go:

func wireApp(ctx context.Context, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	httpServer := server.NewHTTPServer(ctx, bootstrap)
	app := newApp(httpServer, bootstrap)
	return app, func() {
	}, nil
}
