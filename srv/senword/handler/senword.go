package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	senword "github.com/davveo/tsquare/proto/senword"
)

type Senword struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Senword) Call(ctx context.Context, req *senword.Request, rsp *senword.Response) error {
	log.Info("Received Senword.Call request")
	//var tree utils.Trie
	//var filterbyte [][]byte
	//tree.InitRootNode()
	//tree.BuildTrie(filterbyte)

	return nil
}
