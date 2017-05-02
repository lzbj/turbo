// Code generated by protoc-gen-go.
// source: inventoryservice.proto
// DO NOT EDIT!

/*
Package gen is a generated protocol buffer package.

format:
method name: [method_name]
request: [method_name]Request
response: [method_name]Response

It is generated from these files:
	inventoryservice.proto

It has these top-level messages:
	Video
	GetVideoListRequest
	GetVideoListResponse
	GetVideoRequest
	GetVideoResponse
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Video struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Video) Reset()                    { *m = Video{} }
func (m *Video) String() string            { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()               {}
func (*Video) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Video) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Video) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetVideoListRequest struct {
	NetworkId string `protobuf:"bytes,1,opt,name=networkId" json:"networkId,omitempty"`
}

func (m *GetVideoListRequest) Reset()                    { *m = GetVideoListRequest{} }
func (m *GetVideoListRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVideoListRequest) ProtoMessage()               {}
func (*GetVideoListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetVideoListRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

type GetVideoListResponse struct {
	Videos []*Video `protobuf:"bytes,1,rep,name=videos" json:"videos,omitempty"`
}

func (m *GetVideoListResponse) Reset()                    { *m = GetVideoListResponse{} }
func (m *GetVideoListResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVideoListResponse) ProtoMessage()               {}
func (*GetVideoListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetVideoListResponse) GetVideos() []*Video {
	if m != nil {
		return m.Videos
	}
	return nil
}

type GetVideoRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetVideoRequest) Reset()                    { *m = GetVideoRequest{} }
func (m *GetVideoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVideoRequest) ProtoMessage()               {}
func (*GetVideoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetVideoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetVideoResponse struct {
	Video *Video `protobuf:"bytes,1,opt,name=video" json:"video,omitempty"`
}

func (m *GetVideoResponse) Reset()                    { *m = GetVideoResponse{} }
func (m *GetVideoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVideoResponse) ProtoMessage()               {}
func (*GetVideoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetVideoResponse) GetVideo() *Video {
	if m != nil {
		return m.Video
	}
	return nil
}

func init() {
	proto.RegisterType((*Video)(nil), "gen.Video")
	proto.RegisterType((*GetVideoListRequest)(nil), "gen.GetVideoListRequest")
	proto.RegisterType((*GetVideoListResponse)(nil), "gen.GetVideoListResponse")
	proto.RegisterType((*GetVideoRequest)(nil), "gen.GetVideoRequest")
	proto.RegisterType((*GetVideoResponse)(nil), "gen.GetVideoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InventoryService service

type InventoryServiceClient interface {
	GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error)
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error)
}

type inventoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewInventoryServiceClient(cc *grpc.ClientConn) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error) {
	out := new(GetVideoListResponse)
	err := grpc.Invoke(ctx, "/gen.InventoryService/getVideoList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error) {
	out := new(GetVideoResponse)
	err := grpc.Invoke(ctx, "/gen.InventoryService/getVideo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InventoryService service

type InventoryServiceServer interface {
	GetVideoList(context.Context, *GetVideoListRequest) (*GetVideoListResponse, error)
	GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error)
}

func RegisterInventoryServiceServer(s *grpc.Server, srv InventoryServiceServer) {
	s.RegisterService(&_InventoryService_serviceDesc, srv)
}

func _InventoryService_GetVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.InventoryService/GetVideoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetVideoList(ctx, req.(*GetVideoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.InventoryService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InventoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gen.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getVideoList",
			Handler:    _InventoryService_GetVideoList_Handler,
		},
		{
			MethodName: "getVideo",
			Handler:    _InventoryService_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventoryservice.proto",
}

func init() { proto.RegisterFile("inventoryservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0x36, 0xad, 0xbb, 0xd8, 0x51, 0x74, 0x19, 0x57, 0xa9, 0xe2, 0x21, 0xe6, 0x54, 0x10, 0x7a,
	0xd8, 0xf5, 0xa2, 0x77, 0x91, 0x05, 0x4f, 0x11, 0xbc, 0xab, 0x1d, 0x4a, 0x10, 0x93, 0x35, 0x89,
	0x15, 0xdf, 0xc3, 0x07, 0x96, 0x4d, 0x36, 0xd4, 0x4a, 0x6f, 0xe1, 0xfb, 0x9b, 0x2f, 0x33, 0x70,
	0xaa, 0x74, 0x47, 0xda, 0x1b, 0xfb, 0xed, 0xc8, 0x76, 0xea, 0x95, 0xea, 0xb5, 0x35, 0xde, 0x60,
	0xde, 0x92, 0x16, 0x57, 0x30, 0x79, 0x52, 0x0d, 0x19, 0x3c, 0x84, 0x4c, 0x35, 0x25, 0xe3, 0xac,
	0xca, 0x65, 0xa6, 0x1a, 0x44, 0xd8, 0xd5, 0xcf, 0xef, 0x54, 0x66, 0x9c, 0x55, 0x85, 0x0c, 0x6f,
	0xb1, 0x84, 0xe3, 0x7b, 0xf2, 0x41, 0xff, 0xa0, 0x9c, 0x97, 0xf4, 0xf1, 0x49, 0xce, 0xe3, 0x05,
	0x14, 0x9a, 0xfc, 0x97, 0xb1, 0x6f, 0xab, 0x98, 0x50, 0xc8, 0x1e, 0x10, 0xb7, 0x30, 0x1f, 0x9a,
	0xdc, 0xda, 0x68, 0x47, 0x28, 0x60, 0xda, 0x6d, 0x40, 0x57, 0x32, 0x9e, 0x57, 0xfb, 0x0b, 0xa8,
	0x5b, 0xd2, 0x75, 0xd0, 0xc9, 0x2d, 0x23, 0x2e, 0xe1, 0x28, 0x79, 0xd3, 0xb0, 0xbe, 0x67, 0xb1,
	0xe9, 0x29, 0xae, 0x61, 0xd6, 0x4b, 0xb6, 0xd1, 0x1c, 0x26, 0x21, 0x20, 0xc8, 0x86, 0xc9, 0x91,
	0x58, 0xfc, 0x30, 0x98, 0xad, 0xd2, 0x5a, 0x1e, 0xe3, 0x5a, 0xf0, 0x0e, 0x0e, 0xda, 0x3f, 0x4d,
	0xb1, 0x0c, 0xbe, 0x91, 0x1f, 0x9f, 0x9f, 0x8d, 0x30, 0x71, 0xb6, 0xd8, 0xc1, 0x1b, 0xd8, 0x4b,
	0x31, 0x38, 0x1f, 0x08, 0x93, 0xfd, 0xe4, 0x1f, 0x9a, 0xac, 0x2f, 0xd3, 0x70, 0x99, 0xe5, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x8c, 0x45, 0x4f, 0xb3, 0x01, 0x00, 0x00,
}