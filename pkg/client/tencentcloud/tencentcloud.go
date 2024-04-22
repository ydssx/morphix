package tencentcloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/ydssx/morphix/pkg/errors"
)

type TencentCloud struct {
	smsClient *sms.Client
}

func New() *TencentCloud {
	credential := common.NewCredential("your-secret-id", "your-secret-key")
	cpf := profile.NewClientProfile()

	smsClient, err := sms.NewClient(credential, regions.Guangzhou, cpf)
	if err != nil {
		panic("init tencentcloud sms client failed: " + err.Error())
	}

	return &TencentCloud{smsClient: smsClient}
}


// SendSMS sends a sms using tencentcloud sms service.
//
// phoneNumber is the mobile number of the recipient.
//
// templateID is the template ID of the sms.
//
// params is the template parameters of the sms.
func (t *TencentCloud) SendSMS(phoneNumber string, templateID string, params []string) error {
	request := sms.NewSendSmsRequest()
	request.TemplateId = &templateID
	request.SmsSdkAppId = common.StringPtr("your-sms-sdk-appid") // replace with your sms sdk appid
	request.SignName = common.StringPtr("your-sign-name")       // replace with your sign name
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumber})
	request.TemplateParamSet = common.StringPtrs(params)
	_, err := t.smsClient.SendSms(request)
	if err != nil {
		return errors.Wrap(err, "send tencentcloud sms failed")
	}
	return nil
}
