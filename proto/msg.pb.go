// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AddUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Uid                  string   `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (m *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(m, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddUserRequest) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *AddUserRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetAllUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUsersResponse) Reset()         { *m = GetAllUsersResponse{} }
func (m *GetAllUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersResponse) ProtoMessage()    {}
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *GetAllUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersResponse.Unmarshal(m, b)
}
func (m *GetAllUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetAllUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersResponse.Merge(m, src)
}
func (m *GetAllUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersResponse.Size(m)
}
func (m *GetAllUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersResponse proto.InternalMessageInfo

func (m *GetAllUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetSingleUserRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSingleUserRequest) Reset()         { *m = GetSingleUserRequest{} }
func (m *GetSingleUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetSingleUserRequest) ProtoMessage()    {}
func (*GetSingleUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *GetSingleUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSingleUserRequest.Unmarshal(m, b)
}
func (m *GetSingleUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSingleUserRequest.Marshal(b, m, deterministic)
}
func (m *GetSingleUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleUserRequest.Merge(m, src)
}
func (m *GetSingleUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetSingleUserRequest.Size(m)
}
func (m *GetSingleUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleUserRequest proto.InternalMessageInfo

func (m *GetSingleUserRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Uid                  string   `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetAllUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUsersRequest) Reset()         { *m = GetAllUsersRequest{} }
func (m *GetAllUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersRequest) ProtoMessage()    {}
func (*GetAllUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *GetAllUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersRequest.Unmarshal(m, b)
}
func (m *GetAllUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetAllUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersRequest.Merge(m, src)
}
func (m *GetAllUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersRequest.Size(m)
}
func (m *GetAllUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddUserRequest)(nil), "proto.AddUserRequest")
	proto.RegisterType((*GetAllUsersResponse)(nil), "proto.GetAllUsersResponse")
	proto.RegisterType((*GetSingleUserRequest)(nil), "proto.GetSingleUserRequest")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*GetAllUsersRequest)(nil), "proto.GetAllUsersRequest")
}

func init() {
	proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899)
}

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0x27, 0xcd, 0xaa, 0xec, 0x15, 0x87, 0x3c, 0x2b, 0xd4, 0x7a, 0xa9, 0x3d, 0xf5, 0xe2, 0x0e,
	0xf3, 0xa2, 0xc7, 0x82, 0x30, 0xf0, 0x24, 0x19, 0x7e, 0x80, 0x69, 0x1e, 0x25, 0xd0, 0xad, 0xb3,
	0x69, 0xfd, 0x74, 0x7e, 0x38, 0x49, 0x1a, 0x6c, 0xb3, 0xed, 0x94, 0x97, 0xf7, 0x7b, 0xfc, 0xfe,
	0xc1, 0x7c, 0xa7, 0xab, 0xe5, 0xa1, 0x6d, 0xba, 0x06, 0x43, 0xfb, 0xe4, 0x6f, 0xb0, 0x28, 0xa5,
	0xfc, 0xd0, 0xd4, 0x0a, 0xfa, 0xee, 0x49, 0x77, 0x18, 0x43, 0x48, 0xbb, 0xad, 0xaa, 0x13, 0x96,
	0xb1, 0x62, 0x2e, 0x86, 0x0f, 0x5e, 0x03, 0xdf, 0x56, 0x94, 0x04, 0x19, 0x2b, 0xb8, 0x30, 0xa3,
	0xd9, 0xf4, 0x4a, 0x26, 0xdc, 0x5e, 0x99, 0x31, 0x7f, 0x86, 0x9b, 0x35, 0x75, 0x65, 0x5d, 0x1b,
	0x3a, 0x2d, 0x48, 0x1f, 0x9a, 0xbd, 0x26, 0x7c, 0x80, 0xb0, 0x37, 0x8b, 0x84, 0x65, 0xbc, 0x88,
	0x56, 0xd1, 0x60, 0x60, 0x69, 0x35, 0x07, 0x24, 0x2f, 0x20, 0x5e, 0x53, 0xb7, 0x51, 0xfb, 0xaa,
	0xa6, 0xa9, 0x17, 0xa7, 0xc1, 0x46, 0x8d, 0x77, 0x98, 0x99, 0x03, 0x5c, 0x40, 0xe0, 0x00, 0x2e,
	0x02, 0x25, 0xcf, 0xf8, 0xfb, 0xcf, 0xc1, 0x8f, 0x72, 0x18, 0xc6, 0xd9, 0xc8, 0x18, 0x03, 0x7a,
	0xae, 0xad, 0xf2, 0xea, 0x97, 0x01, 0x94, 0x52, 0x6e, 0xa8, 0xfd, 0x51, 0x5f, 0x84, 0x8f, 0x70,
	0xe9, 0x6a, 0xc2, 0x5b, 0xe7, 0xdf, 0xaf, 0x2d, 0x9d, 0xc6, 0xc2, 0x57, 0x88, 0x26, 0x9c, 0x78,
	0xe7, 0xb0, 0x53, 0x9d, 0x34, 0x3d, 0x07, 0xb9, 0xe2, 0x5e, 0xe0, 0xca, 0x6b, 0x05, 0xef, 0xc7,
	0xe3, 0x93, 0xae, 0x3c, 0x03, 0x9f, 0x17, 0x76, 0x7e, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x9d, 0xf0, 0x5a, 0xf1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	GetSingleUser(ctx context.Context, in *GetSingleUserRequest, opts ...grpc.CallOption) (*User, error)
}

type addServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddServiceClient(cc grpc.ClientConnInterface) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.AddService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) GetSingleUser(ctx context.Context, in *GetSingleUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.AddService/GetSingleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*User, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	GetSingleUser(context.Context, *GetSingleUserRequest) (*User, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) AddUser(ctx context.Context, req *AddUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedAddServiceServer) GetAllUsers(ctx context.Context, req *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (*UnimplementedAddServiceServer) GetSingleUser(ctx context.Context, req *GetSingleUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleUser not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_GetSingleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetSingleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/GetSingleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetSingleUser(ctx, req.(*GetSingleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _AddService_AddUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _AddService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetSingleUser",
			Handler:    _AddService_GetSingleUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}
