package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zbrechave/tsquare/basic"
	"github.com/zbrechave/tsquare/basic/common"
	"github.com/zbrechave/tsquare/basic/config"

	sms "github.com/zbrechave/tsquare/proto/sms"
	"github.com/zbrechave/tsquare/srv/sms/handler"
)

var (
	appName = "sms_srv"
	cfg     = &smsCfg{}
)

type smsCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()

	// 初始化短信配置
	//if err := smsconfig.ParseConfigFile("conf/conf.yaml"); err != nil {
	//	fmt.Printf("%s\n%s", "Unable to load config info!", err)
	//	panic(err)
	//}

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = sms.RegisterSmsHandler(service.Server(), new(handler.Sms))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
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

	log.Infof("[initCfg] 配置，cfg：%v", cfg)

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
