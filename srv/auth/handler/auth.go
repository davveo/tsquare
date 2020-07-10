package handler

import (
	"context"
	"strconv"

	log "github.com/micro/go-micro/v2/logger"

	auth "github.com/zbrechave/tsquare/proto/auth"
	"github.com/zbrechave/tsquare/srv/auth/model/access"
)

var (
	accessService access.Service
)

// Init 初始化handler
func Init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}

type Auth struct{}

// MakeAccessToken 生成token
func (s *Auth) MakeAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatInt(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Infof("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
func (s *Auth) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Infof("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}

// GetCachedAccessToken 获取缓存的token
func (s *Auth) GetCachedAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Infof("[GetCachedAccessToken] 获取缓存的token，%d", req.UserId)
	token, err := accessService.GetCachedAccessToken(&access.Subject{
		ID: strconv.FormatInt(req.UserId, 10),
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Errorf("[GetCachedAccessToken] 获取缓存的token失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}
