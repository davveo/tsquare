package main

import (
	"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/zbrechave/tsquare/inventory-srv/handler"

	inventory "inventory-srv/proto/inventory"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.inventory"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	inventory.RegisterInventoryHandler(service.Server(), new(handler.Inventory))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
