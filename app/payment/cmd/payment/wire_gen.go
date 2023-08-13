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
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/common/dapr"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	daprClient, cleanup, err := dapr.NewDaprClient()
	if err != nil {
		return nil, nil, err
	}
	paymentEvents := service.NewEventSender(daprClient)
	paymentService := service.NewPaymentService(paymentEvents)
	grpcServer := server.NewGRPCServer(bootstrap, paymentService, logger)
	app := newApp(grpcServer, bootstrap)
	return app, func() {
		cleanup()
	}, nil
}
