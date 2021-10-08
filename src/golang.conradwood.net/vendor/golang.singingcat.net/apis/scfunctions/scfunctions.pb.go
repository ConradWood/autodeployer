// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/scfunctions/scfunctions.proto
// DO NOT EDIT!

/*
Package scfunctions is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/scfunctions/scfunctions.proto

It has these top-level messages:
	WifiRequest
	SetLedRequest
	SetColourRequest
	SetPinRequest
	SetMosfetRequest
	StrobePinRequest
	StrobeRequest
	LedState
	PinState
	BlinkLedRequest
	SetFlagRequest
	PingResponse
*/
package scfunctions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import singingcat "golang.singingcat.net/apis/singingcat"
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

type WifiRequest struct {
	Module   *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	SSID     string                `protobuf:"bytes,2,opt,name=SSID" json:"SSID,omitempty"`
	Password string                `protobuf:"bytes,3,opt,name=Password" json:"Password,omitempty"`
}

func (m *WifiRequest) Reset()                    { *m = WifiRequest{} }
func (m *WifiRequest) String() string            { return proto.CompactTextString(m) }
func (*WifiRequest) ProtoMessage()               {}
func (*WifiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WifiRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *WifiRequest) GetSSID() string {
	if m != nil {
		return m.SSID
	}
	return ""
}

func (m *WifiRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SetLedRequest struct {
	Module     *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	RGB        uint64                `protobuf:"varint,2,opt,name=RGB" json:"RGB,omitempty"`
	Brightness uint32                `protobuf:"varint,3,opt,name=Brightness" json:"Brightness,omitempty"`
	OnPeriod   uint32                `protobuf:"varint,4,opt,name=OnPeriod" json:"OnPeriod,omitempty"`
	OffPeriod  uint32                `protobuf:"varint,5,opt,name=OffPeriod" json:"OffPeriod,omitempty"`
	SyncAccess bool                  `protobuf:"varint,6,opt,name=SyncAccess" json:"SyncAccess,omitempty"`
}

func (m *SetLedRequest) Reset()                    { *m = SetLedRequest{} }
func (m *SetLedRequest) String() string            { return proto.CompactTextString(m) }
func (*SetLedRequest) ProtoMessage()               {}
func (*SetLedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SetLedRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SetLedRequest) GetRGB() uint64 {
	if m != nil {
		return m.RGB
	}
	return 0
}

func (m *SetLedRequest) GetBrightness() uint32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

func (m *SetLedRequest) GetOnPeriod() uint32 {
	if m != nil {
		return m.OnPeriod
	}
	return 0
}

func (m *SetLedRequest) GetOffPeriod() uint32 {
	if m != nil {
		return m.OffPeriod
	}
	return 0
}

func (m *SetLedRequest) GetSyncAccess() bool {
	if m != nil {
		return m.SyncAccess
	}
	return false
}

type SetColourRequest struct {
	Module     *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	RGB        uint64                `protobuf:"varint,2,opt,name=RGB" json:"RGB,omitempty"`
	Colour     string                `protobuf:"bytes,3,opt,name=Colour" json:"Colour,omitempty"`
	Brightness uint32                `protobuf:"varint,4,opt,name=Brightness" json:"Brightness,omitempty"`
	SyncAccess bool                  `protobuf:"varint,6,opt,name=SyncAccess" json:"SyncAccess,omitempty"`
}

func (m *SetColourRequest) Reset()                    { *m = SetColourRequest{} }
func (m *SetColourRequest) String() string            { return proto.CompactTextString(m) }
func (*SetColourRequest) ProtoMessage()               {}
func (*SetColourRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetColourRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SetColourRequest) GetRGB() uint64 {
	if m != nil {
		return m.RGB
	}
	return 0
}

func (m *SetColourRequest) GetColour() string {
	if m != nil {
		return m.Colour
	}
	return ""
}

func (m *SetColourRequest) GetBrightness() uint32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

func (m *SetColourRequest) GetSyncAccess() bool {
	if m != nil {
		return m.SyncAccess
	}
	return false
}

type SetPinRequest struct {
	Module *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Pin    uint32                `protobuf:"varint,2,opt,name=Pin" json:"Pin,omitempty"`
	State  bool                  `protobuf:"varint,3,opt,name=State" json:"State,omitempty"`
}

func (m *SetPinRequest) Reset()                    { *m = SetPinRequest{} }
func (m *SetPinRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPinRequest) ProtoMessage()               {}
func (*SetPinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetPinRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SetPinRequest) GetPin() uint32 {
	if m != nil {
		return m.Pin
	}
	return 0
}

func (m *SetPinRequest) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

type SetMosfetRequest struct {
	Module *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Mosfet uint32                `protobuf:"varint,2,opt,name=Mosfet" json:"Mosfet,omitempty"`
	State  bool                  `protobuf:"varint,3,opt,name=State" json:"State,omitempty"`
}

func (m *SetMosfetRequest) Reset()                    { *m = SetMosfetRequest{} }
func (m *SetMosfetRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMosfetRequest) ProtoMessage()               {}
func (*SetMosfetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SetMosfetRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SetMosfetRequest) GetMosfet() uint32 {
	if m != nil {
		return m.Mosfet
	}
	return 0
}

func (m *SetMosfetRequest) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

type StrobePinRequest struct {
	Module *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Pin    uint32                `protobuf:"varint,2,opt,name=Pin" json:"Pin,omitempty"`
	On     uint32                `protobuf:"varint,3,opt,name=On" json:"On,omitempty"`
	Off    uint32                `protobuf:"varint,4,opt,name=Off" json:"Off,omitempty"`
	Repeat uint32                `protobuf:"varint,5,opt,name=Repeat" json:"Repeat,omitempty"`
}

func (m *StrobePinRequest) Reset()                    { *m = StrobePinRequest{} }
func (m *StrobePinRequest) String() string            { return proto.CompactTextString(m) }
func (*StrobePinRequest) ProtoMessage()               {}
func (*StrobePinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StrobePinRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *StrobePinRequest) GetPin() uint32 {
	if m != nil {
		return m.Pin
	}
	return 0
}

func (m *StrobePinRequest) GetOn() uint32 {
	if m != nil {
		return m.On
	}
	return 0
}

func (m *StrobePinRequest) GetOff() uint32 {
	if m != nil {
		return m.Off
	}
	return 0
}

func (m *StrobePinRequest) GetRepeat() uint32 {
	if m != nil {
		return m.Repeat
	}
	return 0
}

type StrobeRequest struct {
	Module *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Mosfet uint32                `protobuf:"varint,2,opt,name=Mosfet" json:"Mosfet,omitempty"`
	On     uint32                `protobuf:"varint,3,opt,name=On" json:"On,omitempty"`
	Off    uint32                `protobuf:"varint,4,opt,name=Off" json:"Off,omitempty"`
	Repeat uint32                `protobuf:"varint,5,opt,name=Repeat" json:"Repeat,omitempty"`
}

func (m *StrobeRequest) Reset()                    { *m = StrobeRequest{} }
func (m *StrobeRequest) String() string            { return proto.CompactTextString(m) }
func (*StrobeRequest) ProtoMessage()               {}
func (*StrobeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StrobeRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *StrobeRequest) GetMosfet() uint32 {
	if m != nil {
		return m.Mosfet
	}
	return 0
}

func (m *StrobeRequest) GetOn() uint32 {
	if m != nil {
		return m.On
	}
	return 0
}

func (m *StrobeRequest) GetOff() uint32 {
	if m != nil {
		return m.Off
	}
	return 0
}

func (m *StrobeRequest) GetRepeat() uint32 {
	if m != nil {
		return m.Repeat
	}
	return 0
}

type LedState struct {
	ModuleID   uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
	RGB        uint64 `protobuf:"varint,2,opt,name=RGB" json:"RGB,omitempty"`
	Brightness uint32 `protobuf:"varint,3,opt,name=Brightness" json:"Brightness,omitempty"`
	OnPeriod   uint32 `protobuf:"varint,4,opt,name=OnPeriod" json:"OnPeriod,omitempty"`
	OffPeriod  uint32 `protobuf:"varint,5,opt,name=OffPeriod" json:"OffPeriod,omitempty"`
	Occurred   uint32 `protobuf:"varint,6,opt,name=Occurred" json:"Occurred,omitempty"`
}

func (m *LedState) Reset()                    { *m = LedState{} }
func (m *LedState) String() string            { return proto.CompactTextString(m) }
func (*LedState) ProtoMessage()               {}
func (*LedState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LedState) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *LedState) GetRGB() uint64 {
	if m != nil {
		return m.RGB
	}
	return 0
}

func (m *LedState) GetBrightness() uint32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

func (m *LedState) GetOnPeriod() uint32 {
	if m != nil {
		return m.OnPeriod
	}
	return 0
}

func (m *LedState) GetOffPeriod() uint32 {
	if m != nil {
		return m.OffPeriod
	}
	return 0
}

func (m *LedState) GetOccurred() uint32 {
	if m != nil {
		return m.Occurred
	}
	return 0
}

type PinState struct {
	ID       uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	ModuleID uint64 `protobuf:"varint,2,opt,name=ModuleID" json:"ModuleID,omitempty"`
	Pin      uint32 `protobuf:"varint,3,opt,name=Pin" json:"Pin,omitempty"`
	State    bool   `protobuf:"varint,4,opt,name=State" json:"State,omitempty"`
	Occurred uint32 `protobuf:"varint,5,opt,name=Occurred" json:"Occurred,omitempty"`
}

func (m *PinState) Reset()                    { *m = PinState{} }
func (m *PinState) String() string            { return proto.CompactTextString(m) }
func (*PinState) ProtoMessage()               {}
func (*PinState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PinState) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PinState) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func (m *PinState) GetPin() uint32 {
	if m != nil {
		return m.Pin
	}
	return 0
}

func (m *PinState) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func (m *PinState) GetOccurred() uint32 {
	if m != nil {
		return m.Occurred
	}
	return 0
}

type BlinkLedRequest struct {
	Module     *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	RGB        uint64                `protobuf:"varint,2,opt,name=RGB" json:"RGB,omitempty"`
	Colour     string                `protobuf:"bytes,3,opt,name=Colour" json:"Colour,omitempty"`
	Brightness uint32                `protobuf:"varint,4,opt,name=Brightness" json:"Brightness,omitempty"`
	On         uint32                `protobuf:"varint,5,opt,name=On" json:"On,omitempty"`
	Off        uint32                `protobuf:"varint,6,opt,name=Off" json:"Off,omitempty"`
	DurationMS uint32                `protobuf:"varint,7,opt,name=DurationMS" json:"DurationMS,omitempty"`
}

func (m *BlinkLedRequest) Reset()                    { *m = BlinkLedRequest{} }
func (m *BlinkLedRequest) String() string            { return proto.CompactTextString(m) }
func (*BlinkLedRequest) ProtoMessage()               {}
func (*BlinkLedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BlinkLedRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *BlinkLedRequest) GetRGB() uint64 {
	if m != nil {
		return m.RGB
	}
	return 0
}

func (m *BlinkLedRequest) GetColour() string {
	if m != nil {
		return m.Colour
	}
	return ""
}

func (m *BlinkLedRequest) GetBrightness() uint32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

func (m *BlinkLedRequest) GetOn() uint32 {
	if m != nil {
		return m.On
	}
	return 0
}

func (m *BlinkLedRequest) GetOff() uint32 {
	if m != nil {
		return m.Off
	}
	return 0
}

func (m *BlinkLedRequest) GetDurationMS() uint32 {
	if m != nil {
		return m.DurationMS
	}
	return 0
}

type SetFlagRequest struct {
	Module   *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Flag     uint32                `protobuf:"varint,2,opt,name=Flag" json:"Flag,omitempty"`
	NewState bool                  `protobuf:"varint,3,opt,name=NewState" json:"NewState,omitempty"`
}

func (m *SetFlagRequest) Reset()                    { *m = SetFlagRequest{} }
func (m *SetFlagRequest) String() string            { return proto.CompactTextString(m) }
func (*SetFlagRequest) ProtoMessage()               {}
func (*SetFlagRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetFlagRequest) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SetFlagRequest) GetFlag() uint32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SetFlagRequest) GetNewState() bool {
	if m != nil {
		return m.NewState
	}
	return false
}

type PingResponse struct {
	Module       *singingcat.ModuleRef `protobuf:"bytes,1,opt,name=Module" json:"Module,omitempty"`
	Milliseconds uint64                `protobuf:"varint,2,opt,name=Milliseconds" json:"Milliseconds,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PingResponse) GetModule() *singingcat.ModuleRef {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *PingResponse) GetMilliseconds() uint64 {
	if m != nil {
		return m.Milliseconds
	}
	return 0
}

func init() {
	proto.RegisterType((*WifiRequest)(nil), "scfunctions.WifiRequest")
	proto.RegisterType((*SetLedRequest)(nil), "scfunctions.SetLedRequest")
	proto.RegisterType((*SetColourRequest)(nil), "scfunctions.SetColourRequest")
	proto.RegisterType((*SetPinRequest)(nil), "scfunctions.SetPinRequest")
	proto.RegisterType((*SetMosfetRequest)(nil), "scfunctions.SetMosfetRequest")
	proto.RegisterType((*StrobePinRequest)(nil), "scfunctions.StrobePinRequest")
	proto.RegisterType((*StrobeRequest)(nil), "scfunctions.StrobeRequest")
	proto.RegisterType((*LedState)(nil), "scfunctions.LedState")
	proto.RegisterType((*PinState)(nil), "scfunctions.PinState")
	proto.RegisterType((*BlinkLedRequest)(nil), "scfunctions.BlinkLedRequest")
	proto.RegisterType((*SetFlagRequest)(nil), "scfunctions.SetFlagRequest")
	proto.RegisterType((*PingResponse)(nil), "scfunctions.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SCFunctionsServer service

type SCFunctionsServerClient interface {
	// set a pin to a certain state
	SetPin(ctx context.Context, in *SetPinRequest, opts ...grpc.CallOption) (*PinState, error)
	// set a solid colour on a module, returns previous colour
	SetColour(ctx context.Context, in *SetColourRequest, opts ...grpc.CallOption) (*LedState, error)
	// get current colour of module
	GetColour(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*LedState, error)
	// set the state of one of the mosfets on a module, returns previous state
	SetMosfet(ctx context.Context, in *SetMosfetRequest, opts ...grpc.CallOption) (*PinState, error)
	// blinky
	StrobeMosfet(ctx context.Context, in *StrobeRequest, opts ...grpc.CallOption) (*common.Void, error)
	// blinky pin
	StrobePin(ctx context.Context, in *StrobePinRequest, opts ...grpc.CallOption) (*common.Void, error)
	// blink LED (and return to previous state)
	BlinkLed(ctx context.Context, in *BlinkLedRequest, opts ...grpc.CallOption) (*common.Void, error)
	// reset a module
	Reset(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*common.Void, error)
	// set a single config flag
	SetConfigFlag(ctx context.Context, in *SetFlagRequest, opts ...grpc.CallOption) (*common.Void, error)
	// resend pin/mosfet/led state to module
	RestoreModule(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*common.Void, error)
	// ping and return time it took to ping
	Ping(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*PingResponse, error)
	// set wifi SSID & Password
	SetWifi(ctx context.Context, in *WifiRequest, opts ...grpc.CallOption) (*common.Void, error)
}

type sCFunctionsServerClient struct {
	cc *grpc.ClientConn
}

func NewSCFunctionsServerClient(cc *grpc.ClientConn) SCFunctionsServerClient {
	return &sCFunctionsServerClient{cc}
}

func (c *sCFunctionsServerClient) SetPin(ctx context.Context, in *SetPinRequest, opts ...grpc.CallOption) (*PinState, error) {
	out := new(PinState)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/SetPin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) SetColour(ctx context.Context, in *SetColourRequest, opts ...grpc.CallOption) (*LedState, error) {
	out := new(LedState)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/SetColour", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) GetColour(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*LedState, error) {
	out := new(LedState)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/GetColour", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) SetMosfet(ctx context.Context, in *SetMosfetRequest, opts ...grpc.CallOption) (*PinState, error) {
	out := new(PinState)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/SetMosfet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) StrobeMosfet(ctx context.Context, in *StrobeRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/StrobeMosfet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) StrobePin(ctx context.Context, in *StrobePinRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/StrobePin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) BlinkLed(ctx context.Context, in *BlinkLedRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/BlinkLed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) Reset(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/Reset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) SetConfigFlag(ctx context.Context, in *SetFlagRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/SetConfigFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) RestoreModule(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/RestoreModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) Ping(ctx context.Context, in *singingcat.ModuleRef, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCFunctionsServerClient) SetWifi(ctx context.Context, in *WifiRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/scfunctions.SCFunctionsServer/SetWifi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCFunctionsServer service

type SCFunctionsServerServer interface {
	// set a pin to a certain state
	SetPin(context.Context, *SetPinRequest) (*PinState, error)
	// set a solid colour on a module, returns previous colour
	SetColour(context.Context, *SetColourRequest) (*LedState, error)
	// get current colour of module
	GetColour(context.Context, *singingcat.ModuleRef) (*LedState, error)
	// set the state of one of the mosfets on a module, returns previous state
	SetMosfet(context.Context, *SetMosfetRequest) (*PinState, error)
	// blinky
	StrobeMosfet(context.Context, *StrobeRequest) (*common.Void, error)
	// blinky pin
	StrobePin(context.Context, *StrobePinRequest) (*common.Void, error)
	// blink LED (and return to previous state)
	BlinkLed(context.Context, *BlinkLedRequest) (*common.Void, error)
	// reset a module
	Reset(context.Context, *singingcat.ModuleRef) (*common.Void, error)
	// set a single config flag
	SetConfigFlag(context.Context, *SetFlagRequest) (*common.Void, error)
	// resend pin/mosfet/led state to module
	RestoreModule(context.Context, *singingcat.ModuleRef) (*common.Void, error)
	// ping and return time it took to ping
	Ping(context.Context, *singingcat.ModuleRef) (*PingResponse, error)
	// set wifi SSID & Password
	SetWifi(context.Context, *WifiRequest) (*common.Void, error)
}

func RegisterSCFunctionsServerServer(s *grpc.Server, srv SCFunctionsServerServer) {
	s.RegisterService(&_SCFunctionsServer_serviceDesc, srv)
}

func _SCFunctionsServer_SetPin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).SetPin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/SetPin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).SetPin(ctx, req.(*SetPinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_SetColour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetColourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).SetColour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/SetColour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).SetColour(ctx, req.(*SetColourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_GetColour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(singingcat.ModuleRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).GetColour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/GetColour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).GetColour(ctx, req.(*singingcat.ModuleRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_SetMosfet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMosfetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).SetMosfet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/SetMosfet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).SetMosfet(ctx, req.(*SetMosfetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_StrobeMosfet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrobeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).StrobeMosfet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/StrobeMosfet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).StrobeMosfet(ctx, req.(*StrobeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_StrobePin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrobePinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).StrobePin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/StrobePin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).StrobePin(ctx, req.(*StrobePinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_BlinkLed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlinkLedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).BlinkLed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/BlinkLed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).BlinkLed(ctx, req.(*BlinkLedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(singingcat.ModuleRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).Reset(ctx, req.(*singingcat.ModuleRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_SetConfigFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFlagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).SetConfigFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/SetConfigFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).SetConfigFlag(ctx, req.(*SetFlagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_RestoreModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(singingcat.ModuleRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).RestoreModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/RestoreModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).RestoreModule(ctx, req.(*singingcat.ModuleRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(singingcat.ModuleRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).Ping(ctx, req.(*singingcat.ModuleRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCFunctionsServer_SetWifi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCFunctionsServerServer).SetWifi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scfunctions.SCFunctionsServer/SetWifi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCFunctionsServerServer).SetWifi(ctx, req.(*WifiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCFunctionsServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scfunctions.SCFunctionsServer",
	HandlerType: (*SCFunctionsServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPin",
			Handler:    _SCFunctionsServer_SetPin_Handler,
		},
		{
			MethodName: "SetColour",
			Handler:    _SCFunctionsServer_SetColour_Handler,
		},
		{
			MethodName: "GetColour",
			Handler:    _SCFunctionsServer_GetColour_Handler,
		},
		{
			MethodName: "SetMosfet",
			Handler:    _SCFunctionsServer_SetMosfet_Handler,
		},
		{
			MethodName: "StrobeMosfet",
			Handler:    _SCFunctionsServer_StrobeMosfet_Handler,
		},
		{
			MethodName: "StrobePin",
			Handler:    _SCFunctionsServer_StrobePin_Handler,
		},
		{
			MethodName: "BlinkLed",
			Handler:    _SCFunctionsServer_BlinkLed_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _SCFunctionsServer_Reset_Handler,
		},
		{
			MethodName: "SetConfigFlag",
			Handler:    _SCFunctionsServer_SetConfigFlag_Handler,
		},
		{
			MethodName: "RestoreModule",
			Handler:    _SCFunctionsServer_RestoreModule_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _SCFunctionsServer_Ping_Handler,
		},
		{
			MethodName: "SetWifi",
			Handler:    _SCFunctionsServer_SetWifi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/scfunctions/scfunctions.proto",
}

func init() {
	proto.RegisterFile("golang.singingcat.net/apis/scfunctions/scfunctions.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 782 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x5f, 0x4f, 0xdb, 0x48,
	0x10, 0x97, 0x43, 0x12, 0x92, 0x21, 0xe1, 0x38, 0xeb, 0x40, 0xb9, 0x1c, 0x87, 0x90, 0x1f, 0x4e,
	0x79, 0xb8, 0x0b, 0x3a, 0x5a, 0xa1, 0xb4, 0x7d, 0x22, 0x20, 0x10, 0x12, 0x69, 0xa2, 0xb5, 0xd4,
	0x3e, 0x1b, 0x7b, 0xec, 0xae, 0x6a, 0x76, 0x53, 0xef, 0xa6, 0xa8, 0xfd, 0x0e, 0x95, 0xfa, 0x3d,
	0x2a, 0xf5, 0x83, 0xa0, 0x7e, 0xa8, 0x6a, 0xd7, 0x8b, 0x63, 0x1b, 0xd2, 0x3f, 0x51, 0xd4, 0x27,
	0xef, 0xcc, 0xee, 0xcc, 0xef, 0x37, 0xb3, 0x33, 0xe3, 0x85, 0x41, 0xc4, 0x63, 0x8f, 0x45, 0x7d,
	0x41, 0x59, 0x44, 0x59, 0xe4, 0x7b, 0xb2, 0xcf, 0x50, 0x1e, 0x78, 0x53, 0x2a, 0x0e, 0x84, 0x1f,
	0xce, 0x98, 0x2f, 0x29, 0x67, 0x85, 0x75, 0x7f, 0x9a, 0x70, 0xc9, 0xed, 0x8d, 0x9c, 0xaa, 0x7b,
	0xf4, 0x2d, 0x37, 0x99, 0x2e, 0xb7, 0x4c, 0x9d, 0x74, 0xfb, 0xc6, 0xce, 0xe7, 0x2c, 0xf1, 0x82,
	0x1b, 0xce, 0x83, 0xb9, 0x9d, 0xcf, 0xaf, 0xaf, 0x39, 0x33, 0x9f, 0xf4, 0xbc, 0x13, 0xc3, 0xc6,
	0x4b, 0x1a, 0x52, 0x82, 0x6f, 0x66, 0x28, 0xa4, 0xfd, 0x1f, 0xd4, 0x47, 0x3c, 0x98, 0xc5, 0xd8,
	0xb1, 0xf6, 0xad, 0xde, 0xc6, 0xe1, 0x76, 0x9e, 0x40, 0xba, 0x43, 0x30, 0x24, 0xe6, 0x90, 0x6d,
	0x43, 0xd5, 0x75, 0x2f, 0x4e, 0x3b, 0x95, 0x7d, 0xab, 0xd7, 0x24, 0x7a, 0x6d, 0x77, 0xa1, 0x31,
	0xf1, 0x84, 0xb8, 0xe1, 0x49, 0xd0, 0x59, 0xd3, 0xfa, 0x4c, 0x76, 0x6e, 0x2d, 0x68, 0xbb, 0x28,
	0x2f, 0x31, 0x58, 0x12, 0x70, 0x0b, 0xd6, 0xc8, 0xf9, 0x50, 0xe3, 0x55, 0x89, 0x5a, 0xda, 0x7b,
	0x00, 0xc3, 0x84, 0x46, 0xaf, 0x24, 0x43, 0x21, 0x34, 0x60, 0x9b, 0xe4, 0x34, 0x8a, 0xce, 0x98,
	0x4d, 0x30, 0xa1, 0x3c, 0xe8, 0x54, 0xf5, 0x6e, 0x26, 0xdb, 0xbb, 0xd0, 0x1c, 0x87, 0xa1, 0xd9,
	0xac, 0xe9, 0xcd, 0xb9, 0x42, 0x79, 0x76, 0xdf, 0x31, 0xff, 0xd8, 0xf7, 0x95, 0xe7, 0xfa, 0xbe,
	0xd5, 0x6b, 0x90, 0x9c, 0xc6, 0xf9, 0x64, 0xc1, 0x96, 0x8b, 0xf2, 0x84, 0xc7, 0x7c, 0x96, 0xac,
	0x2c, 0x9e, 0x1d, 0xa8, 0xa7, 0x1e, 0x4d, 0xf2, 0x8c, 0x54, 0x8a, 0xb3, 0x7a, 0x2f, 0xce, 0xef,
	0xb1, 0x0d, 0x75, 0xe6, 0x27, 0x94, 0x2d, 0xcf, 0x74, 0x42, 0x99, 0x66, 0xda, 0x26, 0x6a, 0x69,
	0xff, 0x01, 0x35, 0x57, 0x7a, 0x12, 0x35, 0xd1, 0x06, 0x49, 0x05, 0x87, 0xeb, 0xa4, 0x8c, 0xb8,
	0x08, 0x51, 0x2e, 0x09, 0xb5, 0xa3, 0x8e, 0x2b, 0x7b, 0x83, 0x66, 0xa4, 0x05, 0x80, 0x1f, 0xd4,
	0x35, 0xc8, 0x84, 0x5f, 0xe1, 0x2a, 0x83, 0xdb, 0x84, 0xca, 0x98, 0x99, 0x72, 0xaa, 0x8c, 0x99,
	0x3a, 0x31, 0x0e, 0x43, 0x93, 0x77, 0xb5, 0x54, 0x2c, 0x09, 0x4e, 0xd1, 0x93, 0xa6, 0x72, 0x8c,
	0xe4, 0x7c, 0x54, 0x35, 0xae, 0xf9, 0xac, 0x38, 0xfc, 0xe5, 0x29, 0x7d, 0xb6, 0xa0, 0x71, 0x89,
	0x81, 0xce, 0x97, 0x6a, 0x88, 0x14, 0xe8, 0xe2, 0x54, 0xf3, 0xa9, 0x92, 0x4c, 0xfe, 0xa5, 0xed,
	0xa5, 0x2c, 0x7d, 0x7f, 0x96, 0x24, 0x18, 0xe8, 0x72, 0x55, 0x96, 0x46, 0x76, 0xde, 0x43, 0x63,
	0x42, 0x59, 0xca, 0x77, 0x13, 0x2a, 0x19, 0xd3, 0x4a, 0x3a, 0x5f, 0x32, 0xfe, 0x95, 0xfb, 0xfc,
	0xd5, 0x3d, 0xae, 0x3d, 0x50, 0xa4, 0xd5, 0x5c, 0xcd, 0x14, 0xb0, 0x6b, 0x25, 0xec, 0x2f, 0x16,
	0xfc, 0x36, 0x8c, 0x29, 0x7b, 0xbd, 0xca, 0x29, 0xb5, 0x6c, 0x57, 0xa7, 0x77, 0x5e, 0x2b, 0xdf,
	0x79, 0x7d, 0x7e, 0xe7, 0x7b, 0x00, 0xa7, 0xb3, 0xc4, 0x53, 0x7f, 0x8d, 0x91, 0xdb, 0x59, 0x4f,
	0x3d, 0xcc, 0x35, 0x0e, 0x87, 0x4d, 0x17, 0xe5, 0x59, 0xec, 0x45, 0xcb, 0xcf, 0x78, 0x65, 0x6d,
	0x8a, 0x51, 0xaf, 0x55, 0xfe, 0x9e, 0xe3, 0x4d, 0xbe, 0x19, 0x33, 0xd9, 0xf1, 0xa0, 0x35, 0xa1,
	0x2c, 0x22, 0x28, 0xa6, 0x9c, 0x09, 0xfc, 0x59, 0x38, 0x07, 0x5a, 0x23, 0x1a, 0xc7, 0x54, 0xa0,
	0xcf, 0x59, 0x20, 0x4c, 0x12, 0x0b, 0xba, 0xc3, 0xdb, 0x1a, 0xfc, 0xee, 0x9e, 0x9c, 0xdd, 0xfd,
	0x2c, 0x5d, 0x4c, 0xde, 0x62, 0x62, 0x3f, 0x83, 0x7a, 0x3a, 0xe1, 0xec, 0x6e, 0x3f, 0xff, 0x77,
	0x2d, 0x8c, 0xbd, 0xee, 0x76, 0x61, 0x2f, 0xab, 0xb2, 0x63, 0x68, 0x66, 0xb3, 0xdc, 0xfe, 0xbb,
	0x6c, 0x5f, 0x98, 0xf1, 0x25, 0x17, 0x59, 0x63, 0x3d, 0x81, 0xe6, 0x79, 0xe6, 0xe2, 0xe1, 0x28,
	0x17, 0x99, 0xa6, 0xe8, 0xa6, 0xcf, 0xef, 0xa1, 0x17, 0x86, 0xe9, 0xa2, 0x00, 0x06, 0xd0, 0x4a,
	0xa7, 0x8e, 0xf1, 0x52, 0xca, 0x41, 0x7e, 0x20, 0x75, 0x5b, 0x7d, 0xf3, 0x06, 0x78, 0xc1, 0x69,
	0x60, 0x0f, 0xa0, 0x99, 0xcd, 0xcf, 0x32, 0x78, 0x69, 0xae, 0x96, 0x2c, 0x8f, 0xa0, 0x71, 0xd7,
	0x29, 0xf6, 0x6e, 0xc1, 0xb0, 0xd4, 0x40, 0x25, 0xbb, 0x7f, 0xa1, 0x46, 0x50, 0xa0, 0x5c, 0x94,
	0xa5, 0xe2, 0xe9, 0xa7, 0xfa, 0xcf, 0x75, 0xc2, 0x59, 0x48, 0x23, 0x5d, 0x7d, 0x7f, 0x95, 0x13,
	0x94, 0xab, 0xee, 0x92, 0xed, 0x63, 0x68, 0x13, 0x14, 0x92, 0x27, 0x68, 0xca, 0xeb, 0x87, 0x10,
	0x07, 0x50, 0x55, 0x25, 0xbc, 0xe8, 0xf0, 0x9f, 0xe5, 0x1b, 0x98, 0x17, 0xfb, 0xff, 0xb0, 0xee,
	0xa2, 0x54, 0x2f, 0x2a, 0xbb, 0x53, 0x38, 0x95, 0x7b, 0x64, 0x15, 0xc1, 0x86, 0x3d, 0xf8, 0x87,
	0xa1, 0xcc, 0x23, 0x99, 0x27, 0x9c, 0x7a, 0xb3, 0xe5, 0x7d, 0x5c, 0xd5, 0xf5, 0x93, 0xed, 0xd1,
	0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xc7, 0x1e, 0xb1, 0x63, 0x0a, 0x00, 0x00,
}
