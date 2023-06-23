package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/provider"
	etcdclient "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "f", "../../../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	var c common.Config
	common.MustLoad(&c, flagconf)

	app, cleanup, err := wireApp(&c, logger.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(gs *grpc.Server, c *common.Config) *kratos.App {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: c.Etcd.Endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	r := etcd.New(client)

	tp, _ := provider.InitTraceProvider(c.Jaeger.Addr, c.User.Name)

	mp := provider.InitMeterProvider()

	return kratos.New(
		kratos.Name(c.User.Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.Registrar(r),
		kratos.BeforeStart(func(ctx context.Context) error {
			log.Infow("app.version", "1.0.0")
			return nil
		}),
		kratos.AfterStop(tp.Shutdown),
		kratos.AfterStop(mp.Shutdown),
	)
}
