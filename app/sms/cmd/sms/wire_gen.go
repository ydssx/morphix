// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/sms/internal/biz"
	"github.com/ydssx/morphix/app/sms/internal/server"
	"github.com/ydssx/morphix/app/sms/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	clusterClient := common.NewRedisCluster(bootstrap)
	smsUseCase := biz.NewSmsUseCase(clusterClient)
	smsService := service.NewSMSService(smsUseCase)
	grpcServer := server.NewGRPCServer(bootstrap, smsService)
	app := newApp(grpcServer, bootstrap)
	return app, func() {
	}, nil
}
