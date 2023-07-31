package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/dapr"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/mq"
	"github.com/ydssx/morphix/pkg/provider"
	_ "go.uber.org/automaxprocs"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "f", "./../../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	var bc common.Config
	common.MustLoad(&bc, flagconf)

	app, cleanup, err := wireApp(&bc, logger.DefaultLogger)
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
	r := common.NewEtcdRegistry(c.Etcd)

	tp, _ := provider.InitTraceProvider(c.Otelcol.Addr, c.Sms.Name)

	mp := provider.InitMeterProvider(c.Otelcol.Addr)

	close, _ := mq.InitNats(c.Nats.Addr)
	closeDapr := dapr.Init()

	return kratos.New(
		kratos.Name(c.Payment.Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.Registrar(r),
		kratos.BeforeStart(func(_ context.Context) error {
			log.Infow("app.version", "1.0.0")
			return nil
		}),
		kratos.AfterStop(tp.Shutdown),
		kratos.AfterStop(mp.Shutdown),
		kratos.AfterStop(close),
		kratos.AfterStop(closeDapr),
	)
}
