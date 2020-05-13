package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/zbrechave/tsquare/settlement-srv/handler"

	settlement "github.com/zbrechave/tsquare/settlement-srv/proto/settlement"
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
