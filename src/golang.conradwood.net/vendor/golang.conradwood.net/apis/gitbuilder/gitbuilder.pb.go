// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/gitbuilder/gitbuilder.proto
// DO NOT EDIT!

/*
Package gitbuilder is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/gitbuilder/gitbuilder.proto

It has these top-level messages:
	BuildRequest
	BuildResponse
*/
package gitbuilder

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

type BuildRequest struct {
	GitURL       string   `protobuf:"bytes,1,opt,name=GitURL" json:"GitURL,omitempty"`
	FetchURLS    []string `protobuf:"bytes,2,rep,name=FetchURLS" json:"FetchURLS,omitempty"`
	CommitID     string   `protobuf:"bytes,3,opt,name=CommitID" json:"CommitID,omitempty"`
	BuildNumber  uint64   `protobuf:"varint,4,opt,name=BuildNumber" json:"BuildNumber,omitempty"`
	RepositoryID uint64   `protobuf:"varint,5,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	RepoName     string   `protobuf:"bytes,6,opt,name=RepoName" json:"RepoName,omitempty"`
	ArtefactName string   `protobuf:"bytes,7,opt,name=ArtefactName" json:"ArtefactName,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildRequest) GetGitURL() string {
	if m != nil {
		return m.GitURL
	}
	return ""
}

func (m *BuildRequest) GetFetchURLS() []string {
	if m != nil {
		return m.FetchURLS
	}
	return nil
}

func (m *BuildRequest) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *BuildRequest) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *BuildRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *BuildRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *BuildRequest) GetArtefactName() string {
	if m != nil {
		return m.ArtefactName
	}
	return ""
}

type BuildResponse struct {
	Stdout        []byte `protobuf:"bytes,1,opt,name=Stdout,proto3" json:"Stdout,omitempty"`
	Complete      bool   `protobuf:"varint,2,opt,name=Complete" json:"Complete,omitempty"`
	ResultMessage string `protobuf:"bytes,3,opt,name=ResultMessage" json:"ResultMessage,omitempty"`
	Success       bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
	LogMessage    string `protobuf:"bytes,5,opt,name=LogMessage" json:"LogMessage,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuildResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *BuildResponse) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *BuildResponse) GetResultMessage() string {
	if m != nil {
		return m.ResultMessage
	}
	return ""
}

func (m *BuildResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *BuildResponse) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "gitbuilder.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "gitbuilder.BuildResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GitBuilder service

type GitBuilderClient interface {
	// build something. Note that this RPC might take several minutes to complete
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (GitBuilder_BuildClient, error)
}

type gitBuilderClient struct {
	cc *grpc.ClientConn
}

func NewGitBuilderClient(cc *grpc.ClientConn) GitBuilderClient {
	return &gitBuilderClient{cc}
}

func (c *gitBuilderClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (GitBuilder_BuildClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GitBuilder_serviceDesc.Streams[0], c.cc, "/gitbuilder.GitBuilder/Build", opts...)
	if err != nil {
		return nil, err
	}
	x := &gitBuilderBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GitBuilder_BuildClient interface {
	Recv() (*BuildResponse, error)
	grpc.ClientStream
}

type gitBuilderBuildClient struct {
	grpc.ClientStream
}

func (x *gitBuilderBuildClient) Recv() (*BuildResponse, error) {
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GitBuilder service

type GitBuilderServer interface {
	// build something. Note that this RPC might take several minutes to complete
	Build(*BuildRequest, GitBuilder_BuildServer) error
}

func RegisterGitBuilderServer(s *grpc.Server, srv GitBuilderServer) {
	s.RegisterService(&_GitBuilder_serviceDesc, srv)
}

func _GitBuilder_Build_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GitBuilderServer).Build(m, &gitBuilderBuildServer{stream})
}

type GitBuilder_BuildServer interface {
	Send(*BuildResponse) error
	grpc.ServerStream
}

type gitBuilderBuildServer struct {
	grpc.ServerStream
}

func (x *gitBuilderBuildServer) Send(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GitBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitbuilder.GitBuilder",
	HandlerType: (*GitBuilderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Build",
			Handler:       _GitBuilder_Build_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "golang.conradwood.net/apis/gitbuilder/gitbuilder.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/gitbuilder/gitbuilder.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xff, 0x77, 0x6c, 0x2f, 0x7b, 0x90, 0xb5, 0x48, 0x09, 0xc5, 0x43, 0x4e, 0xa9,
	0x28, 0x78, 0x14, 0xac, 0xc5, 0x52, 0x88, 0x3d, 0x6c, 0xe9, 0x03, 0xa4, 0xc9, 0x18, 0x03, 0x49,
	0x36, 0x66, 0x27, 0x88, 0xcf, 0xe3, 0xbb, 0xf9, 0x1c, 0x92, 0x4d, 0xda, 0x6e, 0x41, 0x3c, 0x65,
	0xbe, 0x6f, 0xf2, 0xed, 0xf0, 0x63, 0x06, 0x1e, 0x22, 0x99, 0xf8, 0x59, 0xe4, 0x06, 0x32, 0x2b,
	0xfc, 0xf0, 0x53, 0xca, 0xd0, 0xcd, 0x90, 0xe6, 0x7e, 0x1e, 0xab, 0x79, 0x14, 0xd3, 0xbe, 0x8c,
	0x93, 0x10, 0x0b, 0xa3, 0x74, 0xf3, 0x42, 0x92, 0x64, 0x70, 0x72, 0x26, 0xee, 0x3f, 0x6f, 0x04,
	0x32, 0x4d, 0x65, 0xd6, 0x7c, 0xea, 0xec, 0xec, 0xc7, 0x82, 0xd1, 0xa2, 0xca, 0x0a, 0xfc, 0x28,
	0x51, 0x11, 0xbb, 0x84, 0xde, 0x2a, 0xa6, 0x9d, 0xf0, 0xb8, 0x65, 0x5b, 0xce, 0x50, 0x34, 0x8a,
	0x5d, 0xc3, 0xf0, 0x05, 0x29, 0x78, 0xdf, 0x09, 0x6f, 0xcb, 0x5b, 0x76, 0xdb, 0x19, 0x8a, 0x93,
	0xc1, 0x26, 0x30, 0x78, 0x96, 0x69, 0x1a, 0xd3, 0x7a, 0xc9, 0xdb, 0x3a, 0x77, 0xd4, 0xcc, 0x86,
	0x0b, 0x3d, 0x61, 0x53, 0xa6, 0x7b, 0x2c, 0x78, 0xc7, 0xb6, 0x9c, 0x8e, 0x30, 0x2d, 0x36, 0x83,
	0x91, 0xc0, 0x5c, 0xaa, 0x98, 0x64, 0xf1, 0xb5, 0x5e, 0xf2, 0xae, 0xfe, 0xe5, 0xcc, 0xab, 0x26,
	0x54, 0x7a, 0xe3, 0xa7, 0xc8, 0x7b, 0xf5, 0x84, 0x83, 0xae, 0xf2, 0x4f, 0x05, 0xe1, 0x9b, 0x1f,
	0x90, 0xee, 0xf7, 0x75, 0xff, 0xcc, 0x9b, 0x7d, 0x5b, 0x30, 0x6e, 0x40, 0x55, 0x2e, 0x33, 0x85,
	0x15, 0xe9, 0x96, 0x42, 0x59, 0x92, 0x26, 0x1d, 0x89, 0x46, 0x35, 0x2c, 0x79, 0x82, 0x84, 0xbc,
	0x65, 0x5b, 0xce, 0x40, 0x1c, 0x35, 0xbb, 0x81, 0xb1, 0x40, 0x55, 0x26, 0xf4, 0x8a, 0x4a, 0xf9,
	0x11, 0x36, 0xb0, 0xe7, 0x26, 0xe3, 0xd0, 0xdf, 0x96, 0x41, 0x80, 0x4a, 0x69, 0xda, 0x81, 0x38,
	0x48, 0x36, 0x05, 0xf0, 0x64, 0x74, 0x08, 0x77, 0x75, 0xd8, 0x70, 0xee, 0x3c, 0x80, 0x55, 0x4c,
	0x8b, 0x7a, 0x99, 0xec, 0x11, 0xba, 0xba, 0x64, 0xdc, 0x35, 0x96, 0x6e, 0xae, 0x6b, 0x72, 0xf5,
	0x47, 0xa7, 0xe6, 0xbb, 0xb5, 0x16, 0x36, 0x4c, 0x33, 0x24, 0xf3, 0x16, 0xaa, 0x3b, 0x30, 0x12,
	0xfb, 0x9e, 0xbe, 0x82, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x07, 0xcb, 0x75, 0x7b,
	0x02, 0x00, 0x00,
}
