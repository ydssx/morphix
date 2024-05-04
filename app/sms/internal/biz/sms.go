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

// SendSMS sends a SMS to the mobile number with the specified scene and code.
func (s *SmsUseCase) SendSMS(ctx context.Context, req *smsv1.SendSMSRequest) (resp *smsv1.SendSMSResponse, err error) {
	// extract the trace span from the context
	span := trace.SpanFromContext(ctx)

	// generate a 6-digit code
	code := util.GenerateCode(6)

	// send the SMS to the user
	_, err = s.repo.SendSMS(ctx, &SendSMSRequest{
		MobileNumber: req.MobileNumber, // the mobile number to send the SMS to
		Scene:        req.Scene,        // the scene of the SMS
		Code:         code,             // the code to send to the user
	})
	if err != nil {
		return nil, err // return the error if something went wrong
	}

	// set the code in the cache with a 10-minute expiration
	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene) // key is the format "<mobile number>-<scene>"
	err = s.cache.Set(key, code, time.Minute*10)             // set the code in the cache
	if err != nil {
		return nil, err // return the error if something went wrong
	}

	// add an event to the span with the sent code and scene
	span.AddEvent("sms code sended", trace.WithAttributes(
		attribute.String("scene", req.Scene.String()), // the scene of the SMS
		attribute.String("code", code),                // the code sent to the user
	))

	// return a successful response
	return &smsv1.SendSMSResponse{Success: true}, nil
}

// CheckSMSStatus checks if the SMS code sent to the specified mobile number and scene is correct.
// It returns an error if the code is incorrect or the cached code does not exist.
func (s *SmsUseCase) CheckSMSStatus(ctx context.Context, req *smsv1.QuerySMSStatusRequest) (*smsv1.QuerySMSStatusResponse, error) {
	// generate the cache key
	key := fmt.Sprintf("%s-%s", req.MobileNumber, req.Scene)

	// get the code from the cache
	var code string
	err := s.cache.Get(key, &code)
	if err != nil {
		return nil, err // return the error if the code does not exist in the cache
	}

	// check if the code is correct
	if code != req.SmsCode {
		return nil, status.Error(codes.Aborted, "验证码错误") // return an Aborted error if the code is incorrect
	}

	// return a successful response if the code is correct
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
