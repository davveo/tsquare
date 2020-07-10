package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	recom "github.com/zbrechave/tsquare/proto/recom"
)

type Recom struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Recom) Call(ctx context.Context, req *recom.Request, rsp *recom.Response) error {
	log.Info("Received Recom.Call request")

	return nil
}
