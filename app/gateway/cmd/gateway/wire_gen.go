// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/ydssx/morphix/app/gateway/internal/server"
	"github.com/ydssx/morphix/common"
)

// Injectors from wire.go:

func wireApp(config *common.Config) (*kratos.App, func(), error) {
	httpServer := server.NewHTTPServer(config)
	app := newApp(httpServer, config)
	return app, func() {
	}, nil
}
