package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
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
	
	var config conf.Bootstrap
	closeConfig := conf.MustLoad(&config, flagconf)
	defer closeConfig()

	application, cleanup, err := wireApp(&config, logger.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// Start the application and wait for stop signal
	if err := application.Run(); err != nil {
		panic(err)
	}
}

// newApp 创建一个新的 Kratos 应用程序实例。它接收 gRPC server 和配置对象作为参数。
// 该函数首先创建服务发现注册表。然后初始化 trace 和 metric 提供者。
// 最后使用提供的参数创建并返回一个 Kratos 应用程序实例。该实例中包含了服务名、元数据、
// gRPC server、注册表、启动前后钩子等信息。
func newApp(gs *grpc.Server, c *conf.Bootstrap) *kratos.App {
	r := common.NewEtcdRegistry(c.Etcd)

	tp, _ := provider.InitTraceProvider(c.Otelcol.Addr, c.ServiceSet.User.Name)

	mp := provider.InitMeterProvider(c.Otelcol.Addr)

	return kratos.New(
		kratos.Name(c.ServiceSet.User.Name),
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

