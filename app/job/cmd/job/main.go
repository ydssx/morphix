package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
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
	var c conf.Bootstrap
	close := conf.MustLoad(&c, flagconf)
	defer close()

	app, cleanup, err := wireApp(&c)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}

func newApp(ls *server.JobServer, gs *grpc.Server, c *conf.Bootstrap) *kratos.App {
	r := common.NewEtcdRegistry(c.Etcd)

	tp, _ := provider.InitTraceProvider(c.Otelcol.Addr, "morphix-job")

	return kratos.New(
		kratos.Name("morphix-job"),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, ls),
		kratos.Registrar(r),
		kratos.BeforeStart(func(_ context.Context) error {
			log.Infow("app.version", "1.0.0")
			return nil
		}),
		kratos.AfterStop(tp.Shutdown),
	)
}
