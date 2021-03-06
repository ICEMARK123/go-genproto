// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/product_bidding_category_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for
// [ProductBiddingCategoryService.GetProductBiddingCategory][].
type GetProductBiddingCategoryConstantRequest struct {
	// Required. Resource name of the Product Bidding Category to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductBiddingCategoryConstantRequest) Reset() {
	*m = GetProductBiddingCategoryConstantRequest{}
}
func (m *GetProductBiddingCategoryConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductBiddingCategoryConstantRequest) ProtoMessage()    {}
func (*GetProductBiddingCategoryConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db720fb3c3f176ea, []int{0}
}

func (m *GetProductBiddingCategoryConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Unmarshal(m, b)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Merge(m, src)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Size(m)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductBiddingCategoryConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductBiddingCategoryConstantRequest proto.InternalMessageInfo

func (m *GetProductBiddingCategoryConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductBiddingCategoryConstantRequest)(nil), "google.ads.googleads.v1.services.GetProductBiddingCategoryConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/product_bidding_category_constant_service.proto", fileDescriptor_db720fb3c3f176ea)
}

var fileDescriptor_db720fb3c3f176ea = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x99, 0x2c, 0x08, 0x16, 0xbd, 0xf4, 0xe2, 0x52, 0x05, 0xc7, 0x65, 0x85, 0x61, 0x0f,
	0x09, 0x55, 0x44, 0x88, 0x78, 0x48, 0xf7, 0x30, 0x2a, 0x28, 0x65, 0x95, 0x39, 0x48, 0xa1, 0x64,
	0x9a, 0x18, 0x03, 0x6d, 0x52, 0x93, 0x4c, 0x41, 0xc4, 0x83, 0x7e, 0x03, 0xf1, 0x1b, 0x78, 0xf4,
	0x93, 0xc8, 0x5e, 0xbd, 0x79, 0xf2, 0xe0, 0xc9, 0x4f, 0xb1, 0x74, 0xd3, 0x74, 0x77, 0x0f, 0xf3,
	0xe7, 0xf6, 0x30, 0xef, 0x33, 0xbf, 0xf7, 0xe9, 0xf3, 0xb6, 0x51, 0x2e, 0xb4, 0x16, 0x35, 0x47,
	0x94, 0x59, 0xe4, 0x65, 0xaf, 0xba, 0x14, 0x59, 0x6e, 0x3a, 0x59, 0x71, 0x8b, 0x5a, 0xa3, 0xd9,
	0xaa, 0x72, 0xe5, 0x52, 0x32, 0x26, 0x95, 0x28, 0x2b, 0xea, 0xb8, 0xd0, 0xe6, 0x63, 0x59, 0x69,
	0x65, 0x1d, 0x55, 0xae, 0x1c, 0xac, 0xb0, 0x35, 0xda, 0xe9, 0x78, 0xea, 0x31, 0x90, 0x32, 0x0b,
	0x47, 0x22, 0xec, 0x52, 0x18, 0x88, 0xc9, 0xf3, 0x75, 0x3b, 0x0d, 0xb7, 0x7a, 0x65, 0x76, 0x5a,
	0xea, 0x97, 0x25, 0x77, 0x02, 0xaa, 0x95, 0x88, 0x2a, 0xa5, 0x1d, 0x75, 0x52, 0x2b, 0x3b, 0x4c,
	0x6f, 0x5d, 0x9a, 0x56, 0xb5, 0xe4, 0xe3, 0xdf, 0xee, 0x5e, 0x1a, 0xbc, 0x93, 0xbc, 0x66, 0xe5,
	0x92, 0xbf, 0xa7, 0x9d, 0xd4, 0xc6, 0x1b, 0x0e, 0xde, 0x44, 0xb3, 0x39, 0x77, 0xb9, 0x4f, 0x91,
	0xf9, 0x10, 0xc7, 0x43, 0x86, 0xe3, 0x21, 0xc2, 0x09, 0xff, 0xb0, 0xe2, 0xd6, 0xc5, 0xb3, 0xe8,
	0x66, 0x08, 0x5e, 0x2a, 0xda, 0xf0, 0xfd, 0xc9, 0x74, 0x32, 0xbb, 0x9e, 0xed, 0xfd, 0x25, 0xe0,
	0xe4, 0x46, 0x98, 0xbc, 0xa2, 0x0d, 0x7f, 0xf0, 0x0b, 0x44, 0xf7, 0x37, 0x33, 0x5f, 0xfb, 0x8e,
	0xe2, 0x2f, 0x20, 0xba, 0xb7, 0x35, 0x40, 0xfc, 0x02, 0x6e, 0xeb, 0x1a, 0xee, 0xfa, 0x14, 0x09,
	0x59, 0xcb, 0x1a, 0xaf, 0x02, 0x37, 0x93, 0x0e, 0x5e, 0xfe, 0x21, 0x57, 0x9b, 0xf8, 0xfa, 0xfb,
	0xdf, 0x77, 0xf0, 0x38, 0x7e, 0xd4, 0xdf, 0xf6, 0xd3, 0x95, 0xc9, 0xd3, 0x76, 0x23, 0xca, 0xa2,
	0xa3, 0xcf, 0xc9, 0xed, 0x53, 0xb2, 0x7f, 0x11, 0x64, 0x50, 0xad, 0xb4, 0xb0, 0xd2, 0x4d, 0xf6,
	0x0d, 0x44, 0x87, 0x95, 0x6e, 0xb6, 0x16, 0x90, 0x1d, 0xed, 0x54, 0x78, 0xde, 0x5f, 0x3d, 0x9f,
	0xbc, 0x7d, 0x36, 0xf0, 0x84, 0xae, 0xa9, 0x12, 0x50, 0x1b, 0x81, 0x04, 0x57, 0xe7, 0xef, 0x04,
	0xba, 0x48, 0xb0, 0xfe, 0x6b, 0x79, 0x12, 0xc4, 0x0f, 0xb0, 0x37, 0x27, 0xe4, 0x27, 0x98, 0xce,
	0x3d, 0x90, 0x30, 0x0b, 0xbd, 0xec, 0xd5, 0x22, 0x85, 0xc3, 0x62, 0x7b, 0x1a, 0x2c, 0x05, 0x61,
	0xb6, 0x18, 0x2d, 0xc5, 0x22, 0x2d, 0x82, 0xe5, 0x3f, 0x38, 0xf4, 0xbf, 0x63, 0x4c, 0x98, 0xc5,
	0x78, 0x34, 0x61, 0xbc, 0x48, 0x31, 0x0e, 0xb6, 0xe5, 0xb5, 0xf3, 0x9c, 0x0f, 0xcf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0xa4, 0x25, 0x1e, 0xd4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductBiddingCategoryConstantServiceClient is the client API for ProductBiddingCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductBiddingCategoryConstantServiceClient interface {
	// Returns the requested Product Bidding Category in full detail.
	GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error)
}

type productBiddingCategoryConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductBiddingCategoryConstantServiceClient(cc grpc.ClientConnInterface) ProductBiddingCategoryConstantServiceClient {
	return &productBiddingCategoryConstantServiceClient{cc}
}

func (c *productBiddingCategoryConstantServiceClient) GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error) {
	out := new(resources.ProductBiddingCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductBiddingCategoryConstantServiceServer is the server API for ProductBiddingCategoryConstantService service.
type ProductBiddingCategoryConstantServiceServer interface {
	// Returns the requested Product Bidding Category in full detail.
	GetProductBiddingCategoryConstant(context.Context, *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error)
}

// UnimplementedProductBiddingCategoryConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductBiddingCategoryConstantServiceServer struct {
}

func (*UnimplementedProductBiddingCategoryConstantServiceServer) GetProductBiddingCategoryConstant(ctx context.Context, req *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductBiddingCategoryConstant not implemented")
}

func RegisterProductBiddingCategoryConstantServiceServer(s *grpc.Server, srv ProductBiddingCategoryConstantServiceServer) {
	s.RegisterService(&_ProductBiddingCategoryConstantService_serviceDesc, srv)
}

func _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBiddingCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, req.(*GetProductBiddingCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductBiddingCategoryConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ProductBiddingCategoryConstantService",
	HandlerType: (*ProductBiddingCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductBiddingCategoryConstant",
			Handler:    _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/product_bidding_category_constant_service.proto",
}
