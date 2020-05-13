package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	order "order-srv/proto/order"
)

type Order struct{}

func (e *Order) Handle(ctx context.Context, msg *order.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *order.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
