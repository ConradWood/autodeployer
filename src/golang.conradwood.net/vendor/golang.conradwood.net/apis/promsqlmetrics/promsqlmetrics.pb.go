// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/promsqlmetrics/promsqlmetrics.proto
// DO NOT EDIT!

/*
Package promsqlmetrics is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/promsqlmetrics/promsqlmetrics.proto

It has these top-level messages:
*/
package promsqlmetrics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PromSQLMetrics service

type PromSQLMetricsClient interface {
}

type promSQLMetricsClient struct {
	cc *grpc.ClientConn
}

func NewPromSQLMetricsClient(cc *grpc.ClientConn) PromSQLMetricsClient {
	return &promSQLMetricsClient{cc}
}

// Server API for PromSQLMetrics service

type PromSQLMetricsServer interface {
}

func RegisterPromSQLMetricsServer(s *grpc.Server, srv PromSQLMetricsServer) {
	s.RegisterService(&_PromSQLMetrics_serviceDesc, srv)
}

var _PromSQLMetrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "promsqlmetrics.PromSQLMetrics",
	HandlerType: (*PromSQLMetricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "golang.conradwood.net/apis/promsqlmetrics/promsqlmetrics.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/promsqlmetrics/promsqlmetrics.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0x4b, 0xce, 0xcf, 0x2b, 0x4a, 0x4c, 0x29, 0xcf, 0xcf, 0x4f, 0xd1, 0xcb, 0x4b,
	0x2d, 0xd1, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x2f, 0x28, 0xca, 0xcf, 0x2d, 0x2e, 0xcc, 0xc9, 0x4d,
	0x2d, 0x29, 0xca, 0x4c, 0x46, 0xe7, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa1, 0x8a,
	0x1a, 0x09, 0x70, 0xf1, 0x05, 0x14, 0xe5, 0xe7, 0x06, 0x07, 0xfa, 0xf8, 0x42, 0x44, 0x9c, 0xb4,
	0xb9, 0x34, 0xf3, 0x52, 0x4b, 0x90, 0x2d, 0x80, 0x5a, 0x09, 0xb2, 0x43, 0x0f, 0x55, 0x7b, 0x12,
	0x1b, 0xd8, 0x54, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xf3, 0xd9, 0x85, 0x97, 0x00,
	0x00, 0x00,
}
