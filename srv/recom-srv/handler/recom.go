package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	recom "recom-srv/proto/recom"
)

type Recom struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Recom) Call(ctx context.Context, req *recom.Request, rsp *recom.Response) error {
	log.Info("Received Recom.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Recom) Stream(ctx context.Context, req *recom.StreamingRequest, stream recom.Recom_StreamStream) error {
	log.Infof("Received Recom.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&recom.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Recom) PingPong(ctx context.Context, stream recom.Recom_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&recom.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
