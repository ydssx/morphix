package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/provider"
)

var configFile = flag.String("f", "../../../../configs/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c conf.Bootstrap
	close := conf.MustLoad(&c, *configFile)
	defer close()

	app, cleanup, err := wireApp(context.Background(), &c)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(hs *khttp.Server, c *conf.Bootstrap) *kratos.App {

	tp, err := provider.InitTraceProvider(c.Otelcol.Addr, c.ServiceSet.Gateway.Name)
	if err != nil {
		panic(err)
	}
	mp := provider.InitMeterProvider(c.Otelcol.Addr)

	app := kratos.New(
		kratos.Name(c.ServiceSet.Gateway.Name),
		kratos.Server(hs),
		kratos.AfterStop(tp.Shutdown),
		kratos.AfterStop(mp.Shutdown),
	)

	return app
}
