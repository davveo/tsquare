package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"question-srv/handler"

	question "question-srv/proto/question"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.question"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	question.RegisterQuestionHandler(service.Server(), new(handler.Question))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
