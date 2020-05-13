package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"user-srv/handler"
	"user-srv/subscriber"

	user "user-srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
