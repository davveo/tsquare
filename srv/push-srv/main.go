package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"push-srv/handler"

	push "push-srv/proto/push"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.push"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	push.RegisterPushHandler(service.Server(), new(handler.Push))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
