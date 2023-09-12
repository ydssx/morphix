package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/ydssx/morphix/pkg/logger"
)

type handler func(ctx context.Context, t *asynq.Task) error

const (
	TestJob     = "job.test"
	CronJobTest = "cronjob.test"
)

var (
	cronJobMap = map[string]string{
		"@every 5s":  CronJobTest,
		"*/1 * * * *": CronJobTest,
	}

	jobHandlerMap = map[string]handler{
		TestJob:     TestJobHandler,
		CronJobTest: TestCronJobHandler,
	}
)

type Payload struct {
	Msg string
}

func TestJobHandler(ctx context.Context, t *asynq.Task) error {
	var p Payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logger.Infof(ctx, "测试job, payload: %v", p)

	return nil
}

func TestCronJobHandler(ctx context.Context, _ *asynq.Task) error {
	logger.Info(ctx, "测试cronjob")

	return nil
}
