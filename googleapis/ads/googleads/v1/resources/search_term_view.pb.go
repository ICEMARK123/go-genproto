// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/search_term_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A search term view with metrics aggregated by search term at the ad group
// level.
type SearchTermView struct {
	// Immutable. The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/searchTermViews/{campaign_id}~{ad_group_id}~{URL-base64_search_term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The search term.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// Output only. The ad group the search term served in.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Output only. Indicates whether the search term is currently one of your
	// targeted or excluded keywords.
	Status               enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *SearchTermView) Reset()         { *m = SearchTermView{} }
func (m *SearchTermView) String() string { return proto.CompactTextString(m) }
func (*SearchTermView) ProtoMessage()    {}
func (*SearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d7523262f43a04, []int{0}
}

func (m *SearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermView.Unmarshal(m, b)
}
func (m *SearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermView.Marshal(b, m, deterministic)
}
func (m *SearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermView.Merge(m, src)
}
func (m *SearchTermView) XXX_Size() int {
	return xxx_messageInfo_SearchTermView.Size(m)
}
func (m *SearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermView proto.InternalMessageInfo

func (m *SearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *SearchTermView) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *SearchTermView) GetStatus() enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus {
	if m != nil {
		return m.Status
	}
	return enums.SearchTermTargetingStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SearchTermView)(nil), "google.ads.googleads.v1.resources.SearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/search_term_view.proto", fileDescriptor_13d7523262f43a04)
}

var fileDescriptor_13d7523262f43a04 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6a, 0xdb, 0x4c,
	0x10, 0xc7, 0x91, 0xf4, 0x7d, 0x69, 0xbb, 0x69, 0x73, 0x50, 0x2f, 0x6e, 0x08, 0xad, 0x13, 0x08,
	0x75, 0x2f, 0xbb, 0x28, 0x3d, 0xb4, 0xa8, 0x97, 0xae, 0xa0, 0x18, 0x7a, 0x28, 0xc1, 0x36, 0x3a,
	0x14, 0x83, 0x58, 0x4b, 0x93, 0x8d, 0xc0, 0xda, 0x15, 0xbb, 0x2b, 0xfb, 0x10, 0x02, 0x7d, 0x96,
	0x1e, 0xfb, 0x28, 0x7d, 0x0a, 0x5f, 0x9b, 0x47, 0xc8, 0xa9, 0x58, 0xd2, 0xca, 0x76, 0xc1, 0x6d,
	0x6e, 0x7f, 0x31, 0xff, 0xf9, 0xcd, 0xec, 0xcc, 0x08, 0xbd, 0xe7, 0x52, 0xf2, 0x39, 0x10, 0x96,
	0x69, 0xd2, 0xc8, 0xb5, 0x5a, 0x04, 0x44, 0x81, 0x96, 0x95, 0x4a, 0x41, 0x13, 0x0d, 0x4c, 0xa5,
	0xd7, 0x89, 0x01, 0x55, 0x24, 0x8b, 0x1c, 0x96, 0xb8, 0x54, 0xd2, 0x48, 0xff, 0xb4, 0xb1, 0x63,
	0x96, 0x69, 0xdc, 0x65, 0xe2, 0x45, 0x80, 0xbb, 0xcc, 0xe3, 0x8f, 0xfb, 0xe0, 0x20, 0xaa, 0x62,
	0x17, 0x6c, 0x98, 0xe2, 0x60, 0x72, 0xc1, 0x13, 0x6d, 0x98, 0xa9, 0x74, 0x53, 0xe4, 0xf8, 0x95,
	0x25, 0x94, 0x39, 0xb9, 0xca, 0x61, 0x9e, 0x25, 0x33, 0xb8, 0x66, 0x8b, 0x5c, 0xaa, 0xd6, 0xf0,
	0x62, 0xcb, 0x60, 0x0b, 0xb7, 0xa1, 0x97, 0x6d, 0xa8, 0xfe, 0x9a, 0x55, 0x57, 0x64, 0xa9, 0x58,
	0x59, 0x82, 0xb2, 0xec, 0x93, 0xad, 0x54, 0x26, 0x84, 0x34, 0xcc, 0xe4, 0x52, 0xb4, 0xd1, 0xb3,
	0x5f, 0x1e, 0x3a, 0x1a, 0xd7, 0x0d, 0x4e, 0x40, 0x15, 0x71, 0x0e, 0x4b, 0x7f, 0x82, 0x9e, 0xd9,
	0x12, 0x89, 0x60, 0x05, 0xf4, 0x9c, 0xbe, 0x33, 0x78, 0x12, 0x91, 0x15, 0xfd, 0xff, 0x9e, 0xbe,
	0x41, 0xaf, 0x37, 0x53, 0x68, 0x55, 0x99, 0x6b, 0x9c, 0xca, 0x82, 0xec, 0x72, 0x46, 0x4f, 0x2d,
	0xe5, 0x0b, 0x2b, 0xc0, 0x8f, 0xd0, 0xe1, 0xd6, 0x20, 0x7a, 0x6e, 0xdf, 0x19, 0x1c, 0x5e, 0x9c,
	0xb4, 0x08, 0x6c, 0x9b, 0xc7, 0x63, 0xa3, 0x72, 0xc1, 0x63, 0x36, 0xaf, 0x20, 0xf2, 0x56, 0xd4,
	0x1b, 0x21, 0xdd, 0x51, 0x7d, 0x86, 0x1e, 0xb3, 0x2c, 0xe1, 0x4a, 0x56, 0x65, 0xcf, 0x7b, 0x00,
	0x60, 0xb0, 0xa2, 0xde, 0x3d, 0x3d, 0x43, 0xfd, 0xbd, 0x2d, 0xd3, 0x6c, 0xb8, 0xa6, 0x8d, 0x1e,
	0xb1, 0x46, 0xf8, 0x1a, 0x1d, 0x34, 0x9b, 0xe9, 0xfd, 0xd7, 0x77, 0x06, 0x47, 0x17, 0x31, 0xde,
	0xb7, 0xff, 0x7a, 0xb9, 0x78, 0xf3, 0xe6, 0x89, 0x5d, 0xed, 0xb8, 0xce, 0xff, 0x24, 0xaa, 0x62,
	0x7f, 0xb4, 0x79, 0x5b, 0x5b, 0x2a, 0x84, 0x3b, 0x3a, 0x7b, 0xf0, 0x5c, 0xfd, 0x77, 0x69, 0xa5,
	0x8d, 0x2c, 0x40, 0x69, 0x72, 0x63, 0xe5, 0x2d, 0xd1, 0x3b, 0x26, 0x4d, 0x6e, 0xfe, 0xbc, 0xe7,
	0xdb, 0xe8, 0x9b, 0x8b, 0xce, 0x53, 0x59, 0xe0, 0x7f, 0x5e, 0x74, 0xf4, 0x7c, 0xb7, 0xe4, 0xe5,
	0x7a, 0xb8, 0x97, 0xce, 0xd7, 0xcf, 0x6d, 0x26, 0x97, 0x73, 0x26, 0x38, 0x96, 0x8a, 0x13, 0x0e,
	0xa2, 0x1e, 0x3d, 0xd9, 0xf4, 0xfc, 0x97, 0x9f, 0xec, 0x43, 0xa7, 0xbe, 0xbb, 0xde, 0x90, 0xd2,
	0x1f, 0xee, 0xe9, 0xb0, 0x41, 0xd2, 0x4c, 0xe3, 0x46, 0xae, 0x55, 0x1c, 0xe0, 0x91, 0x75, 0xfe,
	0xb4, 0x9e, 0x29, 0xcd, 0xf4, 0xb4, 0xf3, 0x4c, 0xe3, 0x60, 0xda, 0x79, 0xee, 0xdc, 0xf3, 0x26,
	0x10, 0x86, 0x34, 0xd3, 0x61, 0xd8, 0xb9, 0xc2, 0x30, 0x0e, 0xc2, 0xb0, 0xf3, 0xcd, 0x0e, 0xea,
	0x66, 0xdf, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x79, 0x1f, 0xbd, 0x8b, 0x10, 0x04, 0x00, 0x00,
}
