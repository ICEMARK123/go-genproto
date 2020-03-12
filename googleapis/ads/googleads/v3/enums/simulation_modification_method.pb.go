// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/simulation_modification_method.proto

package enums

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

// Enum describing the method by which a simulation modifies a field.
type SimulationModificationMethodEnum_SimulationModificationMethod int32

const (
	// Not specified.
	SimulationModificationMethodEnum_UNSPECIFIED SimulationModificationMethodEnum_SimulationModificationMethod = 0
	// Used for return value only. Represents value unknown in this version.
	SimulationModificationMethodEnum_UNKNOWN SimulationModificationMethodEnum_SimulationModificationMethod = 1
	// The values in a simulation were applied to all children of a given
	// resource uniformly. Overrides on child resources were not respected.
	SimulationModificationMethodEnum_UNIFORM SimulationModificationMethodEnum_SimulationModificationMethod = 2
	// The values in a simulation were applied to the given resource.
	// Overrides on child resources were respected, and traffic estimates
	// do not include these resources.
	SimulationModificationMethodEnum_DEFAULT SimulationModificationMethodEnum_SimulationModificationMethod = 3
)

var SimulationModificationMethodEnum_SimulationModificationMethod_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNIFORM",
	3: "DEFAULT",
}

var SimulationModificationMethodEnum_SimulationModificationMethod_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UNIFORM":     2,
	"DEFAULT":     3,
}

func (x SimulationModificationMethodEnum_SimulationModificationMethod) String() string {
	return proto.EnumName(SimulationModificationMethodEnum_SimulationModificationMethod_name, int32(x))
}

func (SimulationModificationMethodEnum_SimulationModificationMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cfaf85ba3d548b87, []int{0, 0}
}

// Container for enum describing the method by which a simulation modifies
// a field.
type SimulationModificationMethodEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimulationModificationMethodEnum) Reset()         { *m = SimulationModificationMethodEnum{} }
func (m *SimulationModificationMethodEnum) String() string { return proto.CompactTextString(m) }
func (*SimulationModificationMethodEnum) ProtoMessage()    {}
func (*SimulationModificationMethodEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfaf85ba3d548b87, []int{0}
}

func (m *SimulationModificationMethodEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimulationModificationMethodEnum.Unmarshal(m, b)
}
func (m *SimulationModificationMethodEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimulationModificationMethodEnum.Marshal(b, m, deterministic)
}
func (m *SimulationModificationMethodEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulationModificationMethodEnum.Merge(m, src)
}
func (m *SimulationModificationMethodEnum) XXX_Size() int {
	return xxx_messageInfo_SimulationModificationMethodEnum.Size(m)
}
func (m *SimulationModificationMethodEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulationModificationMethodEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SimulationModificationMethodEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.SimulationModificationMethodEnum_SimulationModificationMethod", SimulationModificationMethodEnum_SimulationModificationMethod_name, SimulationModificationMethodEnum_SimulationModificationMethod_value)
	proto.RegisterType((*SimulationModificationMethodEnum)(nil), "google.ads.googleads.v3.enums.SimulationModificationMethodEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/simulation_modification_method.proto", fileDescriptor_cfaf85ba3d548b87)
}

var fileDescriptor_cfaf85ba3d548b87 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x7d, 0x4d, 0xa5, 0x87, 0xe4, 0x0e, 0x44, 0x1d, 0x51, 0x2b, 0xd1, 0x7e, 0x80, 0x3d, 0x78,
	0x33, 0x93, 0x4b, 0xd3, 0xaa, 0x82, 0xa4, 0x11, 0x25, 0x41, 0x42, 0x91, 0x50, 0xa8, 0x83, 0xb1,
	0x94, 0xd8, 0x51, 0x9d, 0x74, 0xe0, 0x73, 0x18, 0xf9, 0x14, 0x3e, 0x85, 0x0f, 0x60, 0x46, 0xb1,
	0x9b, 0x88, 0x85, 0x2c, 0xd6, 0xb9, 0xbe, 0xe7, 0x9e, 0x73, 0xef, 0x01, 0x0b, 0xae, 0x14, 0xcf,
	0x33, 0x94, 0x32, 0x8d, 0x2c, 0x6c, 0xd0, 0x11, 0xa3, 0x4c, 0xd6, 0x85, 0x46, 0x5a, 0x14, 0x75,
	0x9e, 0x56, 0x42, 0xc9, 0xa7, 0x42, 0x31, 0xf1, 0x22, 0xf6, 0xa7, 0x22, 0xab, 0x5e, 0x15, 0x83,
	0xe5, 0x41, 0x55, 0x6a, 0x3c, 0xb5, 0x83, 0x30, 0x65, 0x1a, 0x76, 0x1a, 0xf0, 0x88, 0xa1, 0xd1,
	0xb8, 0x98, 0xb4, 0x16, 0xa5, 0x40, 0xa9, 0x94, 0xaa, 0x32, 0x12, 0xda, 0x0e, 0xcf, 0xdf, 0xc0,
	0xe5, 0xae, 0x33, 0xf1, 0x7f, 0x79, 0xf8, 0xc6, 0xc2, 0x93, 0x75, 0x31, 0x8f, 0xc1, 0xa4, 0x8f,
	0x33, 0x3e, 0x07, 0xa3, 0x28, 0xd8, 0x85, 0xde, 0xf5, 0x66, 0xb5, 0xf1, 0x96, 0xee, 0xbf, 0xf1,
	0x08, 0x9c, 0x45, 0xc1, 0x4d, 0xb0, 0x7d, 0x08, 0xdc, 0x81, 0x2d, 0x36, 0xab, 0xed, 0x9d, 0xef,
	0x3a, 0x4d, 0xb1, 0xf4, 0x56, 0x34, 0xba, 0xbd, 0x77, 0x87, 0x8b, 0xef, 0x01, 0x98, 0xed, 0x55,
	0x01, 0x7b, 0xf7, 0x5f, 0xcc, 0xfa, 0xbc, 0xc3, 0xe6, 0x88, 0x70, 0xf0, 0x78, 0xca, 0x11, 0x72,
	0x95, 0xa7, 0x92, 0x43, 0x75, 0xe0, 0x88, 0x67, 0xd2, 0x9c, 0xd8, 0xe6, 0x5a, 0x0a, 0xfd, 0x47,
	0xcc, 0x57, 0xe6, 0x7d, 0x77, 0x86, 0x6b, 0x4a, 0x3f, 0x9c, 0xe9, 0xda, 0x4a, 0x51, 0xa6, 0xa1,
	0x85, 0x0d, 0x8a, 0x31, 0x6c, 0xb2, 0xd0, 0x9f, 0x6d, 0x3f, 0xa1, 0x4c, 0x27, 0x5d, 0x3f, 0x89,
	0x71, 0x62, 0xfa, 0x5f, 0xce, 0xcc, 0x7e, 0x12, 0x42, 0x99, 0x26, 0xa4, 0x63, 0x10, 0x12, 0x63,
	0x42, 0x0c, 0xe7, 0xf9, 0xbf, 0x59, 0x0c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x08, 0xda, 0x85,
	0xbb, 0xfe, 0x01, 0x00, 0x00,
}