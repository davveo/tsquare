package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"
	us "github.com/zbrechave/tsquare/srv/user/model/user"
	user_proto "github.com/zbrechave/tsquare/srv/user/proto/user"
)

type User struct{}

var (
	userService us.Service
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (u *User) QueryUserByName(ctx context.Context, req *user_proto.Request, rsp *user_proto.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Error = &user_proto.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = user
	return nil
}
