// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/order/internal/biz"
	"github.com/ydssx/morphix/app/order/internal/listener"
	"github.com/ydssx/morphix/app/order/internal/server"
	"github.com/ydssx/morphix/app/order/internal/service"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	orderUseCase := biz.NewOrderUseCase()
	orderService := service.NewOrderService(orderUseCase)
	grpcServer := server.NewGRPCServer(bootstrap, orderService)
	conn, cleanup, err := common.NewNatsConn(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	cloudEvent := common.NewCloudEvent(conn)
	listenerServer := listener.NewListenerServer(cloudEvent)
	app := newApp(grpcServer, listenerServer, bootstrap)
	return app, func() {
		cleanup()
	}, nil
}
