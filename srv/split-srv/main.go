package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"split-srv/handler"

	split "split-srv/proto/split"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.split"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	split.RegisterSplitHandler(service.Server(), new(handler.Split))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
