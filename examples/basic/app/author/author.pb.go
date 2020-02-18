// Code generated by protoc-gen-go. DO NOT EDIT.
// source: author/author.proto

package author

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
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

type AuthorType int32

const (
	AuthorType_NORMAL  AuthorType = 0
	AuthorType_SPECIAL AuthorType = 1
)

var AuthorType_name = map[int32]string{
	0: "NORMAL",
	1: "SPECIAL",
}

var AuthorType_value = map[string]int32{
	"NORMAL":  0,
	"SPECIAL": 1,
}

func (x AuthorType) String() string {
	return proto.EnumName(AuthorType_name, int32(x))
}

func (AuthorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{0}
}

type Author struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListAuthorsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAuthorsRequest) Reset()         { *m = ListAuthorsRequest{} }
func (m *ListAuthorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAuthorsRequest) ProtoMessage()    {}
func (*ListAuthorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{1}
}

func (m *ListAuthorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthorsRequest.Unmarshal(m, b)
}
func (m *ListAuthorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthorsRequest.Marshal(b, m, deterministic)
}
func (m *ListAuthorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthorsRequest.Merge(m, src)
}
func (m *ListAuthorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAuthorsRequest.Size(m)
}
func (m *ListAuthorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthorsRequest proto.InternalMessageInfo

type ListAuthorsResponse struct {
	Authors              []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListAuthorsResponse) Reset()         { *m = ListAuthorsResponse{} }
func (m *ListAuthorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAuthorsResponse) ProtoMessage()    {}
func (*ListAuthorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{2}
}

func (m *ListAuthorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthorsResponse.Unmarshal(m, b)
}
func (m *ListAuthorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthorsResponse.Marshal(b, m, deterministic)
}
func (m *ListAuthorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthorsResponse.Merge(m, src)
}
func (m *ListAuthorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAuthorsResponse.Size(m)
}
func (m *ListAuthorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthorsResponse proto.InternalMessageInfo

func (m *ListAuthorsResponse) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

type GetAuthorRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthorRequest) Reset()         { *m = GetAuthorRequest{} }
func (m *GetAuthorRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthorRequest) ProtoMessage()    {}
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{3}
}

func (m *GetAuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorRequest.Unmarshal(m, b)
}
func (m *GetAuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorRequest.Marshal(b, m, deterministic)
}
func (m *GetAuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorRequest.Merge(m, src)
}
func (m *GetAuthorRequest) XXX_Size() int {
	return xxx_messageInfo_GetAuthorRequest.Size(m)
}
func (m *GetAuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorRequest proto.InternalMessageInfo

func (m *GetAuthorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("some.app.v1.author.AuthorType", AuthorType_name, AuthorType_value)
	proto.RegisterType((*Author)(nil), "some.app.v1.author.Author")
	proto.RegisterType((*ListAuthorsRequest)(nil), "some.app.v1.author.ListAuthorsRequest")
	proto.RegisterType((*ListAuthorsResponse)(nil), "some.app.v1.author.ListAuthorsResponse")
	proto.RegisterType((*GetAuthorRequest)(nil), "some.app.v1.author.GetAuthorRequest")
}

func init() { proto.RegisterFile("author/author.proto", fileDescriptor_41efdaee09960382) }

var fileDescriptor_41efdaee09960382 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdd, 0x4e, 0xea, 0x40,
	0x10, 0xc7, 0x4f, 0xcf, 0x39, 0x29, 0x87, 0xe1, 0x68, 0x9a, 0x05, 0x4d, 0xd3, 0x78, 0x41, 0x1a,
	0x45, 0x62, 0xc2, 0x2e, 0xa2, 0x17, 0xc4, 0x3b, 0x24, 0xc4, 0x18, 0xf1, 0x23, 0xc5, 0x2b, 0xbd,
	0x5a, 0x9a, 0x4d, 0xdb, 0xa4, 0x65, 0x97, 0xee, 0x16, 0xe1, 0x15, 0x78, 0x1f, 0x5e, 0xcf, 0x98,
	0xb0, 0x05, 0x3f, 0x20, 0x7a, 0x35, 0x93, 0xdd, 0xdf, 0xce, 0xfc, 0xff, 0x3b, 0x03, 0x65, 0x9a,
	0xa9, 0x90, 0xa7, 0x44, 0x07, 0x2c, 0x52, 0xae, 0x38, 0x42, 0x92, 0x27, 0x0c, 0x53, 0x21, 0xf0,
	0xe4, 0x14, 0xeb, 0x1b, 0x67, 0x2f, 0x48, 0xa9, 0x08, 0xc7, 0x31, 0xc9, 0xa3, 0x46, 0xdd, 0x03,
	0x30, 0x3b, 0x4b, 0x00, 0x21, 0xf8, 0x3b, 0xa2, 0x09, 0xb3, 0x7f, 0x57, 0x8d, 0x7a, 0xd1, 0x5b,
	0xe6, 0x6e, 0x05, 0x50, 0x3f, 0x92, 0x4a, 0x13, 0xd2, 0x63, 0xe3, 0x8c, 0x49, 0xe5, 0xde, 0x40,
	0xf9, 0xd3, 0xa9, 0x14, 0x7c, 0x24, 0x19, 0x3a, 0x87, 0x82, 0xee, 0x25, 0x6d, 0xa3, 0xfa, 0xa7,
	0x5e, 0x6a, 0x39, 0x78, 0x53, 0x07, 0xd6, 0xaf, 0xbc, 0x15, 0xea, 0xd6, 0xc0, 0xba, 0x62, 0x79,
	0xad, 0xbc, 0xc1, 0x5a, 0x8a, 0xf1, 0x2e, 0xe5, 0xe4, 0x08, 0x40, 0x43, 0x8f, 0x33, 0xc1, 0x10,
	0x80, 0x79, 0x77, 0xef, 0xdd, 0x76, 0xfa, 0xd6, 0x2f, 0x54, 0x82, 0xc2, 0xe0, 0xa1, 0xd7, 0xbd,
	0xee, 0xf4, 0x2d, 0xa3, 0xf5, 0x6a, 0xc0, 0x8e, 0xe6, 0x06, 0x2c, 0x9d, 0x44, 0x3e, 0x43, 0x53,
	0x28, 0x7d, 0x50, 0x8b, 0x6a, 0xdb, 0x44, 0x6d, 0x9a, 0x74, 0x8e, 0x7f, 0xe4, 0xb4, 0x6d, 0xd7,
	0x9e, 0x2f, 0xec, 0xca, 0xda, 0xba, 0x5b, 0x74, 0x56, 0x29, 0x7a, 0x86, 0xe2, 0xda, 0x1a, 0x3a,
	0xdc, 0x56, 0xef, 0xab, 0x73, 0xe7, 0x9b, 0x2f, 0x73, 0xff, 0xcf, 0x17, 0xf6, 0x3f, 0x30, 0xf3,
	0x79, 0xee, 0xcf, 0x17, 0x36, 0x82, 0xdd, 0x98, 0xfb, 0x34, 0x0e, 0xb9, 0x54, 0x17, 0xed, 0x66,
	0xbb, 0x69, 0x19, 0x97, 0xbd, 0xa7, 0x6e, 0x10, 0xa9, 0x30, 0x1b, 0x62, 0x9f, 0x27, 0x64, 0x26,
	0xb3, 0x20, 0x4a, 0xb8, 0xe2, 0x24, 0x48, 0x85, 0xdf, 0xc8, 0x67, 0xdf, 0x08, 0xa8, 0x62, 0x2f,
	0x74, 0x46, 0xd8, 0x94, 0x26, 0x22, 0x66, 0x92, 0x0c, 0xa9, 0x8c, 0x7c, 0x42, 0x85, 0xc8, 0x17,
	0x69, 0x68, 0x2e, 0xd7, 0xe3, 0xec, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x8a, 0x3b, 0x1d, 0x60,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorServiceClient interface {
	ListAuthors(ctx context.Context, in *ListAuthorsRequest, opts ...grpc.CallOption) (*ListAuthorsResponse, error)
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error)
}

type authorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorServiceClient(cc *grpc.ClientConn) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) ListAuthors(ctx context.Context, in *ListAuthorsRequest, opts ...grpc.CallOption) (*ListAuthorsResponse, error) {
	out := new(ListAuthorsResponse)
	err := c.cc.Invoke(ctx, "/some.app.v1.author.AuthorService/ListAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/some.app.v1.author.AuthorService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
type AuthorServiceServer interface {
	ListAuthors(context.Context, *ListAuthorsRequest) (*ListAuthorsResponse, error)
	GetAuthor(context.Context, *GetAuthorRequest) (*Author, error)
}

// UnimplementedAuthorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (*UnimplementedAuthorServiceServer) ListAuthors(ctx context.Context, req *ListAuthorsRequest) (*ListAuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthors not implemented")
}
func (*UnimplementedAuthorServiceServer) GetAuthor(ctx context.Context, req *GetAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}

func RegisterAuthorServiceServer(s *grpc.Server, srv AuthorServiceServer) {
	s.RegisterService(&_AuthorService_serviceDesc, srv)
}

func _AuthorService_ListAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).ListAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/some.app.v1.author.AuthorService/ListAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).ListAuthors(ctx, req.(*ListAuthorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/some.app.v1.author.AuthorService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "some.app.v1.author.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuthors",
			Handler:    _AuthorService_ListAuthors_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _AuthorService_GetAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author/author.proto",
}
