// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/blog.proto

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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Heading              string   `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_696fc2b6c579c277, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetHeading() string {
	if m != nil {
		return m.Heading
	}
	return ""
}

func (m *Blog) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type BlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=Blog,proto3" json:"Blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogRequest) Reset()         { *m = BlogRequest{} }
func (m *BlogRequest) String() string { return proto.CompactTextString(m) }
func (*BlogRequest) ProtoMessage()    {}
func (*BlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_696fc2b6c579c277, []int{1}
}

func (m *BlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogRequest.Unmarshal(m, b)
}
func (m *BlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogRequest.Marshal(b, m, deterministic)
}
func (m *BlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogRequest.Merge(m, src)
}
func (m *BlogRequest) XXX_Size() int {
	return xxx_messageInfo_BlogRequest.Size(m)
}
func (m *BlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogRequest proto.InternalMessageInfo

func (m *BlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type BlogResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogResponse) Reset()         { *m = BlogResponse{} }
func (m *BlogResponse) String() string { return proto.CompactTextString(m) }
func (*BlogResponse) ProtoMessage()    {}
func (*BlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_696fc2b6c579c277, []int{2}
}

func (m *BlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogResponse.Unmarshal(m, b)
}
func (m *BlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogResponse.Marshal(b, m, deterministic)
}
func (m *BlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogResponse.Merge(m, src)
}
func (m *BlogResponse) XXX_Size() int {
	return xxx_messageInfo_BlogResponse.Size(m)
}
func (m *BlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogResponse proto.InternalMessageInfo

func (m *BlogResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Blog)(nil), "api.blog")
	proto.RegisterType((*BlogRequest)(nil), "api.BlogRequest")
	proto.RegisterType((*BlogResponse)(nil), "api.BlogResponse")
}

func init() { proto.RegisterFile("api/blog.proto", fileDescriptor_696fc2b6c579c277) }

var fileDescriptor_696fc2b6c579c277 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xdd, 0xca, 0x82, 0x40,
	0x10, 0x40, 0x3f, 0x7f, 0xf8, 0xca, 0x31, 0xa4, 0xe6, 0x6a, 0x11, 0x82, 0xd8, 0xab, 0x88, 0x30,
	0xb0, 0x9e, 0xa0, 0x7a, 0x02, 0x7b, 0x82, 0xb5, 0x5d, 0x6c, 0x41, 0xdc, 0x4d, 0x2d, 0xe8, 0xed,
	0xc3, 0xd1, 0xc2, 0xee, 0xe6, 0x1c, 0x96, 0x33, 0x3b, 0x10, 0x09, 0xab, 0x77, 0x79, 0x69, 0x8a,
	0xc4, 0xd6, 0xa6, 0x35, 0xe8, 0x09, 0xab, 0xf9, 0x19, 0xfc, 0x4e, 0x61, 0x04, 0xae, 0x96, 0xcc,
	0x59, 0x39, 0xeb, 0x20, 0x73, 0xb5, 0x44, 0x06, 0x93, 0x9b, 0x12, 0x52, 0x57, 0x05, 0x73, 0x49,
	0x7e, 0x10, 0x11, 0xfc, 0xdc, 0xc8, 0x17, 0xf3, 0x48, 0xd3, 0xcc, 0xb7, 0x10, 0x1e, 0x4b, 0x53,
	0x64, 0xea, 0xfe, 0x50, 0x4d, 0x8b, 0x4b, 0xf0, 0x3b, 0xa4, 0x5c, 0x98, 0x06, 0x89, 0xb0, 0x3a,
	0xe9, 0xb6, 0x64, 0xa4, 0xf9, 0x06, 0x66, 0xfd, 0xeb, 0xc6, 0x9a, 0xaa, 0x51, 0x18, 0xc3, 0xb4,
	0x1e, 0xe6, 0xe1, 0x07, 0x5f, 0x4e, 0x4f, 0x7d, 0xf9, 0xa2, 0xea, 0xa7, 0xbe, 0x2a, 0x3c, 0xfc,
	0xe2, 0x9c, 0xd2, 0xa3, 0xd5, 0xf1, 0x62, 0x64, 0xfa, 0x04, 0xff, 0xcb, 0xff, 0xe9, 0xe0, 0xfd,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x0c, 0x58, 0x9e, 0x02, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	BlogService(ctx context.Context, in *BlogRequest, opts ...grpc.CallOption) (*BlogResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) BlogService(ctx context.Context, in *BlogRequest, opts ...grpc.CallOption) (*BlogResponse, error) {
	out := new(BlogResponse)
	err := c.cc.Invoke(ctx, "/api.BlogService/BlogService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	BlogService(context.Context, *BlogRequest) (*BlogResponse, error)
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_BlogService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).BlogService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BlogService/BlogService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).BlogService(ctx, req.(*BlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlogService",
			Handler:    _BlogService_BlogService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blog.proto",
}
