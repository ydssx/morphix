package server

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/app/job/internal/handler"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/concurrent"
	"github.com/ydssx/morphix/pkg/logger"
)

type JobServer struct {
	sr  *asynq.Server
	sd  *asynq.Scheduler
	mux *asynq.ServeMux
}

func NewJobServer(c *conf.Bootstrap, clientSet *common.ServiceClientSet) *JobServer {
	opt := common.InitRedisOpt(c)

	// clientSet := common.NewServiceClientSet(c)

	server := asynq.NewServer(opt, asynq.Config{
		Concurrency:  10,
		ErrorHandler: asynq.ErrorHandlerFunc(reportError),
		BaseContext: func() context.Context {
			return common.NewContextWithServiceClientSet(context.Background(), clientSet)
		},
	})

	mux := asynq.NewServeMux()
	handler.RegisterJobHandler(mux)

	scheduler := asynq.NewScheduler(opt, &asynq.SchedulerOpts{Location: time.Local})
	handler.RegisterCronJob(scheduler)

	return &JobServer{sr: server, mux: mux, sd: scheduler}
}

// Start 启动 JobServer,包括启动调度器和服务器。
// 使用 concurrent.Group 同时启动调度器和服务器,如果任一启动失败则立即返回错误。
// 这样可以保证整个 JobServer 要么完全启动成功,要么完全失败。
func (j *JobServer) Start(ctx context.Context) error {
	group := concurrent.NewGroup(ctx, concurrent.WithFastFail(true))
	err := group.Run(
		j.sd.Start,
		func() error { return j.sr.Start(j.mux) },
	)
	return err
}

// Stop 停止 JobServer,包括停止调度器和服务器。
// 依次调用服务器和调度器的 Shutdown 方法进行优雅停止。
func (j *JobServer) Stop(_ context.Context) error {
	j.sr.Stop()
	j.sr.Shutdown()
	j.sd.Shutdown()
	return nil
}

func reportError(ctx context.Context, task *asynq.Task, err error) {
	logger.Errorf(ctx, "执行任务失败,task_type:%s ,err: %v", task.Type(), err)
}
