package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
	"github.com/ydssx/morphix/app/job/internal/handler"
	"github.com/ydssx/morphix/common/conf"
	"github.com/ydssx/morphix/pkg/logger"
)

type JobServer struct {
	sr  *asynq.Server
	mux *asynq.ServeMux
}

func NewJobServer(c *conf.Bootstrap) *JobServer {
	redisConf := c.Redis
	redisClientOpt := asynq.RedisClientOpt{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       1,
	}

	server := asynq.NewServer(redisClientOpt, asynq.Config{Concurrency: 10, ErrorHandler: asynq.ErrorHandlerFunc(reportError)})

	mux := asynq.NewServeMux()
	handler.RegisterJobHandler(mux)
	NewClient(c)
	return &JobServer{sr: server,mux: mux}
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

func NewClient(c *conf.Bootstrap) {
	redisConf := c.Redis
	redisClientOpt := asynq.RedisClientOpt{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       1,
	}
	cli := asynq.NewClient(redisClientOpt)
	payload, _ := json.Marshal(handler.Payload{Msg: "test msg"})
	_, err := cli.Enqueue(asynq.NewTask(handler.TestJob, payload))
	if err != nil {
		log.Print(err)
	}
}
