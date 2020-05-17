package main

import (
	"fmt"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/zbrechave/tsquare/basic/config"
	"github.com/zbrechave/tsquare/web/user-web/handler"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/zbrechave/tsquare/basic"
)

func main() {
	// 初始化操作
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/logout", handler.Logout)
	service.HandleFunc("/user/test", handler.TestSession)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
