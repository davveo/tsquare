package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	settlement "settlement-srv/proto/settlement"
)

type Settlement struct{}

func (e *Settlement) Handle(ctx context.Context, msg *settlement.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *settlement.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
