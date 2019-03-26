// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/server.proto

package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_ddf99cd3866b15c2, []int{0}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (dst *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(dst, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_ddf99cd3866b15c2, []int{1}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (dst *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(dst, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "server.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "server.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	// Version returns PMM Server version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/server.Server/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	// Version returns PMM Server version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Server/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Server_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/server.proto",
}

func init() { proto.RegisterFile("server/server.proto", fileDescriptor_server_ddf99cd3866b15c2) }

var fileDescriptor_server_ddf99cd3866b15c2 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e,
	0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x95, 0x94, 0x0e, 0x98, 0x4a, 0xd6, 0x4d, 0x4f,
	0xcd, 0xd3, 0x2d, 0x2e, 0x4f, 0x4c, 0x4f, 0x4f, 0x2d, 0xd2, 0xcf, 0x2f, 0x00, 0xab, 0xc0, 0x54,
	0xad, 0x24, 0xc0, 0xc5, 0x17, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x17, 0x94, 0x5a, 0x58, 0x9a,
	0x5a, 0x5c, 0xa2, 0xa4, 0xcd, 0xc5, 0x0f, 0x17, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92,
	0xe0, 0x62, 0x2f, 0x83, 0x08, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0x91,
	0x5c, 0x6c, 0xc1, 0x60, 0x47, 0x09, 0xf9, 0x73, 0xb1, 0x43, 0xb5, 0x09, 0x89, 0xe9, 0x41, 0x9d,
	0x8d, 0x6a, 0xb2, 0x94, 0x38, 0x86, 0x38, 0xc4, 0x7c, 0x25, 0xe1, 0xa6, 0xcb, 0x4f, 0x26, 0x33,
	0xf1, 0x0a, 0x71, 0xeb, 0x97, 0x19, 0xea, 0x43, 0x8d, 0x76, 0x0a, 0x9d, 0xe4, 0xe8, 0x15, 0xe4,
	0xc1, 0xc5, 0x9e, 0x92, 0x9a, 0x96, 0x58, 0x9a, 0x53, 0x22, 0x64, 0xcb, 0x25, 0xe4, 0x98, 0xa7,
	0x90, 0x5a, 0x54, 0x94, 0x5f, 0xa4, 0x50, 0x04, 0xd5, 0xa9, 0x27, 0xa4, 0xce, 0xa5, 0x2a, 0xa5,
	0xac, 0xac, 0x9f, 0x92, 0x9a, 0x96, 0x99, 0x97, 0x09, 0xf1, 0x24, 0xc4, 0x92, 0x82, 0x24, 0x57,
	0x90, 0x52, 0x98, 0x1d, 0x51, 0xd0, 0xc0, 0x4b, 0x62, 0x03, 0xfb, 0xdb, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xb7, 0xe1, 0x47, 0x62, 0x01, 0x00, 0x00,
}