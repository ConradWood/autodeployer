// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/scautoupdate/scautoupdate.proto
// DO NOT EDIT!

/*
Package scautoupdate is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/scautoupdate/scautoupdate.proto

It has these top-level messages:
	PingResponse
*/
package scautoupdate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

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

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "scautoupdate.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SCAutoUpdate service

type SCAutoUpdateClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
}

type sCAutoUpdateClient struct {
	cc *grpc.ClientConn
}

func NewSCAutoUpdateClient(cc *grpc.ClientConn) SCAutoUpdateClient {
	return &sCAutoUpdateClient{cc}
}

func (c *sCAutoUpdateClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/scautoupdate.SCAutoUpdate/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCAutoUpdate service

type SCAutoUpdateServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
}

func RegisterSCAutoUpdateServer(s *grpc.Server, srv SCAutoUpdateServer) {
	s.RegisterService(&_SCAutoUpdate_serviceDesc, srv)
}

func _SCAutoUpdate_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCAutoUpdateServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scautoupdate.SCAutoUpdate/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCAutoUpdateServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCAutoUpdate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scautoupdate.SCAutoUpdate",
	HandlerType: (*SCAutoUpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SCAutoUpdate_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/scautoupdate/scautoupdate.proto",
}

func init() {
	proto.RegisterFile("golang.singingcat.net/apis/scautoupdate/scautoupdate.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0xcb, 0xc2, 0x30,
	0x10, 0xc6, 0x29, 0xbc, 0xbc, 0x68, 0xc8, 0x94, 0x49, 0x32, 0x95, 0x4e, 0xe2, 0x90, 0x8a, 0x6e,
	0x4e, 0xfe, 0xf9, 0x02, 0x52, 0xd1, 0x3d, 0x26, 0x21, 0x04, 0xec, 0x5d, 0x68, 0xae, 0xf8, 0xf5,
	0xa5, 0xb5, 0x6a, 0x3b, 0x25, 0xcf, 0xc1, 0xf3, 0xbb, 0xdf, 0xb1, 0x9d, 0xc7, 0x87, 0x06, 0xaf,
	0x52, 0x00, 0x1f, 0xc0, 0x1b, 0x4d, 0x0a, 0x1c, 0x95, 0x3a, 0x86, 0x54, 0x26, 0xa3, 0x5b, 0xc2,
	0x36, 0x5a, 0x4d, 0x6e, 0x12, 0x54, 0x6c, 0x90, 0x50, 0xf0, 0xf1, 0x4c, 0xaa, 0x81, 0x64, 0x10,
	0x1a, 0x6d, 0x9f, 0x88, 0xf6, 0x47, 0x32, 0x58, 0xd7, 0x08, 0xc3, 0xf3, 0x6e, 0x17, 0x2b, 0xc6,
	0xcf, 0x01, 0x7c, 0xe5, 0x52, 0x44, 0x48, 0x4e, 0x48, 0x36, 0xfb, 0xfc, 0x17, 0x59, 0x9e, 0x2d,
	0xe7, 0xd5, 0x37, 0x6f, 0xf6, 0x8c, 0x5f, 0x4e, 0x87, 0x96, 0xf0, 0xda, 0xef, 0x12, 0x6b, 0xf6,
	0xd7, 0x75, 0x05, 0x57, 0x03, 0xf2, 0x86, 0xc1, 0x4a, 0xa9, 0x26, 0x92, 0x63, 0xfa, 0xb1, 0x60,
	0x39, 0x38, 0x1a, 0x9f, 0xd9, 0x89, 0x4d, 0x0a, 0xf7, 0xff, 0x5e, 0x6c, 0xfb, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x60, 0x2b, 0xcb, 0x45, 0x14, 0x01, 0x00, 0x00,
}