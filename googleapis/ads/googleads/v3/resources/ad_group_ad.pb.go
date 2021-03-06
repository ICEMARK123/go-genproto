// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_group_ad.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
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

// An ad group ad.
type AdGroupAd struct {
	// The resource name of the ad.
	// Ad group ad resource names have the form:
	//
	// `customers/{customer_id}/adGroupAds/{ad_group_id}~{ad_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The status of the ad.
	Status enums.AdGroupAdStatusEnum_AdGroupAdStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.AdGroupAdStatusEnum_AdGroupAdStatus" json:"status,omitempty"`
	// The ad group to which the ad belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,4,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The ad.
	Ad *Ad `protobuf:"bytes,5,opt,name=ad,proto3" json:"ad,omitempty"`
	// Policy information for the ad.
	PolicySummary *AdGroupAdPolicySummary `protobuf:"bytes,6,opt,name=policy_summary,json=policySummary,proto3" json:"policy_summary,omitempty"`
	// Overall ad strength for this ad group ad.
	AdStrength           enums.AdStrengthEnum_AdStrength `protobuf:"varint,7,opt,name=ad_strength,json=adStrength,proto3,enum=google.ads.googleads.v3.enums.AdStrengthEnum_AdStrength" json:"ad_strength,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AdGroupAd) Reset()         { *m = AdGroupAd{} }
func (m *AdGroupAd) String() string { return proto.CompactTextString(m) }
func (*AdGroupAd) ProtoMessage()    {}
func (*AdGroupAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ed4b3c949da8cd1, []int{0}
}

func (m *AdGroupAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAd.Unmarshal(m, b)
}
func (m *AdGroupAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAd.Marshal(b, m, deterministic)
}
func (m *AdGroupAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAd.Merge(m, src)
}
func (m *AdGroupAd) XXX_Size() int {
	return xxx_messageInfo_AdGroupAd.Size(m)
}
func (m *AdGroupAd) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAd.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAd proto.InternalMessageInfo

func (m *AdGroupAd) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupAd) GetStatus() enums.AdGroupAdStatusEnum_AdGroupAdStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupAdStatusEnum_UNSPECIFIED
}

func (m *AdGroupAd) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupAd) GetAd() *Ad {
	if m != nil {
		return m.Ad
	}
	return nil
}

func (m *AdGroupAd) GetPolicySummary() *AdGroupAdPolicySummary {
	if m != nil {
		return m.PolicySummary
	}
	return nil
}

func (m *AdGroupAd) GetAdStrength() enums.AdStrengthEnum_AdStrength {
	if m != nil {
		return m.AdStrength
	}
	return enums.AdStrengthEnum_UNSPECIFIED
}

// Contains policy information for an ad.
type AdGroupAdPolicySummary struct {
	// The list of policy findings for this ad.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Where in the review process this ad is.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v3.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// The overall approval status of this ad, calculated based on the status of
	// its individual policy topic entries.
	ApprovalStatus       enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v3.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *AdGroupAdPolicySummary) Reset()         { *m = AdGroupAdPolicySummary{} }
func (m *AdGroupAdPolicySummary) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdPolicySummary) ProtoMessage()    {}
func (*AdGroupAdPolicySummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ed4b3c949da8cd1, []int{1}
}

func (m *AdGroupAdPolicySummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdPolicySummary.Unmarshal(m, b)
}
func (m *AdGroupAdPolicySummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdPolicySummary.Marshal(b, m, deterministic)
}
func (m *AdGroupAdPolicySummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdPolicySummary.Merge(m, src)
}
func (m *AdGroupAdPolicySummary) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdPolicySummary.Size(m)
}
func (m *AdGroupAdPolicySummary) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdPolicySummary.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdPolicySummary proto.InternalMessageInfo

func (m *AdGroupAdPolicySummary) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if m != nil {
		return m.PolicyTopicEntries
	}
	return nil
}

func (m *AdGroupAdPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if m != nil {
		return m.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_UNSPECIFIED
}

func (m *AdGroupAdPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if m != nil {
		return m.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupAd)(nil), "google.ads.googleads.v3.resources.AdGroupAd")
	proto.RegisterType((*AdGroupAdPolicySummary)(nil), "google.ads.googleads.v3.resources.AdGroupAdPolicySummary")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_group_ad.proto", fileDescriptor_1ed4b3c949da8cd1)
}

var fileDescriptor_1ed4b3c949da8cd1 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6a, 0xd4, 0x4c,
	0x14, 0x26, 0x69, 0xff, 0xf6, 0xef, 0xd4, 0xae, 0x10, 0x44, 0x62, 0x29, 0xb2, 0xad, 0x14, 0x16,
	0xc5, 0x19, 0x69, 0x50, 0x6b, 0xbc, 0x4a, 0xa1, 0x54, 0x44, 0xa4, 0x64, 0x65, 0xc1, 0xb2, 0xb0,
	0x4e, 0x93, 0x31, 0x06, 0x36, 0x33, 0xc3, 0xcc, 0x64, 0xcb, 0x52, 0xfa, 0x10, 0xbe, 0x82, 0x97,
	0x3e, 0x8a, 0x8f, 0xe1, 0x65, 0x5f, 0xc1, 0x1b, 0x49, 0x66, 0x26, 0xbb, 0xdb, 0xba, 0xdd, 0xbd,
	0x3b, 0x33, 0xe7, 0xfb, 0xbe, 0xc9, 0xf9, 0xce, 0x39, 0x01, 0x41, 0xc6, 0x58, 0x36, 0x24, 0x08,
	0xa7, 0x12, 0xe9, 0xb0, 0x8a, 0x46, 0x01, 0x12, 0x44, 0xb2, 0x52, 0x24, 0x44, 0x22, 0x9c, 0x0e,
	0x32, 0xc1, 0x4a, 0x3e, 0xc0, 0x29, 0xe4, 0x82, 0x29, 0xe6, 0xed, 0x6a, 0x24, 0xc4, 0xa9, 0x84,
	0x0d, 0x09, 0x8e, 0x02, 0xd8, 0x90, 0xb6, 0x9f, 0xcd, 0xd3, 0x4d, 0x58, 0x51, 0x30, 0x8a, 0x38,
	0x1b, 0xe6, 0xc9, 0x58, 0xeb, 0x6d, 0xbf, 0x9a, 0x07, 0x26, 0xb4, 0x2c, 0x66, 0x3e, 0x60, 0x20,
	0x15, 0x56, 0xa5, 0x34, 0x3c, 0xb4, 0x90, 0x27, 0x95, 0x20, 0x34, 0x53, 0xdf, 0x0c, 0x21, 0xbc,
	0x9b, 0xa0, 0x3f, 0x6a, 0x80, 0x39, 0x17, 0x6c, 0x84, 0x87, 0xb3, 0x8f, 0x1d, 0x2e, 0xc5, 0x15,
	0x64, 0x94, 0x93, 0x8b, 0x59, 0xe6, 0xd3, 0x65, 0x3c, 0x36, 0xd8, 0x47, 0x16, 0xcb, 0xf3, 0x26,
	0x6d, 0x52, 0x8f, 0x4d, 0xaa, 0x3e, 0x9d, 0x97, 0x5f, 0xd1, 0x85, 0xc0, 0x9c, 0x13, 0x61, 0x9f,
	0xd9, 0x99, 0xa2, 0x62, 0x4a, 0x99, 0xc2, 0x2a, 0x67, 0xd4, 0x64, 0xf7, 0xbe, 0xaf, 0x82, 0x8d,
	0x28, 0x3d, 0xa9, 0x7c, 0x8c, 0x52, 0xef, 0x09, 0xd8, 0xb2, 0xea, 0x03, 0x8a, 0x0b, 0xe2, 0x3b,
	0x6d, 0xa7, 0xb3, 0x11, 0xdf, 0xb3, 0x97, 0x1f, 0x71, 0x41, 0xbc, 0x33, 0xb0, 0xa6, 0xeb, 0xf0,
	0x57, 0xda, 0x4e, 0xa7, 0x75, 0x70, 0x04, 0xe7, 0xf5, 0xbd, 0xb6, 0x00, 0x36, 0xf2, 0xdd, 0x9a,
	0x75, 0x4c, 0xcb, 0xe2, 0xe6, 0x5d, 0x6c, 0x14, 0xbd, 0xd7, 0xe0, 0x7f, 0xdb, 0x56, 0x7f, 0xb5,
	0xed, 0x74, 0x36, 0x0f, 0x76, 0xac, 0xba, 0xad, 0x0f, 0x76, 0x95, 0xc8, 0x69, 0xd6, 0xc3, 0xc3,
	0x92, 0xc4, 0xeb, 0x58, 0x0b, 0x79, 0x2f, 0x81, 0x8b, 0x53, 0xff, 0xbf, 0x9a, 0xb2, 0x0f, 0x17,
	0x0e, 0x22, 0x8c, 0xd2, 0xd8, 0xc5, 0xa9, 0xf7, 0x05, 0xb4, 0x4c, 0x87, 0x64, 0x59, 0x14, 0x58,
	0x8c, 0xfd, 0xb5, 0x5a, 0xe2, 0xcd, 0x52, 0x12, 0xa6, 0x86, 0xd3, 0x5a, 0xa1, 0xab, 0x05, 0xe2,
	0x2d, 0x3e, 0x7d, 0xf4, 0x3e, 0x83, 0xcd, 0xa9, 0x81, 0xf3, 0xd7, 0x6b, 0xcb, 0x0e, 0x17, 0x5a,
	0xd6, 0x35, 0x04, 0xe3, 0x96, 0x3d, 0xc6, 0x00, 0x37, 0x71, 0xd8, 0xbb, 0x8e, 0xba, 0x60, 0x6f,
	0x42, 0x37, 0x11, 0xcf, 0x25, 0x4c, 0x58, 0x81, 0x26, 0x6d, 0x7d, 0x9e, 0x94, 0x52, 0xb1, 0x82,
	0x08, 0x89, 0x2e, 0x6d, 0x78, 0x85, 0xb0, 0xcd, 0x4b, 0x74, 0x39, 0xb5, 0x4b, 0x57, 0x7b, 0xbf,
	0x5d, 0xf0, 0xf0, 0xdf, 0xc5, 0x79, 0xe7, 0xe0, 0x81, 0xf1, 0x4b, 0x31, 0x9e, 0x27, 0x03, 0x42,
	0x95, 0xc8, 0x89, 0xf4, 0x9d, 0xf6, 0x4a, 0x67, 0xf3, 0xe0, 0xc5, 0xdc, 0xb2, 0xf4, 0x7a, 0x43,
	0x2d, 0xf6, 0xa9, 0xa2, 0x1e, 0x53, 0x25, 0xc6, 0xb1, 0xc7, 0x67, 0x6f, 0x72, 0x22, 0xbd, 0xa2,
	0x1a, 0xc2, 0xa9, 0x75, 0xf1, 0xdd, 0xda, 0xb3, 0x77, 0x0b, 0x3c, 0xd3, 0xda, 0x71, 0xcd, 0x9c,
	0x9a, 0xb4, 0xdb, 0xd7, 0xd5, 0x38, 0x4f, 0x4e, 0x5e, 0x09, 0xee, 0xdf, 0xd8, 0x6c, 0x33, 0xd7,
	0x1f, 0x96, 0x7a, 0x30, 0x32, 0xdc, 0x5b, 0x4f, 0xce, 0x26, 0xe2, 0x16, 0x9e, 0x39, 0x1f, 0xfd,
	0x71, 0xc0, 0x7e, 0xc2, 0x8a, 0xc5, 0x73, 0x76, 0xd4, 0x9a, 0xf4, 0xa2, 0x5a, 0x81, 0x53, 0xe7,
	0xec, 0xbd, 0x21, 0x65, 0x6c, 0x88, 0x69, 0x06, 0x99, 0xc8, 0x50, 0x46, 0x68, 0xbd, 0x20, 0x68,
	0x32, 0x04, 0x77, 0xfc, 0x58, 0xde, 0x36, 0xd1, 0x0f, 0x77, 0xe5, 0x24, 0x8a, 0x7e, 0xba, 0xbb,
	0x27, 0x5a, 0x32, 0x4a, 0x25, 0xd4, 0x61, 0x15, 0xf5, 0x02, 0x18, 0x5b, 0xe4, 0x2f, 0x8b, 0xe9,
	0x47, 0xa9, 0xec, 0x37, 0x98, 0x7e, 0x2f, 0xe8, 0x37, 0x98, 0x6b, 0x77, 0x5f, 0x27, 0xc2, 0x30,
	0x4a, 0x65, 0x18, 0x36, 0xa8, 0x30, 0xec, 0x05, 0x61, 0xd8, 0xe0, 0xce, 0xd7, 0xea, 0x8f, 0x0d,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x23, 0x16, 0x17, 0x68, 0x06, 0x00, 0x00,
}
