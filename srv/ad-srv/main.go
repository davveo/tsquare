package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"ad-srv/handler"

	ad "ad-srv/proto/ad"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.ad"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	ad.RegisterAdHandler(service.Server(), new(handler.Ad))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
