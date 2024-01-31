package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/pkg/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SmsRepo interface{}

type Transaction interface{}

type SmsUseCase struct {
	rdb *redis.Client
}

func NewSmsUseCase(rdb *redis.Client) *SmsUseCase {
	return &SmsUseCase{rdb: rdb}
}

func (s *SmsUseCase) SendSMS(ctx context.Context, req *smsv1.SendSMSRequest) (resp *smsv1.SendSMSResponse, err error) {
	span := trace.SpanFromContext(ctx)

	code := util.GenerateCode(6)
	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene)
	_, err = s.rdb.Set(ctx, key, code, time.Minute*10).Result()
	if err != nil {
		return nil, err
	}
	span.AddEvent("sms code sended", trace.WithAttributes(attribute.String("scene", req.Scene.String()), attribute.String("code", code)))

	return &smsv1.SendSMSResponse{Success: true}, nil
}

func (s *SmsUseCase) CheckSMSStatus(ctx context.Context, req *smsv1.QuerySMSStatusRequest) (*smsv1.QuerySMSStatusResponse, error) {
	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene)
	code, err := s.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, status.Error(codes.NotFound, "验证码不存在或已失效")
	}
	if err != nil {
		return nil, err
	}
	if code != req.SmsCode {
		return nil, status.Error(codes.Aborted, "验证码错误")
	}
	return &smsv1.QuerySMSStatusResponse{Status: true}, nil
}

// 管理短信模板
func (uc *SmsUseCase) ManageSMSTemplate(ctx context.Context, req *smsv1.TemplateManagementRequest) (res *smsv1.TemplateManagementResponse, err error) {
	res = new(smsv1.TemplateManagementResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 管理短信签名
func (uc *SmsUseCase) ManageSMSSignature(ctx context.Context, req *smsv1.SignatureManagementRequest) (res *smsv1.SignatureManagementResponse, err error) {
	res = new(smsv1.SignatureManagementResponse)

	// TODO:ADD logic here and delete this line.

	return
}
