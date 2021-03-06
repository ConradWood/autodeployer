// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/scapi/scapi.proto
// DO NOT EDIT!

/*
Package scapi is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/scapi/scapi.proto

It has these top-level messages:
	PingRequest
	PingResponse
	ModuleNameUpdate
	SensorConfigUpdateRequest
	UpdateUserAppVersionRequest
	UpdateUserAppRepoRequest
	Version
	UserAppVersions
	UserConfigFlagRequest
*/
package scapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import scweb "golang.singingcat.net/apis/scweb"
import scfunctions "golang.singingcat.net/apis/scfunctions"
import scupdate "golang.singingcat.net/apis/scupdate"
import common "golang.conradwood.net/apis/common"
import singingcat "golang.singingcat.net/apis/singingcat"

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

type ModuleNameUpdate struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *ModuleNameUpdate) Reset()                    { *m = ModuleNameUpdate{} }
func (m *ModuleNameUpdate) String() string            { return proto.CompactTextString(m) }
func (*ModuleNameUpdate) ProtoMessage()               {}
func (*ModuleNameUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ModuleNameUpdate) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *ModuleNameUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SensorConfigUpdateRequest struct {
	SensorDefID       uint64 `protobuf:"varint,1,opt,name=SensorDefID" json:"SensorDefID,omitempty"`
	Enable            bool   `protobuf:"varint,2,opt,name=Enable" json:"Enable,omitempty"`
	ReportingInterval uint32 `protobuf:"varint,3,opt,name=ReportingInterval" json:"ReportingInterval,omitempty"`
	PollingInterval   uint32 `protobuf:"varint,4,opt,name=PollingInterval" json:"PollingInterval,omitempty"`
	FriendlyName      string `protobuf:"bytes,5,opt,name=FriendlyName" json:"FriendlyName,omitempty"`
}

func (m *SensorConfigUpdateRequest) Reset()                    { *m = SensorConfigUpdateRequest{} }
func (m *SensorConfigUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*SensorConfigUpdateRequest) ProtoMessage()               {}
func (*SensorConfigUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SensorConfigUpdateRequest) GetSensorDefID() uint64 {
	if m != nil {
		return m.SensorDefID
	}
	return 0
}

func (m *SensorConfigUpdateRequest) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *SensorConfigUpdateRequest) GetReportingInterval() uint32 {
	if m != nil {
		return m.ReportingInterval
	}
	return 0
}

func (m *SensorConfigUpdateRequest) GetPollingInterval() uint32 {
	if m != nil {
		return m.PollingInterval
	}
	return 0
}

func (m *SensorConfigUpdateRequest) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

type UpdateUserAppVersionRequest struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
	Build    uint64 `protobuf:"varint,2,opt,name=Build" json:"Build,omitempty"`
}

func (m *UpdateUserAppVersionRequest) Reset()                    { *m = UpdateUserAppVersionRequest{} }
func (m *UpdateUserAppVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserAppVersionRequest) ProtoMessage()               {}
func (*UpdateUserAppVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateUserAppVersionRequest) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *UpdateUserAppVersionRequest) GetBuild() uint64 {
	if m != nil {
		return m.Build
	}
	return 0
}

type UpdateUserAppRepoRequest struct {
	NewValue string `protobuf:"bytes,1,opt,name=NewValue" json:"NewValue,omitempty"`
	ModuleID uint64 `protobuf:"varint,2,opt,name=ModuleID" json:"ModuleID,omitempty"`
}

func (m *UpdateUserAppRepoRequest) Reset()                    { *m = UpdateUserAppRepoRequest{} }
func (m *UpdateUserAppRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserAppRepoRequest) ProtoMessage()               {}
func (*UpdateUserAppRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateUserAppRepoRequest) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

func (m *UpdateUserAppRepoRequest) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

type Version struct {
	ID       string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=Selected" json:"Selected,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Version) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type UserAppVersions struct {
	Versions []*Version `protobuf:"bytes,1,rep,name=Versions" json:"Versions,omitempty"`
}

func (m *UserAppVersions) Reset()                    { *m = UserAppVersions{} }
func (m *UserAppVersions) String() string            { return proto.CompactTextString(m) }
func (*UserAppVersions) ProtoMessage()               {}
func (*UserAppVersions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserAppVersions) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type UserConfigFlagRequest struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
	Bit      uint32 `protobuf:"varint,2,opt,name=Bit" json:"Bit,omitempty"`
	NewState bool   `protobuf:"varint,3,opt,name=NewState" json:"NewState,omitempty"`
}

func (m *UserConfigFlagRequest) Reset()                    { *m = UserConfigFlagRequest{} }
func (m *UserConfigFlagRequest) String() string            { return proto.CompactTextString(m) }
func (*UserConfigFlagRequest) ProtoMessage()               {}
func (*UserConfigFlagRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserConfigFlagRequest) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *UserConfigFlagRequest) GetBit() uint32 {
	if m != nil {
		return m.Bit
	}
	return 0
}

func (m *UserConfigFlagRequest) GetNewState() bool {
	if m != nil {
		return m.NewState
	}
	return false
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "scapi.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "scapi.PingResponse")
	proto.RegisterType((*ModuleNameUpdate)(nil), "scapi.ModuleNameUpdate")
	proto.RegisterType((*SensorConfigUpdateRequest)(nil), "scapi.SensorConfigUpdateRequest")
	proto.RegisterType((*UpdateUserAppVersionRequest)(nil), "scapi.UpdateUserAppVersionRequest")
	proto.RegisterType((*UpdateUserAppRepoRequest)(nil), "scapi.UpdateUserAppRepoRequest")
	proto.RegisterType((*Version)(nil), "scapi.Version")
	proto.RegisterType((*UserAppVersions)(nil), "scapi.UserAppVersions")
	proto.RegisterType((*UserConfigFlagRequest)(nil), "scapi.UserConfigFlagRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SCApiService service

type SCApiServiceClient interface {
	// just ping the api
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// update name of a module
	UpdateModuleName(ctx context.Context, in *ModuleNameUpdate, opts ...grpc.CallOption) (*common.Void, error)
	// update friendly name of sensor
	UpdateSensorName(ctx context.Context, in *scweb.SensorNameUpdate, opts ...grpc.CallOption) (*common.Void, error)
	// set a config flag
	SetConfigFlag(ctx context.Context, in *scfunctions.SetFlagRequest, opts ...grpc.CallOption) (*common.Void, error)
	// set a userappconfig flag
	SetUserConfigFlag(ctx context.Context, in *UserConfigFlagRequest, opts ...grpc.CallOption) (*common.Void, error)
	// update sensor configuration
	UpdateSensorConfiguration(ctx context.Context, in *SensorConfigUpdateRequest, opts ...grpc.CallOption) (*common.Void, error)
	// update the user app repo
	UpdateUserAppRepo(ctx context.Context, in *UpdateUserAppRepoRequest, opts ...grpc.CallOption) (*UserAppVersions, error)
	// update the user app version
	UpdateUserAppVersion(ctx context.Context, in *UpdateUserAppVersionRequest, opts ...grpc.CallOption) (*common.Void, error)
	// get all online modules
	GetOnlineModules(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*scweb.ModuleList, error)
	// get a module's current update (flash in progress)
	GetModuleFlashStatus(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*scupdate.UpdateStatusList, error)
}

type sCApiServiceClient struct {
	cc *grpc.ClientConn
}

func NewSCApiServiceClient(cc *grpc.ClientConn) SCApiServiceClient {
	return &sCApiServiceClient{cc}
}

func (c *sCApiServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateModuleName(ctx context.Context, in *ModuleNameUpdate, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateModuleName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateSensorName(ctx context.Context, in *scweb.SensorNameUpdate, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateSensorName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) SetConfigFlag(ctx context.Context, in *scfunctions.SetFlagRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/SetConfigFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) SetUserConfigFlag(ctx context.Context, in *UserConfigFlagRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/SetUserConfigFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateSensorConfiguration(ctx context.Context, in *SensorConfigUpdateRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateSensorConfiguration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateUserAppRepo(ctx context.Context, in *UpdateUserAppRepoRequest, opts ...grpc.CallOption) (*UserAppVersions, error) {
	out := new(UserAppVersions)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateUserAppRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) UpdateUserAppVersion(ctx context.Context, in *UpdateUserAppVersionRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/UpdateUserAppVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) GetOnlineModules(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*scweb.ModuleList, error) {
	out := new(scweb.ModuleList)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/GetOnlineModules", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCApiServiceClient) GetModuleFlashStatus(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*scupdate.UpdateStatusList, error) {
	out := new(scupdate.UpdateStatusList)
	err := grpc.Invoke(ctx, "/scapi.SCApiService/GetModuleFlashStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCApiService service

type SCApiServiceServer interface {
	// just ping the api
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// update name of a module
	UpdateModuleName(context.Context, *ModuleNameUpdate) (*common.Void, error)
	// update friendly name of sensor
	UpdateSensorName(context.Context, *scweb.SensorNameUpdate) (*common.Void, error)
	// set a config flag
	SetConfigFlag(context.Context, *scfunctions.SetFlagRequest) (*common.Void, error)
	// set a userappconfig flag
	SetUserConfigFlag(context.Context, *UserConfigFlagRequest) (*common.Void, error)
	// update sensor configuration
	UpdateSensorConfiguration(context.Context, *SensorConfigUpdateRequest) (*common.Void, error)
	// update the user app repo
	UpdateUserAppRepo(context.Context, *UpdateUserAppRepoRequest) (*UserAppVersions, error)
	// update the user app version
	UpdateUserAppVersion(context.Context, *UpdateUserAppVersionRequest) (*common.Void, error)
	// get all online modules
	GetOnlineModules(context.Context, *common.Void) (*scweb.ModuleList, error)
	// get a module's current update (flash in progress)
	GetModuleFlashStatus(context.Context, *singingcat.ModuleRef) (*scupdate.UpdateStatusList, error)
}

func RegisterSCApiServiceServer(s *grpc.Server, srv SCApiServiceServer) {
	s.RegisterService(&_SCApiService_serviceDesc, srv)
}

func _SCApiService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateModuleName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModuleNameUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateModuleName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateModuleName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateModuleName(ctx, req.(*ModuleNameUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateSensorName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(scweb.SensorNameUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateSensorName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateSensorName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateSensorName(ctx, req.(*scweb.SensorNameUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_SetConfigFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(scfunctions.SetFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).SetConfigFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/SetConfigFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).SetConfigFlag(ctx, req.(*scfunctions.SetFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_SetUserConfigFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).SetUserConfigFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/SetUserConfigFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).SetUserConfigFlag(ctx, req.(*UserConfigFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateSensorConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorConfigUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateSensorConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateSensorConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateSensorConfiguration(ctx, req.(*SensorConfigUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateUserAppRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAppRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateUserAppRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateUserAppRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateUserAppRepo(ctx, req.(*UpdateUserAppRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_UpdateUserAppVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAppVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).UpdateUserAppVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/UpdateUserAppVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).UpdateUserAppVersion(ctx, req.(*UpdateUserAppVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_GetOnlineModules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).GetOnlineModules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/GetOnlineModules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).GetOnlineModules(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCApiService_GetModuleFlashStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(singingcat.ModuleRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCApiServiceServer).GetModuleFlashStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scapi.SCApiService/GetModuleFlashStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCApiServiceServer).GetModuleFlashStatus(ctx, req.(*singingcat.ModuleRef))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scapi.SCApiService",
	HandlerType: (*SCApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SCApiService_Ping_Handler,
		},
		{
			MethodName: "UpdateModuleName",
			Handler:    _SCApiService_UpdateModuleName_Handler,
		},
		{
			MethodName: "UpdateSensorName",
			Handler:    _SCApiService_UpdateSensorName_Handler,
		},
		{
			MethodName: "SetConfigFlag",
			Handler:    _SCApiService_SetConfigFlag_Handler,
		},
		{
			MethodName: "SetUserConfigFlag",
			Handler:    _SCApiService_SetUserConfigFlag_Handler,
		},
		{
			MethodName: "UpdateSensorConfiguration",
			Handler:    _SCApiService_UpdateSensorConfiguration_Handler,
		},
		{
			MethodName: "UpdateUserAppRepo",
			Handler:    _SCApiService_UpdateUserAppRepo_Handler,
		},
		{
			MethodName: "UpdateUserAppVersion",
			Handler:    _SCApiService_UpdateUserAppVersion_Handler,
		},
		{
			MethodName: "GetOnlineModules",
			Handler:    _SCApiService_GetOnlineModules_Handler,
		},
		{
			MethodName: "GetModuleFlashStatus",
			Handler:    _SCApiService_GetModuleFlashStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/scapi/scapi.proto",
}

func init() { proto.RegisterFile("golang.singingcat.net/apis/scapi/scapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xef, 0x4e, 0xdb, 0x48,
	0x10, 0x57, 0x42, 0x02, 0x61, 0x12, 0x20, 0xd9, 0x03, 0x2e, 0x98, 0x93, 0x2e, 0xf2, 0x87, 0x53,
	0x74, 0x42, 0x46, 0x0a, 0xd2, 0xe9, 0xee, 0xa4, 0xb6, 0x22, 0x50, 0x68, 0x54, 0x0a, 0xc8, 0x16,
	0x7c, 0xdf, 0xd8, 0x13, 0x77, 0x25, 0x67, 0xd7, 0xb5, 0xd7, 0x20, 0x9e, 0xa4, 0xcf, 0xd5, 0x37,
	0xaa, 0xec, 0x5d, 0x3b, 0x76, 0x08, 0x51, 0xbf, 0xd8, 0x33, 0xb3, 0xf3, 0xfb, 0xed, 0xfc, 0xdb,
	0x5d, 0x38, 0xf1, 0x45, 0x40, 0xb9, 0x6f, 0xc5, 0x8c, 0xfb, 0x8c, 0xfb, 0x2e, 0x95, 0x16, 0x47,
	0x79, 0x4a, 0x43, 0x16, 0x9f, 0xc6, 0x2e, 0x0d, 0x99, 0xfa, 0x5a, 0x61, 0x24, 0xa4, 0x20, 0xcd,
	0x4c, 0x31, 0xd6, 0x83, 0x9e, 0x71, 0xaa, 0xbe, 0x0a, 0x64, 0xfc, 0xbb, 0xd6, 0x7b, 0x96, 0x70,
	0x57, 0x32, 0xc1, 0x2b, 0xb2, 0x46, 0x8e, 0xd6, 0x22, 0x93, 0xd0, 0xa3, 0x12, 0x0b, 0x41, 0x63,
	0x2c, 0x8d, 0x71, 0x05, 0x8f, 0xa8, 0xf7, 0x2c, 0x84, 0xb7, 0xc0, 0xb8, 0x62, 0x3e, 0x17, 0x5c,
	0xff, 0xb4, 0xff, 0x3f, 0xeb, 0xf6, 0x28, 0x6c, 0x25, 0x51, 0xe1, 0xcc, 0x3b, 0x68, 0xdf, 0x33,
	0xee, 0xdb, 0xf8, 0x2d, 0xc1, 0x58, 0x92, 0x3e, 0x6c, 0xdd, 0xd3, 0x97, 0x40, 0x50, 0xaf, 0x5f,
	0x1f, 0xd4, 0x86, 0xdb, 0x76, 0xae, 0x92, 0xbf, 0x60, 0xd7, 0x49, 0x9d, 0xb8, 0x8b, 0xb7, 0xc9,
	0x7c, 0x8a, 0x51, 0xbf, 0x36, 0xa8, 0x0d, 0x77, 0xec, 0x25, 0xab, 0xf9, 0x1e, 0x3a, 0x8a, 0x30,
	0x0e, 0x05, 0x8f, 0x91, 0x58, 0xd0, 0xca, 0xe5, 0x0c, 0xd1, 0x1e, 0x11, 0x4b, 0xf5, 0xa2, 0xb4,
	0xaf, 0x5d, 0xf8, 0x98, 0x63, 0xe8, 0x7e, 0x11, 0x5e, 0x12, 0xe0, 0x2d, 0x9d, 0xe3, 0x43, 0x56,
	0x12, 0x62, 0x40, 0x4b, 0xd9, 0x26, 0x97, 0x19, 0x47, 0xc3, 0x2e, 0x74, 0x42, 0xa0, 0x91, 0x7a,
	0xea, 0x70, 0x33, 0xd9, 0xfc, 0x51, 0x83, 0x23, 0x07, 0x79, 0x2c, 0xa2, 0x0b, 0xc1, 0x67, 0xcc,
	0x57, 0x34, 0x79, 0x8e, 0x03, 0x68, 0xab, 0xc5, 0x4b, 0x9c, 0x15, 0x84, 0x65, 0x13, 0x39, 0x84,
	0xcd, 0x8f, 0x9c, 0x4e, 0x03, 0xc5, 0xda, 0xb2, 0xb5, 0x46, 0x4e, 0xa0, 0x67, 0x63, 0x28, 0x22,
	0xc9, 0xb8, 0x3f, 0xe1, 0x12, 0xa3, 0x27, 0x1a, 0xf4, 0x37, 0xb2, 0x32, 0xbc, 0x5e, 0x20, 0x43,
	0xd8, 0xbb, 0x17, 0x41, 0x50, 0xf6, 0x6d, 0x64, 0xbe, 0xcb, 0x66, 0x62, 0x42, 0xe7, 0x2a, 0x62,
	0xc8, 0xbd, 0xe0, 0x25, 0xcb, 0xa5, 0x99, 0xe5, 0x52, 0xb1, 0x99, 0x77, 0x70, 0xac, 0xd2, 0x78,
	0x88, 0x31, 0x3a, 0x0f, 0xc3, 0x47, 0x8c, 0x62, 0x26, 0x78, 0x9e, 0xd4, 0xba, 0x12, 0xed, 0x43,
	0x73, 0x9c, 0xb0, 0x40, 0xb5, 0xb4, 0x61, 0x2b, 0xc5, 0xb4, 0xa1, 0x5f, 0x21, 0x4c, 0x13, 0x28,
	0xb1, 0xdd, 0xe2, 0xf3, 0x23, 0x0d, 0x12, 0xd5, 0xb4, 0x6d, 0xbb, 0xd0, 0x2b, 0x3b, 0xd5, 0xab,
	0x3b, 0x99, 0x13, 0xd8, 0xd2, 0x71, 0x91, 0x5d, 0xa8, 0xeb, 0x50, 0xb6, 0xed, 0xfa, 0xea, 0x3e,
	0xa5, 0x54, 0x0e, 0x06, 0xe8, 0x4a, 0xf4, 0xb2, 0x32, 0xb6, 0xec, 0x42, 0x37, 0xdf, 0xc1, 0x5e,
	0x35, 0xd3, 0x98, 0xfc, 0x0d, 0xad, 0x5c, 0xee, 0xd7, 0x06, 0x1b, 0xc3, 0xf6, 0x68, 0x57, 0x8f,
	0x52, 0x5e, 0x8c, 0x62, 0xdd, 0xa4, 0x70, 0x90, 0xc2, 0x55, 0xff, 0xaf, 0x02, 0xea, 0xff, 0x4a,
	0xa1, 0xba, 0xb0, 0x31, 0x66, 0x32, 0x0b, 0x71, 0xc7, 0x4e, 0x45, 0x5d, 0x08, 0x47, 0x52, 0x89,
	0x79, 0x84, 0xb9, 0x3e, 0xfa, 0xde, 0x84, 0x8e, 0x73, 0x71, 0x1e, 0x32, 0x07, 0xa3, 0x27, 0xe6,
	0x22, 0x39, 0x85, 0x46, 0x3a, 0xd3, 0x64, 0xc5, 0x80, 0x1b, 0xbf, 0x55, 0x6c, 0xfa, 0x6c, 0xfc,
	0x07, 0x5d, 0xd5, 0x82, 0xc5, 0xc4, 0x93, 0xdf, 0xb5, 0xe3, 0xf2, 0x21, 0x30, 0x3a, 0x96, 0x3e,
	0xf0, 0x8f, 0x82, 0x79, 0x0b, 0xa8, 0x9a, 0xdb, 0x02, 0x9a, 0xde, 0x57, 0x0b, 0xd3, 0x4a, 0xe8,
	0xff, 0xb0, 0xe3, 0xa0, 0x5c, 0x54, 0x86, 0x1c, 0x5b, 0xe5, 0x3b, 0xcb, 0x41, 0x59, 0xaa, 0xd7,
	0x12, 0xf6, 0x03, 0xf4, 0x1c, 0x94, 0xd5, 0xca, 0x92, 0x3f, 0x74, 0xc8, 0x2b, 0x0b, 0xbe, 0x44,
	0xf0, 0x19, 0x8e, 0xca, 0x71, 0x2b, 0xf7, 0x24, 0xa2, 0xe9, 0xde, 0x64, 0xa0, 0x89, 0xde, 0x3c,
	0xbb, 0x4b, 0x64, 0x37, 0xd0, 0x7b, 0x35, 0xc2, 0xe4, 0xcf, 0x3c, 0x9a, 0x37, 0x86, 0xdb, 0x38,
	0x2c, 0x85, 0x5b, 0x1e, 0xaf, 0x4f, 0xb0, 0xbf, 0xea, 0x84, 0x11, 0x73, 0x15, 0x61, 0xf5, 0xf8,
	0x2d, 0xc5, 0x75, 0x06, 0xdd, 0x6b, 0x94, 0x77, 0x3c, 0x60, 0x5c, 0xb7, 0x36, 0x26, 0x15, 0x0f,
	0xa3, 0xa7, 0x5b, 0xa5, 0x56, 0x6f, 0x58, 0x2c, 0xc9, 0x04, 0xf6, 0xaf, 0x51, 0x2a, 0xc3, 0x55,
	0x40, 0xe3, 0xaf, 0xe9, 0x94, 0x25, 0x31, 0x39, 0x28, 0xdf, 0xe9, 0x6a, 0xd9, 0xc6, 0x99, 0x61,
	0x58, 0xc5, 0x8b, 0xa1, 0x0b, 0x9a, 0xb9, 0xa7, 0x54, 0x63, 0x13, 0x06, 0x1c, 0x65, 0xf9, 0xed,
	0xd0, 0xaf, 0x43, 0xfa, 0x1c, 0xa8, 0x6c, 0xa6, 0x9b, 0xd9, 0xfd, 0x7f, 0xf6, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x50, 0x9d, 0x02, 0xb5, 0x3a, 0x07, 0x00, 0x00,
}
