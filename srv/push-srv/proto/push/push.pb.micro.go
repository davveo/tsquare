// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/push/push.proto

package go_micro_service_push

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

// Client API for Push service

type PushService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Push_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Push_PingPongService, error)
}

type pushService struct {
	c    client.Client
	name string
}

func NewPushService(name string, c client.Client) PushService {
	return &pushService{
		c:    c,
		name: name,
	}
}

func (c *pushService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Push.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Push_StreamService, error) {
	req := c.c.NewRequest(c.name, "Push.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &pushServiceStream{stream}, nil
}

type Push_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type pushServiceStream struct {
	stream client.Stream
}

func (x *pushServiceStream) Close() error {
	return x.stream.Close()
}

func (x *pushServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *pushServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pushServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pushServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushService) PingPong(ctx context.Context, opts ...client.CallOption) (Push_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Push.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &pushServicePingPong{stream}, nil
}

type Push_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type pushServicePingPong struct {
	stream client.Stream
}

func (x *pushServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *pushServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *pushServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pushServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pushServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *pushServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Push service

type PushHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Push_StreamStream) error
	PingPong(context.Context, Push_PingPongStream) error
}

func RegisterPushHandler(s server.Server, hdlr PushHandler, opts ...server.HandlerOption) error {
	type push interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Push struct {
		push
	}
	h := &pushHandler{hdlr}
	return s.Handle(s.NewHandler(&Push{h}, opts...))
}

type pushHandler struct {
	PushHandler
}

func (h *pushHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.PushHandler.Call(ctx, in, out)
}

func (h *pushHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PushHandler.Stream(ctx, m, &pushStreamStream{stream})
}

type Push_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type pushStreamStream struct {
	stream server.Stream
}

func (x *pushStreamStream) Close() error {
	return x.stream.Close()
}

func (x *pushStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *pushStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pushStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pushStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *pushHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.PushHandler.PingPong(ctx, &pushPingPongStream{stream})
}

type Push_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type pushPingPongStream struct {
	stream server.Stream
}

func (x *pushPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *pushPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *pushPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pushPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pushPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *pushPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}