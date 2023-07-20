// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/payment/internal/server"
	"github.com/ydssx/morphix/app/payment/internal/service"
	"github.com/ydssx/morphix/common"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(config *common.Config, logger log.Logger) (*kratos.App, func(), error) {
	paymentService := service.NewPaymentService()
	grpcServer := server.NewGRPCServer(config, paymentService, logger)
	app := newApp(grpcServer, config)
	return app, func() {
	}, nil
}