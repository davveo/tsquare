package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	payment "payment-srv/proto/payment"
)

type Payment struct{}

func (e *Payment) Handle(ctx context.Context, msg *payment.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *payment.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
