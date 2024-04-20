package tencentcloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/ydssx/morphix/pkg/errors"
)

type TencentCloud struct {
	smsClient *sms.Client
}

func New() *TencentCloud {
	credential := common.NewCredential("your-secret-id", "your-secret-key")
	cpf := profile.NewClientProfile()

	smsClient, err := sms.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		panic("init tencentcloud sms client failed: " + err.Error())
	}

	return &TencentCloud{smsClient: smsClient}
}

func (t *TencentCloud) SendSMS(phoneNumber string, templateID string, params []string) error {
	request := sms.NewSendSmsRequest()
	request.TemplateId = &templateID
	request.SmsSdkAppId = common.StringPtr("your-sms-sdk-appid")
	request.SignName = common.StringPtr("your-sign-name")
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumber})
	request.TemplateParamSet = common.StringPtrs(params)
	_, err := t.smsClient.SendSms(request)
	if err != nil {
		return errors.Wrap(err, "send tencentcloud sms failed")
	}
	return nil
}
