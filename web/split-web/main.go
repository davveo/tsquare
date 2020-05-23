package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/zbrechave/tsquare/web/split-web/handler"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.split"),
		web.Version("latest"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	service.HandleFunc("/split/call", handler.SplitCall)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
