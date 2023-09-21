package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	"github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/app/job/internal/handler"
	"github.com/ydssx/morphix/app/job/internal/service"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/util"
	"golang.org/x/sync/errgroup"
)

type JobServer struct {
	sr  *asynq.Server
	sd  *asynq.Scheduler
	mux *asynq.ServeMux
}

func NewJobServer(c *conf.Bootstrap) *JobServer {
	opt := common.InitRedisOpt(c)

	clientSet := common.NewServiceClientSet(c)

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

func (j *JobServer) Start(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(j.sd.Start)
	eg.Go(func() error { return j.sr.Start(j.mux) })
	return eg.Wait()
}

func (j *JobServer) Stop(_ context.Context) error {
	j.sr.Stop()
	j.sr.Shutdown()
	j.sd.Shutdown()
	return nil
}

func reportError(ctx context.Context, task *asynq.Task, err error) {
	logger.Errorf(ctx, "执行任务失败,task_type:%s ,err: %v", task.Type(), err)
}

func NewClient(redisClientOpt asynq.RedisClientOpt) {
	cli := asynq.NewClient(redisClientOpt)
	srv := service.NewJobService(cli,nil)
	for {
		payload, _ := json.Marshal(jobv1.PayLoadTest{Msg: "test msg:" + util.GenerateCode(6)})
		_, err := srv.Enqueue(context.Background(), &jobv1.EnqueueRequest{JobType: jobv1.JobType_TEST_JOB, Payload: payload})
		if err != nil {
			log.Print(err)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
