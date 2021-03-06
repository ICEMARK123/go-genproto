// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/field_error.proto

package errors

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

// Enum describing possible field errors.
type FieldErrorEnum_FieldError int32

const (
	// Enum unspecified.
	FieldErrorEnum_UNSPECIFIED FieldErrorEnum_FieldError = 0
	// The received error code is not known in this version.
	FieldErrorEnum_UNKNOWN FieldErrorEnum_FieldError = 1
	// The required field was not present.
	FieldErrorEnum_REQUIRED FieldErrorEnum_FieldError = 2
	// The field attempted to be mutated is immutable.
	FieldErrorEnum_IMMUTABLE_FIELD FieldErrorEnum_FieldError = 3
	// The field's value is invalid.
	FieldErrorEnum_INVALID_VALUE FieldErrorEnum_FieldError = 4
	// The field cannot be set.
	FieldErrorEnum_VALUE_MUST_BE_UNSET FieldErrorEnum_FieldError = 5
	// The required repeated field was empty.
	FieldErrorEnum_REQUIRED_NONEMPTY_LIST FieldErrorEnum_FieldError = 6
	// The field cannot be cleared.
	FieldErrorEnum_FIELD_CANNOT_BE_CLEARED FieldErrorEnum_FieldError = 7
	// The field's value is on a blacklist for this field.
	FieldErrorEnum_BLACKLISTED_VALUE FieldErrorEnum_FieldError = 8
)

var FieldErrorEnum_FieldError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REQUIRED",
	3: "IMMUTABLE_FIELD",
	4: "INVALID_VALUE",
	5: "VALUE_MUST_BE_UNSET",
	6: "REQUIRED_NONEMPTY_LIST",
	7: "FIELD_CANNOT_BE_CLEARED",
	8: "BLACKLISTED_VALUE",
}

var FieldErrorEnum_FieldError_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"REQUIRED":                2,
	"IMMUTABLE_FIELD":         3,
	"INVALID_VALUE":           4,
	"VALUE_MUST_BE_UNSET":     5,
	"REQUIRED_NONEMPTY_LIST":  6,
	"FIELD_CANNOT_BE_CLEARED": 7,
	"BLACKLISTED_VALUE":       8,
}

func (x FieldErrorEnum_FieldError) String() string {
	return proto.EnumName(FieldErrorEnum_FieldError_name, int32(x))
}

func (FieldErrorEnum_FieldError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94b0ac14134331cc, []int{0, 0}
}

// Container for enum describing possible field errors.
type FieldErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldErrorEnum) Reset()         { *m = FieldErrorEnum{} }
func (m *FieldErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FieldErrorEnum) ProtoMessage()    {}
func (*FieldErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_94b0ac14134331cc, []int{0}
}

func (m *FieldErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldErrorEnum.Unmarshal(m, b)
}
func (m *FieldErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldErrorEnum.Marshal(b, m, deterministic)
}
func (m *FieldErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldErrorEnum.Merge(m, src)
}
func (m *FieldErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FieldErrorEnum.Size(m)
}
func (m *FieldErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FieldErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.FieldErrorEnum_FieldError", FieldErrorEnum_FieldError_name, FieldErrorEnum_FieldError_value)
	proto.RegisterType((*FieldErrorEnum)(nil), "google.ads.googleads.v3.errors.FieldErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/field_error.proto", fileDescriptor_94b0ac14134331cc)
}

var fileDescriptor_94b0ac14134331cc = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0x06, 0xdb, 0xe4, 0x01, 0xcd, 0x3c, 0xc1, 0xa4, 0x81, 0x76, 0x91, 0x07, 0x70,
	0x90, 0x72, 0x67, 0xae, 0x9c, 0xc4, 0x9d, 0xac, 0x26, 0x6e, 0x68, 0xfe, 0x20, 0x50, 0x24, 0x2b,
	0x90, 0x10, 0x45, 0x6a, 0xe3, 0x2a, 0x2e, 0x7d, 0x20, 0x2e, 0x79, 0x13, 0x78, 0x06, 0x9e, 0xa0,
	0x4f, 0x81, 0x12, 0x93, 0xf6, 0x8a, 0x5d, 0xf9, 0xf3, 0xd1, 0xef, 0xfb, 0x8e, 0x7d, 0x0e, 0x78,
	0xd7, 0x48, 0xd9, 0xac, 0x6b, 0xa7, 0xac, 0x94, 0xa3, 0xe5, 0xa0, 0xf6, 0xae, 0x53, 0xf7, 0xbd,
	0xec, 0x95, 0xf3, 0xad, 0xad, 0xd7, 0x95, 0x18, 0x2f, 0x68, 0xdb, 0xcb, 0x9d, 0x84, 0xf7, 0x1a,
	0x43, 0x65, 0xa5, 0xd0, 0xd1, 0x81, 0xf6, 0x2e, 0xd2, 0x8e, 0xbb, 0xb7, 0x53, 0xe2, 0xb6, 0x75,
	0xca, 0xae, 0x93, 0xbb, 0x72, 0xd7, 0xca, 0x4e, 0x69, 0xb7, 0xfd, 0xc7, 0x00, 0x2f, 0xe7, 0x43,
	0x26, 0x1d, 0x68, 0xda, 0x7d, 0xdf, 0xd8, 0xbf, 0x0c, 0x00, 0x4e, 0x25, 0x38, 0x03, 0x57, 0x19,
	0x4f, 0x62, 0xea, 0xb3, 0x39, 0xa3, 0x81, 0xf5, 0x04, 0x5e, 0x81, 0x8b, 0x8c, 0x2f, 0xf8, 0xf2,
	0x23, 0xb7, 0x0c, 0xf8, 0x1c, 0x5c, 0xae, 0xe8, 0x87, 0x8c, 0xad, 0x68, 0x60, 0x99, 0xf0, 0x06,
	0xcc, 0x58, 0x14, 0x65, 0x29, 0xf1, 0x42, 0x2a, 0xe6, 0x8c, 0x86, 0x81, 0x75, 0x06, 0xaf, 0xc1,
	0x0b, 0xc6, 0x73, 0x12, 0xb2, 0x40, 0xe4, 0x24, 0xcc, 0xa8, 0xf5, 0x14, 0xde, 0x82, 0x9b, 0x51,
	0x8a, 0x28, 0x4b, 0x52, 0xe1, 0x51, 0x91, 0xf1, 0x84, 0xa6, 0xd6, 0x33, 0x78, 0x07, 0x5e, 0x4f,
	0x71, 0x82, 0x2f, 0x39, 0x8d, 0xe2, 0xf4, 0x93, 0x08, 0x59, 0x92, 0x5a, 0xe7, 0xf0, 0x0d, 0xb8,
	0x1d, 0x23, 0x85, 0x4f, 0x38, 0x5f, 0x8e, 0x36, 0x3f, 0xa4, 0x64, 0xe8, 0x7c, 0x01, 0x5f, 0x81,
	0x6b, 0x2f, 0x24, 0xfe, 0x62, 0x60, 0xe9, 0xd4, 0xe8, 0xd2, 0x3b, 0x18, 0xc0, 0xfe, 0x2a, 0x37,
	0xe8, 0xf1, 0x19, 0x79, 0xb3, 0xd3, 0x7f, 0xe3, 0x61, 0x2c, 0xb1, 0xf1, 0x39, 0xf8, 0x67, 0x69,
	0xe4, 0xba, 0xec, 0x1a, 0x24, 0xfb, 0xc6, 0x69, 0xea, 0x6e, 0x1c, 0xda, 0xb4, 0x98, 0x6d, 0xab,
	0xfe, 0xb7, 0xa7, 0xf7, 0xfa, 0xf8, 0x61, 0x9e, 0x3d, 0x10, 0xf2, 0xd3, 0xbc, 0x7f, 0xd0, 0x61,
	0xa4, 0x52, 0x48, 0xcb, 0x41, 0xe5, 0x2e, 0x1a, 0x5b, 0xaa, 0xdf, 0x13, 0x50, 0x90, 0x4a, 0x15,
	0x47, 0xa0, 0xc8, 0xdd, 0x42, 0x03, 0x07, 0xd3, 0xd6, 0x55, 0x8c, 0x49, 0xa5, 0x30, 0x3e, 0x22,
	0x18, 0xe7, 0x2e, 0xc6, 0x1a, 0xfa, 0x72, 0x3e, 0xbe, 0xce, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x06, 0x14, 0x1f, 0x44, 0x02, 0x00, 0x00,
}
