package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/provider"
)

var configFile = flag.String("f", "../../../../configs/config.yaml", "the config file")

// main 函数是应用程序的入口点。它解析命令行标志、加载配置、初始化应用程序并运行它。
// 如果在初始化或运行过程中出现任何错误,它会立即panic。
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

// newApp creates a new Kratos application.
//
// It takes an http.Server and a conf.Bootstrap as parameters.
// Returns a pointer to a kratos.App.
func newApp(hs *http.Server, c *conf.Bootstrap) *kratos.App {
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
