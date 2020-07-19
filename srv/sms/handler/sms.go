package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davveo/tsquare/lib/redis"
	"github.com/davveo/tsquare/srv/sms/provider"
	"github.com/davveo/tsquare/srv/sms/utils"

	//"github.com/zbrechave/tsquare/srv/sms-srv/provider"

	smsproto "github.com/davveo/tsquare/proto/sms"
	log "github.com/micro/go-micro/v2/logger"
)

type Sms struct{}

func (s *Sms) Send(ctx context.Context, req *smsproto.Request, rsp *smsproto.Response) error {
	log.Info("Received Sms.Send request")

	rds := redis.Pool.Get()
	code := utils.GenVerificationCode()
	mobileCodeStr := fmt.Sprintf("mobile:%s", req.Mobile)
	if _, err := rds.Do("SET", mobileCodeStr, code); err != nil {
		rsp.Error = &smsproto.Error{
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
		}
		return nil
	}

	sms := provider.SMS{
		Mobile: req.Mobile,
		Code:   code,
	}

	err := provider.Alidayu{}.Send(&sms)
	if err != nil {
		rsp.Error = &smsproto.Error{
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
		}
		return nil
	}
	return nil
}
