package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	inventory "inventory-srv/proto/inventory"
)

type Inventory struct{}

func (e *Inventory) Handle(ctx context.Context, msg *inventory.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *inventory.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
