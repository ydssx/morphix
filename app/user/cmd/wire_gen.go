// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/app/user/internal/data"
	"github.com/ydssx/morphix/app/user/internal/server"
	"github.com/ydssx/morphix/app/user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewGreeterRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	app := newApp(grpcServer, bootstrap)
	return app, func() {
		cleanup()
	}, nil
}
