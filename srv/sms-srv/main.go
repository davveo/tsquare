package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"sms-srv/handler"

	sms "sms-srv/proto/sms"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.sms"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	sms.RegisterSmsHandler(service.Server(), new(handler.Sms))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
