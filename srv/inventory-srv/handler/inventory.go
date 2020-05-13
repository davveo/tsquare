package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	inventory "github.com/zbrechave/tsquare/inventory-srv/proto/inventory"
)

type Inventory struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Inventory) Call(ctx context.Context, req *inventory.Request, rsp *inventory.Response) error {
	log.Info("Received Inventory.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Inventory) Stream(ctx context.Context, req *inventory.StreamingRequest, stream inventory.Inventory_StreamStream) error {
	log.Infof("Received Inventory.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&inventory.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Inventory) PingPong(ctx context.Context, stream inventory.Inventory_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&inventory.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
