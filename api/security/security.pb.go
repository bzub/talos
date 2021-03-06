// Code generated by protoc-gen-go. DO NOT EDIT.
// source: security/security.proto

package security

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the process name.
type CertificateRequest struct {
	Csr                  []byte   `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fd4b7e16002c2e, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}

func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}

func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}

func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}

func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// The response message containing the requested logs.
type CertificateResponse struct {
	Ca                   []byte   `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	Crt                  []byte   `protobuf:"bytes,2,opt,name=crt,proto3" json:"crt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fd4b7e16002c2e, []int{1}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}

func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}

func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}

func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}

func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCa() []byte {
	if m != nil {
		return m.Ca
	}
	return nil
}

func (m *CertificateResponse) GetCrt() []byte {
	if m != nil {
		return m.Crt
	}
	return nil
}

// The request message for reading a file on disk.
type ReadFileRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadFileRequest) Reset()         { *m = ReadFileRequest{} }
func (m *ReadFileRequest) String() string { return proto.CompactTextString(m) }
func (*ReadFileRequest) ProtoMessage()    {}
func (*ReadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fd4b7e16002c2e, []int{2}
}

func (m *ReadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadFileRequest.Unmarshal(m, b)
}

func (m *ReadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadFileRequest.Marshal(b, m, deterministic)
}

func (m *ReadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadFileRequest.Merge(m, src)
}

func (m *ReadFileRequest) XXX_Size() int {
	return xxx_messageInfo_ReadFileRequest.Size(m)
}

func (m *ReadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadFileRequest proto.InternalMessageInfo

func (m *ReadFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// The response message for reading a file on disk.
type ReadFileResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadFileResponse) Reset()         { *m = ReadFileResponse{} }
func (m *ReadFileResponse) String() string { return proto.CompactTextString(m) }
func (*ReadFileResponse) ProtoMessage()    {}
func (*ReadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fd4b7e16002c2e, []int{3}
}

func (m *ReadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadFileResponse.Unmarshal(m, b)
}

func (m *ReadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadFileResponse.Marshal(b, m, deterministic)
}

func (m *ReadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadFileResponse.Merge(m, src)
}

func (m *ReadFileResponse) XXX_Size() int {
	return xxx_messageInfo_ReadFileResponse.Size(m)
}

func (m *ReadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadFileResponse proto.InternalMessageInfo

func (m *ReadFileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request message containing the process name.
type WriteFileRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Perm                 int32    `protobuf:"varint,3,opt,name=perm,proto3" json:"perm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteFileRequest) Reset()         { *m = WriteFileRequest{} }
func (m *WriteFileRequest) String() string { return proto.CompactTextString(m) }
func (*WriteFileRequest) ProtoMessage()    {}
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fd4b7e16002c2e, []int{4}
}

func (m *WriteFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteFileRequest.Unmarshal(m, b)
}

func (m *WriteFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteFileRequest.Marshal(b, m, deterministic)
}

func (m *WriteFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteFileRequest.Merge(m, src)
}

func (m *WriteFileRequest) XXX_Size() int {
	return xxx_messageInfo_WriteFileRequest.Size(m)
}

func (m *WriteFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteFileRequest proto.InternalMessageInfo

func (m *WriteFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WriteFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *WriteFileRequest) GetPerm() int32 {
	if m != nil {
		return m.Perm
	}
	return 0
}

// The response message containing the requested logs.
type WriteFileResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteFileResponse) Reset()         { *m = WriteFileResponse{} }
func (m *WriteFileResponse) String() string { return proto.CompactTextString(m) }
func (*WriteFileResponse) ProtoMessage()    {}
func (*WriteFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45fd4b7e16002c2e, []int{5}
}

func (m *WriteFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteFileResponse.Unmarshal(m, b)
}

func (m *WriteFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteFileResponse.Marshal(b, m, deterministic)
}

func (m *WriteFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteFileResponse.Merge(m, src)
}

func (m *WriteFileResponse) XXX_Size() int {
	return xxx_messageInfo_WriteFileResponse.Size(m)
}

func (m *WriteFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteFileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "securityapi.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "securityapi.CertificateResponse")
	proto.RegisterType((*ReadFileRequest)(nil), "securityapi.ReadFileRequest")
	proto.RegisterType((*ReadFileResponse)(nil), "securityapi.ReadFileResponse")
	proto.RegisterType((*WriteFileRequest)(nil), "securityapi.WriteFileRequest")
	proto.RegisterType((*WriteFileResponse)(nil), "securityapi.WriteFileResponse")
}

func init() { proto.RegisterFile("security/security.proto", fileDescriptor_45fd4b7e16002c2e) }

var fileDescriptor_45fd4b7e16002c2e = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x6d, 0x15, 0x3b, 0x15, 0xbb, 0xa6, 0x07, 0x4b, 0xf1, 0xa3, 0x2c, 0x58, 0x0a,
	0xe2, 0x16, 0xf4, 0xe0, 0xd9, 0x0a, 0x0a, 0x1e, 0x4a, 0xd9, 0x1e, 0x04, 0x6f, 0x69, 0x3a, 0xda,
	0x40, 0xd7, 0xc4, 0x64, 0x2a, 0xf4, 0x3f, 0xf7, 0x28, 0xa6, 0xd9, 0xed, 0x87, 0x2c, 0xde, 0x5e,
	0x98, 0xdf, 0x7b, 0x61, 0x1e, 0x03, 0x27, 0x16, 0xc5, 0xc2, 0x48, 0x5a, 0xf6, 0x73, 0x91, 0x68,
	0xa3, 0x48, 0xb1, 0x7a, 0xfe, 0xe6, 0x5a, 0xc6, 0x5d, 0x60, 0x0f, 0x68, 0x48, 0xbe, 0x49, 0xc1,
	0x09, 0x53, 0xfc, 0x5c, 0xa0, 0x25, 0x16, 0x41, 0x45, 0x58, 0xd3, 0x0a, 0x3a, 0x41, 0xef, 0x30,
	0xfd, 0x95, 0xf1, 0x1d, 0x34, 0xb7, 0x38, 0xab, 0xd5, 0x87, 0x45, 0x76, 0x04, 0xa1, 0xe0, 0x9e,
	0x0b, 0x05, 0x77, 0x46, 0x43, 0xad, 0xd0, 0x1b, 0x0d, 0xc5, 0x97, 0xd0, 0x48, 0x91, 0x4f, 0x1f,
	0xe5, 0xbc, 0x48, 0x67, 0x50, 0xd5, 0x9c, 0x66, 0xce, 0x56, 0x4b, 0x9d, 0x8e, 0xbb, 0x10, 0xad,
	0x31, 0x1f, 0xce, 0xa0, 0x3a, 0xe5, 0x94, 0xc7, 0x3b, 0x1d, 0x0f, 0x21, 0x7a, 0x31, 0x92, 0xf0,
	0x9f, 0xbc, 0xc2, 0x1b, 0xae, 0xbd, 0x8e, 0x43, 0x93, 0xb5, 0x2a, 0x9d, 0xa0, 0xb7, 0x97, 0x3a,
	0x1d, 0x37, 0xe1, 0x78, 0x23, 0x6f, 0xf5, 0xf1, 0xcd, 0x77, 0x00, 0x8d, 0xb1, 0x2f, 0x69, 0x8c,
	0xe6, 0x4b, 0x0a, 0x64, 0x23, 0xa8, 0x6f, 0x14, 0xc0, 0x2e, 0x92, 0x8d, 0x16, 0x93, 0xbf, 0x15,
	0xb6, 0x3b, 0xe5, 0x80, 0x5f, 0xef, 0x09, 0x0e, 0xf2, 0x95, 0xd9, 0xe9, 0x16, 0xbd, 0x53, 0x58,
	0xfb, 0xac, 0x64, 0xea, 0x83, 0x9e, 0xa1, 0x56, 0xec, 0xc0, 0xb6, 0xd9, 0xdd, 0xae, 0xda, 0xe7,
	0x65, 0xe3, 0x55, 0xd6, 0x60, 0x08, 0x91, 0x50, 0x59, 0x01, 0x25, 0x5c, 0xcb, 0x41, 0x3d, 0xef,
	0xe2, 0x5e, 0xcb, 0x51, 0xf0, 0x7a, 0xf5, 0x2e, 0x69, 0xb6, 0x98, 0x24, 0x42, 0x65, 0x7d, 0xe2,
	0x73, 0x65, 0xaf, 0xed, 0xd2, 0x12, 0x66, 0x76, 0xf5, 0xea, 0x73, 0x2d, 0x8b, 0x93, 0x9b, 0xec,
	0xbb, 0x9b, 0xbb, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x87, 0x47, 0xc7, 0x8e, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecurityServiceClient interface {
	Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	ReadFile(ctx context.Context, in *ReadFileRequest, opts ...grpc.CallOption) (*ReadFileResponse, error)
	WriteFile(ctx context.Context, in *WriteFileRequest, opts ...grpc.CallOption) (*WriteFileResponse, error)
}

type securityServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecurityServiceClient(cc *grpc.ClientConn) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/securityapi.SecurityService/Certificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) ReadFile(ctx context.Context, in *ReadFileRequest, opts ...grpc.CallOption) (*ReadFileResponse, error) {
	out := new(ReadFileResponse)
	err := c.cc.Invoke(ctx, "/securityapi.SecurityService/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) WriteFile(ctx context.Context, in *WriteFileRequest, opts ...grpc.CallOption) (*WriteFileResponse, error) {
	out := new(WriteFileResponse)
	err := c.cc.Invoke(ctx, "/securityapi.SecurityService/WriteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
type SecurityServiceServer interface {
	Certificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	ReadFile(context.Context, *ReadFileRequest) (*ReadFileResponse, error)
	WriteFile(context.Context, *WriteFileRequest) (*WriteFileResponse, error)
}

func RegisterSecurityServiceServer(s *grpc.Server, srv SecurityServiceServer) {
	s.RegisterService(&_SecurityService_serviceDesc, srv)
}

func _SecurityService_Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/securityapi.SecurityService/Certificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Certificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/securityapi.SecurityService/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).ReadFile(ctx, req.(*ReadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_WriteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).WriteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/securityapi.SecurityService/WriteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).WriteFile(ctx, req.(*WriteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "securityapi.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificate",
			Handler:    _SecurityService_Certificate_Handler,
		},
		{
			MethodName: "ReadFile",
			Handler:    _SecurityService_ReadFile_Handler,
		},
		{
			MethodName: "WriteFile",
			Handler:    _SecurityService_WriteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/security.proto",
}
