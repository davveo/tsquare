package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	senword "github.com/zbrechave/tsquare/srv/senword/proto/senword"
)

type Senword struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Senword) Call(ctx context.Context, req *senword.Request, rsp *senword.Response) error {
	log.Info("Received Senword.Call request")

	return nil
}
