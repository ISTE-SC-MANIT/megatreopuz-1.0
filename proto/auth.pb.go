// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AcessToken           string               `protobuf:"bytes,1,opt,name=acessToken,proto3" json:"acessToken,omitempty"`
	RefreshToken         string               `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	AccessTokenExpiry    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=accessTokenExpiry,proto3" json:"accessTokenExpiry,omitempty"`
	RefreshTokenExpiry   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=refreshTokenExpiry,proto3" json:"refreshTokenExpiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAcessToken() string {
	if m != nil {
		return m.AcessToken
	}
	return ""
}

func (m *LoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LoginResponse) GetAccessTokenExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.AccessTokenExpiry
	}
	return nil
}

func (m *LoginResponse) GetRefreshTokenExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.RefreshTokenExpiry
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*Empty)(nil), "auth.Empty")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x53, 0x05, 0xd4, 0x01, 0x0f, 0xae, 0x17, 0xd2, 0x83, 0x92, 0x9e, 0xb8, 0xd0, 0x2a,
	0x3e, 0x01, 0x1a, 0x8c, 0x18, 0xf4, 0x00, 0x3d, 0x79, 0x5b, 0xea, 0xd0, 0x36, 0xba, 0x9d, 0x75,
	0xff, 0xa8, 0xf8, 0xc6, 0xbe, 0x85, 0xa1, 0x5b, 0x48, 0x89, 0x26, 0x9e, 0x76, 0xe7, 0x9b, 0xf9,
	0x7e, 0xc9, 0x7c, 0x03, 0xc0, 0xad, 0xc9, 0x42, 0xa9, 0xc8, 0x10, 0x6b, 0xac, 0xff, 0xfe, 0x79,
	0x4a, 0x94, 0xbe, 0x62, 0x54, 0x6a, 0x0b, 0xbb, 0x8c, 0x4c, 0x2e, 0x50, 0x1b, 0x2e, 0xa4, 0x1b,
	0x0b, 0x6e, 0xa1, 0x33, 0xa5, 0x34, 0x2f, 0x66, 0xf8, 0x66, 0x51, 0x1b, 0xe6, 0xc3, 0xa1, 0xd5,
	0xa8, 0x0a, 0x2e, 0xb0, 0xeb, 0xf5, 0xbc, 0xfe, 0xd1, 0x6c, 0x5b, 0xaf, 0x7b, 0x92, 0x6b, 0xfd,
	0x41, 0xea, 0xb9, 0xbb, 0xe7, 0x7a, 0x9b, 0x3a, 0xf8, 0xf6, 0xe0, 0xb8, 0x02, 0x69, 0x49, 0x85,
	0x46, 0x76, 0x06, 0xc0, 0x13, 0xd4, 0x3a, 0xa6, 0x17, 0x2c, 0x2a, 0x56, 0x4d, 0x61, 0x01, 0x74,
	0x14, 0x2e, 0x15, 0xea, 0xcc, 0x4d, 0x38, 0xe2, 0x8e, 0xc6, 0xee, 0xe0, 0x84, 0x27, 0x5b, 0xcb,
	0xf8, 0x53, 0xe6, 0x6a, 0xd5, 0xdd, 0xef, 0x79, 0xfd, 0xf6, 0xd0, 0x0f, 0xdd, 0x6a, 0xe1, 0x66,
	0xb5, 0x30, 0xde, 0xac, 0x36, 0xfb, 0x6d, 0x62, 0xf7, 0xc0, 0xea, 0xe4, 0x0a, 0xd5, 0xf8, 0x17,
	0xf5, 0x87, 0x2b, 0x38, 0x80, 0xe6, 0x58, 0x48, 0xb3, 0x1a, 0x26, 0xd0, 0x1e, 0x59, 0x93, 0xcd,
	0x51, 0xbd, 0xe7, 0x09, 0xb2, 0x0b, 0x68, 0x96, 0x11, 0x30, 0x16, 0x96, 0x87, 0xa8, 0x07, 0xeb,
	0x9f, 0xee, 0x68, 0x55, 0x46, 0x01, 0xb4, 0xa6, 0x94, 0x92, 0x35, 0xac, 0xed, 0xda, 0x25, 0xd7,
	0xaf, 0x17, 0xd7, 0x97, 0x4f, 0x51, 0x9a, 0x9b, 0xcc, 0x2e, 0xc2, 0x84, 0x44, 0x34, 0x99, 0xc7,
	0xe3, 0xc1, 0xfc, 0x66, 0xf0, 0x30, 0x7a, 0x9c, 0xc4, 0x91, 0xc0, 0x94, 0x1b, 0x85, 0x24, 0xed,
	0xd7, 0x60, 0x6d, 0xa9, 0xee, 0xdc, 0x2a, 0x9f, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x89, 0x18, 0x3d, 0x10, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *Empty) (*Empty, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Logout(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
