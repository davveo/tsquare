package main

import (
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/zbrechave/tsquare/web/question-web/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.question"),
		web.Version("latest"),
	)

	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	service.HandleFunc("/question/call", handler.QuestionCall)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
