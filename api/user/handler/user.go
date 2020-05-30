package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/zbrechave/tsquare/basic/common"
	"github.com/zbrechave/tsquare/plugins/session"

	//"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
	hystrixplugins "github.com/zbrechave/tsquare/plugins/breaker/hystrix"

	auth "github.com/zbrechave/tsquare/srv/auth-srv/proto/auth"
	user "github.com/zbrechave/tsquare/srv/user-srv/proto/user"
	"net/http"
	"time"
)

var (
	userService user.UserService
	authService auth.AuthService

)

type User struct{}

type LoginRequest struct {
	UserName string `json:"userName"`
	PassWord    string    `json:"pwd"`
}

func Init()  {
	cl := hystrixplugins.NewClientWrapper()(client.DefaultClient)

	_ = cl.Init(
		client.Retries(3),
		//为了调试看log方便，始终返回true, nil，即会一直重试直至重试次数用尽
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Infof(req.Method(), retryCount, " client retry")
			return true, nil
		}),
	)

	userService = user.NewUserService("go.micro.srv.user", cl)
	authService = auth.NewAuthService("go.micro.srv.auth", cl)
}


func (e *User) Login(ctx *gin.Context) {
	log.Info("Received User.Login request")

	//userClient, ok := srvClient.UserFromContext(ctx)
	//if !ok {
	//	err := errors.InternalServerError("go.micro.srv.user", "user client not found")
	//	ctx.JSON(http.StatusInternalServerError, err)
	//	return
	//}
	//
	//authClient, ok := srvClient.AuthFromContext(ctx)
	//if !ok {
	//	err := errors.InternalServerError("go.micro.srv.auth", "srv client not found")
	//	ctx.JSON(http.StatusInternalServerError, err)
	//	return
	//}
	var loginReq LoginRequest

	if err := ctx.BindJSON(&loginReq); err != nil {
		log.Warnf("bind param err: %v", err)
		return
	}


	fmt.Println(loginReq.UserName, loginReq.PassWord)

	userName := loginReq.UserName
	password := loginReq.PassWord

	resp, err := userService.QueryUserByName(context.TODO(), &user.Request{
		UserName: userName,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if resp.User.Pwd == password {
		resp.User.Pwd = ""
		response["success"] = true
		response["data"] = resp.User

		rsp2, err := authService.MakeAccessToken(
			context.TODO(), &auth.Request{
			UserId:   resp.User.Id,
			UserName: resp.User.Name,
		})

		if err != nil {
			log.Infof("[Login] 创建token失败，err：%s", err)
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		response["token"] = rsp2.Token
		ctx.SetCookie(
			"remember-me-token",
			rsp2.Token, 90000,
			"/", "localhost",
			false, true,
		)

		timeString := time.Now().Format(common.DefaultMsTimeLayout)

		// 同步到session中
		sess := session.GetSession(ctx)
		sess.Values["userId"] = resp.User.Id
		sess.Values["userName"] = resp.User.Name
		sess.Values["login_time"] = timeString
		_ = sess.Save(ctx.Request, ctx.Writer)

	} else {
		response["success"] = false
		response["error"] = "密码错误!"
	}

	ctx.JSON(http.StatusOK, response)

}


func (e *User) Logout(c *gin.Context) {
//	log.Info("Received Say.Hello API request")
//
//	name := c.Param("name")
//
//	response, err := cl.Hello(context.TODO(), &hello.Request{
//		Name: name,
//	})
//
//	if err != nil {
//		c.JSON(500, err)
//	}
//
//	c.JSON(200, response)
//
//	if r.Method != "POST" {
//		log.Infof("非法请求")
//		http.Error(w, "非法请求", 400)
//		return
//	}
//	_ = r.ParseForm()
//
//	rsp, err := userClient.QueryUserByName(context.TODO(), &user.Request{
//		UserName: r.Form.Get("userName"),
//	})
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//	// 返回结果
//	response := map[string]interface{}{
//		"ref": time.Now().UnixNano(),
//	}
//
//	if rsp.User.Pwd == r.Form.Get("pwd") {
//		response["success"] = true
//
//		// 干掉密码返回
//		rsp.User.Pwd = ""
//		response["data"] = rsp.User
//		log.Infof("[Login] 密码校验完成，生成token...")
//
//		// 生成token
//		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
//			UserId:   rsp.User.Id,
//			UserName: rsp.User.Name,
//		})
//		if err != nil {
//			log.Errorf("[Login] 创建token失败，err：%s", err)
//			http.Error(w, err.Error(), 500)
//			return
//		}
//
//		log.Errorf("[Login] token %s", rsp2.Token)
//		response["token"] = rsp2.Token
//
//		// 同时将token写到cookies中
//		w.Header().Add("set-cookie", "application/json; charset=utf-8")
//		// 过期30分钟
//		expire := time.Now().Add(30 * time.Minute)
//		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
//		http.SetCookie(w, &cookie)
//
//		timeString := time.Now().Format(common.DefaultMsTimeLayout)
//
//		// 同步到session中
//		sess := session.GetSession(w, r)
//		sess.Values["userId"] = rsp.User.Id
//		sess.Values["userName"] = rsp.User.Name
//		sess.Values["login_time"] = timeString
//		_ = sess.Save(r, w)
//
//	} else {
//		response["success"] = false
//		response["error"] = &Error{
//			Detail: "密码错误",
//		}
//	}
//	w.Header().Add("Content-Type", "application/json; charset=utf-8")
//	// 返回JSON结构
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
}
