// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/data"
	"github.com/ydssx/morphix/app/user/internal/server"
	"github.com/ydssx/morphix/app/user/internal/service"
	"github.com/ydssx/morphix/common"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(config *common.Config, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewGreeterRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	smsServiceClient := common.NewSMSClient(config)
	userService := service.NewUserService(userUsecase, smsServiceClient)
	grpcServer := server.NewGRPCServer(config, userService, logger)
	app := newApp(grpcServer, config)
	return app, func() {
		cleanup()
	}, nil
}
