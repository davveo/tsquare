package request

type LoginRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"pwd"`
}

type SmsRequest struct {
	Mobile string `json:"mobile"`
}
