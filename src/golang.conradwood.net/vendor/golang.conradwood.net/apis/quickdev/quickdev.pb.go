// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/quickdev/quickdev.proto
// DO NOT EDIT!

/*
Package quickdev is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/quickdev/quickdev.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package quickdev

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

// comment: message pingrequest
type PingRequest struct {
	// comment: payload
	Payload string `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
	// comment: sequencenumber
	SequenceNumber uint32 `protobuf:"varint,1,opt,name=SequenceNumber" json:"SequenceNumber,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PingRequest) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "quickdev.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "quickdev.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for QuickDevService service

type QuickDevServiceClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
}

type quickDevServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuickDevServiceClient(cc *grpc.ClientConn) QuickDevServiceClient {
	return &quickDevServiceClient{cc}
}

func (c *quickDevServiceClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/quickdev.QuickDevService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QuickDevService service

type QuickDevServiceServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
}

func RegisterQuickDevServiceServer(s *grpc.Server, srv QuickDevServiceServer) {
	s.RegisterService(&_QuickDevService_serviceDesc, srv)
}

func _QuickDevService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickDevServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quickdev.QuickDevService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickDevServiceServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuickDevService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quickdev.QuickDevService",
	HandlerType: (*QuickDevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _QuickDevService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/quickdev/quickdev.proto",
}

// Client API for EchoStreamService service

type EchoStreamServiceClient interface {
	// comment: rpc SendToServer
	SendToServer(ctx context.Context, opts ...grpc.CallOption) (EchoStreamService_SendToServerClient, error)
}

type echoStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoStreamServiceClient(cc *grpc.ClientConn) EchoStreamServiceClient {
	return &echoStreamServiceClient{cc}
}

func (c *echoStreamServiceClient) SendToServer(ctx context.Context, opts ...grpc.CallOption) (EchoStreamService_SendToServerClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoStreamService_serviceDesc.Streams[0], c.cc, "/quickdev.EchoStreamService/SendToServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoStreamServiceSendToServerClient{stream}
	return x, nil
}

type EchoStreamService_SendToServerClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingResponse, error)
	grpc.ClientStream
}

type echoStreamServiceSendToServerClient struct {
	grpc.ClientStream
}

func (x *echoStreamServiceSendToServerClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoStreamServiceSendToServerClient) CloseAndRecv() (*PingResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EchoStreamService service

type EchoStreamServiceServer interface {
	// comment: rpc SendToServer
	SendToServer(EchoStreamService_SendToServerServer) error
}

func RegisterEchoStreamServiceServer(s *grpc.Server, srv EchoStreamServiceServer) {
	s.RegisterService(&_EchoStreamService_serviceDesc, srv)
}

func _EchoStreamService_SendToServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoStreamServiceServer).SendToServer(&echoStreamServiceSendToServerServer{stream})
}

type EchoStreamService_SendToServerServer interface {
	SendAndClose(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type echoStreamServiceSendToServerServer struct {
	grpc.ServerStream
}

func (x *echoStreamServiceSendToServerServer) SendAndClose(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoStreamServiceSendToServerServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quickdev.EchoStreamService",
	HandlerType: (*EchoStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendToServer",
			Handler:       _EchoStreamService_SendToServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "golang.conradwood.net/apis/quickdev/quickdev.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/quickdev/quickdev.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xa9, 0x88, 0x6e, 0xd7, 0xaa, 0x18, 0x50, 0x4a, 0x9f, 0xc6, 0x44, 0x29, 0x22, 0x19,
	0xd4, 0x1f, 0x30, 0x10, 0x7d, 0xd5, 0xd9, 0x0e, 0xdf, 0xb3, 0xe4, 0x52, 0x8b, 0x6b, 0x6e, 0x97,
	0xa6, 0x15, 0xff, 0xbd, 0xa4, 0x6b, 0xcb, 0x18, 0xb8, 0xa7, 0xdc, 0x73, 0xc9, 0x77, 0xb8, 0xe7,
	0x40, 0x9c, 0xd1, 0x5a, 0xe8, 0x8c, 0x4b, 0xd2, 0x46, 0xa8, 0x1f, 0x22, 0xc5, 0x35, 0xda, 0x99,
	0x28, 0xf3, 0x6a, 0xb6, 0xa9, 0x73, 0xf9, 0xad, 0xb0, 0x19, 0x06, 0x5e, 0x1a, 0xb2, 0xc4, 0x46,
	0xbd, 0x0e, 0xf9, 0x01, 0x5a, 0x52, 0x51, 0x90, 0xee, 0x9e, 0x2d, 0x39, 0x7d, 0x87, 0xb3, 0x45,
	0xae, 0xb3, 0x04, 0x37, 0x35, 0x56, 0x96, 0x05, 0x70, 0xba, 0x10, 0xbf, 0x6b, 0x12, 0x2a, 0x38,
	0x9a, 0x78, 0xd1, 0x38, 0xe9, 0x25, 0xbb, 0x87, 0x8b, 0xd4, 0x7d, 0xd2, 0x12, 0xdf, 0xea, 0x62,
	0x85, 0x26, 0xf0, 0x26, 0x5e, 0x74, 0x9e, 0xec, 0x6d, 0xa7, 0x0f, 0xe0, 0x6f, 0x0d, 0xab, 0x92,
	0x74, 0x85, 0x2c, 0x84, 0x51, 0x3f, 0xb7, 0xc4, 0x38, 0x19, 0x74, 0x3c, 0x87, 0xcb, 0x0f, 0x77,
	0xf8, 0x0b, 0x36, 0x29, 0x9a, 0x26, 0x97, 0xc8, 0x1e, 0xe1, 0xd8, 0xe1, 0xcc, 0xe7, 0xdd, 0x99,
	0x9f, 0x94, 0xab, 0xf0, 0x86, 0x0f, 0x81, 0x77, 0xcd, 0xe3, 0x25, 0x5c, 0xbd, 0xca, 0x2f, 0x4a,
	0xad, 0x41, 0x51, 0xf4, 0x16, 0x73, 0xf0, 0x53, 0xd4, 0x6a, 0x49, 0x6e, 0x81, 0x86, 0x5d, 0xef,
	0xc3, 0x6d, 0xd4, 0xff, 0x3c, 0x23, 0xef, 0xf9, 0x0e, 0x6e, 0x35, 0xda, 0xdd, 0x0a, 0xbb, 0x52,
	0x5d, 0x8b, 0x03, 0xb5, 0x3a, 0x69, 0x1b, 0x7c, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x11, 0x7e,
	0x0f, 0x1a, 0xb1, 0x01, 0x00, 0x00,
}
