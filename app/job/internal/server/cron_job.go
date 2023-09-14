package server

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/app/job/internal/handler"
	"github.com/ydssx/morphix/common/conf"
)

type CronJobServer struct {
	sd *asynq.Scheduler
}

func NewCronJobServer(c *conf.Bootstrap) *CronJobServer {
	opt := common.InitRedisOpt(c)
	scheduler := asynq.NewScheduler(opt, &asynq.SchedulerOpts{Location: time.Local})
	handler.RegisterCronJob(scheduler)

	return &CronJobServer{sd: scheduler}
}

func (j *CronJobServer) Start(_ context.Context) error {
	return j.sd.Start()
}

func (j *CronJobServer) Stop(_ context.Context) error {
	j.sd.Shutdown()
	return nil
}
