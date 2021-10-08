// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1beta1/source.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Cloud Security Command Center's (Cloud SCC) finding source. A finding source
// is an entity or a mechanism that can produce a finding. A source is like a
// container of findings that come from the same scanner, logger, monitor, etc.
type Source struct {
	// The relative resource name of this source. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/sources/{source_id}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The source's display name.
	// A source's display name must be unique amongst its siblings, for example,
	// two sources with the same parent can't share the same display name.
	// The display name must have a length between 1 and 64 characters
	// (inclusive).
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the source (max of 1024 characters).
	// Example:
	// "Cloud Security Scanner is a web security scanner for common
	// vulnerabilities in App Engine applications. It can automatically
	// scan and detect four common vulnerabilities, including cross-site-scripting
	// (XSS), Flash injection, mixed content (HTTP in HTTPS), and
	// outdated/insecure libraries."
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_96479ffd36eb78cb, []int{0}
}

func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Source) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Source) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Source)(nil), "google.cloud.securitycenter.v1p1beta1.Source")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1beta1/source.proto", fileDescriptor_96479ffd36eb78cb)
}

var fileDescriptor_96479ffd36eb78cb = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xfe, 0x3f, 0x05, 0x53, 0x57, 0xb3, 0xaa, 0xc5, 0x45, 0x15, 0x0b, 0x75, 0x61,
	0xc2, 0xe8, 0x2e, 0xae, 0x6c, 0x17, 0xee, 0xa4, 0x58, 0x2c, 0x22, 0x05, 0x49, 0xd3, 0x10, 0x02,
	0xd3, 0xdc, 0x21, 0x49, 0x85, 0x5a, 0xfa, 0x42, 0xbe, 0x82, 0x6f, 0xe0, 0xa3, 0xf4, 0x1d, 0x04,
	0x99, 0x24, 0xb5, 0x33, 0x6e, 0x74, 0x97, 0xb9, 0xe7, 0x3b, 0xe7, 0x1e, 0xee, 0xa0, 0x4b, 0x09,
	0x20, 0x73, 0x41, 0x78, 0x0e, 0xcb, 0x39, 0xb1, 0x82, 0x2f, 0x8d, 0x72, 0x2b, 0x2e, 0xb4, 0x13,
	0x86, 0xbc, 0x64, 0x45, 0x36, 0x13, 0x8e, 0x65, 0xc4, 0xc2, 0xd2, 0x70, 0x81, 0x0b, 0x03, 0x0e,
	0xd2, 0x5e, 0xf0, 0x60, 0xef, 0xc1, 0x75, 0x0f, 0xfe, 0xf6, 0x74, 0x8e, 0x63, 0x34, 0x2b, 0x14,
	0x61, 0x5a, 0x83, 0x63, 0x4e, 0x81, 0xb6, 0x21, 0xa4, 0x73, 0x54, 0x51, 0x8d, 0xa8, 0xe6, 0x9f,
	0xbe, 0x27, 0xa8, 0x39, 0xf6, 0x83, 0x34, 0x45, 0xff, 0x35, 0x5b, 0x88, 0x76, 0xd2, 0x4d, 0xfa,
	0x07, 0xf7, 0xfe, 0x9d, 0x9e, 0xa0, 0xc3, 0xb9, 0xb2, 0x45, 0xce, 0x56, 0xcf, 0x5e, 0x6b, 0x78,
	0xad, 0x15, 0x67, 0x77, 0x25, 0xd2, 0x45, 0xad, 0xb9, 0xb0, 0xdc, 0xa8, 0xa2, 0x5c, 0xd9, 0xfe,
	0x17, 0x89, 0xfd, 0x88, 0x3e, 0x6e, 0x6f, 0x1e, 0xd0, 0xd9, 0x8f, 0xee, 0xa1, 0x12, 0x2b, 0x94,
	0xc5, 0x1c, 0x16, 0x24, 0x76, 0xb8, 0x00, 0x23, 0x99, 0x56, 0xaf, 0xa1, 0x3e, 0x59, 0x57, 0x3f,
	0x37, 0xf1, 0x32, 0x96, 0xac, 0xc3, 0x63, 0x33, 0xf8, 0x4c, 0xd0, 0x39, 0x87, 0x05, 0xfe, 0xd3,
	0x91, 0x46, 0xc9, 0xd3, 0x38, 0x82, 0x12, 0x72, 0xa6, 0x25, 0x06, 0x23, 0x89, 0x14, 0xda, 0x5f,
	0x82, 0xec, 0x0b, 0xfd, 0xf2, 0x83, 0xae, 0xeb, 0xc2, 0x5b, 0xa3, 0x77, 0x1b, 0x52, 0x87, 0x7e,
	0xfd, 0x38, 0xaa, 0xc3, 0xb0, 0x7e, 0x92, 0x8d, 0xb2, 0x41, 0x69, 0xfb, 0xd8, 0x71, 0x53, 0xcf,
	0x4d, 0xeb, 0xdc, 0x74, 0xb2, 0x8b, 0xdf, 0x36, 0xfa, 0x81, 0xa3, 0xd4, 0x83, 0x94, 0xd6, 0x49,
	0x4a, 0x4b, 0xd4, 0x47, 0xce, 0x9a, 0xbe, 0xfa, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13,
	0xb9, 0xa5, 0xd6, 0x5a, 0x02, 0x00, 0x00,
}
