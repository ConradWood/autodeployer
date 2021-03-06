// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/registrymultiplexer/registrymultiplexer.proto
// DO NOT EDIT!

/*
Package registrymultiplexer is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/registrymultiplexer/registrymultiplexer.proto

It has these top-level messages:
	TargetList
	TokenRequest
	LocalTarget
	EndpointRegistration
	RegisterServiceRequest
	DeregisterRequest
	KeepAliveRequest
	KeepAliveResponse
	UserRegistration
	UserRegistrationList
*/
package registrymultiplexer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
import auth "golang.conradwood.net/apis/auth"
import registry "golang.conradwood.net/apis/registry"

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

type TargetList struct {
	Targets []*registry.Target `protobuf:"bytes,1,rep,name=Targets" json:"Targets,omitempty"`
}

func (m *TargetList) Reset()                    { *m = TargetList{} }
func (m *TargetList) String() string            { return proto.CompactTextString(m) }
func (*TargetList) ProtoMessage()               {}
func (*TargetList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TargetList) GetTargets() []*registry.Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

type TokenRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TokenRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *TokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LocalTarget struct {
	RegistrationID string           `protobuf:"bytes,1,opt,name=RegistrationID" json:"RegistrationID,omitempty"`
	Target         *registry.Target `protobuf:"bytes,2,opt,name=Target" json:"Target,omitempty"`
}

func (m *LocalTarget) Reset()                    { *m = LocalTarget{} }
func (m *LocalTarget) String() string            { return proto.CompactTextString(m) }
func (*LocalTarget) ProtoMessage()               {}
func (*LocalTarget) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LocalTarget) GetRegistrationID() string {
	if m != nil {
		return m.RegistrationID
	}
	return ""
}

func (m *LocalTarget) GetTarget() *registry.Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type EndpointRegistration struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *EndpointRegistration) Reset()                    { *m = EndpointRegistration{} }
func (m *EndpointRegistration) String() string            { return proto.CompactTextString(m) }
func (*EndpointRegistration) ProtoMessage()               {}
func (*EndpointRegistration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EndpointRegistration) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RegisterServiceRequest struct {
	Endpoint *EndpointRegistration `protobuf:"bytes,1,opt,name=Endpoint" json:"Endpoint,omitempty"`
	Target   *registry.Target      `protobuf:"bytes,2,opt,name=Target" json:"Target,omitempty"`
}

func (m *RegisterServiceRequest) Reset()                    { *m = RegisterServiceRequest{} }
func (m *RegisterServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterServiceRequest) ProtoMessage()               {}
func (*RegisterServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterServiceRequest) GetEndpoint() *EndpointRegistration {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *RegisterServiceRequest) GetTarget() *registry.Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type DeregisterRequest struct {
	Endpoint       *EndpointRegistration `protobuf:"bytes,1,opt,name=Endpoint" json:"Endpoint,omitempty"`
	RegistrationID string                `protobuf:"bytes,2,opt,name=RegistrationID" json:"RegistrationID,omitempty"`
}

func (m *DeregisterRequest) Reset()                    { *m = DeregisterRequest{} }
func (m *DeregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregisterRequest) ProtoMessage()               {}
func (*DeregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeregisterRequest) GetEndpoint() *EndpointRegistration {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *DeregisterRequest) GetRegistrationID() string {
	if m != nil {
		return m.RegistrationID
	}
	return ""
}

type KeepAliveRequest struct {
	Endpoint *EndpointRegistration `protobuf:"bytes,1,opt,name=Endpoint" json:"Endpoint,omitempty"`
}

func (m *KeepAliveRequest) Reset()                    { *m = KeepAliveRequest{} }
func (m *KeepAliveRequest) String() string            { return proto.CompactTextString(m) }
func (*KeepAliveRequest) ProtoMessage()               {}
func (*KeepAliveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *KeepAliveRequest) GetEndpoint() *EndpointRegistration {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type KeepAliveResponse struct {
	// if true the dev-server is required to (re)establish the tcp connection to the DC
	TCPRequired bool `protobuf:"varint,1,opt,name=TCPRequired" json:"TCPRequired,omitempty"`
}

func (m *KeepAliveResponse) Reset()                    { *m = KeepAliveResponse{} }
func (m *KeepAliveResponse) String() string            { return proto.CompactTextString(m) }
func (*KeepAliveResponse) ProtoMessage()               {}
func (*KeepAliveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *KeepAliveResponse) GetTCPRequired() bool {
	if m != nil {
		return m.TCPRequired
	}
	return false
}

type UserRegistration struct {
	RegistrationID string                           `protobuf:"bytes,1,opt,name=RegistrationID" json:"RegistrationID,omitempty"`
	Endpoint       *EndpointRegistration            `protobuf:"bytes,2,opt,name=Endpoint" json:"Endpoint,omitempty"`
	Request        *registry.RegisterServiceRequest `protobuf:"bytes,3,opt,name=Request" json:"Request,omitempty"`
	Target         *registry.Target                 `protobuf:"bytes,4,opt,name=Target" json:"Target,omitempty"`
}

func (m *UserRegistration) Reset()                    { *m = UserRegistration{} }
func (m *UserRegistration) String() string            { return proto.CompactTextString(m) }
func (*UserRegistration) ProtoMessage()               {}
func (*UserRegistration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserRegistration) GetRegistrationID() string {
	if m != nil {
		return m.RegistrationID
	}
	return ""
}

func (m *UserRegistration) GetEndpoint() *EndpointRegistration {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *UserRegistration) GetRequest() *registry.RegisterServiceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *UserRegistration) GetTarget() *registry.Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type UserRegistrationList struct {
	Registrations []*UserRegistration `protobuf:"bytes,1,rep,name=Registrations" json:"Registrations,omitempty"`
}

func (m *UserRegistrationList) Reset()                    { *m = UserRegistrationList{} }
func (m *UserRegistrationList) String() string            { return proto.CompactTextString(m) }
func (*UserRegistrationList) ProtoMessage()               {}
func (*UserRegistrationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UserRegistrationList) GetRegistrations() []*UserRegistration {
	if m != nil {
		return m.Registrations
	}
	return nil
}

func init() {
	proto.RegisterType((*TargetList)(nil), "registrymultiplexer.TargetList")
	proto.RegisterType((*TokenRequest)(nil), "registrymultiplexer.TokenRequest")
	proto.RegisterType((*LocalTarget)(nil), "registrymultiplexer.LocalTarget")
	proto.RegisterType((*EndpointRegistration)(nil), "registrymultiplexer.EndpointRegistration")
	proto.RegisterType((*RegisterServiceRequest)(nil), "registrymultiplexer.RegisterServiceRequest")
	proto.RegisterType((*DeregisterRequest)(nil), "registrymultiplexer.DeregisterRequest")
	proto.RegisterType((*KeepAliveRequest)(nil), "registrymultiplexer.KeepAliveRequest")
	proto.RegisterType((*KeepAliveResponse)(nil), "registrymultiplexer.KeepAliveResponse")
	proto.RegisterType((*UserRegistration)(nil), "registrymultiplexer.UserRegistration")
	proto.RegisterType((*UserRegistrationList)(nil), "registrymultiplexer.UserRegistrationList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RegistryMultiplexerService service

type RegistryMultiplexerServiceClient interface {
	// deregister a service
	DeregisterService(ctx context.Context, in *DeregisterRequest, opts ...grpc.CallOption) (*common.Void, error)
	// each devserver needs a unique ID which we get from here
	RegisterMe(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*EndpointRegistration, error)
	// we can run services locally and route datacenter traffic our way
	RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*LocalTarget, error)
	// get services exposed to a user (proxied from datacenter) (based on the calling user this may vary)
	GetServices(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*TargetList, error)
	// convenience method to create a token through dev-server
	CreateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*auth.AuthResponse, error)
	// a dev-server must send keep-alives periodically otherwise all its services will be registered
	KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error)
	// list all registrations a user has active
	ListUserRegistrations(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*UserRegistrationList, error)
}

type registryMultiplexerServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistryMultiplexerServiceClient(cc *grpc.ClientConn) RegistryMultiplexerServiceClient {
	return &registryMultiplexerServiceClient{cc}
}

func (c *registryMultiplexerServiceClient) DeregisterService(ctx context.Context, in *DeregisterRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/DeregisterService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryMultiplexerServiceClient) RegisterMe(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*EndpointRegistration, error) {
	out := new(EndpointRegistration)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/RegisterMe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryMultiplexerServiceClient) RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*LocalTarget, error) {
	out := new(LocalTarget)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/RegisterService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryMultiplexerServiceClient) GetServices(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*TargetList, error) {
	out := new(TargetList)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/GetServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryMultiplexerServiceClient) CreateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*auth.AuthResponse, error) {
	out := new(auth.AuthResponse)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/CreateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryMultiplexerServiceClient) KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error) {
	out := new(KeepAliveResponse)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/KeepAlive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryMultiplexerServiceClient) ListUserRegistrations(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*UserRegistrationList, error) {
	out := new(UserRegistrationList)
	err := grpc.Invoke(ctx, "/registrymultiplexer.RegistryMultiplexerService/ListUserRegistrations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegistryMultiplexerService service

type RegistryMultiplexerServiceServer interface {
	// deregister a service
	DeregisterService(context.Context, *DeregisterRequest) (*common.Void, error)
	// each devserver needs a unique ID which we get from here
	RegisterMe(context.Context, *common.Void) (*EndpointRegistration, error)
	// we can run services locally and route datacenter traffic our way
	RegisterService(context.Context, *RegisterServiceRequest) (*LocalTarget, error)
	// get services exposed to a user (proxied from datacenter) (based on the calling user this may vary)
	GetServices(context.Context, *common.Void) (*TargetList, error)
	// convenience method to create a token through dev-server
	CreateToken(context.Context, *TokenRequest) (*auth.AuthResponse, error)
	// a dev-server must send keep-alives periodically otherwise all its services will be registered
	KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveResponse, error)
	// list all registrations a user has active
	ListUserRegistrations(context.Context, *common.Void) (*UserRegistrationList, error)
}

func RegisterRegistryMultiplexerServiceServer(s *grpc.Server, srv RegistryMultiplexerServiceServer) {
	s.RegisterService(&_RegistryMultiplexerService_serviceDesc, srv)
}

func _RegistryMultiplexerService_DeregisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).DeregisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/DeregisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).DeregisterService(ctx, req.(*DeregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryMultiplexerService_RegisterMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).RegisterMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/RegisterMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).RegisterMe(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryMultiplexerService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).RegisterService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryMultiplexerService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).GetServices(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryMultiplexerService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).CreateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryMultiplexerService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).KeepAlive(ctx, req.(*KeepAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryMultiplexerService_ListUserRegistrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryMultiplexerServiceServer).ListUserRegistrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registrymultiplexer.RegistryMultiplexerService/ListUserRegistrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryMultiplexerServiceServer).ListUserRegistrations(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistryMultiplexerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registrymultiplexer.RegistryMultiplexerService",
	HandlerType: (*RegistryMultiplexerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeregisterService",
			Handler:    _RegistryMultiplexerService_DeregisterService_Handler,
		},
		{
			MethodName: "RegisterMe",
			Handler:    _RegistryMultiplexerService_RegisterMe_Handler,
		},
		{
			MethodName: "RegisterService",
			Handler:    _RegistryMultiplexerService_RegisterService_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _RegistryMultiplexerService_GetServices_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _RegistryMultiplexerService_CreateToken_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _RegistryMultiplexerService_KeepAlive_Handler,
		},
		{
			MethodName: "ListUserRegistrations",
			Handler:    _RegistryMultiplexerService_ListUserRegistrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/registrymultiplexer/registrymultiplexer.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/registrymultiplexer/registrymultiplexer.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x6f, 0xd3, 0x30,
	0x10, 0x57, 0xba, 0x69, 0xeb, 0x2e, 0x63, 0x74, 0x66, 0xa0, 0x2a, 0x2f, 0x94, 0x48, 0x54, 0x65,
	0x48, 0x19, 0x2a, 0x42, 0x42, 0x88, 0x97, 0xb1, 0x16, 0xa8, 0xb6, 0x49, 0x55, 0x28, 0x48, 0xf0,
	0x02, 0xa1, 0x39, 0x75, 0x16, 0xad, 0x1d, 0x6c, 0x77, 0x63, 0xaf, 0xbc, 0xf1, 0xb5, 0xf8, 0x32,
	0x7c, 0x0d, 0x94, 0xc4, 0xc9, 0xdc, 0xcc, 0xda, 0x0a, 0xda, 0x4b, 0x1b, 0xdf, 0xdd, 0xef, 0xe7,
	0xf3, 0xfd, 0x85, 0xb7, 0x13, 0x3e, 0x8d, 0xd8, 0x24, 0x18, 0x73, 0x26, 0xa2, 0xf8, 0x8c, 0xf3,
	0x38, 0x60, 0xa8, 0xf6, 0xa2, 0x84, 0xca, 0x3d, 0x81, 0x13, 0x2a, 0x95, 0x38, 0x9f, 0xcd, 0xa7,
	0x8a, 0x26, 0x53, 0xfc, 0x81, 0xc2, 0x26, 0x0b, 0x12, 0xc1, 0x15, 0x27, 0x77, 0x2c, 0x2a, 0x2f,
	0xb8, 0x82, 0x7e, 0xcc, 0x67, 0x33, 0xce, 0xf4, 0x5f, 0x4e, 0xe2, 0xed, 0x5e, 0x61, 0x1f, 0xcd,
	0xd5, 0x49, 0xf6, 0xa3, 0x6d, 0xbb, 0x4b, 0xb8, 0x5e, 0x7e, 0xe4, 0x18, 0xff, 0x39, 0xc0, 0x28,
	0x12, 0x13, 0x54, 0x47, 0x54, 0x2a, 0xb2, 0x0b, 0xeb, 0xf9, 0x49, 0x36, 0x9d, 0xd6, 0x4a, 0xc7,
	0xed, 0x36, 0x82, 0xd2, 0x3e, 0x57, 0x84, 0x85, 0x81, 0xff, 0x1a, 0x36, 0x47, 0xfc, 0x1b, 0xb2,
	0x10, 0xbf, 0xcf, 0x51, 0x2a, 0xe2, 0x41, 0xfd, 0xbd, 0x44, 0xc1, 0xa2, 0x19, 0x36, 0x9d, 0x96,
	0xd3, 0xd9, 0x08, 0xcb, 0x73, 0xaa, 0x1b, 0x46, 0x52, 0x9e, 0x71, 0x11, 0x37, 0x6b, 0xb9, 0xae,
	0x38, 0xfb, 0x9f, 0xc1, 0x3d, 0xe2, 0xe3, 0x68, 0x9a, 0xf3, 0x92, 0x36, 0x6c, 0x85, 0xf9, 0x95,
	0x91, 0xa2, 0x9c, 0x0d, 0x7a, 0x9a, 0xac, 0x22, 0x25, 0x1d, 0x58, 0xcb, 0x11, 0x19, 0xa1, 0xcd,
	0x53, 0xad, 0xf7, 0xdb, 0xb0, 0xd3, 0x67, 0x71, 0xc2, 0x29, 0x53, 0x26, 0x07, 0xd9, 0x82, 0x5a,
	0xc9, 0x5e, 0x1b, 0xf4, 0xfc, 0x5f, 0x0e, 0xdc, 0xcb, 0x0d, 0x50, 0xbc, 0x43, 0x71, 0x4a, 0xc7,
	0x58, 0xbc, 0xad, 0x0f, 0xf5, 0x82, 0x22, 0x03, 0xb8, 0xdd, 0x47, 0x81, 0x2d, 0xf1, 0xb6, 0x7b,
	0xc2, 0x12, 0xfa, 0x0f, 0x3e, 0xff, 0x74, 0x60, 0xbb, 0x87, 0x42, 0x7b, 0x73, 0xc3, 0x6e, 0x5c,
	0x0e, 0x71, 0xcd, 0x16, 0x62, 0xff, 0x23, 0x34, 0x0e, 0x11, 0x93, 0xfd, 0x29, 0x3d, 0xbd, 0xe1,
	0x48, 0xf8, 0xcf, 0x60, 0xdb, 0xa0, 0x96, 0x09, 0x67, 0x12, 0x49, 0x0b, 0xdc, 0xd1, 0xc1, 0x30,
	0xbd, 0x89, 0x0a, 0x8c, 0x33, 0xfa, 0x7a, 0x68, 0x8a, 0xfc, 0x3f, 0x0e, 0x34, 0xd2, 0xa2, 0x5a,
	0xc8, 0xe3, 0xb2, 0x15, 0x63, 0xba, 0x5e, 0xfb, 0xff, 0xe8, 0xbd, 0x80, 0x75, 0x1d, 0x8c, 0xe6,
	0x4a, 0xc6, 0xd2, 0xba, 0xc8, 0xa2, 0xbd, 0x7c, 0xc2, 0x02, 0x60, 0x14, 0xc0, 0xea, 0x35, 0x05,
	0x30, 0x86, 0x9d, 0xea, 0x43, 0xb3, 0x0e, 0x3d, 0x84, 0x5b, 0xa6, 0xac, 0xe8, 0xd3, 0x87, 0xd6,
	0x97, 0x54, 0x19, 0xc2, 0x45, 0x6c, 0xf7, 0xf7, 0x2a, 0x78, 0x5a, 0x72, 0x7e, 0x7c, 0x81, 0xd3,
	0xde, 0x93, 0x81, 0x59, 0x83, 0x85, 0xb0, 0x6d, 0xbd, 0xe9, 0x52, 0xad, 0x7a, 0x9b, 0x81, 0x9e,
	0x63, 0x1f, 0x38, 0x8d, 0x49, 0x1f, 0xa0, 0x88, 0xcd, 0x31, 0x92, 0x05, 0x9d, 0xb7, 0x7c, 0x16,
	0xc8, 0x17, 0xb8, 0x5d, 0x09, 0x31, 0x79, 0x6c, 0x45, 0xdb, 0x13, 0xe1, 0xb5, 0xac, 0xc6, 0xe6,
	0xf8, 0x79, 0x09, 0xee, 0x1b, 0x54, 0x1a, 0x26, 0x2b, 0x9e, 0xde, 0xb7, 0xc2, 0x8d, 0xf9, 0xd9,
	0x03, 0xf7, 0x40, 0x60, 0xa4, 0x30, 0x9b, 0x8c, 0xe4, 0x81, 0xdd, 0xde, 0x98, 0x9a, 0x1e, 0x09,
	0xb2, 0x01, 0xbe, 0x3f, 0x57, 0x27, 0x65, 0x1f, 0x7c, 0x82, 0x8d, 0xb2, 0x39, 0x88, 0x3d, 0xb3,
	0xd5, 0xbe, 0xf4, 0xda, 0xd7, 0x99, 0x69, 0xee, 0x21, 0xdc, 0x4d, 0x3d, 0xad, 0x56, 0x86, 0x5c,
	0x2a, 0x27, 0xb6, 0x8a, 0x7c, 0xf5, 0x04, 0xd2, 0x2d, 0x63, 0x2e, 0x1d, 0xbd, 0x86, 0xd2, 0xbd,
	0x63, 0xa3, 0xf9, 0xba, 0x96, 0xad, 0x9e, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xc1,
	0xa9, 0x61, 0x6b, 0x07, 0x00, 0x00,
}
