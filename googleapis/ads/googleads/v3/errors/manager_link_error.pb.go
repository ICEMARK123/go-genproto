// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/manager_link_error.proto

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

// Enum describing possible ManagerLink errors.
type ManagerLinkErrorEnum_ManagerLinkError int32

const (
	// Enum unspecified.
	ManagerLinkErrorEnum_UNSPECIFIED ManagerLinkErrorEnum_ManagerLinkError = 0
	// The received error code is not known in this version.
	ManagerLinkErrorEnum_UNKNOWN ManagerLinkErrorEnum_ManagerLinkError = 1
	// The manager and client have incompatible account types.
	ManagerLinkErrorEnum_ACCOUNTS_NOT_COMPATIBLE_FOR_LINKING ManagerLinkErrorEnum_ManagerLinkError = 2
	// Client is already linked to too many managers.
	ManagerLinkErrorEnum_TOO_MANY_MANAGERS ManagerLinkErrorEnum_ManagerLinkError = 3
	// Manager has too many pending invitations.
	ManagerLinkErrorEnum_TOO_MANY_INVITES ManagerLinkErrorEnum_ManagerLinkError = 4
	// Client is already invited by this manager.
	ManagerLinkErrorEnum_ALREADY_INVITED_BY_THIS_MANAGER ManagerLinkErrorEnum_ManagerLinkError = 5
	// The client is already managed by this manager.
	ManagerLinkErrorEnum_ALREADY_MANAGED_BY_THIS_MANAGER ManagerLinkErrorEnum_ManagerLinkError = 6
	// Client is already managed in hierarchy.
	ManagerLinkErrorEnum_ALREADY_MANAGED_IN_HIERARCHY ManagerLinkErrorEnum_ManagerLinkError = 7
	// Manger and sub-manager to be linked have duplicate client.
	ManagerLinkErrorEnum_DUPLICATE_CHILD_FOUND ManagerLinkErrorEnum_ManagerLinkError = 8
	// Client has no active user that can access the client account.
	ManagerLinkErrorEnum_CLIENT_HAS_NO_ADMIN_USER ManagerLinkErrorEnum_ManagerLinkError = 9
	// Adding this link would exceed the maximum hierarchy depth.
	ManagerLinkErrorEnum_MAX_DEPTH_EXCEEDED ManagerLinkErrorEnum_ManagerLinkError = 10
	// Adding this link will create a cycle.
	ManagerLinkErrorEnum_CYCLE_NOT_ALLOWED ManagerLinkErrorEnum_ManagerLinkError = 11
	// Manager account has the maximum number of linked clients.
	ManagerLinkErrorEnum_TOO_MANY_ACCOUNTS ManagerLinkErrorEnum_ManagerLinkError = 12
	// Parent manager account has the maximum number of linked clients.
	ManagerLinkErrorEnum_TOO_MANY_ACCOUNTS_AT_MANAGER ManagerLinkErrorEnum_ManagerLinkError = 13
	// The account is not authorized owner.
	ManagerLinkErrorEnum_NON_OWNER_USER_CANNOT_MODIFY_LINK ManagerLinkErrorEnum_ManagerLinkError = 14
	// Your manager account is suspended, and you are no longer allowed to link
	// to clients.
	ManagerLinkErrorEnum_SUSPENDED_ACCOUNT_CANNOT_ADD_CLIENTS ManagerLinkErrorEnum_ManagerLinkError = 15
	// You are not allowed to move a client to a manager that is not under your
	// current hierarchy.
	ManagerLinkErrorEnum_CLIENT_OUTSIDE_TREE ManagerLinkErrorEnum_ManagerLinkError = 16
)

var ManagerLinkErrorEnum_ManagerLinkError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ACCOUNTS_NOT_COMPATIBLE_FOR_LINKING",
	3:  "TOO_MANY_MANAGERS",
	4:  "TOO_MANY_INVITES",
	5:  "ALREADY_INVITED_BY_THIS_MANAGER",
	6:  "ALREADY_MANAGED_BY_THIS_MANAGER",
	7:  "ALREADY_MANAGED_IN_HIERARCHY",
	8:  "DUPLICATE_CHILD_FOUND",
	9:  "CLIENT_HAS_NO_ADMIN_USER",
	10: "MAX_DEPTH_EXCEEDED",
	11: "CYCLE_NOT_ALLOWED",
	12: "TOO_MANY_ACCOUNTS",
	13: "TOO_MANY_ACCOUNTS_AT_MANAGER",
	14: "NON_OWNER_USER_CANNOT_MODIFY_LINK",
	15: "SUSPENDED_ACCOUNT_CANNOT_ADD_CLIENTS",
	16: "CLIENT_OUTSIDE_TREE",
}

var ManagerLinkErrorEnum_ManagerLinkError_value = map[string]int32{
	"UNSPECIFIED":                          0,
	"UNKNOWN":                              1,
	"ACCOUNTS_NOT_COMPATIBLE_FOR_LINKING":  2,
	"TOO_MANY_MANAGERS":                    3,
	"TOO_MANY_INVITES":                     4,
	"ALREADY_INVITED_BY_THIS_MANAGER":      5,
	"ALREADY_MANAGED_BY_THIS_MANAGER":      6,
	"ALREADY_MANAGED_IN_HIERARCHY":         7,
	"DUPLICATE_CHILD_FOUND":                8,
	"CLIENT_HAS_NO_ADMIN_USER":             9,
	"MAX_DEPTH_EXCEEDED":                   10,
	"CYCLE_NOT_ALLOWED":                    11,
	"TOO_MANY_ACCOUNTS":                    12,
	"TOO_MANY_ACCOUNTS_AT_MANAGER":         13,
	"NON_OWNER_USER_CANNOT_MODIFY_LINK":    14,
	"SUSPENDED_ACCOUNT_CANNOT_ADD_CLIENTS": 15,
	"CLIENT_OUTSIDE_TREE":                  16,
}

func (x ManagerLinkErrorEnum_ManagerLinkError) String() string {
	return proto.EnumName(ManagerLinkErrorEnum_ManagerLinkError_name, int32(x))
}

func (ManagerLinkErrorEnum_ManagerLinkError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aba35d6dfea2fe63, []int{0, 0}
}

// Container for enum describing possible ManagerLink errors.
type ManagerLinkErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManagerLinkErrorEnum) Reset()         { *m = ManagerLinkErrorEnum{} }
func (m *ManagerLinkErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ManagerLinkErrorEnum) ProtoMessage()    {}
func (*ManagerLinkErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba35d6dfea2fe63, []int{0}
}

func (m *ManagerLinkErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagerLinkErrorEnum.Unmarshal(m, b)
}
func (m *ManagerLinkErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagerLinkErrorEnum.Marshal(b, m, deterministic)
}
func (m *ManagerLinkErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagerLinkErrorEnum.Merge(m, src)
}
func (m *ManagerLinkErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ManagerLinkErrorEnum.Size(m)
}
func (m *ManagerLinkErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagerLinkErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ManagerLinkErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.ManagerLinkErrorEnum_ManagerLinkError", ManagerLinkErrorEnum_ManagerLinkError_name, ManagerLinkErrorEnum_ManagerLinkError_value)
	proto.RegisterType((*ManagerLinkErrorEnum)(nil), "google.ads.googleads.v3.errors.ManagerLinkErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/manager_link_error.proto", fileDescriptor_aba35d6dfea2fe63)
}

var fileDescriptor_aba35d6dfea2fe63 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x7f, 0xeb, 0xf6, 0xdb, 0xc0, 0x03, 0x66, 0xcc, 0xc6, 0x3f, 0x4d, 0x03, 0x3a, 0x10,
	0x5c, 0xa5, 0x17, 0xbd, 0x40, 0x0a, 0x57, 0xae, 0x7d, 0xda, 0x58, 0x4b, 0xec, 0x28, 0x71, 0xba,
	0x15, 0x55, 0xb2, 0x02, 0xad, 0xa2, 0x6a, 0x5b, 0x32, 0x25, 0x63, 0x8f, 0xc2, 0x03, 0x70, 0xc9,
	0xa3, 0xf0, 0x28, 0x5c, 0xf1, 0x02, 0x48, 0x28, 0xf1, 0x52, 0x41, 0x27, 0xb8, 0x49, 0x8e, 0xbe,
	0xfe, 0x7c, 0x8f, 0x7d, 0x8e, 0xbe, 0xe8, 0x6d, 0x56, 0x14, 0xd9, 0xd9, 0xbc, 0x97, 0xce, 0xaa,
	0x9e, 0x2d, 0xeb, 0xea, 0xaa, 0xdf, 0x9b, 0x97, 0x65, 0x51, 0x56, 0xbd, 0xf3, 0x34, 0x4f, 0xb3,
	0x79, 0x69, 0xce, 0x16, 0xf9, 0xa9, 0x69, 0x34, 0xe7, 0xa2, 0x2c, 0x2e, 0x0b, 0x72, 0x60, 0x69,
	0x27, 0x9d, 0x55, 0xce, 0xd2, 0xe8, 0x5c, 0xf5, 0x1d, 0x6b, 0x7c, 0xba, 0xdf, 0x36, 0xbe, 0x58,
	0xf4, 0xd2, 0x3c, 0x2f, 0x2e, 0xd3, 0xcb, 0x45, 0x91, 0x57, 0xd6, 0xdd, 0xfd, 0xbc, 0x81, 0x76,
	0x03, 0xdb, 0xda, 0x5f, 0xe4, 0xa7, 0x50, 0x7b, 0x20, 0xff, 0x74, 0xde, 0xfd, 0xb9, 0x8e, 0xf0,
	0xea, 0x01, 0xd9, 0x41, 0xdb, 0x89, 0x8c, 0x43, 0x60, 0x62, 0x28, 0x80, 0xe3, 0xff, 0xc8, 0x36,
	0xda, 0x4a, 0xe4, 0x91, 0x54, 0xc7, 0x12, 0xaf, 0x91, 0xd7, 0xe8, 0x90, 0x32, 0xa6, 0x12, 0xa9,
	0x63, 0x23, 0x95, 0x36, 0x4c, 0x05, 0x21, 0xd5, 0x62, 0xe0, 0x83, 0x19, 0xaa, 0xc8, 0xf8, 0x42,
	0x1e, 0x09, 0x39, 0xc2, 0x1d, 0xb2, 0x87, 0xee, 0x6b, 0xa5, 0x4c, 0x40, 0xe5, 0xa4, 0xfe, 0xd0,
	0x11, 0x44, 0x31, 0x5e, 0x27, 0xbb, 0x08, 0x2f, 0x65, 0x21, 0xc7, 0x42, 0x43, 0x8c, 0x37, 0xc8,
	0x21, 0x7a, 0x46, 0xfd, 0x08, 0x28, 0x6f, 0x45, 0x6e, 0x06, 0x13, 0xa3, 0x3d, 0x11, 0xb7, 0x5e,
	0xfc, 0xff, 0xef, 0x90, 0x15, 0x6f, 0x42, 0x9b, 0xe4, 0x39, 0xda, 0x5f, 0x85, 0x84, 0x34, 0x9e,
	0x80, 0x88, 0x46, 0xcc, 0x9b, 0xe0, 0x2d, 0xf2, 0x04, 0xed, 0xf1, 0x24, 0xf4, 0x05, 0xa3, 0x1a,
	0x0c, 0xf3, 0x84, 0xcf, 0xcd, 0x50, 0x25, 0x92, 0xe3, 0x5b, 0x64, 0x1f, 0x3d, 0x66, 0xbe, 0x00,
	0xa9, 0x8d, 0x47, 0xeb, 0xf1, 0x0c, 0xe5, 0x81, 0x90, 0x26, 0x89, 0x21, 0xc2, 0xb7, 0xc9, 0x43,
	0x44, 0x02, 0x7a, 0x62, 0x38, 0x84, 0xda, 0x33, 0x70, 0xc2, 0x00, 0x38, 0x70, 0x8c, 0xea, 0x49,
	0xd9, 0x84, 0xf9, 0xd0, 0xec, 0x83, 0xfa, 0xbe, 0x3a, 0x06, 0x8e, 0xb7, 0xff, 0x58, 0x40, 0xbb,
	0x32, 0x7c, 0xa7, 0x7e, 0xe0, 0x0d, 0xd9, 0x50, 0xbd, 0x1c, 0xe1, 0x2e, 0x79, 0x85, 0x5e, 0x48,
	0x25, 0x8d, 0x3a, 0x96, 0x10, 0x35, 0x77, 0x1b, 0x46, 0x65, 0xdd, 0x3b, 0x50, 0x5c, 0x0c, 0x27,
	0xcd, 0x8e, 0xf1, 0x3d, 0xf2, 0x06, 0xbd, 0x8c, 0x93, 0x38, 0x04, 0xc9, 0x81, 0xb7, 0x9d, 0x5a,
	0x92, 0x72, 0x6e, 0xec, 0x24, 0x31, 0xde, 0x21, 0x8f, 0xd0, 0x83, 0xeb, 0xb1, 0x54, 0xa2, 0x63,
	0xc1, 0xc1, 0xe8, 0x08, 0x00, 0xe3, 0xc1, 0x8f, 0x35, 0xd4, 0xfd, 0x58, 0x9c, 0x3b, 0xff, 0x4e,
	0xd7, 0x60, 0x6f, 0x35, 0x23, 0x61, 0x1d, 0xab, 0x70, 0xed, 0x3d, 0xbf, 0x36, 0x66, 0xc5, 0x59,
	0x9a, 0x67, 0x4e, 0x51, 0x66, 0xbd, 0x6c, 0x9e, 0x37, 0xa1, 0x6b, 0xf3, 0x7d, 0xb1, 0xa8, 0xfe,
	0x16, 0xf7, 0x77, 0xf6, 0xf7, 0xa5, 0xb3, 0x3e, 0xa2, 0xf4, 0x6b, 0xe7, 0x60, 0x64, 0x9b, 0xd1,
	0x59, 0xe5, 0xd8, 0xb2, 0xae, 0xc6, 0x7d, 0xa7, 0xb9, 0xb2, 0xfa, 0xd6, 0x02, 0x53, 0x3a, 0xab,
	0xa6, 0x4b, 0x60, 0x3a, 0xee, 0x4f, 0x2d, 0xf0, 0xbd, 0xd3, 0xb5, 0xaa, 0xeb, 0xd2, 0x59, 0xe5,
	0xba, 0x4b, 0xc4, 0x75, 0xc7, 0x7d, 0xd7, 0xb5, 0xd0, 0x87, 0xcd, 0xe6, 0x75, 0xfd, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x75, 0x31, 0x98, 0x66, 0x8b, 0x03, 0x00, 0x00,
}