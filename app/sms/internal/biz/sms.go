package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/pkg/util"
)

type SmsUseCase struct {
	rdb *redis.Client
}

func NewSmsUseCase(rdb *redis.Client) *SmsUseCase {
	return &SmsUseCase{rdb: rdb}
}

func (s *SmsUseCase) SendSMS(ctx context.Context, req *smsv1.SendSMSRequest) (resp *smsv1.SendSMSResponse, err error) {
	code := util.GenerateCode(6)
	log.Info("发送短信验证码:", code)
	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene)
	_, err = s.rdb.Set(ctx, key, code, time.Minute*10).Result()
	if err != nil {
		return nil, err
	}
	
	return &smsv1.SendSMSResponse{Success: true}, nil
}
