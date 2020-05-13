package auth

import (
	"github.com/micro/micro/plugin"
	"github.com/zbrechave/tsquare/lib/token"
	"log"
	"net/http"
)

func JWTAuthWrapper(token *token.Token) plugin.Handler  {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth plugin received: " + r.URL.Path)
			// TODO 从配置中心动态获取白名单URL
			if r.URL.Path == "/user/login" || r.URL.Path == "/user/register" || r.URL.Path == "/user/test" || r.URL.Path == "/metrics" {
				handler.ServeHTTP(w, r)
				return
			}

			tokenstr := r.Header.Get("Authorization")
			userFromToken, e := token.Decode(tokenstr)

			if e != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			log.Println("User Name : ", userFromToken.UserName)
			r.Header.Set("X-Example-Username", userFromToken.UserName)
			handler.ServeHTTP(w, r)
		})
	}
}
