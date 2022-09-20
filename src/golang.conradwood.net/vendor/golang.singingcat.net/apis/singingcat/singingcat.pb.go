// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/singingcat/singingcat.proto
// DO NOT EDIT!

/*
Package singingcat is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/singingcat/singingcat.proto

It has these top-level messages:
	Void
	Command
	Status
	ModuleRef
	Com2
	Arg
*/
package singingcat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SensorType int32

const (
	SensorType_SENSORTYPE_UNDEFINED   SensorType = 0
	SensorType_SENSORTYPE_TEMPERATURE SensorType = 1
)

var SensorType_name = map[int32]string{
	0: "SENSORTYPE_UNDEFINED",
	1: "SENSORTYPE_TEMPERATURE",
}
var SensorType_value = map[string]int32{
	"SENSORTYPE_UNDEFINED":   0,
	"SENSORTYPE_TEMPERATURE": 1,
}

func (x SensorType) String() string {
	return proto.EnumName(SensorType_name, int32(x))
}
func (SensorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ModFlags int32

const (
	ModFlags_EXTERNAL_LED          ModFlags = 0
	ModFlags_DISABLE_RADIO         ModFlags = 1
	ModFlags_TEMP_SENSOR_212       ModFlags = 2
	ModFlags_RADIO_REMAPPED        ModFlags = 3
	ModFlags_TEMP_SENSOR_211       ModFlags = 4
	ModFlags_TEMP_SENSOR_207       ModFlags = 5
	ModFlags_SENSORS_DISABLED      ModFlags = 6
	ModFlags_LED_POWER_INVERT      ModFlags = 7
	ModFlags_ROUTING_ENABLED       ModFlags = 8
	ModFlags_POWER_SAVE            ModFlags = 9
	ModFlags_WIFI_UDP              ModFlags = 10
	ModFlags_START_USERAPP_ON_BOOT ModFlags = 11
	ModFlags_WHITE_MOSFET0         ModFlags = 12
	ModFlags_WHITE_MOSFET1         ModFlags = 13
	ModFlags_ESP_COMSYNC           ModFlags = 14
)

var ModFlags_name = map[int32]string{
	0:  "EXTERNAL_LED",
	1:  "DISABLE_RADIO",
	2:  "TEMP_SENSOR_212",
	3:  "RADIO_REMAPPED",
	4:  "TEMP_SENSOR_211",
	5:  "TEMP_SENSOR_207",
	6:  "SENSORS_DISABLED",
	7:  "LED_POWER_INVERT",
	8:  "ROUTING_ENABLED",
	9:  "POWER_SAVE",
	10: "WIFI_UDP",
	11: "START_USERAPP_ON_BOOT",
	12: "WHITE_MOSFET0",
	13: "WHITE_MOSFET1",
	14: "ESP_COMSYNC",
}
var ModFlags_value = map[string]int32{
	"EXTERNAL_LED":          0,
	"DISABLE_RADIO":         1,
	"TEMP_SENSOR_212":       2,
	"RADIO_REMAPPED":        3,
	"TEMP_SENSOR_211":       4,
	"TEMP_SENSOR_207":       5,
	"SENSORS_DISABLED":      6,
	"LED_POWER_INVERT":      7,
	"ROUTING_ENABLED":       8,
	"POWER_SAVE":            9,
	"WIFI_UDP":              10,
	"START_USERAPP_ON_BOOT": 11,
	"WHITE_MOSFET0":         12,
	"WHITE_MOSFET1":         13,
	"ESP_COMSYNC":           14,
}

func (x ModFlags) String() string {
	return proto.EnumName(ModFlags_name, int32(x))
}
func (ModFlags) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ComType int32

const (
	ComType_UNDEFINED             ComType = 0
	ComType_NOOP                  ComType = 1
	ComType_LOG                   ComType = 2
	ComType_ARE_YOU_A_SINGINGCAT  ComType = 3
	ComType_SET_WIRELESS          ComType = 4
	ComType_SET_PIN               ComType = 5
	ComType_RESET                 ComType = 6
	ComType_WIFIRESET             ComType = 7
	ComType_RADIOPING             ComType = 8
	ComType_RADIOPINGLOOP         ComType = 9
	ComType_FREEZE                ComType = 10
	ComType_LIST_MODULES          ComType = 11
	ComType_ROUTE_REQUEST         ComType = 12
	ComType_STROBE                ComType = 13
	ComType_GETPUBKEY             ComType = 14
	ComType_SETCLOUDTOKEN         ComType = 15
	ComType_SET_LOGGING           ComType = 16
	ComType_SET_DAC               ComType = 17
	ComType_SET_SERIAL_PORT       ComType = 18
	ComType_FLASH_APP             ComType = 19
	ComType_MICROSTROBE           ComType = 20
	ComType_SET_SERVER_NAME       ComType = 21
	ComType_SENSOR_VALUE          ComType = 22
	ComType_ANNOUNCE              ComType = 23
	ComType_FACTORY_DEFAULT       ComType = 24
	ComType_SOFTWARE_INFO         ComType = 25
	ComType_BUTTON_PRESSED        ComType = 26
	ComType_STREAMSETUP           ComType = 27
	ComType_STREAMDATA            ComType = 28
	ComType_STREAMCONTROL         ComType = 29
	ComType_BLINKLED              ComType = 30
	ComType_SENSOR_REQUEST        ComType = 31
	ComType_USER                  ComType = 32
	ComType_RADIO_GET_CONFIG      ComType = 33
	ComType_RADIO_SET_CONFIG      ComType = 34
	ComType_STARTAPP              ComType = 35
	ComType_INSTALLED_SOFTWARE    ComType = 36
	ComType_POWERUP               ComType = 37
	ComType_GETCLOUDTOKEN         ComType = 38
	ComType_REQUEST_RADIO_FORWARD ComType = 39
	ComType_RADIO_FORWARD         ComType = 40
	ComType_SET_CONFIG_FLAG       ComType = 41
	ComType_SET_CONFIG_FLAGS      ComType = 42
	ComType_SET_SENSOR_SERVER     ComType = 43
	ComType_WIFI_INFO             ComType = 44
	ComType_GET_CONFIG_FLAGS      ComType = 45
	ComType_WIFI_SCAN             ComType = 46
	ComType_ROUTE                 ComType = 47
	ComType_SENSOR_LIST           ComType = 48
	ComType_SENSOR_CONFIG         ComType = 49
	ComType_SENSOR_REPORT         ComType = 50
	ComType_ROUTE_UPDATE          ComType = 51
	ComType_SETCONFIG             ComType = 52
	ComType_USERAPP_INFO          ComType = 53
	ComType_USERAPP_CONTROL       ComType = 54
	ComType_BT_BLE_WRITE_CHAR     ComType = 55
	ComType_METRIC                ComType = 56
	ComType_BT_BLE_READ_CHAR      ComType = 57
	ComType_BT_PEERS              ComType = 58
	ComType_CLOCK_SYNC            ComType = 59
	ComType_POWERINFO             ComType = 60
)

var ComType_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "NOOP",
	2:  "LOG",
	3:  "ARE_YOU_A_SINGINGCAT",
	4:  "SET_WIRELESS",
	5:  "SET_PIN",
	6:  "RESET",
	7:  "WIFIRESET",
	8:  "RADIOPING",
	9:  "RADIOPINGLOOP",
	10: "FREEZE",
	11: "LIST_MODULES",
	12: "ROUTE_REQUEST",
	13: "STROBE",
	14: "GETPUBKEY",
	15: "SETCLOUDTOKEN",
	16: "SET_LOGGING",
	17: "SET_DAC",
	18: "SET_SERIAL_PORT",
	19: "FLASH_APP",
	20: "MICROSTROBE",
	21: "SET_SERVER_NAME",
	22: "SENSOR_VALUE",
	23: "ANNOUNCE",
	24: "FACTORY_DEFAULT",
	25: "SOFTWARE_INFO",
	26: "BUTTON_PRESSED",
	27: "STREAMSETUP",
	28: "STREAMDATA",
	29: "STREAMCONTROL",
	30: "BLINKLED",
	31: "SENSOR_REQUEST",
	32: "USER",
	33: "RADIO_GET_CONFIG",
	34: "RADIO_SET_CONFIG",
	35: "STARTAPP",
	36: "INSTALLED_SOFTWARE",
	37: "POWERUP",
	38: "GETCLOUDTOKEN",
	39: "REQUEST_RADIO_FORWARD",
	40: "RADIO_FORWARD",
	41: "SET_CONFIG_FLAG",
	42: "SET_CONFIG_FLAGS",
	43: "SET_SENSOR_SERVER",
	44: "WIFI_INFO",
	45: "GET_CONFIG_FLAGS",
	46: "WIFI_SCAN",
	47: "ROUTE",
	48: "SENSOR_LIST",
	49: "SENSOR_CONFIG",
	50: "SENSOR_REPORT",
	51: "ROUTE_UPDATE",
	52: "SETCONFIG",
	53: "USERAPP_INFO",
	54: "USERAPP_CONTROL",
	55: "BT_BLE_WRITE_CHAR",
	56: "METRIC",
	57: "BT_BLE_READ_CHAR",
	58: "BT_PEERS",
	59: "CLOCK_SYNC",
	60: "POWERINFO",
}
var ComType_value = map[string]int32{
	"UNDEFINED": 0,
	"NOOP":      1,
	"LOG":       2,
	"ARE_YOU_A_SINGINGCAT":  3,
	"SET_WIRELESS":          4,
	"SET_PIN":               5,
	"RESET":                 6,
	"WIFIRESET":             7,
	"RADIOPING":             8,
	"RADIOPINGLOOP":         9,
	"FREEZE":                10,
	"LIST_MODULES":          11,
	"ROUTE_REQUEST":         12,
	"STROBE":                13,
	"GETPUBKEY":             14,
	"SETCLOUDTOKEN":         15,
	"SET_LOGGING":           16,
	"SET_DAC":               17,
	"SET_SERIAL_PORT":       18,
	"FLASH_APP":             19,
	"MICROSTROBE":           20,
	"SET_SERVER_NAME":       21,
	"SENSOR_VALUE":          22,
	"ANNOUNCE":              23,
	"FACTORY_DEFAULT":       24,
	"SOFTWARE_INFO":         25,
	"BUTTON_PRESSED":        26,
	"STREAMSETUP":           27,
	"STREAMDATA":            28,
	"STREAMCONTROL":         29,
	"BLINKLED":              30,
	"SENSOR_REQUEST":        31,
	"USER":                  32,
	"RADIO_GET_CONFIG":      33,
	"RADIO_SET_CONFIG":      34,
	"STARTAPP":              35,
	"INSTALLED_SOFTWARE":    36,
	"POWERUP":               37,
	"GETCLOUDTOKEN":         38,
	"REQUEST_RADIO_FORWARD": 39,
	"RADIO_FORWARD":         40,
	"SET_CONFIG_FLAG":       41,
	"SET_CONFIG_FLAGS":      42,
	"SET_SENSOR_SERVER":     43,
	"WIFI_INFO":             44,
	"GET_CONFIG_FLAGS":      45,
	"WIFI_SCAN":             46,
	"ROUTE":                 47,
	"SENSOR_LIST":           48,
	"SENSOR_CONFIG":         49,
	"SENSOR_REPORT":         50,
	"ROUTE_UPDATE":          51,
	"SETCONFIG":             52,
	"USERAPP_INFO":          53,
	"USERAPP_CONTROL":       54,
	"BT_BLE_WRITE_CHAR":     55,
	"METRIC":                56,
	"BT_BLE_READ_CHAR":      57,
	"BT_PEERS":              58,
	"CLOCK_SYNC":            59,
	"POWERINFO":             60,
}

func (x ComType) String() string {
	return proto.EnumName(ComType_name, int32(x))
}
func (ComType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ROUTE_COM_TYPE int32

const (
	ROUTE_COM_TYPE_ROUTE_NONE  ROUTE_COM_TYPE = 0
	ROUTE_COM_TYPE_ROUTE_LIST  ROUTE_COM_TYPE = 1
	ROUTE_COM_TYPE_ROUTE_ADD   ROUTE_COM_TYPE = 2
	ROUTE_COM_TYPE_ROUTE_DEL   ROUTE_COM_TYPE = 3
	ROUTE_COM_TYPE_ROUTE_CLEAR ROUTE_COM_TYPE = 4
)

var ROUTE_COM_TYPE_name = map[int32]string{
	0: "ROUTE_NONE",
	1: "ROUTE_LIST",
	2: "ROUTE_ADD",
	3: "ROUTE_DEL",
	4: "ROUTE_CLEAR",
}
var ROUTE_COM_TYPE_value = map[string]int32{
	"ROUTE_NONE":  0,
	"ROUTE_LIST":  1,
	"ROUTE_ADD":   2,
	"ROUTE_DEL":   3,
	"ROUTE_CLEAR": 4,
}

func (x ROUTE_COM_TYPE) String() string {
	return proto.EnumName(ROUTE_COM_TYPE_name, int32(x))
}
func (ROUTE_COM_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ErrorEnum int32

const (
	ErrorEnum_UNSPECIFIED              ErrorEnum = 0
	ErrorEnum_ACCOUNT_UPGRADE_REQUIRED ErrorEnum = 1
)

var ErrorEnum_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "ACCOUNT_UPGRADE_REQUIRED",
}
var ErrorEnum_value = map[string]int32{
	"UNSPECIFIED":              0,
	"ACCOUNT_UPGRADE_REQUIRED": 1,
}

func (x ErrorEnum) String() string {
	return proto.EnumName(ErrorEnum_name, int32(x))
}
func (ErrorEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// matches 'addressing.h' in firmware
type Device int32

const (
	Device_DEVICE_UNSPECIFIED Device = 0
	Device_WIFI               Device = 1
	Device_RADIO              Device = 2
	Device_SERIAL             Device = 3
	Device_USB                Device = 4
	Device_LORA               Device = 5
	Device_BLUETOOTH          Device = 6
)

var Device_name = map[int32]string{
	0: "DEVICE_UNSPECIFIED",
	1: "WIFI",
	2: "RADIO",
	3: "SERIAL",
	4: "USB",
	5: "LORA",
	6: "BLUETOOTH",
}
var Device_value = map[string]int32{
	"DEVICE_UNSPECIFIED": 0,
	"WIFI":               1,
	"RADIO":              2,
	"SERIAL":             3,
	"USB":                4,
	"LORA":               5,
	"BLUETOOTH":          6,
}

func (x Device) String() string {
	return proto.EnumName(Device_name, int32(x))
}
func (Device) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ArgType int32

const (
	ArgType_ARG_TYPE_INVALID ArgType = 0
	ArgType_ARG_TYPE_UINT8   ArgType = 1
	ArgType_ARG_TYPE_UINT16  ArgType = 2
	ArgType_ARG_TYPE_UINT32  ArgType = 3
	ArgType_ARG_TYPE_UINT64  ArgType = 4
	ArgType_ARG_TYPE_INT8    ArgType = 5
	ArgType_ARG_TYPE_INT16   ArgType = 6
	ArgType_ARG_TYPE_INT32   ArgType = 7
	ArgType_ARG_TYPE_INT64   ArgType = 8
	ArgType_ARG_TYPE_FLOAT   ArgType = 9
)

var ArgType_name = map[int32]string{
	0: "ARG_TYPE_INVALID",
	1: "ARG_TYPE_UINT8",
	2: "ARG_TYPE_UINT16",
	3: "ARG_TYPE_UINT32",
	4: "ARG_TYPE_UINT64",
	5: "ARG_TYPE_INT8",
	6: "ARG_TYPE_INT16",
	7: "ARG_TYPE_INT32",
	8: "ARG_TYPE_INT64",
	9: "ARG_TYPE_FLOAT",
}
var ArgType_value = map[string]int32{
	"ARG_TYPE_INVALID": 0,
	"ARG_TYPE_UINT8":   1,
	"ARG_TYPE_UINT16":  2,
	"ARG_TYPE_UINT32":  3,
	"ARG_TYPE_UINT64":  4,
	"ARG_TYPE_INT8":    5,
	"ARG_TYPE_INT16":   6,
	"ARG_TYPE_INT32":   7,
	"ARG_TYPE_INT64":   8,
	"ARG_TYPE_FLOAT":   9,
}

func (x ArgType) String() string {
	return proto.EnumName(ArgType_name, int32(x))
}
func (ArgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Void should be used when no explicit request or response is required.
type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Command struct {
	Type      ComType  `protobuf:"varint,1,opt,name=Type,enum=singingcat.ComType" json:"Type,omitempty"`
	Args      [][]byte `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty"`
	Sender    uint32   `protobuf:"varint,3,opt,name=Sender" json:"Sender,omitempty"`
	Recipient uint32   `protobuf:"varint,4,opt,name=Recipient" json:"Recipient,omitempty"`
	Target    uint32   `protobuf:"varint,5,opt,name=Target" json:"Target,omitempty"`
	Seq       uint32   `protobuf:"varint,6,opt,name=Seq" json:"Seq,omitempty"`
	Flags     uint32   `protobuf:"varint,7,opt,name=Flags" json:"Flags,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Command) GetType() ComType {
	if m != nil {
		return m.Type
	}
	return ComType_UNDEFINED
}

func (m *Command) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetSender() uint32 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *Command) GetRecipient() uint32 {
	if m != nil {
		return m.Recipient
	}
	return 0
}

func (m *Command) GetTarget() uint32 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *Command) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Command) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

// Status represents whether or not the specified operation
// was successful, and what error occurred if it was not.
type Status struct {
	// Success is set to true or false depending on whether or
	// not the operation was successful.
	// In the event of an error, further details can be found
	// in the `ErrorCode` and `ErrorDescription` fields.
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	// ErrorCode is present if an error has occurred during the
	// operation. ErrorCode mappings will be listed in our
	// documentation.
	ErrorCode int32 `protobuf:"varint,2,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
	// ErrorDescription is present if an error has occurred during
	// the operation. This is intended to be human-readable (machines
	// should use the ErrorCode instead).
	ErrorDescription string `protobuf:"bytes,3,opt,name=ErrorDescription" json:"ErrorDescription,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Status) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type ModuleRef struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
}

func (m *ModuleRef) Reset()                    { *m = ModuleRef{} }
func (m *ModuleRef) String() string            { return proto.CompactTextString(m) }
func (*ModuleRef) ProtoMessage()               {}
func (*ModuleRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ModuleRef) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

type Com2 struct {
	Sender    uint64          `protobuf:"varint,1,opt,name=Sender" json:"Sender,omitempty"`
	Recipient uint64          `protobuf:"varint,2,opt,name=Recipient" json:"Recipient,omitempty"`
	Target    uint64          `protobuf:"varint,3,opt,name=Target" json:"Target,omitempty"`
	Seq       uint32          `protobuf:"varint,4,opt,name=Seq" json:"Seq,omitempty"`
	Flags     uint32          `protobuf:"varint,5,opt,name=Flags" json:"Flags,omitempty"`
	Args      map[uint32]*Arg `protobuf:"bytes,6,rep,name=Args" json:"Args,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Com2) Reset()                    { *m = Com2{} }
func (m *Com2) String() string            { return proto.CompactTextString(m) }
func (*Com2) ProtoMessage()               {}
func (*Com2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Com2) GetSender() uint64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *Com2) GetRecipient() uint64 {
	if m != nil {
		return m.Recipient
	}
	return 0
}

func (m *Com2) GetTarget() uint64 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *Com2) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Com2) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Com2) GetArgs() map[uint32]*Arg {
	if m != nil {
		return m.Args
	}
	return nil
}

type Arg struct {
	Type       ArgType `protobuf:"varint,1,opt,name=Type,enum=singingcat.ArgType" json:"Type,omitempty"`
	IsArray    bool    `protobuf:"varint,2,opt,name=IsArray" json:"IsArray,omitempty"`
	UintValue  uint64  `protobuf:"varint,3,opt,name=UintValue" json:"UintValue,omitempty"`
	IntValue   int64   `protobuf:"varint,4,opt,name=IntValue" json:"IntValue,omitempty"`
	FloatValue float64 `protobuf:"fixed64,5,opt,name=FloatValue" json:"FloatValue,omitempty"`
	ByteArray  []byte  `protobuf:"bytes,6,opt,name=ByteArray,proto3" json:"ByteArray,omitempty"`
}

func (m *Arg) Reset()                    { *m = Arg{} }
func (m *Arg) String() string            { return proto.CompactTextString(m) }
func (*Arg) ProtoMessage()               {}
func (*Arg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Arg) GetType() ArgType {
	if m != nil {
		return m.Type
	}
	return ArgType_ARG_TYPE_INVALID
}

func (m *Arg) GetIsArray() bool {
	if m != nil {
		return m.IsArray
	}
	return false
}

func (m *Arg) GetUintValue() uint64 {
	if m != nil {
		return m.UintValue
	}
	return 0
}

func (m *Arg) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *Arg) GetFloatValue() float64 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

func (m *Arg) GetByteArray() []byte {
	if m != nil {
		return m.ByteArray
	}
	return nil
}

func init() {
	proto.RegisterType((*Void)(nil), "singingcat.Void")
	proto.RegisterType((*Command)(nil), "singingcat.Command")
	proto.RegisterType((*Status)(nil), "singingcat.Status")
	proto.RegisterType((*ModuleRef)(nil), "singingcat.ModuleRef")
	proto.RegisterType((*Com2)(nil), "singingcat.Com2")
	proto.RegisterType((*Arg)(nil), "singingcat.Arg")
	proto.RegisterEnum("singingcat.SensorType", SensorType_name, SensorType_value)
	proto.RegisterEnum("singingcat.ModFlags", ModFlags_name, ModFlags_value)
	proto.RegisterEnum("singingcat.ComType", ComType_name, ComType_value)
	proto.RegisterEnum("singingcat.ROUTE_COM_TYPE", ROUTE_COM_TYPE_name, ROUTE_COM_TYPE_value)
	proto.RegisterEnum("singingcat.ErrorEnum", ErrorEnum_name, ErrorEnum_value)
	proto.RegisterEnum("singingcat.Device", Device_name, Device_value)
	proto.RegisterEnum("singingcat.ArgType", ArgType_name, ArgType_value)
}

func init() {
	proto.RegisterFile("golang.singingcat.net/apis/singingcat/singingcat.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x56, 0x5b, 0x77, 0x1b, 0x3b,
	0x15, 0x3e, 0xbe, 0xdb, 0xca, 0xa5, 0xbb, 0xea, 0x05, 0x9f, 0x50, 0x0e, 0x21, 0x50, 0x1a, 0x0c,
	0xa4, 0x8d, 0x5b, 0x72, 0x4a, 0xe1, 0x45, 0x9e, 0xd9, 0x76, 0x66, 0x75, 0x3c, 0x1a, 0x24, 0x4d,
	0x42, 0x78, 0xd1, 0x32, 0xc9, 0xe0, 0xe5, 0x45, 0x62, 0x07, 0xdb, 0xe9, 0x5a, 0xf9, 0x49, 0x3c,
	0xf3, 0xca, 0x03, 0x7f, 0x82, 0xbf, 0xc2, 0x33, 0x6b, 0x4b, 0x76, 0xec, 0x24, 0x5d, 0xe7, 0x6d,
	0xf6, 0x27, 0xed, 0xdb, 0xb7, 0x3f, 0x69, 0xc4, 0x8e, 0x86, 0x93, 0xcb, 0xc1, 0x78, 0x78, 0x30,
	0x1b, 0x8d, 0x87, 0xa3, 0xf1, 0xf0, 0x7c, 0x30, 0x3f, 0x18, 0xe7, 0xf3, 0xb7, 0x83, 0xeb, 0xd1,
	0xec, 0xed, 0x0a, 0x5b, 0xfb, 0x3c, 0xb8, 0x9e, 0x4e, 0xe6, 0x13, 0xce, 0x56, 0xc8, 0x5e, 0x95,
	0x95, 0x4f, 0x26, 0xa3, 0x8b, 0xbd, 0x7f, 0x17, 0x58, 0x2d, 0x98, 0x5c, 0x5d, 0x0d, 0xc6, 0x17,
	0xfc, 0x0d, 0x2b, 0x9b, 0xdb, 0xeb, 0xbc, 0x59, 0xd8, 0x2d, 0xec, 0x6f, 0xb7, 0x9f, 0xad, 0xc7,
	0x0f, 0x26, 0x57, 0xb4, 0xa4, 0xdc, 0x06, 0xce, 0x59, 0x59, 0x4c, 0x87, 0xb3, 0x66, 0x71, 0xb7,
	0xb4, 0xbf, 0xa9, 0xdc, 0x37, 0x7f, 0xc9, 0xaa, 0x3a, 0x1f, 0x5f, 0xe4, 0xd3, 0x66, 0x69, 0xb7,
	0xb0, 0xbf, 0xa5, 0x16, 0x16, 0x7f, 0xc5, 0x1a, 0x2a, 0x3f, 0x1f, 0x5d, 0x8f, 0xf2, 0xf1, 0xbc,
	0x59, 0x76, 0x4b, 0x2b, 0x80, 0xbc, 0xcc, 0x60, 0x3a, 0xcc, 0xe7, 0xcd, 0x8a, 0xf7, 0xf2, 0x16,
	0x07, 0x56, 0xd2, 0xf9, 0x3f, 0x9a, 0x55, 0x07, 0xd2, 0x27, 0x7f, 0xce, 0x2a, 0xdd, 0xcb, 0xc1,
	0x70, 0xd6, 0xac, 0x39, 0xcc, 0x1b, 0x7b, 0x97, 0xac, 0xaa, 0xe7, 0x83, 0xf9, 0xcd, 0x8c, 0x37,
	0x59, 0x4d, 0xdf, 0x9c, 0x9f, 0xe7, 0xb3, 0x99, 0xab, 0xbf, 0xae, 0x96, 0x26, 0x55, 0x80, 0xd3,
	0xe9, 0x64, 0x1a, 0x4c, 0x2e, 0xf2, 0x66, 0x71, 0xb7, 0xb0, 0x5f, 0x51, 0x2b, 0x80, 0xb7, 0x18,
	0x38, 0x23, 0xcc, 0x67, 0xe7, 0xd3, 0xd1, 0xf5, 0x7c, 0x34, 0x19, 0xbb, 0x0e, 0x1a, 0xea, 0x11,
	0xbe, 0xf7, 0x86, 0x35, 0xfa, 0x93, 0x8b, 0x9b, 0xcb, 0x5c, 0xe5, 0x7f, 0xe3, 0x3b, 0xac, 0xee,
	0x8d, 0x28, 0x74, 0x19, 0xcb, 0xea, 0xce, 0xde, 0xfb, 0x5f, 0x81, 0x95, 0x83, 0xc9, 0x55, 0x7b,
	0x8d, 0x15, 0xbf, 0xe5, 0xab, 0xac, 0x14, 0xdd, 0xd2, 0x57, 0x59, 0x29, 0x79, 0xaf, 0xfb, 0xac,
	0x94, 0xbf, 0xc2, 0x4a, 0x65, 0x8d, 0x15, 0x7e, 0xb0, 0x98, 0x4f, 0x75, 0xb7, 0xb4, 0xbf, 0xd1,
	0xde, 0x79, 0x30, 0xc8, 0xf6, 0x01, 0x2d, 0xe2, 0x78, 0x3e, 0xbd, 0xf5, 0xb3, 0xdb, 0x39, 0x66,
	0x8d, 0x3b, 0x88, 0x92, 0xfc, 0x3d, 0xbf, 0x75, 0xf5, 0x6e, 0x29, 0xfa, 0xe4, 0xaf, 0x59, 0xe5,
	0xcb, 0xe0, 0xf2, 0xc6, 0x93, 0xb7, 0xd1, 0x7e, 0xb2, 0x1e, 0x4f, 0x4c, 0x87, 0xca, 0xaf, 0x7e,
	0x2a, 0x7e, 0x2c, 0xec, 0xfd, 0xa7, 0xc0, 0x4a, 0x62, 0x3a, 0xfc, 0x21, 0x29, 0x89, 0xe9, 0x70,
	0x4d, 0x4a, 0x4d, 0x56, 0x8b, 0x66, 0x62, 0x3a, 0x1d, 0xdc, 0xba, 0xe8, 0x75, 0xb5, 0x34, 0x89,
	0xa2, 0x6c, 0x34, 0x9e, 0x9f, 0xb8, 0xcc, 0x9e, 0x87, 0x15, 0x40, 0xec, 0x47, 0xcb, 0x45, 0xe2,
	0xa3, 0xa4, 0xee, 0x6c, 0xfe, 0x1d, 0x63, 0xdd, 0xcb, 0xc9, 0x60, 0xb1, 0x4a, 0xcc, 0x14, 0xd4,
	0x1a, 0x42, 0x91, 0x3b, 0xb7, 0xf3, 0xdc, 0x67, 0x25, 0x89, 0x6d, 0xaa, 0x15, 0xd0, 0xea, 0x30,
	0xa6, 0xf3, 0xf1, 0x6c, 0x32, 0x5d, 0xd4, 0xf7, 0x5c, 0x63, 0xa2, 0xa5, 0x32, 0x67, 0x29, 0xda,
	0x2c, 0x09, 0xb1, 0x1b, 0x25, 0x18, 0xc2, 0x37, 0x7c, 0x87, 0xbd, 0x5c, 0x5b, 0x31, 0xd8, 0x4f,
	0x51, 0x09, 0x93, 0x29, 0x84, 0x42, 0xeb, 0x5f, 0x45, 0x27, 0x0e, 0x3f, 0x0d, 0x60, 0x9b, 0xf8,
	0x67, 0x83, 0x2a, 0x11, 0xb1, 0x8d, 0x9d, 0xeb, 0x53, 0xb6, 0x15, 0x46, 0x5a, 0x74, 0x62, 0xb4,
	0x4a, 0x84, 0x91, 0x84, 0x02, 0x7f, 0xc6, 0x9e, 0x50, 0x08, 0xeb, 0x43, 0xda, 0xf6, 0x61, 0x1b,
	0x8a, 0x9c, 0xb3, 0x6d, 0xb7, 0x6e, 0x15, 0xf6, 0x45, 0x9a, 0x62, 0x08, 0xa5, 0xc7, 0x1b, 0x0f,
	0xa1, 0xfc, 0x08, 0x7c, 0xf7, 0x3d, 0x54, 0xf8, 0x73, 0x06, 0xde, 0xd6, 0x76, 0x91, 0x2d, 0x84,
	0x2a, 0xa1, 0x31, 0x86, 0x36, 0x95, 0xa7, 0xa8, 0x6c, 0x94, 0x9c, 0xa0, 0x32, 0x50, 0xa3, 0x00,
	0x4a, 0x66, 0x26, 0x4a, 0x7a, 0x16, 0x13, 0xbf, 0xb5, 0xce, 0xb7, 0x19, 0xf3, 0xdb, 0xb4, 0x38,
	0x41, 0x68, 0xf0, 0x4d, 0x56, 0x3f, 0x8d, 0xba, 0x91, 0xcd, 0xc2, 0x14, 0x18, 0xff, 0x96, 0xbd,
	0xd0, 0x46, 0x28, 0x63, 0x33, 0x8d, 0x4a, 0xa4, 0xa9, 0x95, 0x89, 0xed, 0x48, 0x69, 0x60, 0x83,
	0xfa, 0x3b, 0x3d, 0x8e, 0x0c, 0xda, 0xbe, 0xd4, 0x5d, 0x34, 0xef, 0x60, 0xf3, 0x21, 0x74, 0x08,
	0x5b, 0xfc, 0x09, 0xdb, 0x40, 0x9d, 0xda, 0x40, 0xf6, 0xf5, 0x59, 0x12, 0xc0, 0x76, 0xeb, 0x9f,
	0x75, 0x77, 0x17, 0x39, 0xde, 0xb7, 0x58, 0x63, 0x9d, 0xec, 0x3a, 0x2b, 0x27, 0x52, 0xa6, 0x50,
	0xe0, 0x35, 0x56, 0x8a, 0x65, 0x0f, 0x8a, 0x34, 0x19, 0xa1, 0xd0, 0x9e, 0xc9, 0xcc, 0x0a, 0xab,
	0xa3, 0xa4, 0x17, 0x25, 0xbd, 0x40, 0x18, 0x28, 0x11, 0xe1, 0x1a, 0x8d, 0x3d, 0x8d, 0x14, 0xc6,
	0xa8, 0x35, 0x94, 0xf9, 0x06, 0xab, 0x11, 0x92, 0x46, 0x09, 0x54, 0x78, 0x83, 0x55, 0x14, 0x6a,
	0x34, 0x50, 0xa5, 0x2c, 0xd4, 0x91, 0x37, 0x6b, 0x64, 0x3a, 0xbe, 0xd3, 0x28, 0xe9, 0x41, 0x9d,
	0x6a, 0xbe, 0x33, 0x63, 0xca, 0xde, 0xe0, 0x8c, 0x55, 0xbb, 0x0a, 0xf1, 0x2f, 0x08, 0x8c, 0xd2,
	0xc4, 0x91, 0x36, 0xb6, 0x2f, 0xc3, 0x2c, 0x46, 0xed, 0xfb, 0x26, 0x16, 0xd1, 0x2a, 0xfc, 0x53,
	0x86, 0xda, 0xc0, 0x26, 0x39, 0x68, 0xa3, 0x64, 0x07, 0x61, 0x8b, 0xc2, 0xf7, 0xd0, 0xa4, 0x59,
	0xe7, 0x33, 0x9e, 0xc1, 0x36, 0xed, 0xd6, 0x68, 0x82, 0x58, 0x66, 0xa1, 0x91, 0x9f, 0x31, 0x81,
	0x27, 0x44, 0x09, 0xd5, 0x19, 0xcb, 0x1e, 0x75, 0x03, 0xb0, 0x2c, 0x3c, 0x14, 0x01, 0x3c, 0xa5,
	0x21, 0x91, 0xa1, 0x51, 0x45, 0x22, 0xb6, 0xa9, 0x54, 0x06, 0x38, 0x05, 0xed, 0xc6, 0x42, 0x1f,
	0x5b, 0x91, 0xa6, 0xf0, 0x8c, 0x22, 0xf4, 0xa3, 0x40, 0xc9, 0x45, 0xd2, 0xe7, 0x6b, 0x4e, 0x27,
	0xa8, 0x6c, 0x22, 0xfa, 0x08, 0x2f, 0x3c, 0x43, 0x4e, 0x2a, 0x27, 0x22, 0xce, 0x10, 0x5e, 0xd2,
	0x6c, 0x45, 0x92, 0xc8, 0x2c, 0x09, 0x10, 0x7e, 0x44, 0x4e, 0x5d, 0x11, 0x18, 0xa9, 0xce, 0x6c,
	0x88, 0x5d, 0x91, 0xc5, 0x06, 0x9a, 0xae, 0x5e, 0xd9, 0x35, 0xa7, 0xc4, 0x7a, 0x94, 0x74, 0x25,
	0x7c, 0x4b, 0x02, 0xed, 0x64, 0xc6, 0xc8, 0xc4, 0xa6, 0x0a, 0xb5, 0xc6, 0x10, 0x76, 0x5c, 0x0f,
	0x46, 0xa1, 0xe8, 0x6b, 0x34, 0x59, 0x0a, 0x3f, 0x26, 0x19, 0x79, 0x20, 0x14, 0x46, 0xc0, 0x2b,
	0x17, 0xc7, 0xd9, 0x81, 0x4c, 0x8c, 0x92, 0x31, 0xfc, 0x84, 0xb2, 0x77, 0xe2, 0x28, 0xf9, 0x4c,
	0xba, 0xfb, 0x8e, 0xa2, 0x2e, 0xaa, 0x5b, 0xf2, 0xf8, 0x53, 0x12, 0x00, 0xe9, 0x0c, 0x76, 0x49,
	0xc0, 0xfe, 0x50, 0xf4, 0xd0, 0xd8, 0x40, 0x26, 0xdd, 0xa8, 0x07, 0x3f, 0x5b, 0xa1, 0x7a, 0x85,
	0xee, 0x51, 0x5c, 0xa7, 0x51, 0xe2, 0xe6, 0xe7, 0xfc, 0x25, 0xe3, 0x51, 0xa2, 0x8d, 0x88, 0xe9,
	0x00, 0x2c, 0x5b, 0x81, 0x5f, 0x10, 0xc9, 0x4e, 0xe7, 0x59, 0x0a, 0xaf, 0xa9, 0xba, 0xde, 0xbd,
	0xa9, 0xfc, 0x92, 0x94, 0xbe, 0x28, 0xc4, 0x1f, 0x57, 0xdb, 0x95, 0xea, 0x54, 0xa8, 0x10, 0xde,
	0xdc, 0x49, 0xe4, 0x0e, 0xda, 0x5f, 0x12, 0xee, 0x6b, 0xb0, 0xdd, 0x58, 0xf4, 0xe0, 0x57, 0xfe,
	0x2c, 0xde, 0x03, 0x35, 0xb4, 0xf8, 0x0b, 0xf6, 0xd4, 0xcf, 0xc6, 0x35, 0xeb, 0x47, 0x04, 0xbf,
	0x5e, 0xaa, 0xd2, 0x93, 0xfc, 0x1b, 0xf2, 0xed, 0x3d, 0xf4, 0xfd, 0xed, 0xdd, 0x26, 0x1d, 0x88,
	0x04, 0x0e, 0x9c, 0xa8, 0x49, 0x7a, 0xf0, 0xd6, 0x8b, 0xc8, 0x45, 0x24, 0x79, 0xc2, 0x3b, 0x2f,
	0x34, 0x07, 0x2c, 0x88, 0x39, 0x5c, 0x83, 0x14, 0x3a, 0x21, 0xb5, 0x49, 0x13, 0x5e, 0xbc, 0x59,
	0x1a, 0x0a, 0x83, 0xf0, 0x9e, 0x52, 0x90, 0x40, 0xbd, 0xcf, 0x07, 0xda, 0xb0, 0x3c, 0xea, 0xae,
	0xb2, 0xdf, 0x51, 0xab, 0x4b, 0x64, 0x39, 0xcb, 0x23, 0x6a, 0xaa, 0x63, 0x2c, 0xdd, 0x6d, 0xa7,
	0x8a, 0x0e, 0x7c, 0x70, 0x2c, 0x14, 0x7c, 0x4f, 0x07, 0xa1, 0x8f, 0x46, 0x45, 0x01, 0x7c, 0xa4,
	0x8e, 0x16, 0x5b, 0x14, 0x8a, 0xd0, 0xef, 0xf8, 0xbd, 0x13, 0x81, 0xb1, 0x29, 0xa2, 0xd2, 0xf0,
	0x89, 0x54, 0x13, 0xc4, 0x32, 0xf8, 0x6c, 0xdd, 0xe5, 0xf0, 0x07, 0x2a, 0xc6, 0x0d, 0xc9, 0xa5,
	0xfe, 0x63, 0xcb, 0xb2, 0x6d, 0x5f, 0x6d, 0x20, 0xfb, 0x96, 0x6e, 0x60, 0x72, 0xf0, 0x48, 0x22,
	0x13, 0x84, 0x6f, 0x56, 0xb6, 0x63, 0xa1, 0xe0, 0x0e, 0xb7, 0xb3, 0x45, 0x18, 0x42, 0x71, 0x65,
	0x86, 0x18, 0x43, 0x89, 0x48, 0x5b, 0xc4, 0x8b, 0x51, 0x28, 0x28, 0xb7, 0x3e, 0x2d, 0x5e, 0x0d,
	0x38, 0xbe, 0xb9, 0xa2, 0xd5, 0x2c, 0xd1, 0x29, 0x06, 0x51, 0x37, 0x72, 0xf7, 0xd1, 0x2b, 0xd6,
	0x14, 0x41, 0x20, 0xb3, 0xc4, 0xd8, 0x2c, 0xed, 0x29, 0x11, 0xfa, 0x33, 0x1f, 0x29, 0x0c, 0xa1,
	0xd0, 0x3a, 0x67, 0xd5, 0x30, 0xff, 0x32, 0x3a, 0xcf, 0x49, 0x72, 0x21, 0x9e, 0x44, 0x01, 0xfd,
	0x3a, 0xd6, 0xfd, 0xeb, 0xac, 0x4c, 0xd3, 0x83, 0x82, 0x1b, 0x9c, 0xfb, 0x07, 0x14, 0xdd, 0x5d,
	0xe1, 0xce, 0x36, 0x94, 0xe8, 0x9a, 0xcb, 0x74, 0x07, 0xca, 0xb4, 0x33, 0x96, 0x4a, 0x40, 0x85,
	0x2a, 0xee, 0xc4, 0x19, 0x1a, 0x29, 0xcd, 0x31, 0x54, 0x5b, 0xff, 0x2d, 0xb0, 0xda, 0xe2, 0x5f,
	0x4a, 0x84, 0x0a, 0xd5, 0x73, 0x3c, 0xd0, 0x9d, 0x2e, 0xe2, 0x88, 0x92, 0x70, 0xb6, 0x7d, 0x87,
	0x66, 0x51, 0x62, 0x3e, 0xfa, 0xff, 0xcc, 0x3d, 0xec, 0xf0, 0x08, 0x8a, 0x8f, 0xc0, 0xf7, 0x6d,
	0xff, 0xa3, 0xb9, 0x07, 0x1e, 0x7d, 0x80, 0x32, 0xe9, 0x66, 0x2d, 0x91, 0xf9, 0x08, 0x95, 0x7b,
	0x59, 0x7c, 0xc0, 0xea, 0x43, 0xec, 0x7d, 0x1b, 0x6a, 0x0f, 0xb1, 0xa3, 0x0f, 0x50, 0xbf, 0x87,
	0x75, 0x63, 0x29, 0x0c, 0x34, 0x3a, 0x6f, 0xd8, 0xeb, 0x71, 0x3e, 0x5f, 0x7f, 0x31, 0x2c, 0x9e,
	0xbb, 0xf4, 0xbe, 0x5d, 0x83, 0xff, 0x5a, 0x75, 0xaf, 0xda, 0xf7, 0xff, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x1b, 0x34, 0x23, 0x0f, 0x0b, 0x00, 0x00,
}
