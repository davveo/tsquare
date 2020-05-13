package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/zbrechave/tsquare/settlement-srv/handler"
	"github.com/zbrechave/tsquare/settlement-srv/subscriber"

	settlement "settlement-srv/proto/settlement"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.settlement"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	settlement.RegisterSettlementHandler(service.Server(), new(handler.Settlement))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.settlement", service.Server(), new(subscriber.Settlement))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
