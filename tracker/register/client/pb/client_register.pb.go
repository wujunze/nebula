// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client_register.proto

/*
Package register_client_pb is a generated protocol buffer package.

It is generated from these files:
	client_register.proto

It has these top-level messages:
	GetPublicKeyReq
	GetPublicKeyResp
	RegisterReq
	RegisterResp
	VerifyContactEmailReq
	VerifyContactEmailResp
	GetTrackerServerReq
	GetTrackerServerResp
	TrackerServer
*/
package register_client_pb

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

type GetPublicKeyReq struct {
	Version uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
}

func (m *GetPublicKeyReq) Reset()                    { *m = GetPublicKeyReq{} }
func (m *GetPublicKeyReq) String() string            { return proto.CompactTextString(m) }
func (*GetPublicKeyReq) ProtoMessage()               {}
func (*GetPublicKeyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPublicKeyReq) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type GetPublicKeyResp struct {
	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (m *GetPublicKeyResp) Reset()                    { *m = GetPublicKeyResp{} }
func (m *GetPublicKeyResp) String() string            { return proto.CompactTextString(m) }
func (*GetPublicKeyResp) ProtoMessage()               {}
func (*GetPublicKeyResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetPublicKeyResp) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type RegisterReq struct {
	Version         uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	NodeId          []byte `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PublicKeyEnc    []byte `protobuf:"bytes,3,opt,name=publicKeyEnc,proto3" json:"publicKeyEnc,omitempty"`
	ContactEmailEnc []byte `protobuf:"bytes,4,opt,name=contactEmailEnc,proto3" json:"contactEmailEnc,omitempty"`
}

func (m *RegisterReq) Reset()                    { *m = RegisterReq{} }
func (m *RegisterReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()               {}
func (*RegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterReq) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RegisterReq) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *RegisterReq) GetPublicKeyEnc() []byte {
	if m != nil {
		return m.PublicKeyEnc
	}
	return nil
}

func (m *RegisterReq) GetContactEmailEnc() []byte {
	if m != nil {
		return m.ContactEmailEnc
	}
	return nil
}

type RegisterResp struct {
	Code   uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *RegisterResp) Reset()                    { *m = RegisterResp{} }
func (m *RegisterResp) String() string            { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()               {}
func (*RegisterResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterResp) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type VerifyContactEmailReq struct {
	Version    uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	NodeId     []byte `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	VerifyCode string `protobuf:"bytes,3,opt,name=verifyCode" json:"verifyCode,omitempty"`
	Sign       []byte `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *VerifyContactEmailReq) Reset()                    { *m = VerifyContactEmailReq{} }
func (m *VerifyContactEmailReq) String() string            { return proto.CompactTextString(m) }
func (*VerifyContactEmailReq) ProtoMessage()               {}
func (*VerifyContactEmailReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VerifyContactEmailReq) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VerifyContactEmailReq) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *VerifyContactEmailReq) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

func (m *VerifyContactEmailReq) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type VerifyContactEmailResp struct {
	Code   uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *VerifyContactEmailResp) Reset()                    { *m = VerifyContactEmailResp{} }
func (m *VerifyContactEmailResp) String() string            { return proto.CompactTextString(m) }
func (*VerifyContactEmailResp) ProtoMessage()               {}
func (*VerifyContactEmailResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VerifyContactEmailResp) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *VerifyContactEmailResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type GetTrackerServerReq struct {
	Version   uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	NodeId    []byte `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Sign      []byte `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *GetTrackerServerReq) Reset()                    { *m = GetTrackerServerReq{} }
func (m *GetTrackerServerReq) String() string            { return proto.CompactTextString(m) }
func (*GetTrackerServerReq) ProtoMessage()               {}
func (*GetTrackerServerReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetTrackerServerReq) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetTrackerServerReq) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *GetTrackerServerReq) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *GetTrackerServerReq) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type GetTrackerServerResp struct {
	Server []*TrackerServer `protobuf:"bytes,1,rep,name=server" json:"server,omitempty"`
}

func (m *GetTrackerServerResp) Reset()                    { *m = GetTrackerServerResp{} }
func (m *GetTrackerServerResp) String() string            { return proto.CompactTextString(m) }
func (*GetTrackerServerResp) ProtoMessage()               {}
func (*GetTrackerServerResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetTrackerServerResp) GetServer() []*TrackerServer {
	if m != nil {
		return m.Server
	}
	return nil
}

type TrackerServer struct {
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	Port   uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *TrackerServer) Reset()                    { *m = TrackerServer{} }
func (m *TrackerServer) String() string            { return proto.CompactTextString(m) }
func (*TrackerServer) ProtoMessage()               {}
func (*TrackerServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TrackerServer) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *TrackerServer) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPublicKeyReq)(nil), "register.client.pb.GetPublicKeyReq")
	proto.RegisterType((*GetPublicKeyResp)(nil), "register.client.pb.GetPublicKeyResp")
	proto.RegisterType((*RegisterReq)(nil), "register.client.pb.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "register.client.pb.RegisterResp")
	proto.RegisterType((*VerifyContactEmailReq)(nil), "register.client.pb.VerifyContactEmailReq")
	proto.RegisterType((*VerifyContactEmailResp)(nil), "register.client.pb.VerifyContactEmailResp")
	proto.RegisterType((*GetTrackerServerReq)(nil), "register.client.pb.GetTrackerServerReq")
	proto.RegisterType((*GetTrackerServerResp)(nil), "register.client.pb.GetTrackerServerResp")
	proto.RegisterType((*TrackerServer)(nil), "register.client.pb.TrackerServer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClientRegisterService service

type ClientRegisterServiceClient interface {
	GetPublicKey(ctx context.Context, in *GetPublicKeyReq, opts ...grpc.CallOption) (*GetPublicKeyResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	VerifyContactEmail(ctx context.Context, in *VerifyContactEmailReq, opts ...grpc.CallOption) (*VerifyContactEmailResp, error)
	GetTrackerServer(ctx context.Context, in *GetTrackerServerReq, opts ...grpc.CallOption) (*GetTrackerServerResp, error)
}

type clientRegisterServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientRegisterServiceClient(cc *grpc.ClientConn) ClientRegisterServiceClient {
	return &clientRegisterServiceClient{cc}
}

func (c *clientRegisterServiceClient) GetPublicKey(ctx context.Context, in *GetPublicKeyReq, opts ...grpc.CallOption) (*GetPublicKeyResp, error) {
	out := new(GetPublicKeyResp)
	err := grpc.Invoke(ctx, "/register.client.pb.ClientRegisterService/GetPublicKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegisterServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := grpc.Invoke(ctx, "/register.client.pb.ClientRegisterService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegisterServiceClient) VerifyContactEmail(ctx context.Context, in *VerifyContactEmailReq, opts ...grpc.CallOption) (*VerifyContactEmailResp, error) {
	out := new(VerifyContactEmailResp)
	err := grpc.Invoke(ctx, "/register.client.pb.ClientRegisterService/VerifyContactEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientRegisterServiceClient) GetTrackerServer(ctx context.Context, in *GetTrackerServerReq, opts ...grpc.CallOption) (*GetTrackerServerResp, error) {
	out := new(GetTrackerServerResp)
	err := grpc.Invoke(ctx, "/register.client.pb.ClientRegisterService/GetTrackerServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClientRegisterService service

type ClientRegisterServiceServer interface {
	GetPublicKey(context.Context, *GetPublicKeyReq) (*GetPublicKeyResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	VerifyContactEmail(context.Context, *VerifyContactEmailReq) (*VerifyContactEmailResp, error)
	GetTrackerServer(context.Context, *GetTrackerServerReq) (*GetTrackerServerResp, error)
}

func RegisterClientRegisterServiceServer(s *grpc.Server, srv ClientRegisterServiceServer) {
	s.RegisterService(&_ClientRegisterService_serviceDesc, srv)
}

func _ClientRegisterService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegisterServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.client.pb.ClientRegisterService/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegisterServiceServer).GetPublicKey(ctx, req.(*GetPublicKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegisterService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegisterServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.client.pb.ClientRegisterService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegisterServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegisterService_VerifyContactEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyContactEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegisterServiceServer).VerifyContactEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.client.pb.ClientRegisterService/VerifyContactEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegisterServiceServer).VerifyContactEmail(ctx, req.(*VerifyContactEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientRegisterService_GetTrackerServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrackerServerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientRegisterServiceServer).GetTrackerServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.client.pb.ClientRegisterService/GetTrackerServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientRegisterServiceServer).GetTrackerServer(ctx, req.(*GetTrackerServerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientRegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "register.client.pb.ClientRegisterService",
	HandlerType: (*ClientRegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKey",
			Handler:    _ClientRegisterService_GetPublicKey_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ClientRegisterService_Register_Handler,
		},
		{
			MethodName: "VerifyContactEmail",
			Handler:    _ClientRegisterService_VerifyContactEmail_Handler,
		},
		{
			MethodName: "GetTrackerServer",
			Handler:    _ClientRegisterService_GetTrackerServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_register.proto",
}

func init() { proto.RegisterFile("client_register.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x8b, 0xd3, 0x40,
	0x10, 0xb7, 0xb6, 0x54, 0x3b, 0xd7, 0x72, 0x32, 0xda, 0x23, 0x94, 0x43, 0xeb, 0x2a, 0x58, 0x15,
	0x82, 0x9c, 0x4f, 0xea, 0xe3, 0x79, 0x1c, 0x22, 0xa2, 0xae, 0xe2, 0x8b, 0x0f, 0x92, 0x6e, 0xc6,
	0xb0, 0xd8, 0x24, 0xeb, 0xee, 0x5a, 0xe8, 0x83, 0x1f, 0xc1, 0x8f, 0xe0, 0x77, 0x95, 0xcc, 0x25,
	0x6d, 0xd2, 0xc6, 0xf3, 0xbc, 0xb7, 0x9d, 0xd9, 0xdf, 0xfc, 0xfe, 0xb4, 0x93, 0x85, 0xb1, 0x5a,
	0x68, 0xca, 0xfc, 0x17, 0x4b, 0x89, 0x76, 0x9e, 0x6c, 0x68, 0x6c, 0xee, 0x73, 0xc4, 0x75, 0x7d,
	0x76, 0x1f, 0x9a, 0xb9, 0x78, 0x0c, 0xfb, 0xa7, 0xe4, 0xdf, 0xfd, 0x98, 0x2f, 0xb4, 0x7a, 0x4d,
	0x2b, 0x49, 0xdf, 0x31, 0x80, 0x6b, 0x4b, 0xb2, 0x4e, 0xe7, 0x59, 0xd0, 0x99, 0x76, 0x66, 0x23,
	0x59, 0x95, 0xe2, 0x09, 0xdc, 0x68, 0x82, 0x9d, 0xc1, 0x43, 0x18, 0x98, 0xaa, 0xc1, 0xf8, 0xa1,
	0xdc, 0x34, 0xc4, 0xaf, 0x0e, 0xec, 0xc9, 0x52, 0xf5, 0x5c, 0x6e, 0x3c, 0x80, 0x7e, 0x96, 0xc7,
	0xf4, 0x2a, 0x0e, 0xae, 0x32, 0x49, 0x59, 0xa1, 0x80, 0xe1, 0x9a, 0xee, 0x24, 0x53, 0x41, 0x97,
	0x6f, 0x1b, 0x3d, 0x9c, 0xc1, 0xbe, 0xca, 0x33, 0x1f, 0x29, 0x7f, 0x92, 0x46, 0x7a, 0x51, 0xc0,
	0x7a, 0x0c, 0xdb, 0x6e, 0x8b, 0xe7, 0x30, 0xdc, 0xd8, 0x71, 0x06, 0x11, 0x7a, 0x2a, 0x8f, 0xa9,
	0x34, 0xc3, 0xe7, 0xc2, 0x09, 0x59, 0xfb, 0xc6, 0x25, 0xec, 0x64, 0x20, 0xcb, 0x4a, 0xfc, 0x84,
	0xf1, 0x27, 0xb2, 0xfa, 0xeb, 0xea, 0xb8, 0x46, 0x7a, 0xb9, 0x50, 0xb7, 0x01, 0x96, 0x25, 0x55,
	0x4c, 0x1c, 0x69, 0x20, 0x6b, 0x9d, 0xc2, 0x96, 0xd3, 0x49, 0x56, 0xa6, 0xe0, 0xb3, 0x78, 0x09,
	0x07, 0x6d, 0xf2, 0xff, 0x19, 0x62, 0x05, 0x37, 0x4f, 0xc9, 0x7f, 0xb4, 0x91, 0xfa, 0x46, 0xf6,
	0x03, 0xd9, 0xe5, 0x65, 0xff, 0x97, 0x43, 0x18, 0x78, 0x9d, 0x92, 0xf3, 0x51, 0x6a, 0x38, 0x41,
	0x4f, 0x6e, 0x1a, 0xad, 0x01, 0xde, 0xc3, 0xad, 0x5d, 0x69, 0x67, 0xf0, 0x19, 0xf4, 0x1d, 0x57,
	0x41, 0x67, 0xda, 0x9d, 0xed, 0x1d, 0xdd, 0x0d, 0x77, 0xf7, 0x34, 0x6c, 0x8e, 0x95, 0x03, 0xe2,
	0x05, 0x8c, 0x1a, 0x17, 0x85, 0xdb, 0x35, 0x17, 0xc7, 0x3e, 0xab, 0x0a, 0x3f, 0x26, 0xb7, 0x9e,
	0x33, 0x8c, 0x24, 0x9f, 0x8f, 0x7e, 0x77, 0x61, 0x7c, 0xcc, 0x02, 0xd5, 0x4a, 0x14, 0x24, 0x5a,
	0x11, 0x7e, 0x86, 0x61, 0x7d, 0xcf, 0xf1, 0x5e, 0x9b, 0xa3, 0xad, 0xcf, 0x66, 0x72, 0xff, 0xdf,
	0x20, 0x67, 0xc4, 0x15, 0x7c, 0x0b, 0xd7, 0x2b, 0x3d, 0xbc, 0xd3, 0x36, 0x53, 0xfb, 0x5e, 0x26,
	0xd3, 0xf3, 0x01, 0x4c, 0x98, 0x02, 0xee, 0x2e, 0x06, 0x3e, 0x6c, 0x9b, 0x6c, 0xdd, 0xdf, 0xc9,
	0xa3, 0x8b, 0x42, 0x59, 0x2e, 0xe1, 0x47, 0xa0, 0xf9, 0xb3, 0x3f, 0xf8, 0x4b, 0xf6, 0xed, 0x3d,
	0x9b, 0xcc, 0x2e, 0x06, 0x2c, 0x84, 0xe6, 0x7d, 0x7e, 0xb5, 0x9e, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x85, 0xe4, 0xea, 0xeb, 0xce, 0x04, 0x00, 0x00,
}
