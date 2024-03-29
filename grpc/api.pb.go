// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Role int32

const (
	Role_GUEST  Role = 0
	Role_MEMBER Role = 1
)

var Role_name = map[int32]string{
	0: "GUEST",
	1: "MEMBER",
}

var Role_value = map[string]int32{
	"GUEST":  0,
	"MEMBER": 1,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type CheckMessage struct {
	Check                string   `protobuf:"bytes,1,opt,name=check,proto3" json:"check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckMessage) Reset()         { *m = CheckMessage{} }
func (m *CheckMessage) String() string { return proto.CompactTextString(m) }
func (*CheckMessage) ProtoMessage()    {}
func (*CheckMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *CheckMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckMessage.Unmarshal(m, b)
}
func (m *CheckMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckMessage.Marshal(b, m, deterministic)
}
func (m *CheckMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckMessage.Merge(m, src)
}
func (m *CheckMessage) XXX_Size() int {
	return xxx_messageInfo_CheckMessage.Size(m)
}
func (m *CheckMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CheckMessage proto.InternalMessageInfo

func (m *CheckMessage) GetCheck() string {
	if m != nil {
		return m.Check
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type User struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 Role     `protobuf:"varint,2,opt,name=role,proto3,enum=api.Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

func init() {
	proto.RegisterEnum("api.Role", Role_name, Role_value)
	proto.RegisterType((*CheckMessage)(nil), "api.CheckMessage")
	proto.RegisterType((*Void)(nil), "api.Void")
	proto.RegisterType((*User)(nil), "api.User")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xdb, 0xda, 0x56, 0x3a, 0xea, 0xb2, 0x0e, 0x22, 0x4b, 0x65, 0x61, 0x09, 0x22, 0x8b,
	0x87, 0x54, 0x2a, 0xfe, 0x00, 0x95, 0xe2, 0xc5, 0x82, 0x74, 0x5d, 0xef, 0xdd, 0x66, 0xac, 0xc1,
	0x6a, 0x6a, 0x53, 0x05, 0xff, 0xbd, 0x24, 0x41, 0xf4, 0xb2, 0xb7, 0x79, 0x93, 0x97, 0x37, 0xdf,
	0x83, 0xa4, 0xee, 0x25, 0xef, 0x07, 0x35, 0x2a, 0xdc, 0xa9, 0x7b, 0x99, 0x9e, 0xb4, 0x4a, 0xb5,
	0x1d, 0x65, 0x76, 0xb5, 0xf9, 0x7c, 0xce, 0xe8, 0xad, 0x1f, 0xbf, 0x9d, 0x83, 0x9d, 0xc2, 0xfe,
	0xed, 0x0b, 0x35, 0xaf, 0x25, 0x69, 0x5d, 0xb7, 0x84, 0x47, 0x10, 0x35, 0x46, 0xcf, 0xfc, 0x85,
	0xbf, 0x4c, 0x2a, 0x27, 0x58, 0x0c, 0xe1, 0x93, 0x92, 0x82, 0x5d, 0x41, 0xb8, 0xd6, 0x34, 0xe0,
	0x04, 0x02, 0x29, 0xac, 0xe5, 0xa0, 0x0a, 0xa4, 0xc0, 0x39, 0x84, 0x83, 0xea, 0x68, 0x16, 0x2c,
	0xfc, 0xe5, 0x24, 0x4f, 0xb8, 0x21, 0xa8, 0x54, 0x47, 0x95, 0x5d, 0x9f, 0xcf, 0x21, 0x34, 0x0a,
	0x13, 0x88, 0xee, 0xd6, 0xc5, 0xea, 0x71, 0xea, 0x21, 0x40, 0x5c, 0x16, 0xe5, 0x4d, 0x51, 0x4d,
	0xfd, 0x3c, 0x83, 0xc8, 0x32, 0xe0, 0x19, 0x84, 0x0f, 0xf2, 0xbd, 0x45, 0x17, 0x60, 0x2e, 0xa6,
	0x87, 0x76, 0xfc, 0x8f, 0xc8, 0xbc, 0xfc, 0x03, 0xf6, 0x0c, 0xc6, 0x8a, 0x86, 0x2f, 0xd9, 0x10,
	0x72, 0xd8, 0xbd, 0x16, 0xc2, 0x82, 0xb9, 0x9f, 0x66, 0x4c, 0x8f, 0xb9, 0xeb, 0xcd, 0x7f, 0x7b,
	0xf3, 0xc2, 0xf4, 0x66, 0x1e, 0xe6, 0x90, 0xdc, 0x4b, 0x3d, 0x1a, 0x97, 0xc6, 0x2d, 0xb6, 0xf4,
	0x2f, 0x89, 0x79, 0x17, 0xfe, 0x26, 0xb6, 0xcf, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0x86, 0x9f, 0xf1, 0x5d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckClient interface {
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*CheckMessage, error)
}

type checkClient struct {
	cc *grpc.ClientConn
}

func NewCheckClient(cc *grpc.ClientConn) CheckClient {
	return &checkClient{cc}
}

func (c *checkClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*CheckMessage, error) {
	out := new(CheckMessage)
	err := c.cc.Invoke(ctx, "/api.Check/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServer is the server API for Check service.
type CheckServer interface {
	Ping(context.Context, *Void) (*CheckMessage, error)
}

// UnimplementedCheckServer can be embedded to have forward compatible implementations.
type UnimplementedCheckServer struct {
}

func (*UnimplementedCheckServer) Ping(ctx context.Context, req *Void) (*CheckMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterCheckServer(s *grpc.Server, srv CheckServer) {
	s.RegisterService(&_Check_serviceDesc, srv)
}

func _Check_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Check/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Check_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Check",
	HandlerType: (*CheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Check_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
	ListUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (UserService_ListUsersClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (UserService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/api.UserService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *User) (*empty.Empty, error)
	ListUsers(*empty.Empty, UserService_ListUsersServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) AddUser(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedUserServiceServer) ListUsers(req *empty.Empty, srv UserService_ListUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).ListUsers(m, &userServiceListUsersServer{stream})
}

type UserService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _UserService_ListUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
