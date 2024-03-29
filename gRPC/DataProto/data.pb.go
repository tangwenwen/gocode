// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package DataProto

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

type Data struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "DataProto.Data")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x74, 0x49, 0x2c, 0x49, 0x0c, 0x00, 0x31, 0x95,
	0xa4, 0xb8, 0x58, 0x40, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xd4, 0x8a, 0x12, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0xc8, 0x8a, 0x8b, 0xc3, 0x2d, 0xbf, 0x28, 0x17, 0x2c, 0xaf,
	0xc7, 0xc5, 0xe1, 0x92, 0x0f, 0xe2, 0x25, 0x96, 0x08, 0xf1, 0xeb, 0xc1, 0xf5, 0x83, 0x59, 0x52,
	0xe8, 0x02, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x9b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86,
	0x19, 0xf0, 0xb3, 0x77, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FormDataClient is the client API for FormData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FormDataClient interface {
	DoFormat(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error)
}

type formDataClient struct {
	cc *grpc.ClientConn
}

func NewFormDataClient(cc *grpc.ClientConn) FormDataClient {
	return &formDataClient{cc}
}

func (c *formDataClient) DoFormat(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/DataProto.FormData/DoFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormDataServer is the server API for FormData service.
type FormDataServer interface {
	DoFormat(context.Context, *Data) (*Data, error)
}

// UnimplementedFormDataServer can be embedded to have forward compatible implementations.
type UnimplementedFormDataServer struct {
}

func (*UnimplementedFormDataServer) DoFormat(ctx context.Context, req *Data) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoFormat not implemented")
}

func RegisterFormDataServer(s *grpc.Server, srv FormDataServer) {
	s.RegisterService(&_FormData_serviceDesc, srv)
}

func _FormData_DoFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormDataServer).DoFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataProto.FormData/DoFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormDataServer).DoFormat(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _FormData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DataProto.FormData",
	HandlerType: (*FormDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoFormat",
			Handler:    _FormData_DoFormat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
