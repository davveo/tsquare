package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"uuid-srv/handler"

	uuid "uuid-srv/proto/uuid"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.uuid"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	uuid.RegisterUuidHandler(service.Server(), new(handler.Uuid))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
