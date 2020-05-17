package main

import (
	"fmt"

	"github.com/micro/cli/v2"
	"github.com/zbrechave/tsquare/srv/auth-srv/model"

	"github.com/zbrechave/tsquare/srv/auth-srv/handler"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	auth "github.com/zbrechave/tsquare/srv/auth-srv/proto/auth"

	"github.com/zbrechave/tsquare/basic"
	"github.com/zbrechave/tsquare/basic/config"
)

func main() {
	// 初始化配置
	basic.Init()

	// 注册服务
	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) error {
			// 初始化model
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	// Register Handler
	_ = auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
