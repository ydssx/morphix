package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	"github.com/ydssx/morphix/pkg/logger"
)

type handler func(ctx context.Context, t *asynq.Task) error

var (
	cronJobMap = map[string]jobv1.JobType{
		"@every 5s":   jobv1.JobType_TEST_CRON_JOB,
		"*/1 * * * *": jobv1.JobType_TEST_CRON_JOB,
	}

	jobHandlerMap = map[jobv1.JobType]handler{
		jobv1.JobType_TEST_JOB:      TestJobHandler,
		jobv1.JobType_TEST_CRON_JOB: TestCronJobHandler,
	}
)

func TestJobHandler(ctx context.Context, t *asynq.Task) error {
	var p jobv1.PayLoadTest
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logger.Infof(ctx, "测试job, payload: %v", string(t.Payload()))

	return nil
}

func TestCronJobHandler(ctx context.Context, _ *asynq.Task) error {
	logger.Info(ctx, "测试cronjob")

	return nil
}
