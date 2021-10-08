// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/phonebook/phonebook.proto
// DO NOT EDIT!

/*
Package phonebook is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/phonebook/phonebook.proto

It has these top-level messages:
*/
package phonebook

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
import goasterisk "golang.conradwood.net/apis/goasterisk"

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

// Client API for PhoneBookService service

type PhoneBookServiceClient interface {
	// get entries in phonebook
	GetEntries(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*goasterisk.ContactList, error)
	// set up a "dialthrough" to a contact
	Dialthrough(ctx context.Context, in *goasterisk.Contact, opts ...grpc.CallOption) (*goasterisk.DialThroughResponse, error)
}

type phoneBookServiceClient struct {
	cc *grpc.ClientConn
}

func NewPhoneBookServiceClient(cc *grpc.ClientConn) PhoneBookServiceClient {
	return &phoneBookServiceClient{cc}
}

func (c *phoneBookServiceClient) GetEntries(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*goasterisk.ContactList, error) {
	out := new(goasterisk.ContactList)
	err := grpc.Invoke(ctx, "/phonebook.PhoneBookService/GetEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phoneBookServiceClient) Dialthrough(ctx context.Context, in *goasterisk.Contact, opts ...grpc.CallOption) (*goasterisk.DialThroughResponse, error) {
	out := new(goasterisk.DialThroughResponse)
	err := grpc.Invoke(ctx, "/phonebook.PhoneBookService/Dialthrough", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhoneBookService service

type PhoneBookServiceServer interface {
	// get entries in phonebook
	GetEntries(context.Context, *common.Void) (*goasterisk.ContactList, error)
	// set up a "dialthrough" to a contact
	Dialthrough(context.Context, *goasterisk.Contact) (*goasterisk.DialThroughResponse, error)
}

func RegisterPhoneBookServiceServer(s *grpc.Server, srv PhoneBookServiceServer) {
	s.RegisterService(&_PhoneBookService_serviceDesc, srv)
}

func _PhoneBookService_GetEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServiceServer).GetEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonebook.PhoneBookService/GetEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServiceServer).GetEntries(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhoneBookService_Dialthrough_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(goasterisk.Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneBookServiceServer).Dialthrough(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phonebook.PhoneBookService/Dialthrough",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneBookServiceServer).Dialthrough(ctx, req.(*goasterisk.Contact))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhoneBookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "phonebook.PhoneBookService",
	HandlerType: (*PhoneBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntries",
			Handler:    _PhoneBookService_GetEntries_Handler,
		},
		{
			MethodName: "Dialthrough",
			Handler:    _PhoneBookService_Dialthrough_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/phonebook/phonebook.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/phonebook/phonebook.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xb1, 0x11, 0x8c, 0x16, 0x12, 0x0b, 0x21, 0x8d, 0x8d, 0x58, 0xe6, 0xc0, 0x13, 0x7f,
	0xc0, 0x9d, 0x62, 0x63, 0x21, 0x2a, 0xf6, 0xb9, 0xec, 0x90, 0x0d, 0x7b, 0x3b, 0x2f, 0x24, 0xa3,
	0xfe, 0x09, 0x7f, 0xb4, 0xac, 0xbb, 0xac, 0x29, 0x8e, 0xad, 0x32, 0x0c, 0xdf, 0x7b, 0x2f, 0x6f,
	0xd4, 0x5d, 0xc0, 0xde, 0x71, 0xb0, 0x1e, 0x9c, 0x5d, 0xf3, 0x0d, 0x34, 0x96, 0x49, 0x56, 0x2e,
	0xc5, 0xb2, 0x4a, 0x2d, 0x98, 0x76, 0x40, 0xf7, 0x3f, 0xd9, 0x94, 0x21, 0xd0, 0x27, 0xf3, 0xc2,
	0xd8, 0x05, 0x03, 0x8f, 0xbe, 0x07, 0x4f, 0xcf, 0x28, 0x35, 0xf7, 0x0b, 0x7c, 0x80, 0x2b, 0x42,
	0x39, 0x96, 0xae, 0x1a, 0x47, 0xdd, 0xed, 0xcf, 0x91, 0x3a, 0x7f, 0x19, 0x52, 0x37, 0x40, 0xf7,
	0x46, 0xf9, 0x2b, 0x7a, 0xd2, 0x6b, 0xa5, 0x9e, 0x48, 0x1e, 0x59, 0x72, 0xa4, 0xa2, 0xcf, 0xec,
	0x94, 0xf4, 0x81, 0xd8, 0x98, 0x4b, 0x5b, 0x79, 0x6c, 0xc1, 0xe2, 0xbc, 0x3c, 0xc7, 0x22, 0x7a,
	0xab, 0x4e, 0x1f, 0xa2, 0xdb, 0x4b, 0x9b, 0xf1, 0x19, 0x5a, 0x7d, 0x71, 0x80, 0x33, 0x57, 0xf5,
	0x72, 0xa0, 0xdf, 0x47, 0xfa, 0x95, 0x4a, 0x02, 0x17, 0xda, 0xdc, 0xa8, 0x6b, 0x26, 0xa9, 0x5b,
	0x4c, 0xbd, 0x86, 0x22, 0x76, 0x3e, 0xcf, 0xee, 0xf8, 0xef, 0xf7, 0xeb, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x63, 0x85, 0x68, 0x5c, 0x68, 0x01, 0x00, 0x00,
}