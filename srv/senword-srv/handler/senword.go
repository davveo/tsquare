package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	senword "senword-srv/proto/senword"
)

type Senword struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Senword) Call(ctx context.Context, req *senword.Request, rsp *senword.Response) error {
	log.Info("Received Senword.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Senword) Stream(ctx context.Context, req *senword.StreamingRequest, stream senword.Senword_StreamStream) error {
	log.Infof("Received Senword.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&senword.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Senword) PingPong(ctx context.Context, stream senword.Senword_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&senword.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
