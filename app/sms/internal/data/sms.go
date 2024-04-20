package data

import (
	"context"

	"github.com/ydssx/morphix/app/sms/internal/biz"
	"github.com/ydssx/morphix/pkg/client/tencentcloud"
)

type smsRepo struct {
	data          *Data
	tencentClient *tencentcloud.TencentCloud
}

var _ biz.SmsRepo = (*smsRepo)(nil)

func NewSmsRepo(data *Data, tencentClient *tencentcloud.TencentCloud) biz.SmsRepo {
	return &smsRepo{
		data:          data,
		tencentClient: tencentClient,
	}
}

func (s *smsRepo) SendSMS(ctx context.Context, req *biz.SendSMSRequest) (string, error) {
	return "", nil
}
