package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	senword "senword-srv/proto/senword"
)

type Senword struct{}

func (e *Senword) Handle(ctx context.Context, msg *senword.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *senword.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
