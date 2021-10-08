// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/frontendconfigurator/frontendconfigurator.proto
// DO NOT EDIT!

/*
Package frontendconfigurator is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/frontendconfigurator/frontendconfigurator.proto

It has these top-level messages:
	Application
	CreateApplicationRequest
	CreateApplicationResponse
	DeleteApplicationRequest
	GetApplicationsRequest
	GetApplicationsResponse
*/
package frontendconfigurator

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

type Application struct {
	ID        uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Feature   string `protobuf:"bytes,3,opt,name=Feature" json:"Feature,omitempty"`
	UserGroup string `protobuf:"bytes,4,opt,name=UserGroup" json:"UserGroup,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Application) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetFeature() string {
	if m != nil {
		return m.Feature
	}
	return ""
}

func (m *Application) GetUserGroup() string {
	if m != nil {
		return m.UserGroup
	}
	return ""
}

// create an application
type CreateApplicationRequest struct {
	Application *Application `protobuf:"bytes,1,opt,name=Application" json:"Application,omitempty"`
}

func (m *CreateApplicationRequest) Reset()                    { *m = CreateApplicationRequest{} }
func (m *CreateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationRequest) ProtoMessage()               {}
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateApplicationRequest) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

// return the createdapplication's id
type CreateApplicationResponse struct {
	ID uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *CreateApplicationResponse) Reset()                    { *m = CreateApplicationResponse{} }
func (m *CreateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationResponse) ProtoMessage()               {}
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateApplicationResponse) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// delete an application by id
type DeleteApplicationRequest struct {
	ID uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *DeleteApplicationRequest) Reset()                    { *m = DeleteApplicationRequest{} }
func (m *DeleteApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationRequest) ProtoMessage()               {}
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteApplicationRequest) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// get application by group name
type GetApplicationsRequest struct {
	UserGroup string `protobuf:"bytes,1,opt,name=UserGroup" json:"UserGroup,omitempty"`
}

func (m *GetApplicationsRequest) Reset()                    { *m = GetApplicationsRequest{} }
func (m *GetApplicationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationsRequest) ProtoMessage()               {}
func (*GetApplicationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetApplicationsRequest) GetUserGroup() string {
	if m != nil {
		return m.UserGroup
	}
	return ""
}

type GetApplicationsResponse struct {
	Applications []*Application `protobuf:"bytes,1,rep,name=Applications" json:"Applications,omitempty"`
}

func (m *GetApplicationsResponse) Reset()                    { *m = GetApplicationsResponse{} }
func (m *GetApplicationsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationsResponse) ProtoMessage()               {}
func (*GetApplicationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetApplicationsResponse) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

func init() {
	proto.RegisterType((*Application)(nil), "frontendconfigurator.Application")
	proto.RegisterType((*CreateApplicationRequest)(nil), "frontendconfigurator.CreateApplicationRequest")
	proto.RegisterType((*CreateApplicationResponse)(nil), "frontendconfigurator.CreateApplicationResponse")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "frontendconfigurator.DeleteApplicationRequest")
	proto.RegisterType((*GetApplicationsRequest)(nil), "frontendconfigurator.GetApplicationsRequest")
	proto.RegisterType((*GetApplicationsResponse)(nil), "frontendconfigurator.GetApplicationsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FrontEndConfiguratorService service

type FrontEndConfiguratorServiceClient interface {
	CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error)
	DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*common.Void, error)
	GetApplications(ctx context.Context, in *GetApplicationsRequest, opts ...grpc.CallOption) (*GetApplicationsResponse, error)
}

type frontEndConfiguratorServiceClient struct {
	cc *grpc.ClientConn
}

func NewFrontEndConfiguratorServiceClient(cc *grpc.ClientConn) FrontEndConfiguratorServiceClient {
	return &frontEndConfiguratorServiceClient{cc}
}

func (c *frontEndConfiguratorServiceClient) CreateApplication(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error) {
	out := new(CreateApplicationResponse)
	err := grpc.Invoke(ctx, "/frontendconfigurator.FrontEndConfiguratorService/CreateApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontEndConfiguratorServiceClient) DeleteApplication(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/frontendconfigurator.FrontEndConfiguratorService/DeleteApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontEndConfiguratorServiceClient) GetApplications(ctx context.Context, in *GetApplicationsRequest, opts ...grpc.CallOption) (*GetApplicationsResponse, error) {
	out := new(GetApplicationsResponse)
	err := grpc.Invoke(ctx, "/frontendconfigurator.FrontEndConfiguratorService/GetApplications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontEndConfiguratorService service

type FrontEndConfiguratorServiceServer interface {
	CreateApplication(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error)
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*common.Void, error)
	GetApplications(context.Context, *GetApplicationsRequest) (*GetApplicationsResponse, error)
}

func RegisterFrontEndConfiguratorServiceServer(s *grpc.Server, srv FrontEndConfiguratorServiceServer) {
	s.RegisterService(&_FrontEndConfiguratorService_serviceDesc, srv)
}

func _FrontEndConfiguratorService_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontEndConfiguratorServiceServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontendconfigurator.FrontEndConfiguratorService/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontEndConfiguratorServiceServer).CreateApplication(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontEndConfiguratorService_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontEndConfiguratorServiceServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontendconfigurator.FrontEndConfiguratorService/DeleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontEndConfiguratorServiceServer).DeleteApplication(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrontEndConfiguratorService_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontEndConfiguratorServiceServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontendconfigurator.FrontEndConfiguratorService/GetApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontEndConfiguratorServiceServer).GetApplications(ctx, req.(*GetApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FrontEndConfiguratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontendconfigurator.FrontEndConfiguratorService",
	HandlerType: (*FrontEndConfiguratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _FrontEndConfiguratorService_CreateApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _FrontEndConfiguratorService_DeleteApplication_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _FrontEndConfiguratorService_GetApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/frontendconfigurator/frontendconfigurator.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/frontendconfigurator/frontendconfigurator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x4f, 0xea, 0x40,
	0x10, 0x4d, 0x0b, 0xb9, 0x37, 0x0c, 0xdc, 0x7b, 0xc3, 0xe6, 0x46, 0xd7, 0xea, 0x03, 0xf6, 0x89,
	0xf8, 0xd1, 0x46, 0x4c, 0x7c, 0x57, 0xbe, 0x82, 0x0f, 0x26, 0x62, 0xf4, 0x55, 0xd7, 0x76, 0x20,
	0x4d, 0x60, 0xa7, 0x6e, 0xb7, 0xfa, 0x4f, 0xfd, 0x3d, 0x86, 0x52, 0x62, 0x29, 0x8b, 0xe1, 0xa9,
	0xed, 0xcc, 0x39, 0x73, 0xe6, 0xcc, 0x74, 0xe0, 0x76, 0x4a, 0x33, 0x21, 0xa7, 0x5e, 0x40, 0x52,
	0x89, 0xf0, 0x83, 0x28, 0xf4, 0x24, 0x6a, 0x5f, 0xc4, 0x51, 0xe2, 0x4f, 0x14, 0x49, 0x8d, 0x32,
	0x0c, 0x48, 0x4e, 0xa2, 0x69, 0xaa, 0x84, 0x26, 0x65, 0x0c, 0x7a, 0xb1, 0x22, 0x4d, 0xec, 0xbf,
	0x29, 0xe7, 0x78, 0x3f, 0x28, 0x04, 0x34, 0x9f, 0x93, 0xcc, 0x1f, 0xcb, 0x2a, 0x6e, 0x04, 0xf5,
	0xeb, 0x38, 0x9e, 0x45, 0x81, 0xd0, 0x11, 0x49, 0xf6, 0x17, 0xec, 0x51, 0x8f, 0x5b, 0x2d, 0xab,
	0xfd, 0x67, 0x6c, 0x8f, 0x7a, 0x8c, 0x41, 0xf5, 0x4e, 0xcc, 0x91, 0xdb, 0x2d, 0xab, 0x5d, 0x1b,
	0x67, 0xef, 0x8c, 0xc3, 0xef, 0x01, 0x0a, 0x9d, 0x2a, 0xe4, 0x95, 0x2c, 0xbc, 0xfa, 0x64, 0x47,
	0x50, 0x7b, 0x4c, 0x50, 0x0d, 0x15, 0xa5, 0x31, 0xaf, 0x66, 0xb9, 0xef, 0x80, 0xfb, 0x0c, 0xbc,
	0xab, 0x50, 0x68, 0x2c, 0x08, 0x8e, 0xf1, 0x2d, 0xc5, 0x44, 0xb3, 0xee, 0x5a, 0x1b, 0x59, 0x03,
	0xf5, 0xce, 0xb1, 0x67, 0xb4, 0x5f, 0xa4, 0x17, 0x59, 0xee, 0x29, 0x1c, 0x18, 0x04, 0x92, 0x98,
	0x64, 0x82, 0x65, 0x67, 0xee, 0x09, 0xf0, 0x1e, 0xce, 0xd0, 0xd8, 0x4d, 0x19, 0x7b, 0x05, 0x7b,
	0x43, 0xd4, 0x05, 0x60, 0xb2, 0x42, 0xae, 0x39, 0xb6, 0xca, 0x8e, 0x5f, 0x60, 0x7f, 0x83, 0x97,
	0xb7, 0xd3, 0x87, 0x46, 0x31, 0xce, 0xad, 0x56, 0x65, 0x37, 0xc7, 0x6b, 0xb4, 0xce, 0xa7, 0x0d,
	0x87, 0x83, 0x05, 0xa5, 0x2f, 0xc3, 0x6e, 0x81, 0xf2, 0x80, 0xea, 0x3d, 0x0a, 0x90, 0x69, 0x68,
	0x6e, 0x8c, 0x84, 0x79, 0x66, 0x95, 0x6d, 0xcb, 0x71, 0xfc, 0x9d, 0xf1, 0xb9, 0xb9, 0x7b, 0x68,
	0x6e, 0xcc, 0x76, 0x9b, 0xea, 0xb6, 0x25, 0x38, 0x0d, 0x2f, 0xff, 0x51, 0x9f, 0x28, 0x0a, 0x99,
	0x84, 0x7f, 0xa5, 0x51, 0xb2, 0x33, 0x73, 0x41, 0xf3, 0xa6, 0x9c, 0xf3, 0x1d, 0xd1, 0x4b, 0x0b,
	0x37, 0x17, 0xe0, 0x4b, 0xd4, 0xc5, 0x33, 0xca, 0x0f, 0x6b, 0x71, 0x49, 0xc6, 0x52, 0xaf, 0xbf,
	0xb2, 0x8b, 0xba, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x85, 0x9f, 0x20, 0x6b, 0xe5, 0x03, 0x00,
	0x00,
}
