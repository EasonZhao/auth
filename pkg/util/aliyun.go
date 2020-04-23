package util

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

const (
	regionID     = "cn-hangzhou"
	accessKeyID  = "LTAI4G9ZrfzVCmWKZ3n4zw3a"
	accessSecret = "JAFwnlk4Rc2flBkRYSLw17ww66VdyH"
	templateCode = "SMS_188642192"
	signName     = "SecurityIn"
)

// SendCodeByAli 通过阿里云发送验证码
func SendCodeByAli(phone string, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey(regionID, accessKeyID, accessSecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	_, err = client.SendSms(request)
	return err
}
