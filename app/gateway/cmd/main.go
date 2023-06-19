package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ydssx/morphix/app/gateway/conf"
	kmiddleware "github.com/ydssx/morphix/pkg/middleware/kratos"
	"github.com/ydssx/morphix/pkg/provider"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var configFile = flag.String("f", "../configs/config.yaml", "the config file")

func main() {
	var config conf.Config
	conf.MustLoad(*configFile, &config)

	if err := Run(context.Background(), config); err != nil {
		panic(err)
	}
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, c conf.Config) error {
	registerRpcServer(c)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tp, err := provider.InitTraceProvider("http://localhost:14268/api/traces", "gateway")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())
	server.Any("/metrics", gin.WrapH(promhttp.Handler()))

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

	httpSrv := khttp.NewServer(khttp.Address(c.Addr), khttp.Middleware(kmiddleware.MetricServer()))
	openAPIhandler := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", openAPIhandler)
	httpSrv.HandlePrefix("/", server)

	app := kratos.New(
		kratos.Name(c.Name),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}
