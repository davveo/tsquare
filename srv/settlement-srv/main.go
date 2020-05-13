package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/zbrechave/tsquare/settlement-srv/handler"

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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
