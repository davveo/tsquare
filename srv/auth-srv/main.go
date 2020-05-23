package main

import (
	"fmt"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zbrechave/tsquare/basic/common"
	"go.uber.org/zap"

	"github.com/micro/cli/v2"
	"github.com/zbrechave/tsquare/srv/auth-srv/model"

	"github.com/zbrechave/tsquare/srv/auth-srv/handler"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	z "github.com/zbrechave/tsquare/plugins/zap"

	"github.com/micro/go-micro/v2"

	auth "github.com/zbrechave/tsquare/srv/auth-srv/proto/auth"

	"github.com/zbrechave/tsquare/basic"
	"github.com/zbrechave/tsquare/basic/config"
)

var (
	log     = z.GetLogger()
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()

	// 注册服务
	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
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
		panic(err)
	}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Info("[initCfg] 配置", zap.Any("cfg", cfg))

	return
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}
