// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/keyword_plan_negative_keyword_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for
// [KeywordPlanNegativeKeywordService.GetKeywordPlanNegativeKeyword][google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService.GetKeywordPlanNegativeKeyword].
type GetKeywordPlanNegativeKeywordRequest struct {
	// Required. The resource name of the plan to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanNegativeKeywordRequest) Reset()         { *m = GetKeywordPlanNegativeKeywordRequest{} }
func (m *GetKeywordPlanNegativeKeywordRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanNegativeKeywordRequest) ProtoMessage()    {}
func (*GetKeywordPlanNegativeKeywordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6100cff1a3faa6ff, []int{0}
}

func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Merge(m, src)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.Size(m)
}
func (m *GetKeywordPlanNegativeKeywordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanNegativeKeywordRequest proto.InternalMessageInfo

func (m *GetKeywordPlanNegativeKeywordRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [KeywordPlanNegativeKeywordService.MutateKeywordPlanNegativeKeywords][google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService.MutateKeywordPlanNegativeKeywords].
type MutateKeywordPlanNegativeKeywordsRequest struct {
	// Required. The ID of the customer whose negative keywords are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual Keyword Plan negative
	// keywords.
	Operations []*KeywordPlanNegativeKeywordOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateKeywordPlanNegativeKeywordsRequest) Reset() {
	*m = MutateKeywordPlanNegativeKeywordsRequest{}
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6100cff1a3faa6ff, []int{1}
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Merge(m, src)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetOperations() []*KeywordPlanNegativeKeywordOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateKeywordPlanNegativeKeywordsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a Keyword Plan negative
// keyword.
type KeywordPlanNegativeKeywordOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanNegativeKeywordOperation_Create
	//	*KeywordPlanNegativeKeywordOperation_Update
	//	*KeywordPlanNegativeKeywordOperation_Remove
	Operation            isKeywordPlanNegativeKeywordOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *KeywordPlanNegativeKeywordOperation) Reset()         { *m = KeywordPlanNegativeKeywordOperation{} }
func (m *KeywordPlanNegativeKeywordOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNegativeKeywordOperation) ProtoMessage()    {}
func (*KeywordPlanNegativeKeywordOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6100cff1a3faa6ff, []int{2}
}

func (m *KeywordPlanNegativeKeywordOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Unmarshal(m, b)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Merge(m, src)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNegativeKeywordOperation.Size(m)
}
func (m *KeywordPlanNegativeKeywordOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNegativeKeywordOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNegativeKeywordOperation proto.InternalMessageInfo

func (m *KeywordPlanNegativeKeywordOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanNegativeKeywordOperation_Operation interface {
	isKeywordPlanNegativeKeywordOperation_Operation()
}

type KeywordPlanNegativeKeywordOperation_Create struct {
	Create *resources.KeywordPlanNegativeKeyword `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanNegativeKeywordOperation_Update struct {
	Update *resources.KeywordPlanNegativeKeyword `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanNegativeKeywordOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanNegativeKeywordOperation_Create) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (*KeywordPlanNegativeKeywordOperation_Update) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (*KeywordPlanNegativeKeywordOperation_Remove) isKeywordPlanNegativeKeywordOperation_Operation() {}

func (m *KeywordPlanNegativeKeywordOperation) GetOperation() isKeywordPlanNegativeKeywordOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetCreate() *resources.KeywordPlanNegativeKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetUpdate() *resources.KeywordPlanNegativeKeyword {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanNegativeKeywordOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanNegativeKeywordOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanNegativeKeywordOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanNegativeKeywordOperation_Create)(nil),
		(*KeywordPlanNegativeKeywordOperation_Update)(nil),
		(*KeywordPlanNegativeKeywordOperation_Remove)(nil),
	}
}

// Response message for a Keyword Plan negative keyword mutate.
type MutateKeywordPlanNegativeKeywordsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateKeywordPlanNegativeKeywordResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) Reset() {
	*m = MutateKeywordPlanNegativeKeywordsResponse{}
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6100cff1a3faa6ff, []int{3}
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Merge(m, src)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateKeywordPlanNegativeKeywordsResponse) GetResults() []*MutateKeywordPlanNegativeKeywordResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan negative keyword mutate.
type MutateKeywordPlanNegativeKeywordResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanNegativeKeywordResult) Reset() {
	*m = MutateKeywordPlanNegativeKeywordResult{}
}
func (m *MutateKeywordPlanNegativeKeywordResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanNegativeKeywordResult) ProtoMessage()    {}
func (*MutateKeywordPlanNegativeKeywordResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6100cff1a3faa6ff, []int{4}
}

func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Merge(m, src)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.Size(m)
}
func (m *MutateKeywordPlanNegativeKeywordResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanNegativeKeywordResult proto.InternalMessageInfo

func (m *MutateKeywordPlanNegativeKeywordResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanNegativeKeywordRequest)(nil), "google.ads.googleads.v2.services.GetKeywordPlanNegativeKeywordRequest")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordsRequest)(nil), "google.ads.googleads.v2.services.MutateKeywordPlanNegativeKeywordsRequest")
	proto.RegisterType((*KeywordPlanNegativeKeywordOperation)(nil), "google.ads.googleads.v2.services.KeywordPlanNegativeKeywordOperation")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordsResponse)(nil), "google.ads.googleads.v2.services.MutateKeywordPlanNegativeKeywordsResponse")
	proto.RegisterType((*MutateKeywordPlanNegativeKeywordResult)(nil), "google.ads.googleads.v2.services.MutateKeywordPlanNegativeKeywordResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/keyword_plan_negative_keyword_service.proto", fileDescriptor_6100cff1a3faa6ff)
}

var fileDescriptor_6100cff1a3faa6ff = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6b, 0xe3, 0x46,
	0x14, 0xc0, 0x6b, 0x39, 0xa4, 0xcd, 0x38, 0x69, 0x61, 0x4a, 0x5b, 0xe3, 0xb6, 0xd4, 0x51, 0x4c,
	0xea, 0x9a, 0x22, 0x81, 0x7a, 0x53, 0x30, 0xad, 0x0c, 0x76, 0xd2, 0xa6, 0x49, 0x8c, 0x42, 0x53,
	0x28, 0x06, 0x31, 0x96, 0x26, 0xae, 0x1a, 0x49, 0xa3, 0xce, 0x8c, 0x5c, 0x42, 0xc8, 0xa5, 0xc7,
	0x85, 0x3d, 0xed, 0x37, 0xc8, 0x71, 0xbf, 0xc2, 0x7e, 0x83, 0x5c, 0xf7, 0xb0, 0x90, 0xd3, 0x1e,
	0x76, 0x2f, 0xfb, 0x1d, 0x16, 0x16, 0x69, 0x34, 0xfe, 0x13, 0xe2, 0xc8, 0x90, 0xdc, 0x9e, 0xde,
	0x7b, 0xfe, 0xbd, 0xbf, 0xf3, 0x0c, 0x7e, 0x1f, 0x11, 0x32, 0x0a, 0xb0, 0x8e, 0x3c, 0xa6, 0x0b,
	0x31, 0x95, 0xc6, 0x86, 0xce, 0x30, 0x1d, 0xfb, 0x2e, 0x66, 0xfa, 0x19, 0x3e, 0xff, 0x8f, 0x50,
	0xcf, 0x89, 0x03, 0x14, 0x39, 0x11, 0x1e, 0x21, 0xee, 0x8f, 0xb1, 0x23, 0xb5, 0xb9, 0x9b, 0x16,
	0x53, 0xc2, 0x09, 0xac, 0x0b, 0x84, 0x86, 0x3c, 0xa6, 0x4d, 0x68, 0xda, 0xd8, 0xd0, 0x24, 0xad,
	0xd6, 0x5d, 0x14, 0x8f, 0x62, 0x46, 0x12, 0x5a, 0x18, 0x50, 0x04, 0xaa, 0x7d, 0x23, 0x31, 0xb1,
	0xaf, 0xa3, 0x28, 0x22, 0x1c, 0x71, 0x9f, 0x44, 0x2c, 0xb7, 0x7e, 0x35, 0x63, 0x75, 0x03, 0x1f,
	0x47, 0x3c, 0x37, 0x7c, 0x37, 0x63, 0x38, 0xf5, 0x71, 0xe0, 0x39, 0x43, 0xfc, 0x37, 0x1a, 0xfb,
	0x84, 0xe6, 0x0e, 0x79, 0x01, 0x7a, 0xf6, 0x35, 0x4c, 0x4e, 0x73, 0xaf, 0x10, 0xb1, 0xb3, 0x5b,
	0x6c, 0x1a, 0xbb, 0x3a, 0xe3, 0x88, 0x27, 0x79, 0x50, 0xb5, 0x0f, 0x1a, 0xbb, 0x98, 0xef, 0x8b,
	0x34, 0xfb, 0x01, 0x8a, 0x0e, 0xf3, 0xd4, 0x73, 0x95, 0x8d, 0xff, 0x4d, 0x30, 0xe3, 0xb0, 0x09,
	0x36, 0x64, 0xad, 0x4e, 0x84, 0x42, 0x5c, 0x2d, 0xd5, 0x4b, 0xcd, 0xb5, 0x4e, 0xf9, 0xb5, 0xa5,
	0xd8, 0xeb, 0xd2, 0x72, 0x88, 0x42, 0xac, 0x3e, 0x51, 0x40, 0xf3, 0x20, 0xe1, 0x88, 0xe3, 0xc5,
	0x54, 0x26, 0xb1, 0x0d, 0x50, 0x71, 0x13, 0xc6, 0x49, 0x88, 0xa9, 0xe3, 0x7b, 0xb3, 0x50, 0x20,
	0xf5, 0xbf, 0x7a, 0xf0, 0x1f, 0x00, 0x48, 0x8c, 0xa9, 0xe8, 0x56, 0x55, 0xa9, 0x97, 0x9b, 0x15,
	0xa3, 0xab, 0x15, 0x4d, 0x4d, 0x5b, 0x1c, 0xff, 0x48, 0xd2, 0xf2, 0x58, 0x53, 0x3a, 0xfc, 0x1e,
	0x7c, 0x16, 0x23, 0xca, 0x7d, 0x14, 0x38, 0xa7, 0xc8, 0x0f, 0x12, 0x8a, 0xab, 0xe5, 0x7a, 0xa9,
	0xf9, 0x89, 0xfd, 0x69, 0xae, 0xee, 0x09, 0x2d, 0xdc, 0x02, 0x1b, 0x63, 0x14, 0xf8, 0x1e, 0xe2,
	0xd8, 0x21, 0x51, 0x70, 0x5e, 0x5d, 0xc9, 0xdc, 0xd6, 0xa5, 0xf2, 0x28, 0x0a, 0xce, 0xd5, 0x17,
	0x0a, 0xd8, 0x5a, 0x22, 0x0d, 0xb8, 0x03, 0x2a, 0x49, 0x9c, 0xa1, 0xd2, 0xa1, 0x65, 0xa8, 0x8a,
	0x51, 0x93, 0x25, 0xca, 0xb9, 0x6a, 0xbd, 0x74, 0xae, 0x07, 0x88, 0x9d, 0xd9, 0x40, 0xb8, 0xa7,
	0x32, 0xfc, 0x13, 0xac, 0xba, 0x14, 0x23, 0x2e, 0x86, 0x52, 0x31, 0xda, 0x0b, 0x5b, 0x33, 0x59,
	0xd7, 0x7b, 0x7a, 0xb3, 0xf7, 0x91, 0x9d, 0xe3, 0x52, 0xb0, 0x08, 0x53, 0x55, 0x1e, 0x09, 0x2c,
	0x70, 0xb0, 0x0a, 0x56, 0x29, 0x0e, 0xc9, 0x58, 0xf4, 0x76, 0x2d, 0xb5, 0x88, 0xef, 0x4e, 0x05,
	0xac, 0x4d, 0x86, 0xa1, 0xbe, 0x2a, 0x81, 0x1f, 0x96, 0x58, 0x25, 0x16, 0x93, 0x88, 0x61, 0xd8,
	0x03, 0x5f, 0xdc, 0x9a, 0x9c, 0x83, 0x29, 0x25, 0x34, 0x8b, 0x51, 0x31, 0xa0, 0x4c, 0x9e, 0xc6,
	0xae, 0x76, 0x9c, 0xbd, 0x01, 0xfb, 0xf3, 0xf9, 0x99, 0x76, 0x53, 0x77, 0x38, 0x04, 0x1f, 0x53,
	0xcc, 0x92, 0x80, 0xcb, 0x55, 0xdb, 0x2b, 0x5e, 0xb5, 0xa2, 0x2c, 0xed, 0x0c, 0x68, 0x4b, 0xb0,
	0x7a, 0x00, 0xb6, 0x97, 0xfb, 0x49, 0xba, 0x66, 0x77, 0x3c, 0xbc, 0xf9, 0x37, 0x67, 0xbc, 0x5d,
	0x01, 0x9b, 0x8b, 0x49, 0xc7, 0x22, 0x4b, 0xf8, 0xbe, 0x04, 0xbe, 0xbd, 0xf7, 0xb1, 0xc3, 0x5e,
	0x71, 0xa5, 0xcb, 0x5c, 0x8b, 0xda, 0xc3, 0x16, 0x45, 0xfd, 0xe3, 0xc6, 0x9a, 0x2f, 0xfa, 0xff,
	0x97, 0x6f, 0x9e, 0x29, 0x3f, 0xc3, 0x76, 0x7a, 0x72, 0x2f, 0xe6, 0x2c, 0x6d, 0x79, 0x26, 0x98,
	0xde, 0x92, 0x37, 0xf8, 0xae, 0x2d, 0xd1, 0x5b, 0x97, 0xf0, 0x4a, 0x01, 0x9b, 0x85, 0xeb, 0x04,
	0x7f, 0x7b, 0xf8, 0xb4, 0xe5, 0x79, 0xab, 0xed, 0x3f, 0x0a, 0x4b, 0xec, 0xb7, 0xea, 0xdd, 0x58,
	0x5f, 0xce, 0x1c, 0xcb, 0x1f, 0xa7, 0x47, 0x2b, 0x6b, 0x4f, 0x57, 0xfd, 0x25, 0x6d, 0xcf, 0xb4,
	0x1f, 0x17, 0x33, 0xce, 0xed, 0xd6, 0xe5, 0x7d, 0xdd, 0x31, 0xc3, 0x2c, 0x0b, 0xb3, 0xd4, 0xaa,
	0x7d, 0x7d, 0x6d, 0x55, 0xa7, 0x99, 0xe6, 0x52, 0xec, 0x33, 0xcd, 0x25, 0x61, 0xe7, 0xa9, 0x02,
	0x1a, 0x2e, 0x09, 0x0b, 0xab, 0xea, 0x6c, 0x17, 0x6e, 0x63, 0x3f, 0xbd, 0x69, 0xfd, 0xd2, 0x5f,
	0x7b, 0x39, 0x6b, 0x44, 0x02, 0x14, 0x8d, 0x34, 0x42, 0x47, 0xfa, 0x08, 0x47, 0xd9, 0xc5, 0xd3,
	0xa7, 0xd1, 0x17, 0xff, 0xd3, 0xef, 0x48, 0xe1, 0x4a, 0x29, 0xef, 0x5a, 0xd6, 0x73, 0xa5, 0xbe,
	0x2b, 0x80, 0x96, 0xc7, 0x34, 0x21, 0xa6, 0xd2, 0x89, 0xa1, 0xe5, 0x81, 0xd9, 0xb5, 0x74, 0x19,
	0x58, 0x1e, 0x1b, 0x4c, 0x5c, 0x06, 0x27, 0xc6, 0x40, 0xba, 0xbc, 0x53, 0x1a, 0x42, 0x6f, 0x9a,
	0x96, 0xc7, 0x4c, 0x73, 0xe2, 0x64, 0x9a, 0x27, 0x86, 0x69, 0x4a, 0xb7, 0xe1, 0x6a, 0x96, 0xe7,
	0x4f, 0x1f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x99, 0x18, 0xca, 0x90, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanNegativeKeywordServiceClient is the client API for KeywordPlanNegativeKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanNegativeKeywordServiceClient interface {
	// Returns the requested plan in full detail.
	GetKeywordPlanNegativeKeyword(ctx context.Context, in *GetKeywordPlanNegativeKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanNegativeKeyword, error)
	// Creates, updates, or removes Keyword Plan negative keywords. Operation
	// statuses are returned.
	MutateKeywordPlanNegativeKeywords(ctx context.Context, in *MutateKeywordPlanNegativeKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanNegativeKeywordsResponse, error)
}

type keywordPlanNegativeKeywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanNegativeKeywordServiceClient(cc grpc.ClientConnInterface) KeywordPlanNegativeKeywordServiceClient {
	return &keywordPlanNegativeKeywordServiceClient{cc}
}

func (c *keywordPlanNegativeKeywordServiceClient) GetKeywordPlanNegativeKeyword(ctx context.Context, in *GetKeywordPlanNegativeKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanNegativeKeyword, error) {
	out := new(resources.KeywordPlanNegativeKeyword)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService/GetKeywordPlanNegativeKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanNegativeKeywordServiceClient) MutateKeywordPlanNegativeKeywords(ctx context.Context, in *MutateKeywordPlanNegativeKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanNegativeKeywordsResponse, error) {
	out := new(MutateKeywordPlanNegativeKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService/MutateKeywordPlanNegativeKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanNegativeKeywordServiceServer is the server API for KeywordPlanNegativeKeywordService service.
type KeywordPlanNegativeKeywordServiceServer interface {
	// Returns the requested plan in full detail.
	GetKeywordPlanNegativeKeyword(context.Context, *GetKeywordPlanNegativeKeywordRequest) (*resources.KeywordPlanNegativeKeyword, error)
	// Creates, updates, or removes Keyword Plan negative keywords. Operation
	// statuses are returned.
	MutateKeywordPlanNegativeKeywords(context.Context, *MutateKeywordPlanNegativeKeywordsRequest) (*MutateKeywordPlanNegativeKeywordsResponse, error)
}

// UnimplementedKeywordPlanNegativeKeywordServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanNegativeKeywordServiceServer struct {
}

func (*UnimplementedKeywordPlanNegativeKeywordServiceServer) GetKeywordPlanNegativeKeyword(ctx context.Context, req *GetKeywordPlanNegativeKeywordRequest) (*resources.KeywordPlanNegativeKeyword, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetKeywordPlanNegativeKeyword not implemented")
}
func (*UnimplementedKeywordPlanNegativeKeywordServiceServer) MutateKeywordPlanNegativeKeywords(ctx context.Context, req *MutateKeywordPlanNegativeKeywordsRequest) (*MutateKeywordPlanNegativeKeywordsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateKeywordPlanNegativeKeywords not implemented")
}

func RegisterKeywordPlanNegativeKeywordServiceServer(s *grpc.Server, srv KeywordPlanNegativeKeywordServiceServer) {
	s.RegisterService(&_KeywordPlanNegativeKeywordService_serviceDesc, srv)
}

func _KeywordPlanNegativeKeywordService_GetKeywordPlanNegativeKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanNegativeKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanNegativeKeywordServiceServer).GetKeywordPlanNegativeKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService/GetKeywordPlanNegativeKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanNegativeKeywordServiceServer).GetKeywordPlanNegativeKeyword(ctx, req.(*GetKeywordPlanNegativeKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanNegativeKeywordService_MutateKeywordPlanNegativeKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanNegativeKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanNegativeKeywordServiceServer).MutateKeywordPlanNegativeKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService/MutateKeywordPlanNegativeKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanNegativeKeywordServiceServer).MutateKeywordPlanNegativeKeywords(ctx, req.(*MutateKeywordPlanNegativeKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanNegativeKeywordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordPlanNegativeKeywordService",
	HandlerType: (*KeywordPlanNegativeKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanNegativeKeyword",
			Handler:    _KeywordPlanNegativeKeywordService_GetKeywordPlanNegativeKeyword_Handler,
		},
		{
			MethodName: "MutateKeywordPlanNegativeKeywords",
			Handler:    _KeywordPlanNegativeKeywordService_MutateKeywordPlanNegativeKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_plan_negative_keyword_service.proto",
}
