package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"answer-srv/handler"

	answer "answer-srv/proto/answer"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.answer"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	answer.RegisterAnswerHandler(service.Server(), new(handler.Answer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
