// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/ad_network_type.proto

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

// Enumerates Google Ads network types.
type AdNetworkTypeEnum_AdNetworkType int32

const (
	// Not specified.
	AdNetworkTypeEnum_UNSPECIFIED AdNetworkTypeEnum_AdNetworkType = 0
	// The value is unknown in this version.
	AdNetworkTypeEnum_UNKNOWN AdNetworkTypeEnum_AdNetworkType = 1
	// Google search.
	AdNetworkTypeEnum_SEARCH AdNetworkTypeEnum_AdNetworkType = 2
	// Search partners.
	AdNetworkTypeEnum_SEARCH_PARTNERS AdNetworkTypeEnum_AdNetworkType = 3
	// Display Network.
	AdNetworkTypeEnum_CONTENT AdNetworkTypeEnum_AdNetworkType = 4
	// YouTube Search.
	AdNetworkTypeEnum_YOUTUBE_SEARCH AdNetworkTypeEnum_AdNetworkType = 5
	// YouTube Videos
	AdNetworkTypeEnum_YOUTUBE_WATCH AdNetworkTypeEnum_AdNetworkType = 6
	// Cross-network.
	AdNetworkTypeEnum_MIXED AdNetworkTypeEnum_AdNetworkType = 7
)

var AdNetworkTypeEnum_AdNetworkType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH",
	3: "SEARCH_PARTNERS",
	4: "CONTENT",
	5: "YOUTUBE_SEARCH",
	6: "YOUTUBE_WATCH",
	7: "MIXED",
}

var AdNetworkTypeEnum_AdNetworkType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"SEARCH":          2,
	"SEARCH_PARTNERS": 3,
	"CONTENT":         4,
	"YOUTUBE_SEARCH":  5,
	"YOUTUBE_WATCH":   6,
	"MIXED":           7,
}

func (x AdNetworkTypeEnum_AdNetworkType) String() string {
	return proto.EnumName(AdNetworkTypeEnum_AdNetworkType_name, int32(x))
}

func (AdNetworkTypeEnum_AdNetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a418102a00fe2b95, []int{0, 0}
}

// Container for enumeration of Google Ads network types.
type AdNetworkTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdNetworkTypeEnum) Reset()         { *m = AdNetworkTypeEnum{} }
func (m *AdNetworkTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdNetworkTypeEnum) ProtoMessage()    {}
func (*AdNetworkTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a418102a00fe2b95, []int{0}
}

func (m *AdNetworkTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdNetworkTypeEnum.Unmarshal(m, b)
}
func (m *AdNetworkTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdNetworkTypeEnum.Marshal(b, m, deterministic)
}
func (m *AdNetworkTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdNetworkTypeEnum.Merge(m, src)
}
func (m *AdNetworkTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdNetworkTypeEnum.Size(m)
}
func (m *AdNetworkTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdNetworkTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdNetworkTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdNetworkTypeEnum_AdNetworkType", AdNetworkTypeEnum_AdNetworkType_name, AdNetworkTypeEnum_AdNetworkType_value)
	proto.RegisterType((*AdNetworkTypeEnum)(nil), "google.ads.googleads.v3.enums.AdNetworkTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/ad_network_type.proto", fileDescriptor_a418102a00fe2b95)
}

var fileDescriptor_a418102a00fe2b95 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x6e, 0xe2, 0x30,
	0x18, 0x9d, 0x84, 0x01, 0x34, 0x46, 0x0c, 0xc1, 0xb3, 0x1b, 0x0d, 0x0b, 0x38, 0x80, 0xb3, 0xc8,
	0xce, 0xb3, 0x72, 0x82, 0x0b, 0xa8, 0xaa, 0x89, 0x20, 0x81, 0xb6, 0x8a, 0x14, 0xa5, 0x4d, 0x14,
	0xa1, 0x82, 0x1d, 0xe1, 0x40, 0xc5, 0x21, 0x7a, 0x89, 0x76, 0xd7, 0xa3, 0xf4, 0x22, 0x95, 0x7a,
	0x8a, 0x2a, 0x31, 0x41, 0x62, 0xd1, 0x6e, 0xac, 0xa7, 0xef, 0xfd, 0xc8, 0xef, 0x01, 0x2b, 0x15,
	0x22, 0x5d, 0x27, 0x66, 0x14, 0x4b, 0x53, 0xc1, 0x02, 0xed, 0x2d, 0x33, 0xe1, 0xbb, 0x8d, 0x34,
	0xa3, 0x38, 0xe4, 0x49, 0xfe, 0x28, 0xb6, 0x0f, 0x61, 0x7e, 0xc8, 0x12, 0x94, 0x6d, 0x45, 0x2e,
	0x60, 0x4f, 0x29, 0x51, 0x14, 0x4b, 0x74, 0x32, 0xa1, 0xbd, 0x85, 0x4a, 0xd3, 0xdf, 0x7f, 0x55,
	0x66, 0xb6, 0x32, 0x23, 0xce, 0x45, 0x1e, 0xe5, 0x2b, 0xc1, 0xa5, 0x32, 0x0f, 0x5e, 0x34, 0xd0,
	0x25, 0x31, 0x53, 0xa9, 0xde, 0x21, 0x4b, 0x28, 0xdf, 0x6d, 0x06, 0x4f, 0x1a, 0x68, 0x9f, 0x5d,
	0x61, 0x07, 0xb4, 0x7c, 0x36, 0x77, 0xa9, 0x33, 0xb9, 0x98, 0xd0, 0xa1, 0xf1, 0x03, 0xb6, 0x40,
	0xd3, 0x67, 0x97, 0x6c, 0xba, 0x64, 0x86, 0x06, 0x01, 0x68, 0xcc, 0x29, 0x99, 0x39, 0x63, 0x43,
	0x87, 0x7f, 0x40, 0x47, 0xe1, 0xd0, 0x25, 0x33, 0x8f, 0xd1, 0xd9, 0xdc, 0xa8, 0x15, 0x6a, 0x67,
	0xca, 0x3c, 0xca, 0x3c, 0xe3, 0x27, 0x84, 0xe0, 0xf7, 0xcd, 0xd4, 0xf7, 0x7c, 0x9b, 0x86, 0x47,
	0x57, 0x1d, 0x76, 0x41, 0xbb, 0xba, 0x2d, 0x89, 0xe7, 0x8c, 0x8d, 0x06, 0xfc, 0x05, 0xea, 0x57,
	0x93, 0x6b, 0x3a, 0x34, 0x9a, 0xf6, 0xbb, 0x06, 0xfa, 0xf7, 0x62, 0x83, 0xbe, 0x6d, 0x6a, 0xc3,
	0xb3, 0x2f, 0xbb, 0x45, 0x3f, 0x57, 0xbb, 0xb5, 0x8f, 0xa6, 0x54, 0xac, 0x23, 0x9e, 0x22, 0xb1,
	0x4d, 0xcd, 0x34, 0xe1, 0x65, 0xfb, 0x6a, 0xe3, 0x6c, 0x25, 0xbf, 0x98, 0xfc, 0x7f, 0xf9, 0x3e,
	0xeb, 0xb5, 0x11, 0x21, 0xaf, 0x7a, 0x6f, 0xa4, 0xa2, 0x48, 0x2c, 0x91, 0x82, 0x05, 0x5a, 0x58,
	0xa8, 0x18, 0x4d, 0xbe, 0x55, 0x7c, 0x40, 0x62, 0x19, 0x9c, 0xf8, 0x60, 0x61, 0x05, 0x25, 0xff,
	0xa1, 0xf7, 0xd5, 0x11, 0x63, 0x12, 0x4b, 0x8c, 0x4f, 0x0a, 0x8c, 0x17, 0x16, 0xc6, 0xa5, 0xe6,
	0xae, 0x51, 0x7e, 0xcc, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xda, 0xad, 0xab, 0x55, 0x0a, 0x02,
	0x00, 0x00,
}
