// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graphql.proto

package graphql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type GraphqlService struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Insecure             bool     `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlService) Reset()         { *m = GraphqlService{} }
func (m *GraphqlService) String() string { return proto.CompactTextString(m) }
func (*GraphqlService) ProtoMessage()    {}
func (*GraphqlService) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{0}
}

func (m *GraphqlService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlService.Unmarshal(m, b)
}
func (m *GraphqlService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlService.Marshal(b, m, deterministic)
}
func (m *GraphqlService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlService.Merge(m, src)
}
func (m *GraphqlService) XXX_Size() int {
	return xxx_messageInfo_GraphqlService.Size(m)
}
func (m *GraphqlService) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlService.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlService proto.InternalMessageInfo

func (m *GraphqlService) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GraphqlService) GetInsecure() bool {
	if m != nil {
		return m.Insecure
	}
	return false
}

type GraphqlRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Expose               string   `protobuf:"bytes,2,opt,name=expose,proto3" json:"expose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlRequest) Reset()         { *m = GraphqlRequest{} }
func (m *GraphqlRequest) String() string { return proto.CompactTextString(m) }
func (*GraphqlRequest) ProtoMessage()    {}
func (*GraphqlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{1}
}

func (m *GraphqlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlRequest.Unmarshal(m, b)
}
func (m *GraphqlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlRequest.Marshal(b, m, deterministic)
}
func (m *GraphqlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlRequest.Merge(m, src)
}
func (m *GraphqlRequest) XXX_Size() int {
	return xxx_messageInfo_GraphqlRequest.Size(m)
}
func (m *GraphqlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlRequest proto.InternalMessageInfo

func (m *GraphqlRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphqlRequest) GetExpose() string {
	if m != nil {
		return m.Expose
	}
	return ""
}

type GraphqlResponse struct {
	Repeated             bool     `protobuf:"varint,1,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Optional             bool     `protobuf:"varint,2,opt,name=optional,proto3" json:"optional,omitempty"`
	Expose               string   `protobuf:"bytes,3,opt,name=expose,proto3" json:"expose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlResponse) Reset()         { *m = GraphqlResponse{} }
func (m *GraphqlResponse) String() string { return proto.CompactTextString(m) }
func (*GraphqlResponse) ProtoMessage()    {}
func (*GraphqlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{2}
}

func (m *GraphqlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlResponse.Unmarshal(m, b)
}
func (m *GraphqlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlResponse.Marshal(b, m, deterministic)
}
func (m *GraphqlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlResponse.Merge(m, src)
}
func (m *GraphqlResponse) XXX_Size() int {
	return xxx_messageInfo_GraphqlResponse.Size(m)
}
func (m *GraphqlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlResponse proto.InternalMessageInfo

func (m *GraphqlResponse) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *GraphqlResponse) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *GraphqlResponse) GetExpose() string {
	if m != nil {
		return m.Expose
	}
	return ""
}

type GraphqlQuery struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Field                string           `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Response             *GraphqlResponse `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GraphqlQuery) Reset()         { *m = GraphqlQuery{} }
func (m *GraphqlQuery) String() string { return proto.CompactTextString(m) }
func (*GraphqlQuery) ProtoMessage()    {}
func (*GraphqlQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{3}
}

func (m *GraphqlQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlQuery.Unmarshal(m, b)
}
func (m *GraphqlQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlQuery.Marshal(b, m, deterministic)
}
func (m *GraphqlQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlQuery.Merge(m, src)
}
func (m *GraphqlQuery) XXX_Size() int {
	return xxx_messageInfo_GraphqlQuery.Size(m)
}
func (m *GraphqlQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlQuery proto.InternalMessageInfo

func (m *GraphqlQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphqlQuery) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *GraphqlQuery) GetResponse() *GraphqlResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type GraphqlMutation struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Field                string           `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Request              *GraphqlRequest  `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Response             *GraphqlResponse `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GraphqlMutation) Reset()         { *m = GraphqlMutation{} }
func (m *GraphqlMutation) String() string { return proto.CompactTextString(m) }
func (*GraphqlMutation) ProtoMessage()    {}
func (*GraphqlMutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{4}
}

func (m *GraphqlMutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlMutation.Unmarshal(m, b)
}
func (m *GraphqlMutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlMutation.Marshal(b, m, deterministic)
}
func (m *GraphqlMutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlMutation.Merge(m, src)
}
func (m *GraphqlMutation) XXX_Size() int {
	return xxx_messageInfo_GraphqlMutation.Size(m)
}
func (m *GraphqlMutation) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlMutation.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlMutation proto.InternalMessageInfo

func (m *GraphqlMutation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphqlMutation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *GraphqlMutation) GetRequest() *GraphqlRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *GraphqlMutation) GetResponse() *GraphqlResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type GraphqlField struct {
	Optional             bool     `protobuf:"varint,1,opt,name=optional,proto3" json:"optional,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlField) Reset()         { *m = GraphqlField{} }
func (m *GraphqlField) String() string { return proto.CompactTextString(m) }
func (*GraphqlField) ProtoMessage()    {}
func (*GraphqlField) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{5}
}

func (m *GraphqlField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlField.Unmarshal(m, b)
}
func (m *GraphqlField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlField.Marshal(b, m, deterministic)
}
func (m *GraphqlField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlField.Merge(m, src)
}
func (m *GraphqlField) XXX_Size() int {
	return xxx_messageInfo_GraphqlField.Size(m)
}
func (m *GraphqlField) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlField.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlField proto.InternalMessageInfo

func (m *GraphqlField) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *GraphqlField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*GraphqlService)(nil),
	Field:         50001,
	Name:          "graphql.service",
	Tag:           "bytes,50001,opt,name=service",
	Filename:      "graphql.proto",
}

var E_Query = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*GraphqlQuery)(nil),
	Field:         50001,
	Name:          "graphql.query",
	Tag:           "bytes,50001,opt,name=query",
	Filename:      "graphql.proto",
}

var E_Mutation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*GraphqlMutation)(nil),
	Field:         50002,
	Name:          "graphql.mutation",
	Tag:           "bytes,50002,opt,name=mutation",
	Filename:      "graphql.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*GraphqlField)(nil),
	Field:         50004,
	Name:          "graphql.field",
	Tag:           "bytes,50004,opt,name=field",
	Filename:      "graphql.proto",
}

func init() {
	proto.RegisterType((*GraphqlService)(nil), "graphql.GraphqlService")
	proto.RegisterType((*GraphqlRequest)(nil), "graphql.GraphqlRequest")
	proto.RegisterType((*GraphqlResponse)(nil), "graphql.GraphqlResponse")
	proto.RegisterType((*GraphqlQuery)(nil), "graphql.GraphqlQuery")
	proto.RegisterType((*GraphqlMutation)(nil), "graphql.GraphqlMutation")
	proto.RegisterType((*GraphqlField)(nil), "graphql.GraphqlField")
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Query)
	proto.RegisterExtension(E_Mutation)
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("graphql.proto", fileDescriptor_3ce0fc368bdc1a51) }

var fileDescriptor_3ce0fc368bdc1a51 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x8e, 0x94, 0x40,
	0x14, 0x86, 0xc3, 0x38, 0x33, 0xcd, 0xd4, 0x78, 0x49, 0x88, 0x17, 0x32, 0x89, 0x4a, 0x58, 0xf5,
	0x66, 0x20, 0xd3, 0xed, 0x8a, 0x18, 0x63, 0x5c, 0xe8, 0xc6, 0x8e, 0xb1, 0x34, 0x2e, 0xdc, 0x55,
	0xc3, 0x69, 0x20, 0x01, 0xaa, 0xba, 0xaa, 0x50, 0xfb, 0x05, 0x7c, 0x12, 0x5f, 0x48, 0xe3, 0x03,
	0x19, 0xea, 0x66, 0x0b, 0x76, 0x8c, 0xb3, 0xe3, 0x50, 0xa7, 0xbe, 0xff, 0xaf, 0xff, 0x1c, 0x74,
	0xab, 0xe4, 0x84, 0x55, 0xdb, 0x26, 0x61, 0x9c, 0x4a, 0x1a, 0xcc, 0x4c, 0x79, 0x11, 0x95, 0x94,
	0x96, 0x0d, 0xa4, 0xea, 0xf7, 0xba, 0xdf, 0xa4, 0x05, 0x88, 0x9c, 0xd7, 0x4c, 0x52, 0xae, 0x5b,
	0xe3, 0xe7, 0xe8, 0xf6, 0x2b, 0xdd, 0xfc, 0x0e, 0xf8, 0xa7, 0x3a, 0x87, 0x20, 0x40, 0xc7, 0x15,
	0x15, 0x32, 0xf4, 0x22, 0x6f, 0x7e, 0x86, 0xd5, 0x77, 0x70, 0x81, 0xfc, 0xba, 0x13, 0x90, 0xf7,
	0x1c, 0xc2, 0xa3, 0xc8, 0x9b, 0xfb, 0xd8, 0xd5, 0xf1, 0x53, 0x47, 0xc0, 0xb0, 0xed, 0x41, 0xc8,
	0x81, 0xd0, 0x91, 0x16, 0x2c, 0x61, 0xf8, 0x0e, 0xee, 0xa3, 0x53, 0xf8, 0xc2, 0xa8, 0xd0, 0xf7,
	0xcf, 0xb0, 0xa9, 0x62, 0x82, 0xee, 0xb8, 0xdb, 0x82, 0xd1, 0x4e, 0xc0, 0x20, 0xc6, 0x81, 0x01,
	0x91, 0x50, 0x28, 0x84, 0x8f, 0x5d, 0x3d, 0x9c, 0x51, 0x26, 0x6b, 0xda, 0x91, 0xc6, 0x1a, 0xb1,
	0xf5, 0x9e, 0xc4, 0x8d, 0x3f, 0x24, 0x3a, 0x74, 0xd3, 0x48, 0xbc, 0xed, 0x81, 0xef, 0xfe, 0x6a,
	0xef, 0x2e, 0x3a, 0xd9, 0xd4, 0xd0, 0x14, 0xc6, 0x9d, 0x2e, 0x82, 0x27, 0x83, 0x13, 0xed, 0x2a,
	0x3c, 0x8e, 0xbc, 0xf9, 0xf9, 0x22, 0x4c, 0x6c, 0xd2, 0x23, 0xd7, 0xd8, 0x75, 0xc6, 0xdf, 0x3c,
	0xf7, 0xa6, 0x55, 0x2f, 0xc9, 0xe0, 0xee, 0x3f, 0x34, 0xaf, 0xd0, 0x8c, 0xeb, 0x1c, 0xd5, 0x33,
	0xce, 0x17, 0x0f, 0xa6, 0x92, 0xea, 0x18, 0xdb, 0xbe, 0x6b, 0xda, 0x7c, 0xe6, 0x62, 0x79, 0xa9,
	0x84, 0xf7, 0xa3, 0xf5, 0x46, 0xd1, 0x5a, 0xfb, 0x47, 0xbf, 0xed, 0x67, 0xef, 0xd1, 0x4c, 0x98,
	0x95, 0x79, 0x9c, 0xe8, 0x3d, 0x4b, 0xec, 0x9e, 0x25, 0x66, 0x99, 0xde, 0x28, 0x80, 0x08, 0xbf,
	0x7f, 0x3d, 0xf0, 0x16, 0xd3, 0x87, 0x2d, 0x2a, 0x5b, 0xa1, 0x93, 0xad, 0x9a, 0xd2, 0xa3, 0x09,
	0x73, 0x05, 0xb2, 0xa2, 0xc5, 0x18, 0x79, 0x6f, 0x8c, 0x54, 0x43, 0xc6, 0x9a, 0x92, 0x7d, 0x40,
	0x7e, 0x6b, 0x67, 0xf0, 0x2f, 0xe2, 0x0f, 0x43, 0x9c, 0x84, 0x67, 0xa7, 0x88, 0x1d, 0x2b, 0x7b,
	0x6d, 0x66, 0x17, 0x3c, 0x9c, 0x40, 0x55, 0x9a, 0x96, 0xf9, 0xf3, 0x90, 0x4b, 0xd5, 0x65, 0x66,
	0xfe, 0x62, 0xf9, 0xf1, 0xaa, 0xac, 0x65, 0xd5, 0xaf, 0x93, 0x9c, 0xb6, 0xe9, 0x4e, 0xf4, 0x65,
	0xdd, 0x52, 0x49, 0xd3, 0x92, 0xb3, 0xfc, 0xd2, 0x5c, 0xbc, 0x2c, 0x89, 0x84, 0xcf, 0x64, 0x97,
	0x9a, 0x7a, 0x7d, 0xaa, 0x14, 0x97, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0xb5, 0xee, 0x3f,
	0xfc, 0x03, 0x00, 0x00,
}