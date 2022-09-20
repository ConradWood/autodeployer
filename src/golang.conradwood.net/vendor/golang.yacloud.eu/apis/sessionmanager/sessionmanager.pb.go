// Code generated by protoc-gen-go.
// source: golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto
// DO NOT EDIT!

/*
Package sessionmanager is a generated protocol buffer package.

It is generated from these files:
	golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto

It has these top-level messages:
	PingResponse
	SessionLog
	NewSessionRequest
	SessionResponse
	SessionVerifyResponse
	SessionToken
*/
package sessionmanager

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

type SessionLog struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	UserID       string `protobuf:"bytes,2,opt,name=UserID" json:"UserID,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=Username" json:"Username,omitempty"`
	Useremail    string `protobuf:"bytes,4,opt,name=Useremail" json:"Useremail,omitempty"`
	IP           string `protobuf:"bytes,5,opt,name=IP" json:"IP,omitempty"`
	UserAgent    string `protobuf:"bytes,6,opt,name=UserAgent" json:"UserAgent,omitempty"`
	Created      uint32 `protobuf:"varint,7,opt,name=Created" json:"Created,omitempty"`
	BrowserID    string `protobuf:"bytes,8,opt,name=BrowserID" json:"BrowserID,omitempty"`
	SessionToken string `protobuf:"bytes,9,opt,name=SessionToken" json:"SessionToken,omitempty"`
	LastUsed     uint32 `protobuf:"varint,10,opt,name=LastUsed" json:"LastUsed,omitempty"`
}

func (m *SessionLog) Reset()                    { *m = SessionLog{} }
func (m *SessionLog) String() string            { return proto.CompactTextString(m) }
func (*SessionLog) ProtoMessage()               {}
func (*SessionLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionLog) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SessionLog) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *SessionLog) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SessionLog) GetUseremail() string {
	if m != nil {
		return m.Useremail
	}
	return ""
}

func (m *SessionLog) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *SessionLog) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *SessionLog) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SessionLog) GetBrowserID() string {
	if m != nil {
		return m.BrowserID
	}
	return ""
}

func (m *SessionLog) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *SessionLog) GetLastUsed() uint32 {
	if m != nil {
		return m.LastUsed
	}
	return 0
}

type NewSessionRequest struct {
	IPAddress string `protobuf:"bytes,1,opt,name=IPAddress" json:"IPAddress,omitempty"`
	UserAgent string `protobuf:"bytes,2,opt,name=UserAgent" json:"UserAgent,omitempty"`
	BrowserID string `protobuf:"bytes,3,opt,name=BrowserID" json:"BrowserID,omitempty"`
	UserID    string `protobuf:"bytes,4,opt,name=UserID" json:"UserID,omitempty"`
	Username  string `protobuf:"bytes,5,opt,name=Username" json:"Username,omitempty"`
	Useremail string `protobuf:"bytes,6,opt,name=Useremail" json:"Useremail,omitempty"`
}

func (m *NewSessionRequest) Reset()                    { *m = NewSessionRequest{} }
func (m *NewSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*NewSessionRequest) ProtoMessage()               {}
func (*NewSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewSessionRequest) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *NewSessionRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *NewSessionRequest) GetBrowserID() string {
	if m != nil {
		return m.BrowserID
	}
	return ""
}

func (m *NewSessionRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *NewSessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewSessionRequest) GetUseremail() string {
	if m != nil {
		return m.Useremail
	}
	return ""
}

type SessionResponse struct {
	Token                string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	LastSessionTimestamp uint32 `protobuf:"varint,2,opt,name=LastSessionTimestamp" json:"LastSessionTimestamp,omitempty"`
	NewDevice            bool   `protobuf:"varint,3,opt,name=NewDevice" json:"NewDevice,omitempty"`
}

func (m *SessionResponse) Reset()                    { *m = SessionResponse{} }
func (m *SessionResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()               {}
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SessionResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SessionResponse) GetLastSessionTimestamp() uint32 {
	if m != nil {
		return m.LastSessionTimestamp
	}
	return 0
}

func (m *SessionResponse) GetNewDevice() bool {
	if m != nil {
		return m.NewDevice
	}
	return false
}

type SessionVerifyResponse struct {
	IsValid    bool        `protobuf:"varint,1,opt,name=IsValid" json:"IsValid,omitempty"`
	SessionLog *SessionLog `protobuf:"bytes,2,opt,name=SessionLog" json:"SessionLog,omitempty"`
}

func (m *SessionVerifyResponse) Reset()                    { *m = SessionVerifyResponse{} }
func (m *SessionVerifyResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionVerifyResponse) ProtoMessage()               {}
func (*SessionVerifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SessionVerifyResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *SessionVerifyResponse) GetSessionLog() *SessionLog {
	if m != nil {
		return m.SessionLog
	}
	return nil
}

type SessionToken struct {
	Token string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
}

func (m *SessionToken) Reset()                    { *m = SessionToken{} }
func (m *SessionToken) String() string            { return proto.CompactTextString(m) }
func (*SessionToken) ProtoMessage()               {}
func (*SessionToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SessionToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "sessionmanager.PingResponse")
	proto.RegisterType((*SessionLog)(nil), "sessionmanager.SessionLog")
	proto.RegisterType((*NewSessionRequest)(nil), "sessionmanager.NewSessionRequest")
	proto.RegisterType((*SessionResponse)(nil), "sessionmanager.SessionResponse")
	proto.RegisterType((*SessionVerifyResponse)(nil), "sessionmanager.SessionVerifyResponse")
	proto.RegisterType((*SessionToken)(nil), "sessionmanager.SessionToken")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SessionManager service

type SessionManagerClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
	// create a new session
	NewSession(ctx context.Context, in *NewSessionRequest, opts ...grpc.CallOption) (*SessionResponse, error)
	// verify a session (by token)
	VerifySession(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionVerifyResponse, error)
}

type sessionManagerClient struct {
	cc *grpc.ClientConn
}

func NewSessionManagerClient(cc *grpc.ClientConn) SessionManagerClient {
	return &sessionManagerClient{cc}
}

func (c *sessionManagerClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) NewSession(ctx context.Context, in *NewSessionRequest, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/NewSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) VerifySession(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionVerifyResponse, error) {
	out := new(SessionVerifyResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/VerifySession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionManager service

type SessionManagerServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
	// create a new session
	NewSession(context.Context, *NewSessionRequest) (*SessionResponse, error)
	// verify a session (by token)
	VerifySession(context.Context, *SessionToken) (*SessionVerifyResponse, error)
}

func RegisterSessionManagerServer(s *grpc.Server, srv SessionManagerServer) {
	s.RegisterService(&_SessionManager_serviceDesc, srv)
}

func _SessionManager_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_NewSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).NewSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/NewSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).NewSession(ctx, req.(*NewSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_VerifySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).VerifySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/VerifySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).VerifySession(ctx, req.(*SessionToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sessionmanager.SessionManager",
	HandlerType: (*SessionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SessionManager_Ping_Handler,
		},
		{
			MethodName: "NewSession",
			Handler:    _SessionManager_NewSession_Handler,
		},
		{
			MethodName: "VerifySession",
			Handler:    _SessionManager_VerifySession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto",
}

func init() {
	proto.RegisterFile("golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0xdc, 0x34, 0x3f, 0x43, 0x12, 0xc4, 0xaa, 0xa0, 0xc5, 0x8a, 0x20, 0x58, 0x20, 0x55,
	0x08, 0xb9, 0x52, 0xb8, 0xf5, 0xd6, 0x90, 0x8b, 0xa5, 0x52, 0x45, 0xa6, 0xcd, 0x7d, 0x89, 0x07,
	0xcb, 0x22, 0xde, 0x0d, 0x5e, 0x9b, 0x28, 0x57, 0x1e, 0x84, 0x07, 0xe0, 0x25, 0x78, 0x21, 0x1e,
	0x02, 0xed, 0x8f, 0xed, 0xda, 0x34, 0x9c, 0xbc, 0x33, 0xfb, 0xcd, 0xcf, 0xf7, 0x79, 0x66, 0xe1,
	0x32, 0x16, 0x5b, 0xc6, 0x63, 0xff, 0xc0, 0x36, 0x5b, 0x51, 0x44, 0x3e, 0x16, 0x17, 0x6c, 0x97,
	0xc8, 0x0b, 0x89, 0x52, 0x26, 0x82, 0xa7, 0x8c, 0xb3, 0x18, 0xb3, 0x96, 0xe9, 0xef, 0x32, 0x91,
	0x0b, 0x32, 0x69, 0x7a, 0x5d, 0xdf, 0xe6, 0xda, 0x08, 0x9e, 0xb1, 0x68, 0x2f, 0x44, 0xe4, 0x73,
	0xcc, 0x4d, 0xbe, 0x8d, 0x48, 0x53, 0xc1, 0xed, 0xc7, 0xc4, 0x7b, 0x6f, 0x61, 0xb4, 0x4a, 0x78,
	0x1c, 0xa2, 0xdc, 0x09, 0x2e, 0x91, 0xb8, 0x30, 0x28, 0xcf, 0xb4, 0x33, 0xeb, 0x9c, 0x0f, 0xc3,
	0xca, 0xf6, 0x7e, 0x3a, 0x00, 0x9f, 0x4c, 0xb9, 0x6b, 0x11, 0x93, 0x09, 0x38, 0xc1, 0x52, 0x83,
	0xba, 0xa1, 0x13, 0x2c, 0xc9, 0x33, 0xe8, 0xdd, 0x49, 0xcc, 0x82, 0x25, 0x75, 0x74, 0xa0, 0xb5,
	0x54, 0x4a, 0x75, 0xe2, 0x2c, 0x45, 0x7a, 0x62, 0x52, 0x96, 0x36, 0x99, 0xc2, 0x50, 0x9d, 0x31,
	0x65, 0xc9, 0x96, 0x76, 0xf5, 0x65, 0xed, 0xd0, 0x15, 0x56, 0xf4, 0x54, 0xbb, 0x9d, 0x60, 0x55,
	0xa2, 0xaf, 0x62, 0xe4, 0x39, 0xed, 0xd5, 0x68, 0xed, 0x20, 0x14, 0xfa, 0x1f, 0x32, 0x64, 0x39,
	0x46, 0xb4, 0x3f, 0xeb, 0x9c, 0x8f, 0xc3, 0xd2, 0x54, 0x71, 0x8b, 0x4c, 0xec, 0x4d, 0x73, 0x03,
	0x13, 0x57, 0x39, 0xc8, 0x3b, 0x18, 0x59, 0x56, 0xb7, 0xe2, 0x2b, 0x72, 0x3a, 0x54, 0x80, 0xc5,
	0xe0, 0xd7, 0x8f, 0xe7, 0xdd, 0x3c, 0x2b, 0x30, 0x6c, 0xdc, 0x2a, 0x36, 0xd7, 0x4c, 0xe6, 0x77,
	0x12, 0x23, 0x0a, 0xba, 0x4c, 0x65, 0x7b, 0xbf, 0x3b, 0xf0, 0xe4, 0x06, 0xf7, 0x16, 0x1f, 0xe2,
	0xb7, 0x02, 0x65, 0xae, 0xaa, 0x07, 0xab, 0xab, 0x28, 0xca, 0x50, 0x4a, 0xab, 0x69, 0xed, 0x68,
	0x72, 0x72, 0xda, 0x9c, 0x1a, 0x9d, 0x9f, 0xb4, 0x3b, 0xaf, 0x15, 0xef, 0x1e, 0x55, 0xfc, 0xf4,
	0x7f, 0x8a, 0xf7, 0x5a, 0x8a, 0x7b, 0x07, 0x78, 0x5c, 0x75, 0x6f, 0x27, 0xe2, 0x0c, 0x4e, 0x8d,
	0x2e, 0xa6, 0x75, 0x63, 0x90, 0x39, 0x9c, 0x29, 0xda, 0xa5, 0x34, 0x49, 0x8a, 0x32, 0x67, 0xe9,
	0x4e, 0x33, 0x18, 0x87, 0x0f, 0xde, 0xa9, 0xd2, 0x37, 0xb8, 0x5f, 0xe2, 0xf7, 0x64, 0x63, 0x26,
	0x61, 0x10, 0xd6, 0x0e, 0x2f, 0x85, 0xa7, 0x36, 0x62, 0x8d, 0x59, 0xf2, 0xe5, 0x50, 0x35, 0x40,
	0xa1, 0x1f, 0xc8, 0x35, 0xdb, 0x26, 0x91, 0x6e, 0x61, 0x10, 0x96, 0x26, 0xb9, 0xbc, 0x3f, 0x8f,
	0xba, 0xf4, 0xa3, 0xb9, 0xeb, 0xb7, 0xf6, 0xa4, 0x46, 0x84, 0xf7, 0xd0, 0xde, 0xeb, 0xe6, 0x5f,
	0x7f, 0x98, 0xe6, 0xfc, 0x4f, 0x07, 0x26, 0x16, 0xf6, 0xd1, 0xe4, 0x23, 0x73, 0xe8, 0xaa, 0x8d,
	0x21, 0x23, 0xdf, 0x2e, 0xd2, 0x5a, 0x24, 0x91, 0x3b, 0x6d, 0x97, 0x6d, 0x6c, 0xd5, 0x0a, 0xa0,
	0x9e, 0x0b, 0xf2, 0xaa, 0x8d, 0xfd, 0x67, 0x66, 0xdc, 0x97, 0x47, 0x58, 0x54, 0x19, 0x6f, 0x61,
	0x6c, 0x64, 0x2a, 0x93, 0x4e, 0x8f, 0x44, 0x68, 0x1e, 0xee, 0x9b, 0x23, 0xb7, 0x4d, 0xa9, 0x17,
	0x33, 0x78, 0x81, 0x45, 0xf5, 0x0e, 0xa9, 0x47, 0xa3, 0x15, 0xf7, 0xb9, 0xa7, 0x9f, 0x8d, 0xf7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xac, 0xeb, 0x8b, 0x62, 0xb4, 0x04, 0x00, 0x00,
}
