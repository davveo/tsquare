package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"recom-srv/handler"
	
	recom "recom-srv/proto/recom"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.recom"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	recom.RegisterRecomHandler(service.Server(), new(handler.Recom))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
