// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/dnsconfigurator/dnsconfigurator.proto
// DO NOT EDIT!

/*
Package dnsconfigurator is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/dnsconfigurator/dnsconfigurator.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package dnsconfigurator

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
	Response *PingRequest `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetResponse() *PingRequest {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "dnsconfigurator.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "dnsconfigurator.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DNSConfiguratorService service

type DNSConfiguratorServiceClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type dNSConfiguratorServiceClient struct {
	cc *grpc.ClientConn
}

func NewDNSConfiguratorServiceClient(cc *grpc.ClientConn) DNSConfiguratorServiceClient {
	return &dNSConfiguratorServiceClient{cc}
}

func (c *dNSConfiguratorServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/dnsconfigurator.DNSConfiguratorService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DNSConfiguratorService service

type DNSConfiguratorServiceServer interface {
	// comment: rpc ping
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterDNSConfiguratorServiceServer(s *grpc.Server, srv DNSConfiguratorServiceServer) {
	s.RegisterService(&_DNSConfiguratorService_serviceDesc, srv)
}

func _DNSConfiguratorService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSConfiguratorServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dnsconfigurator.DNSConfiguratorService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSConfiguratorServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DNSConfiguratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dnsconfigurator.DNSConfiguratorService",
	HandlerType: (*DNSConfiguratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DNSConfiguratorService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/dnsconfigurator/dnsconfigurator.proto",
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
	stream, err := grpc.NewClientStream(ctx, &_EchoStreamService_serviceDesc.Streams[0], c.cc, "/dnsconfigurator.EchoStreamService/SendToServer", opts...)
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
	ServiceName: "dnsconfigurator.EchoStreamService",
	HandlerType: (*EchoStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendToServer",
			Handler:       _EchoStreamService_SendToServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "golang.conradwood.net/apis/dnsconfigurator/dnsconfigurator.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/dnsconfigurator/dnsconfigurator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x88, 0x7f, 0xa6, 0x55, 0x71, 0x0f, 0x12, 0x44, 0xa1, 0xf4, 0x20, 0x41, 0x24,
	0x85, 0x7a, 0xf1, 0x28, 0x56, 0xc1, 0x8b, 0xb5, 0x64, 0xbd, 0x7a, 0xd8, 0xec, 0x8e, 0x31, 0xd0,
	0xce, 0xc4, 0xdd, 0x8d, 0xe2, 0xb7, 0x97, 0xd4, 0x44, 0xc2, 0x82, 0x5e, 0xbc, 0xed, 0xbc, 0x7d,
	0xef, 0xcd, 0xc0, 0x0f, 0xae, 0x0b, 0x5e, 0x2a, 0x2a, 0x52, 0xcd, 0x64, 0x95, 0xf9, 0x60, 0x36,
	0x29, 0xa1, 0x9f, 0xa8, 0xaa, 0x74, 0x13, 0x43, 0x4e, 0x33, 0xbd, 0x94, 0x45, 0x6d, 0x95, 0x67,
	0x1b, 0xce, 0x69, 0x65, 0xd9, 0xb3, 0x38, 0x08, 0xe4, 0xf1, 0x23, 0x0c, 0x16, 0x25, 0x15, 0x19,
	0xbe, 0xd5, 0xe8, 0xbc, 0x88, 0x61, 0x7b, 0xa1, 0x3e, 0x97, 0xac, 0x4c, 0xbc, 0x31, 0x8a, 0x92,
	0xdd, 0xac, 0x1b, 0xc5, 0x19, 0xec, 0xcb, 0xc6, 0x44, 0x1a, 0xe7, 0xf5, 0x2a, 0x47, 0x1b, 0x47,
	0xa3, 0x28, 0xd9, 0xcb, 0x02, 0x75, 0x7c, 0x0f, 0xc3, 0xef, 0x42, 0x57, 0x31, 0x39, 0x14, 0x57,
	0xb0, 0xd3, 0xbd, 0xd7, 0x89, 0xc1, 0xf4, 0x24, 0x0d, 0x6f, 0xeb, 0x5d, 0x90, 0xfd, 0xb8, 0xa7,
	0xcf, 0x70, 0x74, 0x3b, 0x97, 0xb3, 0x9e, 0x51, 0xa2, 0x7d, 0x2f, 0x35, 0x8a, 0x19, 0x6c, 0x36,
	0x11, 0xf1, 0x67, 0xd3, 0xf1, 0xe9, 0x2f, 0xbf, 0x6d, 0x7d, 0x0e, 0x87, 0x77, 0xfa, 0x95, 0xa5,
	0xb7, 0xa8, 0x56, 0x5d, 0xf3, 0x03, 0x0c, 0x25, 0x92, 0x79, 0xe2, 0x46, 0x40, 0xfb, 0xaf, 0x0d,
	0x49, 0x74, 0x73, 0x01, 0xe7, 0x84, 0xbe, 0xcf, 0xab, 0x25, 0xd8, 0x20, 0x0b, 0xc3, 0xf9, 0xd6,
	0x9a, 0xd1, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x86, 0x37, 0x6c, 0xe7, 0x01, 0x00,
	0x00,
}
