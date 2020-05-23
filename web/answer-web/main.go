package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/zbrechave/tsquare/web/answer-web/handler"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.answer"),
		web.Version("latest"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	service.HandleFunc("/answer/call", handler.AnswerCall)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
