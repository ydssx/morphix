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
	registerJobHandler(mux)

	scheduler := asynq.NewScheduler(opt, &asynq.SchedulerOpts{Location: time.Local})
	registerCronJob(scheduler)

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

// registerJobHandler 注册 jobHandlerMap 中定义的所有 job 的处理函数到 ServeMux。
// 它会遍历 jobHandlerMap,并为每个 job 注册对应的处理函数到 mux。
// mux 会根据请求中的 job name 来路由到相应的处理函数。
func registerJobHandler(mux *asynq.ServeMux) {
	for k, v := range handler.JobHandlerMap {
		mux.HandleFunc(k.String(), v)
	}
}

// registerCronJob 注册定时任务的处理函数。
// 它会遍历 cronJobMap 中定义的所有定时任务,并在调度器 sd 中注册对应的处理函数。
// 如果某个定时任务在 jobHandlerMap 中没有找到对应的处理函数,会 panic。
// 注册成功后,定时任务会按照 cronJobMap 中定义的时间表定期执行。
func registerCronJob(sd *asynq.Scheduler) {
	for k, jobType := range handler.CronJobMap {
		err := handler.ValidateTask(jobType)
		if err != nil {
			panic(err)
		}
		_, err = sd.Register(k, asynq.NewTask(jobType.String(), nil))
		if err != nil {
			panic(err)
		}
	}
}
