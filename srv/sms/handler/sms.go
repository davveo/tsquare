package handler

import (
	"context"
	"fmt"
	"github.com/zbrechave/tsquare/lib/redis"
	"github.com/zbrechave/tsquare/srv/sms/provider"
	"github.com/zbrechave/tsquare/srv/sms/utils"
	"net/http"

	//"github.com/zbrechave/tsquare/srv/sms-srv/provider"

	log "github.com/micro/go-micro/v2/logger"
	sms_proto "github.com/zbrechave/tsquare/srv/sms/proto/sms"
)

type Sms struct{}

func (s *Sms) Send(ctx context.Context, req *sms_proto.Request, rsp *sms_proto.Response) error {
	log.Info("Received Sms.Send request")

	rds := redis.RedisPool.Get()
	code := utils.GenVerificationCode()
	mobileCodeStr := fmt.Sprintf("mobile:%s" , req.Mobile)
	if _, err := rds.Do("SET", mobileCodeStr, code); err != nil {
		rsp.Error = &sms_proto.Error{
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
		}
		return nil
	}

	sms := provider.SMS{
		Mobile: req.Mobile,
		Code: code,
	}

	err := provider.Alidayu{}.Send(&sms)
	if err != nil {
		rsp.Error = &sms_proto.Error{
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
		}
		return nil
	}
	return nil
}
