// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ad/ad.proto

package go_micro_service_ad

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Ad service

type AdService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Ad_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Ad_PingPongService, error)
}

type adService struct {
	c    client.Client
	name string
}

func NewAdService(name string, c client.Client) AdService {
	return &adService{
		c:    c,
		name: name,
	}
}

func (c *adService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Ad.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Ad_StreamService, error) {
	req := c.c.NewRequest(c.name, "Ad.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &adServiceStream{stream}, nil
}

type Ad_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type adServiceStream struct {
	stream client.Stream
}

func (x *adServiceStream) Close() error {
	return x.stream.Close()
}

func (x *adServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *adServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *adServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *adServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adService) PingPong(ctx context.Context, opts ...client.CallOption) (Ad_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Ad.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &adServicePingPong{stream}, nil
}

type Ad_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type adServicePingPong struct {
	stream client.Stream
}

func (x *adServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *adServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *adServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *adServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *adServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *adServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ad service

type AdHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Ad_StreamStream) error
	PingPong(context.Context, Ad_PingPongStream) error
}

func RegisterAdHandler(s server.Server, hdlr AdHandler, opts ...server.HandlerOption) error {
	type ad interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Ad struct {
		ad
	}
	h := &adHandler{hdlr}
	return s.Handle(s.NewHandler(&Ad{h}, opts...))
}

type adHandler struct {
	AdHandler
}

func (h *adHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.AdHandler.Call(ctx, in, out)
}

func (h *adHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.AdHandler.Stream(ctx, m, &adStreamStream{stream})
}

type Ad_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type adStreamStream struct {
	stream server.Stream
}

func (x *adStreamStream) Close() error {
	return x.stream.Close()
}

func (x *adStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *adStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *adStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *adStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *adHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.AdHandler.PingPong(ctx, &adPingPongStream{stream})
}

type Ad_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type adPingPongStream struct {
	stream server.Stream
}

func (x *adPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *adPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *adPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *adPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *adPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *adPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
