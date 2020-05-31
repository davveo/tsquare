package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zbrechave/tsquare/api/question/handler"
	"github.com/zbrechave/tsquare/basic"
	"github.com/zbrechave/tsquare/basic/common"
	"github.com/zbrechave/tsquare/basic/config"
	"time"
)


var (
	appName = "question_api"
	cfg     = &questionCfg{}
)

type questionCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()

	// 使用etcd注册中心
	micReg := etcd.NewRegistry(registryOptions)

	// 初始化gin
	router := gin.Default()
	question := new(handler.Question)
	router.POST("/user/login", question.Call)


	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
		web.Handler(router),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
	)
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

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

