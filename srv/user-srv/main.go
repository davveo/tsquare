package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/zbrechave/tsquare/user-srv/handler"
	"github.com/zbrechave/tsquare/user-srv/subscriber"

	user "github.com/zbrechave/tsquare/user-srv/proto/user"
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
