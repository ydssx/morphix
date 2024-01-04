package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/util"
)

type jobHandler func(ctx context.Context, t *asynq.Task) error

var (
	// 定时任务注册
	cronJobMap = map[string]jobv1.JobType{
		"@every 5s":   jobv1.JobType_TEST_CRON_JOB,
		"*/1 * * * *": jobv1.JobType_TEST_CRON_JOB,
	}

	// 任务处理函数注册
	jobHandlerMap = map[jobv1.JobType]jobHandler{
		jobv1.JobType_TEST_JOB:      TestJobHandler,
		jobv1.JobType_TEST_CRON_JOB: TestCronJobHandler,
	}
)

// =======================================================
//
// =======================================================

func TestJobHandler(ctx context.Context, t *asynq.Task) error {
	var p jobv1.PayLoadTest
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logger.Infof(ctx, "测试job, payload: %v", string(t.Payload()))

	return nil
}

func TestCronJobHandler(ctx context.Context, _ *asynq.Task) error {
	logger.Info(ctx, "测试cronjob:"+util.GenerateCode(6))
	res, err := common.ClientSetFromContext(ctx).SendSMS(ctx, &smsv1.SendSMSRequest{Scene: smsv1.SmsScene_USER_LOGIN, MobileNumber: "123456"})
	if err != nil {
		logger.Error(ctx, err.Error())
		return err
	}
	logger.Infof(ctx, "测试测试：%v", res.String())
	return nil
}
