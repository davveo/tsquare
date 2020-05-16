package handler

import (
	"auth-srv/model/access"
	"auth-srv/proto/auth"
	"context"
	"strconv"

	"github.com/micro/go-micro/util/log"
)

var (
	accessService access.Service
)

func init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}

type AuthService struct{}

func (as *AuthService) MakeAccessToken(ctx context.Context, in *auth.Request, out *auth.Response) error {
	log.Log("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatInt(in.UserId, 10),
		Name: in.UserName,
	})
	if err != nil {
		out.Error = &auth.Error{
			Detail: err.Error(),
		}
		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}
	out.Token = token
	return nil
}

func (as *AuthService) DelUserAccessToken(ctx context.Context, in *auth.Request, out *auth.Response) error {
	log.Log("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(in.Token)
	if err != nil {
		out.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}

func (as *AuthService) GetCachedAccessToken(ctx context.Context, in *auth.Request, out *auth.Response) error {
	log.Logf("[GetCachedAccessToken] 获取缓存的token，%d", in.UserId)
	token, err := accessService.GetCachedAccessToken(&access.Subject{
		ID: strconv.FormatInt(in.UserId, 10),
	})
	if err != nil {
		out.Error = &auth.Error{
			Detail: err.Error(),
		}

		log.Logf("[GetCachedAccessToken] 获取缓存的token失败，err：%s", err)
		return err
	}

	out.Token = token
	return nil
}
