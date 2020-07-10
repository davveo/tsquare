package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	answer "github.com/zbrechave/tsquare/proto/answer"
)

type Answer struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Answer) Call(ctx context.Context, req *answer.Request, rsp *answer.Response) error {
	log.Info("Received Answer.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Answer) Stream(ctx context.Context, req *answer.StreamingRequest, stream answer.Answer_StreamStream) error {
	log.Infof("Received Answer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&answer.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Answer) PingPong(ctx context.Context, stream answer.Answer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&answer.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
