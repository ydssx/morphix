package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/logger"
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

func newApp(c *conf.Bootstrap, srv ...transport.Server) *kratos.App {
	options := []kratos.Option{
		kratos.Name(c.ServiceSet.Aiart.Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(srv...),
		kratos.BeforeStart(func(ctx context.Context) error {
			logger.Infof(ctx, "service %s is starting...", c.ServiceSet.Aiart.Name)
			return nil
		}),
	}

	if c.ServiceSet.Aiart.EnableRegistry {
		registry := common.NewEtcdRegistry(c.Etcd)
		options = append(options, kratos.Registrar(registry))
	}
	if c.ServiceSet.Aiart.EnableTracing {
		traceProvider, _ := provider.InitTraceProvider(c.Otelcol.Addr, c.ServiceSet.Aiart.Name)
		options = append(options, kratos.AfterStop(traceProvider.Shutdown))
	}
	if c.ServiceSet.Aiart.EnableMetric {
		meterProvider := provider.InitMeterProvider(c.Otelcol.Addr)
		options = append(options, kratos.AfterStop(meterProvider.Shutdown))
	}

	return kratos.New(options...)
}
