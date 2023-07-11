package service

import (
	"context"

	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	"github.com/ydssx/morphix/app/sms/internal/biz"
)

var _ smsv1.SMSServiceServer = (*SMSService)(nil)

type SMSService struct {
	smsv1.UnimplementedSMSServiceServer

	uc *biz.SmsUseCase
}

func NewSMSService(uc *biz.SmsUseCase) *SMSService {
	return &SMSService{uc: uc}
}

// ManageSMSSignature implements smsv1.SMSServiceServer.
func (*SMSService) ManageSMSSignature(context.Context, *smsv1.SignatureManagementRequest) (*smsv1.SignatureManagementResponse, error) {
	panic("unimplemented")
}

// ManageSMSTemplate implements smsv1.SMSServiceServer.
func (*SMSService) ManageSMSTemplate(context.Context, *smsv1.TemplateManagementRequest) (*smsv1.TemplateManagementResponse, error) {
	panic("unimplemented")
}

// QuerySMSStatus implements smsv1.SMSServiceServer.
func (*SMSService) QuerySMSStatus(context.Context, *smsv1.QuerySMSStatusRequest) (*smsv1.QuerySMSStatusResponse, error) {
	panic("unimplemented")
}

// SendSMS implements smsv1.SMSServiceServer.
func (s *SMSService) SendSMS(ctx context.Context, req *smsv1.SendSMSRequest) (resp *smsv1.SendSMSResponse, err error) {
	return s.uc.SendSMS(ctx, req)
}
