package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	ad "ad-srv/proto/ad"
)

type Ad struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Ad) Call(ctx context.Context, req *ad.Request, rsp *ad.Response) error {
	log.Info("Received Ad.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Ad) Stream(ctx context.Context, req *ad.StreamingRequest, stream ad.Ad_StreamStream) error {
	log.Infof("Received Ad.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&ad.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Ad) PingPong(ctx context.Context, stream ad.Ad_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&ad.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
