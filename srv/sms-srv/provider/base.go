package provider

import (
	"time"
)

type Sender interface {
	Send(sms *SMS) error
}

type SMS struct {
	Mobile  string
	Code    string
	NowTime time.Time
}
