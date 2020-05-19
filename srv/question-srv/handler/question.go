package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	question "question-srv/proto/question"
)

type Question struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Question) Call(ctx context.Context, req *question.Request, rsp *question.Response) error {
	log.Info("Received Question.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Question) Stream(ctx context.Context, req *question.StreamingRequest, stream question.Question_StreamStream) error {
	log.Infof("Received Question.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&question.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Question) PingPong(ctx context.Context, stream question.Question_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&question.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
