package main

import (
	"fmt"

	"github.com/davveo/tsquare/basic"
	"github.com/davveo/tsquare/basic/common"
	"github.com/davveo/tsquare/basic/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc/v2"

	ad "github.com/davveo/tsquare/proto/ad"
	"github.com/davveo/tsquare/srv/ad/handler"
)

var (
	appName = "ad_srv"
	cfg     = &adCfg{}
)

type adCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()
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
	_ = ad.RegisterAdHandler(service.Server(), new(handler.Ad))

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
