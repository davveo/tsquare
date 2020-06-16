package provider

/*
腾讯云短信
*/

type QcloudSms struct {
}

func (qcloud QcloudSms) Send(sms *SMS) error {
	return nil
}
