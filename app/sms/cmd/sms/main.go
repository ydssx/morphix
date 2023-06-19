package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/app/sms/internal/conf"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/provider"
	etcdclient "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "-f", "../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c := config.New(config.WithSource(file.NewSource(flagconf)))
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger.DefaultLogger, &bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}

func newApp(gs *grpc.Server, c *conf.Bootstrap) *kratos.App {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: c.Etcd.Endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	r := etcd.New(client)

	tp, _ := provider.InitTraceProvider(c.Jeager.Addr, c.Name)

	mp := provider.InitMeterProvider()

	return kratos.New(
		kratos.Name(c.Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.Registrar(r),
		kratos.BeforeStart(func(_ context.Context) error {
			log.Infow("app.version", "1.0.0")
			return nil
		}),
		kratos.AfterStop(tp.Shutdown),
		kratos.AfterStop(mp.Shutdown),
	)
}
