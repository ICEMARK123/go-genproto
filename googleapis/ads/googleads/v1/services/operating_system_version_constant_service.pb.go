// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/operating_system_version_constant_service.proto

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
// [OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant][google.ads.googleads.v1.services.OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant].
type GetOperatingSystemVersionConstantRequest struct {
	// Required. Resource name of the OS version to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperatingSystemVersionConstantRequest) Reset() {
	*m = GetOperatingSystemVersionConstantRequest{}
}
func (m *GetOperatingSystemVersionConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperatingSystemVersionConstantRequest) ProtoMessage()    {}
func (*GetOperatingSystemVersionConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_97a0893223528a69, []int{0}
}

func (m *GetOperatingSystemVersionConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Unmarshal(m, b)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Merge(m, src)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Size(m)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatingSystemVersionConstantRequest proto.InternalMessageInfo

func (m *GetOperatingSystemVersionConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperatingSystemVersionConstantRequest)(nil), "google.ads.googleads.v1.services.GetOperatingSystemVersionConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/operating_system_version_constant_service.proto", fileDescriptor_97a0893223528a69)
}

var fileDescriptor_97a0893223528a69 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x69, 0x16, 0x04, 0x83, 0x5e, 0x72, 0x71, 0xa9, 0x82, 0x75, 0x59, 0xa1, 0xec, 0x61,
	0x86, 0x28, 0x22, 0x8c, 0x78, 0x98, 0x7a, 0xa8, 0x0a, 0x6a, 0xd9, 0x95, 0x1e, 0xa4, 0x10, 0x66,
	0x93, 0xd7, 0x38, 0xd0, 0xcc, 0x5b, 0xe7, 0x9d, 0x0d, 0x88, 0x78, 0xd0, 0x6f, 0x20, 0x7e, 0x03,
	0x8f, 0x7e, 0x12, 0xd9, 0xab, 0x37, 0x4f, 0x1e, 0x3c, 0xf9, 0x29, 0x24, 0x9d, 0x4c, 0x76, 0x7b,
	0xe8, 0x9f, 0xdb, 0x43, 0xdf, 0xa7, 0xbf, 0xf7, 0xc9, 0xf3, 0x26, 0xf1, 0xa4, 0x44, 0x2c, 0xe7,
	0xc0, 0x55, 0x41, 0xdc, 0xcb, 0x46, 0xd5, 0x29, 0x27, 0xb0, 0xb5, 0xce, 0x81, 0x38, 0x2e, 0xc0,
	0x2a, 0xa7, 0x4d, 0x99, 0xd1, 0x07, 0x72, 0x50, 0x65, 0x35, 0x58, 0xd2, 0x68, 0xb2, 0x1c, 0x0d,
	0x39, 0x65, 0x5c, 0xd6, 0x5a, 0xd9, 0xc2, 0xa2, 0xc3, 0x64, 0xe0, 0x31, 0x4c, 0x15, 0xc4, 0x3a,
	0x22, 0xab, 0x53, 0x16, 0x88, 0xfd, 0x67, 0xeb, 0x76, 0x5a, 0x20, 0x3c, 0xb3, 0x3b, 0x2d, 0xf5,
	0xcb, 0xfa, 0xb7, 0x02, 0x6a, 0xa1, 0xb9, 0x32, 0x06, 0x9d, 0x72, 0x1a, 0x0d, 0xb5, 0xd3, 0x1b,
	0x97, 0xa6, 0xf9, 0x5c, 0x43, 0xf7, 0xb7, 0xdb, 0x97, 0x06, 0x6f, 0x35, 0xcc, 0x8b, 0xec, 0x14,
	0xde, 0xa9, 0x5a, 0xa3, 0xf5, 0x86, 0x83, 0xd7, 0xf1, 0x70, 0x0c, 0xee, 0x55, 0x48, 0x71, 0xb2,
	0x0c, 0x31, 0xf5, 0x19, 0x9e, 0xb4, 0x11, 0x8e, 0xe1, 0xfd, 0x19, 0x90, 0x4b, 0x86, 0xf1, 0xf5,
	0x10, 0x3c, 0x33, 0xaa, 0x82, 0xfd, 0xde, 0xa0, 0x37, 0xbc, 0x3a, 0xda, 0xfb, 0x23, 0xa3, 0xe3,
	0x6b, 0x61, 0xf2, 0x52, 0x55, 0x70, 0xef, 0x67, 0x14, 0xdf, 0xdd, 0xcc, 0x3c, 0xf1, 0x1d, 0x25,
	0x9f, 0xa3, 0xf8, 0xce, 0xd6, 0x00, 0xc9, 0x73, 0xb6, 0xad, 0x6b, 0xb6, 0xeb, 0x53, 0xf4, 0xe5,
	0x5a, 0x56, 0x77, 0x15, 0xb6, 0x99, 0x74, 0xf0, 0xe2, 0xb7, 0x5c, 0x6d, 0xe2, 0xcb, 0xaf, 0xbf,
	0xdf, 0xa2, 0x87, 0xc9, 0x83, 0xe6, 0xb6, 0x1f, 0x57, 0x26, 0x8f, 0x71, 0x23, 0x8a, 0xf8, 0xd1,
	0xa7, 0xfe, 0xcd, 0x73, 0xb9, 0x7f, 0x11, 0xa4, 0x55, 0x0b, 0x4d, 0x2c, 0xc7, 0x6a, 0xf4, 0x35,
	0x8a, 0x0f, 0x73, 0xac, 0xb6, 0x16, 0x30, 0x3a, 0xda, 0xa9, 0xf0, 0x49, 0x73, 0xf5, 0x49, 0xef,
	0xcd, 0xd3, 0x96, 0x57, 0xe2, 0x5c, 0x99, 0x92, 0xa1, 0x2d, 0x79, 0x09, 0x66, 0xf9, 0x4e, 0xf0,
	0x8b, 0x04, 0xeb, 0xbf, 0x96, 0x47, 0x41, 0x7c, 0x8f, 0xf6, 0xc6, 0x52, 0xfe, 0x88, 0x06, 0x63,
	0x0f, 0x94, 0x05, 0x31, 0x2f, 0x1b, 0x35, 0x4d, 0x59, 0xbb, 0x98, 0xce, 0x83, 0x65, 0x26, 0x0b,
	0x9a, 0x75, 0x96, 0xd9, 0x34, 0x9d, 0x05, 0xcb, 0xbf, 0xe8, 0xd0, 0xff, 0x2e, 0x84, 0x2c, 0x48,
	0x88, 0xce, 0x24, 0xc4, 0x34, 0x15, 0x22, 0xd8, 0x4e, 0xaf, 0x2c, 0x73, 0xde, 0xff, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xbe, 0x16, 0xe4, 0x3c, 0xd4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatingSystemVersionConstantServiceClient(cc grpc.ClientConnInterface) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
}

// UnimplementedOperatingSystemVersionConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperatingSystemVersionConstantServiceServer struct {
}

func (*UnimplementedOperatingSystemVersionConstantServiceServer) GetOperatingSystemVersionConstant(ctx context.Context, req *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatingSystemVersionConstant not implemented")
}

func RegisterOperatingSystemVersionConstantServiceServer(s *grpc.Server, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&_OperatingSystemVersionConstantService_serviceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatingSystemVersionConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/operating_system_version_constant_service.proto",
}
