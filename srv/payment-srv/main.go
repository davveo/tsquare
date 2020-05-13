package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/zbrechave/tsquare/payment-srv/handler"

	payment "payment-srv/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
