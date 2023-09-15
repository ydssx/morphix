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
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/logger"
)

type JobServer struct {
	sr  *asynq.Server
	mux *asynq.ServeMux
}

func NewJobServer(c *conf.Bootstrap) *JobServer {
	opt := common.InitRedisOpt(c)
	server := asynq.NewServer(opt, asynq.Config{Concurrency: 10, ErrorHandler: asynq.ErrorHandlerFunc(reportError)})

	mux := asynq.NewServeMux()
	handler.RegisterJobHandler(mux)
	go NewClient(opt)
	return &JobServer{sr: server, mux: mux}
}

func (j *JobServer) Start(_ context.Context) error {
	return j.sr.Start(j.mux)
}

func (j *JobServer) Stop(_ context.Context) error {
	j.sr.Stop()
	j.sr.Shutdown()
	return nil
}

func reportError(ctx context.Context, task *asynq.Task, err error) {
	logger.Errorf(ctx, "执行任务失败,task_type:%s ,err: %v", task.Type(), err)
}

func NewClient(redisClientOpt asynq.RedisClientOpt) {
	cli := asynq.NewClient(redisClientOpt)
	for {
		payload, _ := json.Marshal(jobv1.PayLoadTest{Msg: "test msg"})
		_, err := cli.Enqueue(asynq.NewTask(jobv1.JobType_TEST_JOB.String(), payload))
		if err != nil {
			log.Print(err)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
