package sendsms

import (
	"happylemon/conf"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

// SendSms 发送短信接口
// phoneNumbers 短信发送的号码列表，必填
// templateCode 申请的短信模板编码,必填
// templateParam 短信模板变量参数,json字符串
//返回状态码
func SendSms(mobile, templateCode, templateParam string) (int, string) {
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(conf.Config.Sms.AccessKeyId, conf.Config.Sms.AccessKeySecret)
	respSendSms, err := dysms.SendSms(uuid.New(), mobile, conf.Config.Sms.Signname, templateCode, templateParam).DoActionWithException()
	if err != nil {
		return conf.CodeSmsSendErr, respSendSms.Error()
	}
	if *respSendSms.Code != "OK" {
		//小时级限制：每小时5条
		if *respSendSms.Message == "触发小时级流控Permits:5" {
			return conf.CodeSmsLimitHour, respSendSms.Error()
		} else if *respSendSms.Message == "触发分钟级流控Permits:1" {
			return conf.CodeSmsLimitMinute, respSendSms.Error()
		} else if *respSendSms.Message == "触发天级流控Permits:10" {
			return conf.CodeSmsLimitDay, respSendSms.Error()
		}
		return conf.CodeSmsSendErr, respSendSms.Error()
	}

	return conf.CodeOk, respSendSms.GetRequestID()
}
