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
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/provider"
	etcdclient "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
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

	x := log.With(logger.DefaultLogger, "caller", log.Caller(5), "ts", log.Timestamp("2006-01-02 15:04:05"))
	log.SetLogger(x)
	
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
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	r := etcd.New(client)

	tp, _ := provider.InitTraceProvider("http://localhost:14268/api/traces", "user-rpc")

	mp := provider.InitMeterProvider()

	return kratos.New(
		kratos.Name(c.Name),
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
