// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/geographic_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A geographic view.
//
// Geographic View includes all metrics aggregated at the country level,
// one row per country. It reports metrics at either actual physical location of
// the user or an area of interest. If other segment fields are used, you may
// get more than one row per country.
type GeographicView struct {
	// The resource name of the geographic view.
	// Geographic view resource names have the form:
	//
	// `customers/{customer_id}/geographicViews/{country_criterion_id}~{location_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Type of the geo targeting of the campaign.
	LocationType enums.GeoTargetingTypeEnum_GeoTargetingType `protobuf:"varint,3,opt,name=location_type,json=locationType,proto3,enum=google.ads.googleads.v3.enums.GeoTargetingTypeEnum_GeoTargetingType" json:"location_type,omitempty"`
	// Criterion Id for the country.
	CountryCriterionId   *wrappers.Int64Value `protobuf:"bytes,4,opt,name=country_criterion_id,json=countryCriterionId,proto3" json:"country_criterion_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GeographicView) Reset()         { *m = GeographicView{} }
func (m *GeographicView) String() string { return proto.CompactTextString(m) }
func (*GeographicView) ProtoMessage()    {}
func (*GeographicView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9d21ad3e852168, []int{0}
}

func (m *GeographicView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeographicView.Unmarshal(m, b)
}
func (m *GeographicView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeographicView.Marshal(b, m, deterministic)
}
func (m *GeographicView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeographicView.Merge(m, src)
}
func (m *GeographicView) XXX_Size() int {
	return xxx_messageInfo_GeographicView.Size(m)
}
func (m *GeographicView) XXX_DiscardUnknown() {
	xxx_messageInfo_GeographicView.DiscardUnknown(m)
}

var xxx_messageInfo_GeographicView proto.InternalMessageInfo

func (m *GeographicView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GeographicView) GetLocationType() enums.GeoTargetingTypeEnum_GeoTargetingType {
	if m != nil {
		return m.LocationType
	}
	return enums.GeoTargetingTypeEnum_UNSPECIFIED
}

func (m *GeographicView) GetCountryCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CountryCriterionId
	}
	return nil
}

func init() {
	proto.RegisterType((*GeographicView)(nil), "google.ads.googleads.v3.resources.GeographicView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/geographic_view.proto", fileDescriptor_9e9d21ad3e852168)
}

var fileDescriptor_9e9d21ad3e852168 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x65, 0xa6, 0x22, 0x38, 0xb6, 0x7d, 0x18, 0x7d, 0x58, 0xab, 0xc8, 0x56, 0x29, 0xee, 0x53,
	0x02, 0x1d, 0x59, 0x21, 0x3e, 0x4d, 0x55, 0x96, 0x0a, 0x4a, 0x59, 0xca, 0x3c, 0xc8, 0xc2, 0x90,
	0xce, 0x5c, 0x63, 0x60, 0x27, 0x77, 0x48, 0x32, 0xbb, 0x2c, 0xa5, 0xe0, 0x6f, 0xf1, 0xd1, 0x9f,
	0xe2, 0xbb, 0x7f, 0xa2, 0xbf, 0x42, 0xe6, 0x23, 0x69, 0xad, 0xd4, 0xbe, 0x9d, 0xe4, 0x9e, 0x73,
	0xee, 0x3d, 0xc9, 0x8d, 0xde, 0x08, 0x44, 0xb1, 0x04, 0xca, 0x4b, 0x43, 0x7b, 0xd8, 0xa2, 0x55,
	0x42, 0x35, 0x18, 0x6c, 0x74, 0x01, 0x86, 0x0a, 0x40, 0xa1, 0x79, 0xfd, 0x4d, 0x16, 0xf9, 0x4a,
	0xc2, 0x9a, 0xd4, 0x1a, 0x2d, 0xc6, 0xfb, 0x3d, 0x9b, 0xf0, 0xd2, 0x10, 0x2f, 0x24, 0xab, 0x84,
	0x78, 0xe1, 0xde, 0xf4, 0x36, 0x6f, 0x50, 0x4d, 0xd5, 0xf9, 0xe6, 0x96, 0x6b, 0x01, 0x56, 0x2a,
	0x91, 0xdb, 0x4d, 0x0d, 0xbd, 0xf5, 0xde, 0x13, 0xa7, 0xab, 0xa5, 0x1f, 0x63, 0x28, 0x3d, 0x1f,
	0x4a, 0xdd, 0xe9, 0xac, 0xf9, 0x4a, 0xd7, 0x9a, 0xd7, 0x35, 0x68, 0x33, 0xd4, 0x9f, 0x5d, 0x93,
	0x72, 0xa5, 0xd0, 0x72, 0x2b, 0x51, 0x0d, 0xd5, 0x17, 0xbf, 0xc3, 0x68, 0x77, 0xe6, 0xd3, 0x64,
	0x12, 0xd6, 0xf1, 0xcb, 0x68, 0xc7, 0xb5, 0xc8, 0x15, 0xaf, 0x60, 0x14, 0x8c, 0x83, 0xc9, 0x83,
	0xf9, 0xb6, 0xbb, 0xfc, 0xcc, 0x2b, 0x88, 0x65, 0xb4, 0xb3, 0xc4, 0xa2, 0xb3, 0xea, 0xe6, 0x1c,
	0x6d, 0x8d, 0x83, 0xc9, 0xee, 0xe1, 0x7b, 0x72, 0xdb, 0x1b, 0x74, 0x01, 0xc9, 0x0c, 0xf0, 0xd4,
	0xe5, 0x3b, 0xdd, 0xd4, 0xf0, 0x41, 0x35, 0xd5, 0x3f, 0x97, 0xf3, 0x6d, 0x67, 0xdd, 0x9e, 0xe2,
	0x4f, 0xd1, 0xe3, 0x02, 0x1b, 0x65, 0xf5, 0x26, 0x2f, 0xb4, 0xb4, 0xa0, 0xdb, 0x9e, 0xb2, 0x1c,
	0xdd, 0x1b, 0x07, 0x93, 0x87, 0x87, 0x4f, 0x5d, 0x47, 0x97, 0x9f, 0x1c, 0x2b, 0x3b, 0x7d, 0x9d,
	0xf1, 0x65, 0x03, 0xf3, 0x78, 0x10, 0xbe, 0x73, 0xba, 0xe3, 0x92, 0x95, 0x97, 0x29, 0x8f, 0x5e,
	0x5d, 0xcd, 0x36, 0xa0, 0x5a, 0x1a, 0x52, 0x60, 0x45, 0x6f, 0x3c, 0xc6, 0xb4, 0x68, 0x8c, 0xc5,
	0x0a, 0xb4, 0xa1, 0xe7, 0x0e, 0x5e, 0x5c, 0xfb, 0xff, 0x96, 0x64, 0xe8, 0xf9, 0x8d, 0x85, 0xb8,
	0x38, 0xfa, 0x1e, 0x46, 0x07, 0x05, 0x56, 0xe4, 0xce, 0x95, 0x38, 0x7a, 0xf4, 0x77, 0xc7, 0x93,
	0x36, 0xc6, 0x49, 0xf0, 0xe5, 0xe3, 0xa0, 0x14, 0xb8, 0xe4, 0x4a, 0x10, 0xd4, 0x82, 0x0a, 0x50,
	0x5d, 0x48, 0x7a, 0x35, 0xf2, 0x7f, 0x96, 0xf4, 0xad, 0x47, 0x3f, 0xc2, 0xad, 0x59, 0x9a, 0xfe,
	0x0c, 0xf7, 0x67, 0xbd, 0x65, 0x5a, 0x1a, 0xd2, 0xc3, 0x16, 0x65, 0x09, 0x99, 0x3b, 0xe6, 0x2f,
	0xc7, 0x59, 0xa4, 0xa5, 0x59, 0x78, 0xce, 0x22, 0x4b, 0x16, 0x9e, 0x73, 0x19, 0x1e, 0xf4, 0x05,
	0xc6, 0xd2, 0xd2, 0x30, 0xe6, 0x59, 0x8c, 0x65, 0x09, 0x63, 0x9e, 0x77, 0x76, 0xbf, 0x1b, 0x36,
	0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd7, 0x44, 0x65, 0x50, 0x03, 0x00, 0x00,
}
