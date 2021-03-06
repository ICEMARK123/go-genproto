// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/group_placement_view_service.proto

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

// Request message for [GroupPlacementViewService.GetGroupPlacementView][google.ads.googleads.v1.services.GroupPlacementViewService.GetGroupPlacementView].
type GetGroupPlacementViewRequest struct {
	// Required. The resource name of the Group Placement view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupPlacementViewRequest) Reset()         { *m = GetGroupPlacementViewRequest{} }
func (m *GetGroupPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupPlacementViewRequest) ProtoMessage()    {}
func (*GetGroupPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_722a806b0135cef4, []int{0}
}

func (m *GetGroupPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetGroupPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupPlacementViewRequest.Merge(m, src)
}
func (m *GetGroupPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Size(m)
}
func (m *GetGroupPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupPlacementViewRequest proto.InternalMessageInfo

func (m *GetGroupPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGroupPlacementViewRequest)(nil), "google.ads.googleads.v1.services.GetGroupPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/group_placement_view_service.proto", fileDescriptor_722a806b0135cef4)
}

var fileDescriptor_722a806b0135cef4 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6a, 0xd5, 0x40,
	0x18, 0x25, 0x29, 0x08, 0x06, 0xdd, 0x04, 0xc4, 0x1a, 0x8b, 0x5e, 0x4a, 0x17, 0xa5, 0x8b, 0x19,
	0xa2, 0x14, 0x61, 0xfc, 0x81, 0xb9, 0x2e, 0xd2, 0x8d, 0x72, 0xa9, 0x90, 0x85, 0x04, 0xc2, 0x34,
	0xf9, 0x8c, 0x03, 0xc9, 0x4c, 0x9c, 0x99, 0xa4, 0x0b, 0x71, 0x23, 0xf8, 0x04, 0x6e, 0x5c, 0xbb,
	0xf4, 0x51, 0xba, 0x75, 0x27, 0x08, 0x2e, 0x5c, 0xf9, 0x14, 0x92, 0x4e, 0x26, 0xed, 0xa5, 0xc6,
	0xbb, 0x3b, 0xcc, 0x39, 0xdf, 0x39, 0xdf, 0xcf, 0x04, 0xcf, 0x2b, 0x29, 0xab, 0x1a, 0x30, 0x2b,
	0x35, 0xb6, 0x70, 0x40, 0x7d, 0x8c, 0x35, 0xa8, 0x9e, 0x17, 0xa0, 0x71, 0xa5, 0x64, 0xd7, 0xe6,
	0x6d, 0xcd, 0x0a, 0x68, 0x40, 0x98, 0xbc, 0xe7, 0x70, 0x9a, 0x8f, 0x2c, 0x6a, 0x95, 0x34, 0x32,
	0x5c, 0xd8, 0x4a, 0xc4, 0x4a, 0x8d, 0x26, 0x13, 0xd4, 0xc7, 0xc8, 0x99, 0x44, 0x4f, 0xe6, 0x62,
	0x14, 0x68, 0xd9, 0xa9, 0xb9, 0x1c, 0xeb, 0x1f, 0xed, 0xb8, 0xea, 0x96, 0x63, 0x26, 0x84, 0x34,
	0xcc, 0x70, 0x29, 0xf4, 0xc8, 0xde, 0xbe, 0xc4, 0x16, 0x35, 0x07, 0x61, 0x46, 0xe2, 0xfe, 0x25,
	0xe2, 0x0d, 0x87, 0xba, 0xcc, 0x4f, 0xe0, 0x2d, 0xeb, 0xb9, 0x54, 0x56, 0xb0, 0x7b, 0x14, 0xec,
	0x24, 0x60, 0x92, 0x21, 0x78, 0xe5, 0x72, 0x53, 0x0e, 0xa7, 0xc7, 0xf0, 0xae, 0x03, 0x6d, 0xc2,
	0xfd, 0xe0, 0xa6, 0xeb, 0x2f, 0x17, 0xac, 0x81, 0x6d, 0x6f, 0xe1, 0xed, 0x5f, 0x5f, 0x6e, 0xfd,
	0xa2, 0xfe, 0xf1, 0x0d, 0xc7, 0xbc, 0x64, 0x0d, 0x3c, 0xf8, 0xe2, 0x07, 0x77, 0xae, 0xfa, 0xbc,
	0xb2, 0xe3, 0x87, 0x3f, 0xbd, 0xe0, 0xd6, 0x3f, 0x83, 0xc2, 0x67, 0x68, 0xd3, 0xea, 0xd0, 0xff,
	0x3a, 0x8c, 0x0e, 0x67, 0xeb, 0xa7, 0xc5, 0xa2, 0xab, 0xd5, 0xbb, 0x2f, 0x7e, 0xd0, 0xf5, 0xc9,
	0x3e, 0x7e, 0xff, 0xfd, 0xd9, 0x7f, 0x14, 0x1e, 0x0e, 0x27, 0x79, 0xbf, 0xc6, 0x3c, 0x2d, 0x3a,
	0x6d, 0x64, 0x03, 0x4a, 0xe3, 0x03, 0x7b, 0xa3, 0x35, 0x2b, 0x8d, 0x0f, 0x3e, 0x44, 0x77, 0xcf,
	0xe8, 0xf6, 0x45, 0xf8, 0x88, 0x5a, 0xae, 0x51, 0x21, 0x9b, 0xe5, 0x27, 0x3f, 0xd8, 0x2b, 0x64,
	0xb3, 0x71, 0xd0, 0xe5, 0xbd, 0xd9, 0x05, 0xae, 0x86, 0x6b, 0xad, 0xbc, 0xd7, 0x47, 0xa3, 0x47,
	0x25, 0x6b, 0x26, 0x2a, 0x24, 0x55, 0x85, 0x2b, 0x10, 0xe7, 0xb7, 0xc4, 0x17, 0xa9, 0xf3, 0x7f,
	0xf9, 0xb1, 0x03, 0x5f, 0xfd, 0xad, 0x84, 0xd2, 0x6f, 0xfe, 0x22, 0xb1, 0x86, 0xb4, 0xd4, 0xc8,
	0xc2, 0x01, 0xa5, 0x31, 0x1a, 0x83, 0xf5, 0x99, 0x93, 0x64, 0xb4, 0xd4, 0xd9, 0x24, 0xc9, 0xd2,
	0x38, 0x73, 0x92, 0x3f, 0xfe, 0x9e, 0x7d, 0x27, 0x84, 0x96, 0x9a, 0x90, 0x49, 0x44, 0x48, 0x1a,
	0x13, 0xe2, 0x64, 0x27, 0xd7, 0xce, 0xfb, 0x7c, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x02, 0x37,
	0x00, 0xec, 0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupPlacementViewServiceClient is the client API for GroupPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupPlacementViewServiceClient interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error)
}

type groupPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupPlacementViewServiceClient(cc grpc.ClientConnInterface) GroupPlacementViewServiceClient {
	return &groupPlacementViewServiceClient{cc}
}

func (c *groupPlacementViewServiceClient) GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error) {
	out := new(resources.GroupPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.GroupPlacementViewService/GetGroupPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupPlacementViewServiceServer is the server API for GroupPlacementViewService service.
type GroupPlacementViewServiceServer interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(context.Context, *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error)
}

// UnimplementedGroupPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroupPlacementViewServiceServer struct {
}

func (*UnimplementedGroupPlacementViewServiceServer) GetGroupPlacementView(ctx context.Context, req *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupPlacementView not implemented")
}

func RegisterGroupPlacementViewServiceServer(s *grpc.Server, srv GroupPlacementViewServiceServer) {
	s.RegisterService(&_GroupPlacementViewService_serviceDesc, srv)
}

func _GroupPlacementViewService_GetGroupPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.GroupPlacementViewService/GetGroupPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, req.(*GetGroupPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.GroupPlacementViewService",
	HandlerType: (*GroupPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupPlacementView",
			Handler:    _GroupPlacementViewService_GetGroupPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/group_placement_view_service.proto",
}
