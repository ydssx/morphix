// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/aiart/internal/biz"
	"github.com/ydssx/morphix/app/aiart/internal/server"
	"github.com/ydssx/morphix/app/aiart/internal/service"
	"github.com/ydssx/morphix/common/conf"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	aiartUseCase := biz.NewAiartUseCase()
	artService := service.NewArtService(aiartUseCase)
	httpServer := server.NewHTTPServer(bootstrap, artService)
	grpcServer := server.NewGRPCServer(bootstrap, artService)
	v := server.NewServer(httpServer, grpcServer)
	app := newApp(bootstrap, v...)
	return app, func() {
	}, nil
}
