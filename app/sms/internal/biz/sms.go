package biz

import (
	"context"
	"fmt"
	"time"

	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SmsRepo interface {
	SendSMS(ctx context.Context, req *SendSMSRequest) (string, error)
}

type SendSMSRequest struct {
	MobileNumber string
	Scene        smsv1.SmsScene
	Code         string
	Provider     string // 短信服务商，如：aliyun, tencent, etc.
}

type Transaction interface{}

type SmsUseCase struct {
	cache cache.Cache
	repo  SmsRepo
}

func NewSmsUseCase(cache cache.Cache, repo SmsRepo) *SmsUseCase {
	return &SmsUseCase{cache: cache, repo: repo}
}

func (s *SmsUseCase) SendSMS(ctx context.Context, req *smsv1.SendSMSRequest) (resp *smsv1.SendSMSResponse, err error) {
	span := trace.SpanFromContext(ctx)

	code := util.GenerateCode(6)

	_, err = s.repo.SendSMS(ctx, &SendSMSRequest{
		MobileNumber: req.MobileNumber,
		Scene:        req.Scene,
		Code:         code,
	})
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene)
	err = s.cache.Set(key, code, time.Minute*10)
	if err != nil {
		return nil, err
	}

	span.AddEvent("sms code sended", trace.WithAttributes(attribute.String("scene", req.Scene.String()), attribute.String("code", code)))

	return &smsv1.SendSMSResponse{Success: true}, nil
}

func (s *SmsUseCase) CheckSMSStatus(ctx context.Context, req *smsv1.QuerySMSStatusRequest) (*smsv1.QuerySMSStatusResponse, error) {
	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene)
	var code string
	err := s.cache.Get(key, &code)
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
