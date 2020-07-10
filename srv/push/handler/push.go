package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	push "github.com/zbrechave/tsquare/proto/push"
)

type Push struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Push) Call(ctx context.Context, req *push.Request, rsp *push.Response) error {
	log.Info("Received Push.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
