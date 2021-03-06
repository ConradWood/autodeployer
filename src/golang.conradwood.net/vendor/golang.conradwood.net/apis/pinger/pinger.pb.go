// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/pinger/pinger.proto
// DO NOT EDIT!

/*
Package pinger is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/pinger/pinger.proto

It has these top-level messages:
	PingResult
	PingRequest
	PingListRequest
	PingList
	PingEntry
*/
package pinger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "golang.conradwood.net/apis/common"

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

// result of a single ping
type PingResult struct {
	IP           string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Milliseconds uint32 `protobuf:"varint,2,opt,name=Milliseconds" json:"Milliseconds,omitempty"`
	Success      bool   `protobuf:"varint,3,opt,name=Success" json:"Success,omitempty"`
}

func (m *PingResult) Reset()                    { *m = PingResult{} }
func (m *PingResult) String() string            { return proto.CompactTextString(m) }
func (*PingResult) ProtoMessage()               {}
func (*PingResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResult) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *PingResult) GetMilliseconds() uint32 {
	if m != nil {
		return m.Milliseconds
	}
	return 0
}

func (m *PingResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// ping something once
type PingRequest struct {
	IP string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

type PingListRequest struct {
	PingerID string `protobuf:"bytes,1,opt,name=PingerID" json:"PingerID,omitempty"`
}

func (m *PingListRequest) Reset()                    { *m = PingListRequest{} }
func (m *PingListRequest) String() string            { return proto.CompactTextString(m) }
func (*PingListRequest) ProtoMessage()               {}
func (*PingListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingListRequest) GetPingerID() string {
	if m != nil {
		return m.PingerID
	}
	return ""
}

type PingList struct {
	Entries []*PingEntry `protobuf:"bytes,1,rep,name=Entries" json:"Entries,omitempty"`
}

func (m *PingList) Reset()                    { *m = PingList{} }
func (m *PingList) String() string            { return proto.CompactTextString(m) }
func (*PingList) ProtoMessage()               {}
func (*PingList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PingList) GetEntries() []*PingEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type PingEntry struct {
	ID             uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	IP             string `protobuf:"bytes,2,opt,name=IP" json:"IP,omitempty"`
	Interval       uint32 `protobuf:"varint,3,opt,name=Interval" json:"Interval,omitempty"`
	MetricHostName string `protobuf:"bytes,4,opt,name=MetricHostName" json:"MetricHostName,omitempty"`
	PingerID       string `protobuf:"bytes,5,opt,name=PingerID" json:"PingerID,omitempty"`
}

func (m *PingEntry) Reset()                    { *m = PingEntry{} }
func (m *PingEntry) String() string            { return proto.CompactTextString(m) }
func (*PingEntry) ProtoMessage()               {}
func (*PingEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PingEntry) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PingEntry) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *PingEntry) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *PingEntry) GetMetricHostName() string {
	if m != nil {
		return m.MetricHostName
	}
	return ""
}

func (m *PingEntry) GetPingerID() string {
	if m != nil {
		return m.PingerID
	}
	return ""
}

func init() {
	proto.RegisterType((*PingResult)(nil), "pinger.PingResult")
	proto.RegisterType((*PingRequest)(nil), "pinger.PingRequest")
	proto.RegisterType((*PingListRequest)(nil), "pinger.PingListRequest")
	proto.RegisterType((*PingList)(nil), "pinger.PingList")
	proto.RegisterType((*PingEntry)(nil), "pinger.PingEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pinger service

type PingerClient interface {
	// execute a single ping
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResult, error)
}

type pingerClient struct {
	cc *grpc.ClientConn
}

func NewPingerClient(cc *grpc.ClientConn) PingerClient {
	return &pingerClient{cc}
}

func (c *pingerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResult, error) {
	out := new(PingResult)
	err := grpc.Invoke(ctx, "/pinger.Pinger/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pinger service

type PingerServer interface {
	// execute a single ping
	Ping(context.Context, *PingRequest) (*PingResult, error)
}

func RegisterPingerServer(s *grpc.Server, srv PingerServer) {
	s.RegisterService(&_Pinger_serviceDesc, srv)
}

func _Pinger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinger.Pinger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pinger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pinger.Pinger",
	HandlerType: (*PingerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Pinger_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/pinger/pinger.proto",
}

// Client API for PingerList service

type PingerListClient interface {
	// get ping config (list of ip addresses to ping regularly).
	GetPingList(ctx context.Context, in *PingListRequest, opts ...grpc.CallOption) (*PingList, error)
}

type pingerListClient struct {
	cc *grpc.ClientConn
}

func NewPingerListClient(cc *grpc.ClientConn) PingerListClient {
	return &pingerListClient{cc}
}

func (c *pingerListClient) GetPingList(ctx context.Context, in *PingListRequest, opts ...grpc.CallOption) (*PingList, error) {
	out := new(PingList)
	err := grpc.Invoke(ctx, "/pinger.PingerList/GetPingList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PingerList service

type PingerListServer interface {
	// get ping config (list of ip addresses to ping regularly).
	GetPingList(context.Context, *PingListRequest) (*PingList, error)
}

func RegisterPingerListServer(s *grpc.Server, srv PingerListServer) {
	s.RegisterService(&_PingerList_serviceDesc, srv)
}

func _PingerList_GetPingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingerListServer).GetPingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pinger.PingerList/GetPingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingerListServer).GetPingList(ctx, req.(*PingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingerList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pinger.PingerList",
	HandlerType: (*PingerListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPingList",
			Handler:    _PingerList_GetPingList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/pinger/pinger.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/pinger/pinger.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x69, 0xed, 0xc7, 0xd4, 0x56, 0x5d, 0x0f, 0x86, 0xa0, 0x10, 0x72, 0x90, 0x80, 0x98,
	0x42, 0x3d, 0xa8, 0x57, 0xa9, 0x1f, 0x01, 0x2b, 0x61, 0xbd, 0x79, 0x8b, 0xe9, 0x12, 0x16, 0xd2,
	0xdd, 0xba, 0xbb, 0x55, 0xfc, 0x0d, 0xfe, 0x69, 0xd9, 0xec, 0x36, 0x24, 0x15, 0x3c, 0xed, 0xce,
	0xcc, 0x7b, 0xc3, 0xbc, 0x37, 0x03, 0x71, 0xc1, 0xcb, 0x8c, 0x15, 0x71, 0xce, 0x99, 0xc8, 0x96,
	0x5f, 0x9c, 0x2f, 0x63, 0x46, 0xd4, 0x34, 0x5b, 0x53, 0x39, 0x5d, 0x53, 0x56, 0x10, 0x61, 0x9f,
	0x78, 0x2d, 0xb8, 0xe2, 0xa8, 0x67, 0x22, 0xff, 0x3f, 0x5e, 0xce, 0x57, 0x2b, 0xce, 0xec, 0x63,
	0x78, 0xe1, 0x1b, 0x40, 0x4a, 0x59, 0x81, 0x89, 0xdc, 0x94, 0x0a, 0x4d, 0xc0, 0x4d, 0x52, 0xcf,
	0x09, 0x9c, 0x68, 0x88, 0xdd, 0x24, 0x45, 0x21, 0xec, 0x2f, 0x68, 0x59, 0x52, 0x49, 0x72, 0xce,
	0x96, 0xd2, 0x73, 0x03, 0x27, 0x1a, 0xe3, 0x56, 0x0e, 0x79, 0xd0, 0x7f, 0xdd, 0xe4, 0x39, 0x91,
	0xd2, 0xeb, 0x04, 0x4e, 0x34, 0xc0, 0xdb, 0x30, 0x3c, 0x83, 0x91, 0xe9, 0xfd, 0xb1, 0x21, 0xf2,
	0x4f, 0xf3, 0xf0, 0x12, 0x0e, 0x74, 0xf9, 0x99, 0x4a, 0xb5, 0x85, 0xf8, 0x30, 0x48, 0x2b, 0x1d,
	0xc9, 0xdc, 0x02, 0xeb, 0x38, 0xbc, 0x36, 0x35, 0x0d, 0x47, 0x17, 0xd0, 0xbf, 0x67, 0x4a, 0x50,
	0x22, 0x3d, 0x27, 0xe8, 0x44, 0xa3, 0xd9, 0x51, 0x6c, 0xdd, 0xd0, 0x10, 0x5d, 0xfa, 0xc6, 0x5b,
	0x44, 0xf8, 0xe3, 0xc0, 0xb0, 0x4e, 0x57, 0x53, 0x98, 0xe6, 0x5d, 0xec, 0x26, 0x73, 0x3b, 0x95,
	0x5b, 0x4b, 0xf6, 0x61, 0x90, 0x30, 0x45, 0xc4, 0x67, 0x56, 0x56, 0x7a, 0xc6, 0xb8, 0x8e, 0xd1,
	0x39, 0x4c, 0x16, 0x44, 0x09, 0x9a, 0x3f, 0x71, 0xa9, 0x5e, 0xb2, 0x15, 0xf1, 0xba, 0x15, 0x6f,
	0x27, 0xdb, 0x92, 0xb1, 0xd7, 0x96, 0x31, 0xbb, 0x85, 0x9e, 0xf9, 0xa3, 0x29, 0x74, 0xf5, 0x0f,
	0x1d, 0x37, 0x67, 0xb7, 0x4e, 0xf8, 0xa8, 0x9d, 0xd4, 0xdb, 0x99, 0x3d, 0x98, 0x5d, 0x11, 0x51,
	0x79, 0x70, 0x03, 0xa3, 0x47, 0xa2, 0x6a, 0x4b, 0x4e, 0x9a, 0x84, 0x86, 0xa7, 0xfe, 0xe1, 0x6e,
	0xe1, 0xee, 0x14, 0x7c, 0x46, 0x54, 0xf3, 0x44, 0xf4, 0x79, 0x58, 0xd8, 0x7b, 0xaf, 0x3a, 0x8c,
	0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x24, 0x33, 0x60, 0x82, 0x02, 0x00, 0x00,
}
