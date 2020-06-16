package handler

import (
	"context"
	"github.com/zbrechave/tsquare/srv/sms/provider"
	"github.com/zbrechave/tsquare/srv/sms/utils"
	//"github.com/zbrechave/tsquare/srv/sms-srv/provider"

	log "github.com/micro/go-micro/v2/logger"

	sms_proto "github.com/zbrechave/tsquare/srv/sms/proto/sms"
)

type Sms struct{}

func (s *Sms) Send(ctx context.Context, req *sms_proto.Request, rsp *sms_proto.Response) error {
	log.Info("Received Sms.Send request")

	code := utils.GenVerificationCode()

	sms := provider.SMS{
		Mobile: req.Mobile,
		Code: code,
	}

	err := provider.Alidayu{}.Send(&sms)
	if err != nil {
		rsp.Error = &sms_proto.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}
	return nil
}
