package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ydssx/morphix/app/user/internal/conf"
	"github.com/ydssx/morphix/pkg/provider"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
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

	tp, _ := provider.InitTraceProvider("http://localhost:14268/api/traces", "user-rpc")
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Errorf("Error shutting down tracer provider: %v", err)
		}
	}()

	mp := provider.InitMeterProvider()
	defer func() {
		if err := mp.Shutdown(context.Background()); err != nil {
			log.Errorf("Error shutting down meter provider: %v", err)
		}
	}()
	
	app, cleanup, err := wireApp(bc.Server, bc.Data, log.DefaultLogger, zap.NewExample())
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.BeforeStart(func(ctx context.Context) error {
			log.Infow("app.version", "1.0.0")
			return nil
		}),
	)
}
