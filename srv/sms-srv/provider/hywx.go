package provider

/**
互亿无线 短信通道
http://www.ihuyi.cn/
**/

type Hywx struct {
	sms *SMS
}

func (h *Hywx) Send(sms *SMS) error {

	return nil
}
