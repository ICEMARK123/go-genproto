// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/keyword_plan_campaign_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [KeywordPlanCampaignService.GetKeywordPlanCampaign][google.ads.googleads.v3.services.KeywordPlanCampaignService.GetKeywordPlanCampaign].
type GetKeywordPlanCampaignRequest struct {
	// Required. The resource name of the Keyword Plan campaign to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanCampaignRequest) Reset()         { *m = GetKeywordPlanCampaignRequest{} }
func (m *GetKeywordPlanCampaignRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanCampaignRequest) ProtoMessage()    {}
func (*GetKeywordPlanCampaignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f01e9912da6d3c86, []int{0}
}

func (m *GetKeywordPlanCampaignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanCampaignRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanCampaignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanCampaignRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordPlanCampaignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanCampaignRequest.Merge(m, src)
}
func (m *GetKeywordPlanCampaignRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanCampaignRequest.Size(m)
}
func (m *GetKeywordPlanCampaignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanCampaignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanCampaignRequest proto.InternalMessageInfo

func (m *GetKeywordPlanCampaignRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [KeywordPlanCampaignService.MutateKeywordPlanCampaigns][google.ads.googleads.v3.services.KeywordPlanCampaignService.MutateKeywordPlanCampaigns].
type MutateKeywordPlanCampaignsRequest struct {
	// Required. The ID of the customer whose Keyword Plan campaigns are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual Keyword Plan campaigns.
	Operations []*KeywordPlanCampaignOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanCampaignsRequest) Reset()         { *m = MutateKeywordPlanCampaignsRequest{} }
func (m *MutateKeywordPlanCampaignsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanCampaignsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanCampaignsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f01e9912da6d3c86, []int{1}
}

func (m *MutateKeywordPlanCampaignsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanCampaignsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanCampaignsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanCampaignsRequest.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanCampaignsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanCampaignsRequest.Merge(m, src)
}
func (m *MutateKeywordPlanCampaignsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanCampaignsRequest.Size(m)
}
func (m *MutateKeywordPlanCampaignsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanCampaignsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanCampaignsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanCampaignsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanCampaignsRequest) GetOperations() []*KeywordPlanCampaignOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateKeywordPlanCampaignsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateKeywordPlanCampaignsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a Keyword Plan campaign.
type KeywordPlanCampaignOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanCampaignOperation_Create
	//	*KeywordPlanCampaignOperation_Update
	//	*KeywordPlanCampaignOperation_Remove
	Operation            isKeywordPlanCampaignOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *KeywordPlanCampaignOperation) Reset()         { *m = KeywordPlanCampaignOperation{} }
func (m *KeywordPlanCampaignOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaignOperation) ProtoMessage()    {}
func (*KeywordPlanCampaignOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f01e9912da6d3c86, []int{2}
}

func (m *KeywordPlanCampaignOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaignOperation.Unmarshal(m, b)
}
func (m *KeywordPlanCampaignOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaignOperation.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaignOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaignOperation.Merge(m, src)
}
func (m *KeywordPlanCampaignOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaignOperation.Size(m)
}
func (m *KeywordPlanCampaignOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaignOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaignOperation proto.InternalMessageInfo

func (m *KeywordPlanCampaignOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanCampaignOperation_Operation interface {
	isKeywordPlanCampaignOperation_Operation()
}

type KeywordPlanCampaignOperation_Create struct {
	Create *resources.KeywordPlanCampaign `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanCampaignOperation_Update struct {
	Update *resources.KeywordPlanCampaign `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanCampaignOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanCampaignOperation_Create) isKeywordPlanCampaignOperation_Operation() {}

func (*KeywordPlanCampaignOperation_Update) isKeywordPlanCampaignOperation_Operation() {}

func (*KeywordPlanCampaignOperation_Remove) isKeywordPlanCampaignOperation_Operation() {}

func (m *KeywordPlanCampaignOperation) GetOperation() isKeywordPlanCampaignOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanCampaignOperation) GetCreate() *resources.KeywordPlanCampaign {
	if x, ok := m.GetOperation().(*KeywordPlanCampaignOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanCampaignOperation) GetUpdate() *resources.KeywordPlanCampaign {
	if x, ok := m.GetOperation().(*KeywordPlanCampaignOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanCampaignOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanCampaignOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanCampaignOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanCampaignOperation_Create)(nil),
		(*KeywordPlanCampaignOperation_Update)(nil),
		(*KeywordPlanCampaignOperation_Remove)(nil),
	}
}

// Response message for a Keyword Plan campaign mutate.
type MutateKeywordPlanCampaignsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateKeywordPlanCampaignResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *MutateKeywordPlanCampaignsResponse) Reset()         { *m = MutateKeywordPlanCampaignsResponse{} }
func (m *MutateKeywordPlanCampaignsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanCampaignsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanCampaignsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f01e9912da6d3c86, []int{3}
}

func (m *MutateKeywordPlanCampaignsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanCampaignsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanCampaignsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanCampaignsResponse.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanCampaignsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanCampaignsResponse.Merge(m, src)
}
func (m *MutateKeywordPlanCampaignsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanCampaignsResponse.Size(m)
}
func (m *MutateKeywordPlanCampaignsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanCampaignsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanCampaignsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanCampaignsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateKeywordPlanCampaignsResponse) GetResults() []*MutateKeywordPlanCampaignResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan campaign mutate.
type MutateKeywordPlanCampaignResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanCampaignResult) Reset()         { *m = MutateKeywordPlanCampaignResult{} }
func (m *MutateKeywordPlanCampaignResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanCampaignResult) ProtoMessage()    {}
func (*MutateKeywordPlanCampaignResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f01e9912da6d3c86, []int{4}
}

func (m *MutateKeywordPlanCampaignResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanCampaignResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanCampaignResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanCampaignResult.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanCampaignResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanCampaignResult.Merge(m, src)
}
func (m *MutateKeywordPlanCampaignResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanCampaignResult.Size(m)
}
func (m *MutateKeywordPlanCampaignResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanCampaignResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanCampaignResult proto.InternalMessageInfo

func (m *MutateKeywordPlanCampaignResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanCampaignRequest)(nil), "google.ads.googleads.v3.services.GetKeywordPlanCampaignRequest")
	proto.RegisterType((*MutateKeywordPlanCampaignsRequest)(nil), "google.ads.googleads.v3.services.MutateKeywordPlanCampaignsRequest")
	proto.RegisterType((*KeywordPlanCampaignOperation)(nil), "google.ads.googleads.v3.services.KeywordPlanCampaignOperation")
	proto.RegisterType((*MutateKeywordPlanCampaignsResponse)(nil), "google.ads.googleads.v3.services.MutateKeywordPlanCampaignsResponse")
	proto.RegisterType((*MutateKeywordPlanCampaignResult)(nil), "google.ads.googleads.v3.services.MutateKeywordPlanCampaignResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/keyword_plan_campaign_service.proto", fileDescriptor_f01e9912da6d3c86)
}

var fileDescriptor_f01e9912da6d3c86 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0xc7, 0x1f, 0x3b, 0x88, 0xe7, 0x61, 0x03, 0x4f, 0xa5, 0xad, 0x4a, 0x23, 0x97, 0x8a, 0xd4,
	0x45, 0x6a, 0x14, 0x55, 0xb6, 0x94, 0x48, 0xa8, 0x35, 0x82, 0xca, 0x81, 0x06, 0x50, 0x05, 0x44,
	0x46, 0xe2, 0xd0, 0xa6, 0xb2, 0x16, 0x7b, 0x49, 0x2d, 0x6c, 0xaf, 0xbb, 0x6b, 0xa7, 0x42, 0x88,
	0x4b, 0x0f, 0x7c, 0x81, 0x7e, 0x83, 0x1e, 0x38, 0xf4, 0x9b, 0x94, 0x6b, 0x6f, 0x9c, 0x5a, 0xa9,
	0xa7, 0x7e, 0x86, 0x1e, 0x2a, 0x7b, 0xbd, 0x79, 0x41, 0x09, 0x91, 0xca, 0x6d, 0xb2, 0xf3, 0xf7,
	0x6f, 0x66, 0x67, 0x66, 0x27, 0x60, 0xa3, 0x43, 0x48, 0xc7, 0xc7, 0x3a, 0x72, 0x99, 0xce, 0xcd,
	0xd4, 0xea, 0xd6, 0x75, 0x86, 0x69, 0xd7, 0x73, 0x30, 0xd3, 0x8f, 0xf1, 0xc9, 0x07, 0x42, 0x5d,
	0x3b, 0xf2, 0x51, 0x68, 0x3b, 0x28, 0x88, 0x90, 0xd7, 0x09, 0xed, 0xdc, 0xad, 0x45, 0x94, 0xc4,
	0x04, 0x96, 0xf9, 0xa7, 0x1a, 0x72, 0x99, 0xd6, 0xa3, 0x68, 0xdd, 0xba, 0x26, 0x28, 0xca, 0xea,
	0xb8, 0x38, 0x14, 0x33, 0x92, 0xd0, 0xb1, 0x81, 0x78, 0x00, 0x65, 0x41, 0x7c, 0x1e, 0x79, 0x3a,
	0x0a, 0x43, 0x12, 0xa3, 0xd8, 0x23, 0x21, 0xcb, 0xbd, 0xf7, 0x07, 0xbc, 0x8e, 0xef, 0xe1, 0x30,
	0xce, 0x1d, 0x8b, 0x03, 0x8e, 0x23, 0x0f, 0xfb, 0xae, 0x7d, 0x88, 0xdf, 0xa1, 0xae, 0x47, 0x68,
	0x2e, 0xc8, 0x13, 0xd7, 0xb3, 0x5f, 0x87, 0xc9, 0x51, 0xae, 0x0a, 0x10, 0x3b, 0xbe, 0xc6, 0xa6,
	0x91, 0xa3, 0xb3, 0x18, 0xc5, 0x49, 0x1e, 0x54, 0xdd, 0x06, 0x0f, 0x37, 0x71, 0xfc, 0x8a, 0x27,
	0xdd, 0xf2, 0x51, 0xb8, 0x9e, 0xa7, 0x6c, 0xe1, 0xf7, 0x09, 0x66, 0x31, 0xac, 0x80, 0x39, 0x71,
	0x39, 0x3b, 0x44, 0x01, 0x2e, 0x49, 0x65, 0xa9, 0x32, 0xd3, 0x28, 0x7c, 0x37, 0x65, 0x6b, 0x56,
	0x78, 0x76, 0x51, 0x80, 0xd5, 0xdf, 0x12, 0x78, 0xb4, 0x93, 0xc4, 0x28, 0xc6, 0x23, 0x70, 0x4c,
	0xf0, 0x96, 0x40, 0xd1, 0x49, 0x58, 0x4c, 0x02, 0x4c, 0x6d, 0xcf, 0x1d, 0xa4, 0x01, 0x71, 0xbe,
	0xed, 0x42, 0x07, 0x00, 0x12, 0x61, 0xca, 0xeb, 0x53, 0x92, 0xcb, 0x85, 0x4a, 0xb1, 0xb6, 0xa6,
	0x4d, 0xea, 0x8f, 0x36, 0x22, 0xf0, 0x9e, 0xc0, 0xe4, 0x41, 0xfa, 0x58, 0xf8, 0x04, 0xdc, 0x89,
	0x10, 0x8d, 0x3d, 0xe4, 0xdb, 0x47, 0xc8, 0xf3, 0x13, 0x8a, 0x4b, 0x85, 0xb2, 0x54, 0xf9, 0xcf,
	0xfa, 0x3f, 0x3f, 0x6e, 0xf2, 0x53, 0xf8, 0x18, 0xcc, 0x75, 0x91, 0xef, 0xb9, 0x28, 0xc6, 0x36,
	0x09, 0xfd, 0x93, 0xd2, 0x54, 0x26, 0x9b, 0x15, 0x87, 0x7b, 0xa1, 0x7f, 0xa2, 0x5e, 0xc8, 0x60,
	0xe1, 0xa6, 0xf8, 0x70, 0x05, 0x14, 0x93, 0x28, 0x63, 0xa4, 0x8d, 0xc9, 0x18, 0xc5, 0x9a, 0x22,
	0x2e, 0x25, 0x7a, 0xa7, 0x35, 0xd3, 0xde, 0xed, 0x20, 0x76, 0x6c, 0x01, 0x2e, 0x4f, 0x6d, 0xd8,
	0x02, 0xd3, 0x0e, 0xc5, 0x28, 0xe6, 0xf5, 0x2f, 0xd6, 0x96, 0xc7, 0x16, 0xa3, 0x37, 0x8a, 0xa3,
	0xaa, 0xb1, 0xf5, 0x8f, 0x95, 0x73, 0x52, 0x22, 0xe7, 0x97, 0xe4, 0xdb, 0x12, 0x39, 0x07, 0x96,
	0xc0, 0x34, 0xc5, 0x01, 0xe9, 0xf2, 0x32, 0xce, 0xa4, 0x1e, 0xfe, 0xbb, 0x51, 0x04, 0x33, 0xbd,
	0xba, 0xab, 0x5f, 0x25, 0xa0, 0xde, 0x34, 0x27, 0x2c, 0x22, 0x21, 0xc3, 0xb0, 0x09, 0xee, 0x5d,
	0xeb, 0x8e, 0x8d, 0x29, 0x25, 0x34, 0x83, 0x17, 0x6b, 0x50, 0xa4, 0x4b, 0x23, 0x47, 0xdb, 0xcf,
	0x46, 0xda, 0xba, 0x3b, 0xdc, 0xb7, 0x97, 0xa9, 0x1c, 0xbe, 0x01, 0xff, 0x52, 0xcc, 0x12, 0x3f,
	0x16, 0x73, 0x64, 0x4e, 0x9e, 0xa3, 0xb1, 0xe9, 0x59, 0x19, 0xc9, 0x12, 0x44, 0xb5, 0x09, 0x16,
	0x27, 0x68, 0xd3, 0xe1, 0x19, 0xf1, 0x80, 0x86, 0xdf, 0x4e, 0xed, 0x62, 0x0a, 0x28, 0x23, 0x10,
	0xfb, 0x3c, 0x21, 0xf8, 0x43, 0x02, 0xf3, 0xa3, 0x9f, 0x29, 0x7c, 0x31, 0xf9, 0x36, 0x37, 0x3e,
	0x70, 0xe5, 0x2f, 0xfb, 0xae, 0xee, 0x5e, 0x99, 0xc3, 0x17, 0xfb, 0xf8, 0xed, 0xe7, 0x27, 0xf9,
	0x19, 0x5c, 0x4e, 0xf7, 0xe1, 0xe9, 0x90, 0x67, 0x55, 0xbc, 0x6c, 0xa6, 0x57, 0xc5, 0x82, 0x1c,
	0x6a, 0xba, 0x5e, 0x3d, 0x83, 0xe7, 0x32, 0x50, 0xc6, 0x8f, 0x05, 0x5c, 0xbf, 0x45, 0xd7, 0xc4,
	0xf2, 0x51, 0x36, 0x6e, 0x07, 0xe1, 0x93, 0xa9, 0xbe, 0xbd, 0x32, 0xe7, 0x07, 0x76, 0xd8, 0xd3,
	0xfe, 0x4a, 0xc9, 0x4a, 0xb0, 0xa6, 0x3e, 0x4f, 0x4b, 0xd0, 0xbf, 0xf3, 0xe9, 0x80, 0x78, 0xb5,
	0x7a, 0x36, 0xb2, 0x02, 0x46, 0x90, 0xc5, 0x35, 0xa4, 0xaa, 0xf2, 0xe0, 0xd2, 0x2c, 0xf5, 0x73,
	0xcb, 0xad, 0xc8, 0x63, 0x9a, 0x43, 0x82, 0xc6, 0xb9, 0x0c, 0x96, 0x1c, 0x12, 0x4c, 0xbc, 0x47,
	0x63, 0x71, 0xfc, 0x38, 0xb5, 0xd2, 0x55, 0xd3, 0x92, 0x5e, 0x6f, 0xe5, 0x90, 0x0e, 0xf1, 0x51,
	0xd8, 0xd1, 0x08, 0xed, 0xe8, 0x1d, 0x1c, 0x66, 0x8b, 0x48, 0xef, 0x87, 0x1d, 0xff, 0xa7, 0xba,
	0x22, 0x8c, 0xcf, 0x72, 0x61, 0xd3, 0x34, 0xbf, 0xc8, 0xe5, 0x4d, 0x0e, 0x34, 0x5d, 0xa6, 0x71,
	0x33, 0xb5, 0x0e, 0xea, 0x5a, 0x1e, 0x98, 0x5d, 0x0a, 0x49, 0xdb, 0x74, 0x59, 0xbb, 0x27, 0x69,
	0x1f, 0xd4, 0xdb, 0x42, 0xf2, 0x4b, 0x5e, 0xe2, 0xe7, 0x86, 0x61, 0xba, 0xcc, 0x30, 0x7a, 0x22,
	0xc3, 0x38, 0xa8, 0x1b, 0x86, 0x90, 0x1d, 0x4e, 0x67, 0x79, 0xd6, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0xd4, 0x40, 0xb3, 0xfb, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanCampaignServiceClient is the client API for KeywordPlanCampaignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanCampaignServiceClient interface {
	// Returns the requested Keyword Plan campaign in full detail.
	GetKeywordPlanCampaign(ctx context.Context, in *GetKeywordPlanCampaignRequest, opts ...grpc.CallOption) (*resources.KeywordPlanCampaign, error)
	// Creates, updates, or removes Keyword Plan campaigns. Operation statuses are
	// returned.
	MutateKeywordPlanCampaigns(ctx context.Context, in *MutateKeywordPlanCampaignsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanCampaignsResponse, error)
}

type keywordPlanCampaignServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanCampaignServiceClient(cc grpc.ClientConnInterface) KeywordPlanCampaignServiceClient {
	return &keywordPlanCampaignServiceClient{cc}
}

func (c *keywordPlanCampaignServiceClient) GetKeywordPlanCampaign(ctx context.Context, in *GetKeywordPlanCampaignRequest, opts ...grpc.CallOption) (*resources.KeywordPlanCampaign, error) {
	out := new(resources.KeywordPlanCampaign)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordPlanCampaignService/GetKeywordPlanCampaign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanCampaignServiceClient) MutateKeywordPlanCampaigns(ctx context.Context, in *MutateKeywordPlanCampaignsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanCampaignsResponse, error) {
	out := new(MutateKeywordPlanCampaignsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordPlanCampaignService/MutateKeywordPlanCampaigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanCampaignServiceServer is the server API for KeywordPlanCampaignService service.
type KeywordPlanCampaignServiceServer interface {
	// Returns the requested Keyword Plan campaign in full detail.
	GetKeywordPlanCampaign(context.Context, *GetKeywordPlanCampaignRequest) (*resources.KeywordPlanCampaign, error)
	// Creates, updates, or removes Keyword Plan campaigns. Operation statuses are
	// returned.
	MutateKeywordPlanCampaigns(context.Context, *MutateKeywordPlanCampaignsRequest) (*MutateKeywordPlanCampaignsResponse, error)
}

// UnimplementedKeywordPlanCampaignServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanCampaignServiceServer struct {
}

func (*UnimplementedKeywordPlanCampaignServiceServer) GetKeywordPlanCampaign(ctx context.Context, req *GetKeywordPlanCampaignRequest) (*resources.KeywordPlanCampaign, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetKeywordPlanCampaign not implemented")
}
func (*UnimplementedKeywordPlanCampaignServiceServer) MutateKeywordPlanCampaigns(ctx context.Context, req *MutateKeywordPlanCampaignsRequest) (*MutateKeywordPlanCampaignsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateKeywordPlanCampaigns not implemented")
}

func RegisterKeywordPlanCampaignServiceServer(s *grpc.Server, srv KeywordPlanCampaignServiceServer) {
	s.RegisterService(&_KeywordPlanCampaignService_serviceDesc, srv)
}

func _KeywordPlanCampaignService_GetKeywordPlanCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanCampaignServiceServer).GetKeywordPlanCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordPlanCampaignService/GetKeywordPlanCampaign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanCampaignServiceServer).GetKeywordPlanCampaign(ctx, req.(*GetKeywordPlanCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanCampaignService_MutateKeywordPlanCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanCampaignServiceServer).MutateKeywordPlanCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordPlanCampaignService/MutateKeywordPlanCampaigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanCampaignServiceServer).MutateKeywordPlanCampaigns(ctx, req.(*MutateKeywordPlanCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanCampaignService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.KeywordPlanCampaignService",
	HandlerType: (*KeywordPlanCampaignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanCampaign",
			Handler:    _KeywordPlanCampaignService_GetKeywordPlanCampaign_Handler,
		},
		{
			MethodName: "MutateKeywordPlanCampaigns",
			Handler:    _KeywordPlanCampaignService_MutateKeywordPlanCampaigns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/keyword_plan_campaign_service.proto",
}
