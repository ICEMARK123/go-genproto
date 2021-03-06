// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/parental_status_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [ParentalStatusViewService.GetParentalStatusView][google.ads.googleads.v3.services.ParentalStatusViewService.GetParentalStatusView].
type GetParentalStatusViewRequest struct {
	// Required. The resource name of the parental status view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParentalStatusViewRequest) Reset()         { *m = GetParentalStatusViewRequest{} }
func (m *GetParentalStatusViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetParentalStatusViewRequest) ProtoMessage()    {}
func (*GetParentalStatusViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1481d5b1312f10e7, []int{0}
}

func (m *GetParentalStatusViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParentalStatusViewRequest.Unmarshal(m, b)
}
func (m *GetParentalStatusViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParentalStatusViewRequest.Marshal(b, m, deterministic)
}
func (m *GetParentalStatusViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParentalStatusViewRequest.Merge(m, src)
}
func (m *GetParentalStatusViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetParentalStatusViewRequest.Size(m)
}
func (m *GetParentalStatusViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParentalStatusViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParentalStatusViewRequest proto.InternalMessageInfo

func (m *GetParentalStatusViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetParentalStatusViewRequest)(nil), "google.ads.googleads.v3.services.GetParentalStatusViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/parental_status_view_service.proto", fileDescriptor_1481d5b1312f10e7)
}

var fileDescriptor_1481d5b1312f10e7 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6b, 0xd4, 0x40,
	0x18, 0x65, 0x53, 0x10, 0x0c, 0x7a, 0x09, 0x88, 0x75, 0x2d, 0xba, 0x94, 0x1e, 0x4a, 0x0f, 0x33,
	0x60, 0x28, 0xc2, 0xf8, 0x03, 0x66, 0x3d, 0x6c, 0x2f, 0xca, 0xd2, 0xc2, 0x1e, 0x24, 0x10, 0xa6,
	0xc9, 0x67, 0x1c, 0x48, 0x66, 0xe2, 0x7c, 0x93, 0xf4, 0x20, 0x5e, 0x04, 0xff, 0x02, 0x2f, 0x9e,
	0x3d, 0xfa, 0xa7, 0xf4, 0xea, 0x4d, 0x10, 0x3c, 0x78, 0xf2, 0xaf, 0x90, 0xec, 0x64, 0xb2, 0x2d,
	0x35, 0xf6, 0xf6, 0x98, 0xf7, 0xbe, 0xf7, 0xbe, 0x1f, 0x13, 0xbe, 0x28, 0xb4, 0x2e, 0x4a, 0xa0,
	0x22, 0x47, 0xea, 0x60, 0x87, 0xda, 0x98, 0x22, 0x98, 0x56, 0x66, 0x80, 0xb4, 0x16, 0x06, 0x94,
	0x15, 0x65, 0x8a, 0x56, 0xd8, 0x06, 0xd3, 0x56, 0xc2, 0x59, 0xda, 0xb3, 0xa4, 0x36, 0xda, 0xea,
	0x68, 0xe6, 0x2a, 0x89, 0xc8, 0x91, 0x0c, 0x26, 0xa4, 0x8d, 0x89, 0x37, 0x99, 0x3e, 0x1d, 0x8b,
	0x31, 0x80, 0xba, 0x31, 0x63, 0x39, 0xce, 0x7f, 0xba, 0xe3, 0xab, 0x6b, 0x49, 0x85, 0x52, 0xda,
	0x0a, 0x2b, 0xb5, 0xc2, 0x9e, 0xbd, 0x7b, 0x81, 0xcd, 0x4a, 0x09, 0xca, 0xf6, 0xc4, 0xc3, 0x0b,
	0xc4, 0x1b, 0x09, 0x65, 0x9e, 0x9e, 0xc2, 0x5b, 0xd1, 0x4a, 0x6d, 0x9c, 0x60, 0xf7, 0x28, 0xdc,
	0x59, 0x80, 0x5d, 0xf6, 0xc1, 0x27, 0xeb, 0xdc, 0x95, 0x84, 0xb3, 0x63, 0x78, 0xd7, 0x00, 0xda,
	0x68, 0x3f, 0xbc, 0xed, 0xfb, 0x4b, 0x95, 0xa8, 0x60, 0x7b, 0x32, 0x9b, 0xec, 0xdf, 0x9c, 0x6f,
	0xfd, 0xe2, 0xc1, 0xf1, 0x2d, 0xcf, 0xbc, 0x12, 0x15, 0x3c, 0xfa, 0x12, 0x84, 0xf7, 0xae, 0xfa,
	0x9c, 0xb8, 0xf1, 0xa3, 0x9f, 0x93, 0xf0, 0xce, 0x3f, 0x83, 0xa2, 0xe7, 0xe4, 0xba, 0xd5, 0x91,
	0xff, 0x75, 0x38, 0x3d, 0x1c, 0xad, 0x1f, 0x16, 0x4b, 0xae, 0x56, 0xef, 0xbe, 0xfc, 0xc1, 0x2f,
	0x4f, 0xf6, 0xf1, 0xfb, 0xef, 0xcf, 0xc1, 0xe3, 0xe8, 0xb0, 0x3b, 0xc9, 0xfb, 0x4b, 0xcc, 0xb3,
	0xac, 0x41, 0xab, 0x2b, 0x30, 0x48, 0x0f, 0x86, 0x1b, 0x6d, 0xac, 0x90, 0x1e, 0x7c, 0x98, 0xde,
	0x3f, 0xe7, 0xdb, 0x9b, 0xf0, 0x1e, 0xd5, 0x12, 0x49, 0xa6, 0xab, 0xf9, 0xa7, 0x20, 0xdc, 0xcb,
	0x74, 0x75, 0xed, 0xa0, 0xf3, 0x07, 0xa3, 0x0b, 0x5c, 0x76, 0xd7, 0x5a, 0x4e, 0x5e, 0x1f, 0xf5,
	0x1e, 0x85, 0x2e, 0x85, 0x2a, 0x88, 0x36, 0x05, 0x2d, 0x40, 0xad, 0x6f, 0x49, 0x37, 0xa9, 0xe3,
	0x7f, 0xf9, 0x89, 0x07, 0x5f, 0x83, 0xad, 0x05, 0xe7, 0xdf, 0x82, 0xd9, 0xc2, 0x19, 0xf2, 0x1c,
	0x89, 0x83, 0x1d, 0x5a, 0xc5, 0xa4, 0x0f, 0xc6, 0x73, 0x2f, 0x49, 0x78, 0x8e, 0xc9, 0x20, 0x49,
	0x56, 0x71, 0xe2, 0x25, 0x7f, 0x82, 0x3d, 0xf7, 0xce, 0x18, 0xcf, 0x91, 0xb1, 0x41, 0xc4, 0xd8,
	0x2a, 0x66, 0xcc, 0xcb, 0x4e, 0x6f, 0xac, 0xfb, 0x8c, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x37, 0x7a, 0x88, 0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParentalStatusViewServiceClient is the client API for ParentalStatusViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParentalStatusViewServiceClient interface {
	// Returns the requested parental status view in full detail.
	GetParentalStatusView(ctx context.Context, in *GetParentalStatusViewRequest, opts ...grpc.CallOption) (*resources.ParentalStatusView, error)
}

type parentalStatusViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParentalStatusViewServiceClient(cc grpc.ClientConnInterface) ParentalStatusViewServiceClient {
	return &parentalStatusViewServiceClient{cc}
}

func (c *parentalStatusViewServiceClient) GetParentalStatusView(ctx context.Context, in *GetParentalStatusViewRequest, opts ...grpc.CallOption) (*resources.ParentalStatusView, error) {
	out := new(resources.ParentalStatusView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.ParentalStatusViewService/GetParentalStatusView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParentalStatusViewServiceServer is the server API for ParentalStatusViewService service.
type ParentalStatusViewServiceServer interface {
	// Returns the requested parental status view in full detail.
	GetParentalStatusView(context.Context, *GetParentalStatusViewRequest) (*resources.ParentalStatusView, error)
}

// UnimplementedParentalStatusViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParentalStatusViewServiceServer struct {
}

func (*UnimplementedParentalStatusViewServiceServer) GetParentalStatusView(ctx context.Context, req *GetParentalStatusViewRequest) (*resources.ParentalStatusView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParentalStatusView not implemented")
}

func RegisterParentalStatusViewServiceServer(s *grpc.Server, srv ParentalStatusViewServiceServer) {
	s.RegisterService(&_ParentalStatusViewService_serviceDesc, srv)
}

func _ParentalStatusViewService_GetParentalStatusView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParentalStatusViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParentalStatusViewServiceServer).GetParentalStatusView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.ParentalStatusViewService/GetParentalStatusView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParentalStatusViewServiceServer).GetParentalStatusView(ctx, req.(*GetParentalStatusViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParentalStatusViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.ParentalStatusViewService",
	HandlerType: (*ParentalStatusViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParentalStatusView",
			Handler:    _ParentalStatusViewService_GetParentalStatusView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/parental_status_view_service.proto",
}
