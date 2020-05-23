// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/uuid/uuid.proto

package go_micro_service_uuid

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

// Client API for Uuid service

type UuidService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Uuid_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Uuid_PingPongService, error)
}

type uuidService struct {
	c    client.Client
	name string
}

func NewUuidService(name string, c client.Client) UuidService {
	return &uuidService{
		c:    c,
		name: name,
	}
}

func (c *uuidService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Uuid.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uuidService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Uuid_StreamService, error) {
	req := c.c.NewRequest(c.name, "Uuid.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &uuidServiceStream{stream}, nil
}

type Uuid_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type uuidServiceStream struct {
	stream client.Stream
}

func (x *uuidServiceStream) Close() error {
	return x.stream.Close()
}

func (x *uuidServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uuidServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uuidServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uuidServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uuidService) PingPong(ctx context.Context, opts ...client.CallOption) (Uuid_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Uuid.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &uuidServicePingPong{stream}, nil
}

type Uuid_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type uuidServicePingPong struct {
	stream client.Stream
}

func (x *uuidServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *uuidServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *uuidServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uuidServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uuidServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *uuidServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Uuid service

type UuidHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Uuid_StreamStream) error
	PingPong(context.Context, Uuid_PingPongStream) error
}

func RegisterUuidHandler(s server.Server, hdlr UuidHandler, opts ...server.HandlerOption) error {
	type uuid interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Uuid struct {
		uuid
	}
	h := &uuidHandler{hdlr}
	return s.Handle(s.NewHandler(&Uuid{h}, opts...))
}

type uuidHandler struct {
	UuidHandler
}

func (h *uuidHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UuidHandler.Call(ctx, in, out)
}

func (h *uuidHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UuidHandler.Stream(ctx, m, &uuidStreamStream{stream})
}

type Uuid_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type uuidStreamStream struct {
	stream server.Stream
}

func (x *uuidStreamStream) Close() error {
	return x.stream.Close()
}

func (x *uuidStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uuidStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uuidStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uuidStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *uuidHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.UuidHandler.PingPong(ctx, &uuidPingPongStream{stream})
}

type Uuid_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type uuidPingPongStream struct {
	stream server.Stream
}

func (x *uuidPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *uuidPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uuidPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uuidPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uuidPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *uuidPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
