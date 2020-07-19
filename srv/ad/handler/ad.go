package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	ad "github.com/davveo/tsquare/proto/ad"
)

type Ad struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Ad) Call(ctx context.Context, req *ad.Request, rsp *ad.Response) error {
	log.Info("Received Ad.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
