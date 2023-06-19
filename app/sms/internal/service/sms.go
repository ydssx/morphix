package service

import (
	"context"

	smsv1 "github.com/ydssx/morphix/api/sms/v1"
)

var _ smsv1.SMSServiceServer = (*SMSService)(nil)

type SMSService struct {
	smsv1.UnimplementedSMSServiceServer
}

func NewSMSService() *SMSService {
	return &SMSService{}
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
func (*SMSService) SendSMS(ctx context.Context, req *smsv1.SendSMSRequest) (resp *smsv1.SendSMSResponse, err error) {
	return &smsv1.SendSMSResponse{Success: false}, nil
}
