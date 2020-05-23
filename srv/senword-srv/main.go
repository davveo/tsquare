package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"senword-srv/handler"
	"senword-srv/subscriber"

	senword "senword-srv/proto/senword"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.senword"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	senword.RegisterSenwordHandler(service.Server(), new(handler.Senword))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.senword", service.Server(), new(subscriber.Senword))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
