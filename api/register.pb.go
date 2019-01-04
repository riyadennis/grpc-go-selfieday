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

type RegisterStreamRequest struct {
	Register             *Register `protobuf:"bytes,1,opt,name=register,proto3" json:"register,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterStreamRequest) Reset()         { *m = RegisterStreamRequest{} }
func (m *RegisterStreamRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterStreamRequest) ProtoMessage()    {}
func (*RegisterStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{3}
}

func (m *RegisterStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterStreamRequest.Unmarshal(m, b)
}
func (m *RegisterStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterStreamRequest.Marshal(b, m, deterministic)
}
func (m *RegisterStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterStreamRequest.Merge(m, src)
}
func (m *RegisterStreamRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterStreamRequest.Size(m)
}
func (m *RegisterStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterStreamRequest proto.InternalMessageInfo

func (m *RegisterStreamRequest) GetRegister() *Register {
	if m != nil {
		return m.Register
	}
	return nil
}

type RegisterStreamResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterStreamResponse) Reset()         { *m = RegisterStreamResponse{} }
func (m *RegisterStreamResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterStreamResponse) ProtoMessage()    {}
func (*RegisterStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{4}
}

func (m *RegisterStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterStreamResponse.Unmarshal(m, b)
}
func (m *RegisterStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterStreamResponse.Marshal(b, m, deterministic)
}
func (m *RegisterStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterStreamResponse.Merge(m, src)
}
func (m *RegisterStreamResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterStreamResponse.Size(m)
}
func (m *RegisterStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterStreamResponse proto.InternalMessageInfo

func (m *RegisterStreamResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type Chat struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Heading              string   `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{5}
}

func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (m *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(m, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Chat) GetHeading() string {
	if m != nil {
		return m.Heading
	}
	return ""
}

func (m *Chat) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type ChatRequest struct {
	Chat                 *Chat    `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatRequest) Reset()         { *m = ChatRequest{} }
func (m *ChatRequest) String() string { return proto.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()    {}
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{6}
}

func (m *ChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRequest.Unmarshal(m, b)
}
func (m *ChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRequest.Marshal(b, m, deterministic)
}
func (m *ChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRequest.Merge(m, src)
}
func (m *ChatRequest) XXX_Size() int {
	return xxx_messageInfo_ChatRequest.Size(m)
}
func (m *ChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRequest proto.InternalMessageInfo

func (m *ChatRequest) GetChat() *Chat {
	if m != nil {
		return m.Chat
	}
	return nil
}

type ChatResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatResponse) Reset()         { *m = ChatResponse{} }
func (m *ChatResponse) String() string { return proto.CompactTextString(m) }
func (*ChatResponse) ProtoMessage()    {}
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ca244b99eeb2808, []int{7}
}

func (m *ChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatResponse.Unmarshal(m, b)
}
func (m *ChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatResponse.Marshal(b, m, deterministic)
}
func (m *ChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatResponse.Merge(m, src)
}
func (m *ChatResponse) XXX_Size() int {
	return xxx_messageInfo_ChatResponse.Size(m)
}
func (m *ChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChatResponse proto.InternalMessageInfo

func (m *ChatResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Register)(nil), "api.Register")
	proto.RegisterType((*RegisterRequest)(nil), "api.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "api.RegisterResponse")
	proto.RegisterType((*RegisterStreamRequest)(nil), "api.RegisterStreamRequest")
	proto.RegisterType((*RegisterStreamResponse)(nil), "api.RegisterStreamResponse")
	proto.RegisterType((*Chat)(nil), "api.chat")
	proto.RegisterType((*ChatRequest)(nil), "api.ChatRequest")
	proto.RegisterType((*ChatResponse)(nil), "api.ChatResponse")
}

func init() { proto.RegisterFile("api/register.proto", fileDescriptor_7ca244b99eeb2808) }

var fileDescriptor_7ca244b99eeb2808 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x4f, 0xea, 0x50,
	0x10, 0xc6, 0x6f, 0x79, 0xdc, 0x4b, 0x07, 0xb8, 0x97, 0x3b, 0x01, 0xd3, 0x94, 0x90, 0x98, 0xba,
	0x41, 0x63, 0xaa, 0x41, 0xe3, 0xca, 0xb8, 0x50, 0xb7, 0xb2, 0x28, 0xba, 0x26, 0xa7, 0x74, 0x84,
	0x93, 0xd0, 0x87, 0xed, 0xc1, 0x84, 0xa5, 0xff, 0xa9, 0x7f, 0x8a, 0x61, 0xda, 0xc3, 0x2b, 0x26,
	0x1a, 0x77, 0x9d, 0xdf, 0x37, 0x33, 0xe7, 0x3b, 0xdf, 0x29, 0xa0, 0x48, 0xe4, 0x59, 0x4a, 0x53,
	0x99, 0x29, 0x4a, 0xdd, 0x24, 0x8d, 0x55, 0x8c, 0x65, 0x91, 0x48, 0xe7, 0xcd, 0x80, 0x9a, 0x57,
	0x70, 0xec, 0x01, 0x3c, 0xcb, 0x34, 0x53, 0xe3, 0x48, 0x84, 0x64, 0x19, 0x87, 0x46, 0xdf, 0xf4,
	0x4c, 0x26, 0x43, 0x11, 0x12, 0x76, 0xc1, 0x9c, 0x0b, 0xad, 0x96, 0x58, 0xad, 0xad, 0x00, 0x8b,
	0x6d, 0xa8, 0x52, 0x28, 0xe4, 0xdc, 0x2a, 0xb3, 0x90, 0x17, 0x78, 0x04, 0xcd, 0x30, 0xf6, 0xe5,
	0x9c, 0xc6, 0xd1, 0x22, 0xf4, 0x29, 0xb5, 0x2a, 0xac, 0x36, 0x72, 0x38, 0x64, 0xe6, 0x5c, 0xc3,
	0x3f, 0x6d, 0xc1, 0xa3, 0x97, 0x05, 0x65, 0x0a, 0x8f, 0xa1, 0xa6, 0xdd, 0xb2, 0x8f, 0xfa, 0xa0,
	0xe9, 0x8a, 0x44, 0xba, 0xeb, 0xbe, 0xb5, 0xec, 0xb8, 0xd0, 0xda, 0x4c, 0x67, 0x49, 0x1c, 0x65,
	0x84, 0xf6, 0x6a, 0x3c, 0xff, 0x2e, 0xae, 0xb1, 0xae, 0x9d, 0x5b, 0xe8, 0xe8, 0xfe, 0x91, 0x4a,
	0x49, 0x84, 0x3f, 0x38, 0xf3, 0x12, 0x0e, 0xf6, 0x77, 0x7c, 0xe3, 0xe4, 0x7b, 0xa8, 0x4c, 0x66,
	0x42, 0xe1, 0x5f, 0x28, 0xc9, 0x80, 0xd5, 0xaa, 0x57, 0x92, 0x01, 0x5a, 0xf0, 0x67, 0x46, 0x22,
	0x90, 0xd1, 0xb4, 0x48, 0x55, 0x97, 0x88, 0x50, 0xf1, 0xe3, 0x60, 0x59, 0x64, 0xca, 0xdf, 0xce,
	0x29, 0xd4, 0xef, 0x66, 0x42, 0x69, 0xd7, 0xbd, 0x7c, 0x69, 0xe1, 0xd8, 0x64, 0xc7, 0x2b, 0xe0,
	0x31, 0x76, 0x4e, 0xa0, 0x91, 0x77, 0x7f, 0xed, 0x6f, 0xf0, 0x6e, 0x6c, 0x1e, 0x62, 0x44, 0xe9,
	0xab, 0x9c, 0x10, 0xde, 0x40, 0x53, 0x23, 0xa1, 0x64, 0x1c, 0x61, 0x7b, 0x37, 0x93, 0xdc, 0x85,
	0xdd, 0xd9, 0xa3, 0xc5, 0x8d, 0x7f, 0xe1, 0xd3, 0x26, 0x29, 0x9e, 0x7f, 0x10, 0xd1, 0xf2, 0x51,
	0x86, 0x94, 0xa1, 0xbd, 0x33, 0xb2, 0xf3, 0x14, 0x76, 0xf7, 0x53, 0x4d, 0x2f, 0x3d, 0x37, 0xf0,
	0x2a, 0x0f, 0x41, 0xbb, 0x6c, 0x71, 0xff, 0x56, 0x2c, 0xf6, 0xff, 0x2d, 0xa2, 0xe7, 0xfa, 0x86,
	0xff, 0x9b, 0x7f, 0xfd, 0x8b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x2c, 0x90, 0x71, 0x10,
	0x03, 0x00, 0x00,
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
	// Registeration unary call
	Registeration(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// RegisterationManyTimes server streaming
	RegisterationManyTimes(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (RegisterService_RegisterationManyTimesClient, error)
	// ChatService client streaming
	ChatService(ctx context.Context, opts ...grpc.CallOption) (RegisterService_ChatServiceClient, error)
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

func (c *registerServiceClient) RegisterationManyTimes(ctx context.Context, in *RegisterStreamRequest, opts ...grpc.CallOption) (RegisterService_RegisterationManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RegisterService_serviceDesc.Streams[0], "/api.RegisterService/RegisterationManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &registerServiceRegisterationManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RegisterService_RegisterationManyTimesClient interface {
	Recv() (*RegisterStreamResponse, error)
	grpc.ClientStream
}

type registerServiceRegisterationManyTimesClient struct {
	grpc.ClientStream
}

func (x *registerServiceRegisterationManyTimesClient) Recv() (*RegisterStreamResponse, error) {
	m := new(RegisterStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registerServiceClient) ChatService(ctx context.Context, opts ...grpc.CallOption) (RegisterService_ChatServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RegisterService_serviceDesc.Streams[1], "/api.RegisterService/ChatService", opts...)
	if err != nil {
		return nil, err
	}
	x := &registerServiceChatServiceClient{stream}
	return x, nil
}

type RegisterService_ChatServiceClient interface {
	Send(*ChatRequest) error
	CloseAndRecv() (*ChatResponse, error)
	grpc.ClientStream
}

type registerServiceChatServiceClient struct {
	grpc.ClientStream
}

func (x *registerServiceChatServiceClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registerServiceChatServiceClient) CloseAndRecv() (*ChatResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterServiceServer is the server API for RegisterService service.
type RegisterServiceServer interface {
	// Registeration unary call
	Registeration(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// RegisterationManyTimes server streaming
	RegisterationManyTimes(*RegisterStreamRequest, RegisterService_RegisterationManyTimesServer) error
	// ChatService client streaming
	ChatService(RegisterService_ChatServiceServer) error
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

func _RegisterService_RegisterationManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegisterServiceServer).RegisterationManyTimes(m, &registerServiceRegisterationManyTimesServer{stream})
}

type RegisterService_RegisterationManyTimesServer interface {
	Send(*RegisterStreamResponse) error
	grpc.ServerStream
}

type registerServiceRegisterationManyTimesServer struct {
	grpc.ServerStream
}

func (x *registerServiceRegisterationManyTimesServer) Send(m *RegisterStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RegisterService_ChatService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegisterServiceServer).ChatService(&registerServiceChatServiceServer{stream})
}

type RegisterService_ChatServiceServer interface {
	SendAndClose(*ChatResponse) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type registerServiceChatServiceServer struct {
	grpc.ServerStream
}

func (x *registerServiceChatServiceServer) SendAndClose(m *ChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registerServiceChatServiceServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterationManyTimes",
			Handler:       _RegisterService_RegisterationManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChatService",
			Handler:       _RegisterService_ChatService_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/register.proto",
}
