package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli/v2"
	"github.com/zbrechave/tsquare/srv/user-srv/model"

	"github.com/micro/go-micro/v2/registry"
	"github.com/zbrechave/tsquare/basic/config"
	"github.com/zbrechave/tsquare/srv/user-srv/handler"

	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/zbrechave/tsquare/basic"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	user "github.com/zbrechave/tsquare/srv/user-srv/proto/user"
)

func main() {
	// 初始化配置
	basic.Init()
	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) error {
		// 初始化模型层
		model.Init()
		// 初始化handler
		handler.Init()

		return nil
	}))

	// Register Handler
	_ = user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
