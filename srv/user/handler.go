package user

import (
	"context"
	"net/http"

	log "github.com/micro/go-micro/v2/logger"
	userproto "github.com/zbrechave/tsquare/proto/user"
)

type User struct{}

var (
	userService Service
)

// Init 初始化handler
func InitHandler() {
	var err error
	userService, err = GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (u *User) QueryUserByName(ctx context.Context, req *userproto.Request, rsp *userproto.Response) error {
	username := req.UserName
	if username == "" {
		rsp.Error = &userproto.Error{
			Code:   http.StatusBadRequest,
			Detail: "username can't be empty",
		}
		return nil
	}

	user, err := userService.QueryUserByName(username)
	if err != nil {
		rsp.Error = &userproto.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = &userproto.User{
		Id:   0,
		Name: user.UserName,
	}
	return nil
}

func (u *User) CreateUser(ctx context.Context, req *userproto.Request, rsp *userproto.Response) error {
	// 获取用户信息
	username := req.UserName
	password := req.UserPwd
	if username == "" || password == "" {
		rsp.Error = &userproto.Error{
			Code:   http.StatusBadRequest,
			Detail: "username or password can't be empty",
		}
		return nil
	}
	// create user
	user, err := userService.CreateUser(username, password)
	if err != nil {
		rsp.Error = &userproto.Error{
			Code:   http.StatusBadRequest,
			Detail: err.Error(),
		}

		return nil
	}
	rsp.User = user
	rsp.Success = true

}
