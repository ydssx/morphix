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
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/cache"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewRedisCLient(bootstrap)
	dataData, cleanup, err := data.NewData(logger, client)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	cacheCache := cache.NewRedisCache(client)
	userRepoWithCache := data.NewUserRepoCacheDecorator(userRepo, cacheCache)
	userUsecase := biz.NewUserUsecase(userRepoWithCache, logger)
	smsServiceClient := common.NewSMSClient(bootstrap)
	userService := service.NewUserService(userUsecase, smsServiceClient)
	grpcServer := server.NewGRPCServer(bootstrap, userService)
	app := newApp(grpcServer, bootstrap)
	return app, func() {
		cleanup()
	}, nil
}
