package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/app/job/internal/server"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/provider"
	_ "go.uber.org/automaxprocs"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "f", "../../../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	// Load configuration
	var config conf.Bootstrap
	closeConfig := conf.MustLoad(&config, flagconf)
	defer closeConfig()

	// Create and initialize the application
	app, cleanup, err := wireApp(&config)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// Run the application
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(ls *server.JobServer, gs *grpc.Server, ps *server.ListenerServer, c *conf.Bootstrap) *kratos.App {
	etcdRegistry := common.NewEtcdRegistry(c.Etcd)
	traceProvider, _ := provider.InitTraceProvider(c.Otelcol.Addr, "morphix-job")

	return kratos.New(
		kratos.Name("morphix-job"),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, ls, ps),
		kratos.Registrar(etcdRegistry),
		kratos.BeforeStart(func(_ context.Context) error {
			return nil
		}),
		kratos.AfterStop(traceProvider.Shutdown),
	)
}
