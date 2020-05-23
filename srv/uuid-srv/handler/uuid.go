package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	uuid "uuid-srv/proto/uuid"
)

type Uuid struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Uuid) Call(ctx context.Context, req *uuid.Request, rsp *uuid.Response) error {
	log.Info("Received Uuid.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Uuid) Stream(ctx context.Context, req *uuid.StreamingRequest, stream uuid.Uuid_StreamStream) error {
	log.Infof("Received Uuid.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&uuid.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Uuid) PingPong(ctx context.Context, stream uuid.Uuid_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&uuid.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
