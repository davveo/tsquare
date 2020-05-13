package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	settlement "github.com/zbrechave/tsquare/settlement-srv/proto/settlement"
)

type Settlement struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Settlement) Call(ctx context.Context, req *settlement.Request, rsp *settlement.Response) error {
	log.Info("Received Settlement.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Settlement) Stream(ctx context.Context, req *settlement.StreamingRequest, stream settlement.Settlement_StreamStream) error {
	log.Infof("Received Settlement.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&settlement.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Settlement) PingPong(ctx context.Context, stream settlement.Settlement_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&settlement.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
