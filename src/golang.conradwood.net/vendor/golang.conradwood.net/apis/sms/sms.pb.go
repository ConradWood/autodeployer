// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/sms/sms.proto
// DO NOT EDIT!

/*
Package sms is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/sms/sms.proto

It has these top-level messages:
	NewSMSRequest
	SendSMSRequest
*/
package sms

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

// new sms received
type NewSMSRequest struct {
	Sender  string `protobuf:"bytes,1,opt,name=Sender" json:"Sender,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content" json:"Content,omitempty"`
}

func (m *NewSMSRequest) Reset()                    { *m = NewSMSRequest{} }
func (m *NewSMSRequest) String() string            { return proto.CompactTextString(m) }
func (*NewSMSRequest) ProtoMessage()               {}
func (*NewSMSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewSMSRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *NewSMSRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SendSMSRequest struct {
	Sender    string `protobuf:"bytes,1,opt,name=Sender" json:"Sender,omitempty"`
	Recipient string `protobuf:"bytes,2,opt,name=Recipient" json:"Recipient,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=Content" json:"Content,omitempty"`
}

func (m *SendSMSRequest) Reset()                    { *m = SendSMSRequest{} }
func (m *SendSMSRequest) String() string            { return proto.CompactTextString(m) }
func (*SendSMSRequest) ProtoMessage()               {}
func (*SendSMSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SendSMSRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SendSMSRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *SendSMSRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*NewSMSRequest)(nil), "sms.NewSMSRequest")
	proto.RegisterType((*SendSMSRequest)(nil), "sms.SendSMSRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SMSService service

type SMSServiceClient interface {
	// a new SMS was received
	Received(ctx context.Context, in *NewSMSRequest, opts ...grpc.CallOption) (*common.Void, error)
	// send out an SMS (via kapow)
	Send(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type sMSServiceClient struct {
	cc *grpc.ClientConn
}

func NewSMSServiceClient(cc *grpc.ClientConn) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) Received(ctx context.Context, in *NewSMSRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/sms.SMSService/Received", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSServiceClient) Send(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/sms.SMSService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SMSService service

type SMSServiceServer interface {
	// a new SMS was received
	Received(context.Context, *NewSMSRequest) (*common.Void, error)
	// send out an SMS (via kapow)
	Send(context.Context, *SendSMSRequest) (*common.Void, error)
}

func RegisterSMSServiceServer(s *grpc.Server, srv SMSServiceServer) {
	s.RegisterService(&_SMSService_serviceDesc, srv)
}

func _SMSService_Received_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).Received(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMSService/Received",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).Received(ctx, req.(*NewSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMSService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).Send(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Received",
			Handler:    _SMSService_Received_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _SMSService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/sms/sms.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/sms/sms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xa9, 0x95, 0x6a, 0x07, 0xf5, 0x30, 0x82, 0x2c, 0x45, 0xa4, 0xf4, 0x54, 0x41, 0x52,
	0xd0, 0x5f, 0xa0, 0x9e, 0xeb, 0x61, 0x03, 0x9e, 0x5d, 0x93, 0xa1, 0x04, 0xdc, 0x99, 0x35, 0x89,
	0xed, 0xdf, 0x97, 0xc4, 0x48, 0x9b, 0x8b, 0x78, 0x08, 0x61, 0x92, 0x37, 0xf3, 0xe6, 0x7d, 0xb0,
	0xdc, 0xc8, 0x47, 0xc7, 0x1b, 0x65, 0x84, 0x7d, 0x67, 0x77, 0x22, 0x56, 0x31, 0xc5, 0x55, 0x37,
	0xb8, 0xb0, 0x0a, 0x7d, 0x3e, 0x6a, 0xf0, 0x12, 0x05, 0xc7, 0xa1, 0x0f, 0x33, 0xf5, 0x87, 0xdc,
	0x48, 0xdf, 0x0b, 0x97, 0xeb, 0xa7, 0x69, 0xf1, 0x08, 0xe7, 0x2f, 0xb4, 0xd3, 0x6b, 0xdd, 0xd2,
	0xe7, 0x17, 0x85, 0x88, 0x57, 0x30, 0xd1, 0xc4, 0x96, 0x7c, 0x33, 0x9a, 0x8f, 0x96, 0xd3, 0xb6,
	0x54, 0xd8, 0xc0, 0xc9, 0xb3, 0x70, 0x24, 0x8e, 0xcd, 0x51, 0xfe, 0xf8, 0x2d, 0x17, 0x6f, 0x70,
	0x91, 0x34, 0xff, 0x98, 0x71, 0x0d, 0xd3, 0x96, 0x8c, 0x1b, 0xdc, 0x7e, 0xca, 0xfe, 0xe1, 0xd0,
	0x61, 0x5c, 0x39, 0xdc, 0x13, 0x80, 0x5e, 0x6b, 0x4d, 0x7e, 0xeb, 0x0c, 0xe1, 0x1d, 0x9c, 0xb6,
	0x64, 0xc8, 0x6d, 0xc9, 0x22, 0xaa, 0x94, 0xbf, 0x4a, 0x30, 0x3b, 0x53, 0x25, 0xe1, 0xab, 0x38,
	0x8b, 0xb7, 0x70, 0x9c, 0xdc, 0xf1, 0x32, 0x2b, 0xeb, 0x45, 0x6b, 0xe9, 0xd3, 0x1c, 0x6e, 0x98,
	0xe2, 0x21, 0xba, 0x02, 0x33, 0xd1, 0x4b, 0xed, 0xef, 0x93, 0x0c, 0xed, 0xe1, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x3e, 0x70, 0x70, 0x95, 0x01, 0x00, 0x00,
}