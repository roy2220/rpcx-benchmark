// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro_benchmark.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	micro_benchmark.proto

It has these top-level messages:
	BenchmarkMessage
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Hello service

type HelloService interface {
	// Sends a greeting
	Say(ctx context.Context, in *BenchmarkMessage, opts ...client.CallOption) (*BenchmarkMessage, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pb"
	}
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) Say(ctx context.Context, in *BenchmarkMessage, opts ...client.CallOption) (*BenchmarkMessage, error) {
	req := c.c.NewRequest(c.name, "Hello.Say", in)
	out := new(BenchmarkMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloHandler interface {
	// Sends a greeting
	Say(context.Context, *BenchmarkMessage, *BenchmarkMessage) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) error {
	type hello interface {
		Say(ctx context.Context, in *BenchmarkMessage, out *BenchmarkMessage) error
	}
	type Hello struct {
		hello
	}
	h := &helloHandler{hdlr}
	return s.Handle(s.NewHandler(&Hello{h}, opts...))
}

type helloHandler struct {
	HelloHandler
}

func (h *helloHandler) Say(ctx context.Context, in *BenchmarkMessage, out *BenchmarkMessage) error {
	return h.HelloHandler.Say(ctx, in, out)
}
