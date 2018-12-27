// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/register.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Register struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	MobileNumber         string   `protobuf:"bytes,4,opt,name=mobile_number,json=mobileNumber,proto3" json:"mobile_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{0}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Register) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Register) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Register) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

type RegisterRequest struct {
	Register             *Register `protobuf:"bytes,1,opt,name=register,proto3" json:"register,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetRegister() *Register {
	if m != nil {
		return m.Register
	}
	return nil
}

type RegisterResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{2}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Register)(nil), "api.Register")
	proto.RegisterType((*RegisterRequest)(nil), "api.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "api.RegisterResponse")
}

func init() { proto.RegisterFile("api/register.proto", fileDescriptor_7ca244b99eeb2808) }

var fileDescriptor_7ca244b99eeb2808 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0xad, 0x53, 0x69, 0x4f, 0x8b, 0x72, 0x4c, 0x28, 0x13, 0x41, 0xea, 0x8b, 0xbe, 0x54,
	0x98, 0xaf, 0xe2, 0x47, 0x18, 0x58, 0x3f, 0xc0, 0xb8, 0xca, 0x29, 0x07, 0x4d, 0x13, 0x93, 0xcc,
	0x77, 0xbf, 0xb9, 0xec, 0xb2, 0x76, 0xb8, 0xb7, 0xdc, 0xef, 0x97, 0xcb, 0xfd, 0x73, 0x80, 0xe4,
	0xe4, 0xc9, 0xf3, 0x97, 0x84, 0xc8, 0xbe, 0x71, 0xde, 0x46, 0x8b, 0x33, 0x72, 0x52, 0xff, 0x66,
	0x90, 0xb7, 0x3b, 0x8e, 0xb7, 0x00, 0x9f, 0xe2, 0x43, 0x5c, 0x0f, 0x64, 0xb8, 0xca, 0xee, 0xb2,
	0x87, 0xa2, 0x2d, 0x94, 0xac, 0xc8, 0x30, 0xde, 0x40, 0xd1, 0xd3, 0x68, 0x8f, 0xd5, 0xe6, 0x5b,
	0xa0, 0x72, 0x0e, 0xa7, 0x6c, 0x48, 0xfa, 0x6a, 0xa6, 0x22, 0x15, 0x78, 0x0f, 0xa5, 0xb1, 0x9d,
	0xf4, 0xbc, 0x1e, 0x36, 0xa6, 0x63, 0x5f, 0x9d, 0xa8, 0xbd, 0x48, 0x70, 0xa5, 0xac, 0x7e, 0x81,
	0xcb, 0x31, 0x42, 0xcb, 0xdf, 0x1b, 0x0e, 0x11, 0x1f, 0x21, 0x1f, 0xd3, 0x6a, 0x8e, 0xf3, 0x65,
	0xd9, 0x90, 0x93, 0x66, 0xba, 0x37, 0xe9, 0xba, 0x81, 0xab, 0x7d, 0x77, 0x70, 0x76, 0x08, 0x8c,
	0x8b, 0x6d, 0x7b, 0x3a, 0xef, 0xbe, 0x31, 0xd5, 0xcb, 0xb7, 0xfd, 0xb4, 0x77, 0xf6, 0x3f, 0xf2,
	0xc1, 0xf8, 0x0a, 0xe5, 0x88, 0x28, 0x8a, 0x1d, 0x70, 0xfe, 0x7f, 0x58, 0x0a, 0xb5, 0xb8, 0x3e,
	0xa0, 0xe9, 0xc1, 0xfa, 0xa8, 0x3b, 0xd3, 0x85, 0x3e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x6c, 0xea, 0x4d, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterServiceClient interface {
	Registeration(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registerServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegisterServiceClient(cc *grpc.ClientConn) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) Registeration(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/api.RegisterService/Registeration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
type RegisterServiceServer interface {
	Registeration(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

func RegisterRegisterServiceServer(s *grpc.Server, srv RegisterServiceServer) {
	s.RegisterService(&_RegisterService_serviceDesc, srv)
}

func _RegisterService_Registeration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).Registeration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RegisterService/Registeration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).Registeration(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registeration",
			Handler:    _RegisterService_Registeration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/register.proto",
}