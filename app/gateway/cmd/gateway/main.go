package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ydssx/morphix/common"
	kmiddleware "github.com/ydssx/morphix/pkg/middleware/kratos"
	"github.com/ydssx/morphix/pkg/provider"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var configFile = flag.String("f", "../../../../configs/config.yaml", "the config file")

func main() {
	flag.Parse()

	var config common.Config
	common.MustLoad(&config, *configFile)

	if err := Run(context.Background(), config); err != nil {
		panic(err)
	}
}

func Run(ctx context.Context, c common.Config) error {
	registerRpcServer(c)

	tp, err := provider.InitTraceProvider(c.Jaeger.Addr, c.Gateway.Name)
	if err != nil {
		panic(err)
	}
	mp := provider.InitMeterProvider()

	server := gin.New()
	server.Use(gin.Logger(), ginprom.PromMiddleware(nil), gin.Recovery())
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	opts := []gwruntime.ServeMuxOption{}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Endpoints,
		DialTimeout: time.Second * time.Duration(c.Etcd.Timeout),
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)

	gw, err := newGateway(ctx, opts, r)
	if err != nil {
		return err
	}
	server.Any("/api/*any", gin.WrapH(gw))

	httpSrv := khttp.NewServer(khttp.Address(c.Gateway.Addr), khttp.Middleware(kmiddleware.MetricServer()))
	openAPIhandler := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", openAPIhandler)

	httpSrv.HandlePrefix("/", server)

	app := kratos.New(
		kratos.Name(c.Gateway.Name),
		kratos.Context(ctx),
		kratos.Server(
			httpSrv,
		),
		kratos.AfterStop(tp.Shutdown),
		kratos.AfterStop(mp.Shutdown),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}
