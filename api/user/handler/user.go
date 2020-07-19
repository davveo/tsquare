package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/davveo/tsquare/api/user/model/request"
	"github.com/davveo/tsquare/basic/common"
	"github.com/davveo/tsquare/plugins/session"
	//"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
	hystrixplugins "github.com/davveo/tsquare/plugins/breaker/hystrix"
	auth "github.com/davveo/tsquare/srv/auth/proto/auth"
	sms "github.com/davveo/tsquare/srv/sms/proto/sms"
	user "github.com/davveo/tsquare/srv/user/proto/user"
	"net/http"
	"time"
)

var (
	smsService sms.SmsService
	userService user.UserService
	authService auth.AuthService
)

type User struct{}

func Init() {
	cl := hystrixplugins.NewClientWrapper()(client.DefaultClient)

	_ = cl.Init(
		client.Retries(3),
		//为了调试看log方便，始终返回true, nil，即会一直重试直至重试次数用尽
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Infof(req.Method(), retryCount, " client retry")
			return true, nil
		}),
	)

	smsService = sms.NewSmsService("go.micro.srv.sms", cl)
	userService = user.NewUserService("go.micro.srv.user", cl)
	authService = auth.NewAuthService("go.micro.srv.auth", cl)
}

func (u *User) Login(ctx *gin.Context) {
	log.Info("Received User.Login request")

	var loginReq request.LoginRequest
	if err := ctx.BindJSON(&loginReq); err != nil {
		log.Errorf("bind param err: %v", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resp, err := userService.QueryUserByName(context.TODO(), &user.Request{
		UserName: loginReq.UserName,
	})

	if err != nil {
		log.Errorf("call userService.QueryUserByName err: %v", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if resp.User.Pwd == loginReq.PassWord {
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

func (u *User) Logout(ctx *gin.Context) {
	log.Info("Received User.Logout request")

	tokenVal, err := ctx.Cookie("remember-me-token")
	if err != nil {
		log.Error("token获取失败")
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = authService.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenVal,
	})
	if err != nil {
		log.Errorf("call authService.DelUserAccessToken err: %v", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.SetCookie(
		"remember-me-token",
		"", 0,
		"/", "localhost",
		false, true,
	)
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}
	ctx.JSON(http.StatusOK, response)
}

func (u *User) SmsCode(ctx *gin.Context) {
	log.Info("Received User.SmsCode request")
	var smsReq request.SmsRequest
	if err := ctx.BindJSON(&smsReq); err != nil {
		log.Errorf("bind param err: %v", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if _, err := smsService.Send(context.TODO(), &sms.Request{Mobile: smsReq.Mobile}); err != nil {
		log.Errorf("call smsService.Send err: %v", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
