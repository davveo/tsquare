package main

import (
	"auth-srv/handler"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	auth "auth-srv/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
