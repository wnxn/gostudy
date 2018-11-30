// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

package time

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRequest) Reset()         { *m = TimeRequest{} }
func (m *TimeRequest) String() string { return proto.CompactTextString(m) }
func (*TimeRequest) ProtoMessage()    {}
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49a92d779a28c7fd, []int{0}
}

func (m *TimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRequest.Unmarshal(m, b)
}
func (m *TimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRequest.Marshal(b, m, deterministic)
}
func (m *TimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRequest.Merge(m, src)
}
func (m *TimeRequest) XXX_Size() int {
	return xxx_messageInfo_TimeRequest.Size(m)
}
func (m *TimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRequest proto.InternalMessageInfo

type TimeResponse struct {
	ServerTime           string   `protobuf:"bytes,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeResponse) Reset()         { *m = TimeResponse{} }
func (m *TimeResponse) String() string { return proto.CompactTextString(m) }
func (*TimeResponse) ProtoMessage()    {}
func (*TimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49a92d779a28c7fd, []int{1}
}

func (m *TimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeResponse.Unmarshal(m, b)
}
func (m *TimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeResponse.Marshal(b, m, deterministic)
}
func (m *TimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeResponse.Merge(m, src)
}
func (m *TimeResponse) XXX_Size() int {
	return xxx_messageInfo_TimeResponse.Size(m)
}
func (m *TimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimeResponse proto.InternalMessageInfo

func (m *TimeResponse) GetServerTime() string {
	if m != nil {
		return m.ServerTime
	}
	return ""
}

func init() {
	proto.RegisterType((*TimeRequest)(nil), "time.TimeRequest")
	proto.RegisterType((*TimeResponse)(nil), "time.TimeResponse")
}

func init() { proto.RegisterFile("time.proto", fileDescriptor_49a92d779a28c7fd) }

var fileDescriptor_49a92d779a28c7fd = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x78, 0xb9, 0xb8, 0x43, 0x32,
	0x73, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xf4, 0xb9, 0x78, 0x20, 0xdc, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x79, 0x2e, 0xee, 0xe2, 0xd4, 0xa2, 0xb2, 0xd4, 0xa2, 0x78,
	0x90, 0x6a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x2e, 0x88, 0x10, 0x48, 0xa1, 0x91, 0x03,
	0x17, 0x17, 0x88, 0x0e, 0x06, 0x8b, 0x08, 0x19, 0x71, 0xb1, 0xbb, 0xa7, 0x96, 0x80, 0x04, 0x84,
	0x04, 0xf5, 0xc0, 0x76, 0x21, 0x19, 0x2e, 0x25, 0x84, 0x2c, 0x04, 0xb1, 0x40, 0x89, 0x21, 0x89,
	0x0d, 0xec, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xf7, 0x39, 0x30, 0x9c, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeServerClient is the client API for TimeServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeServerClient interface {
	GetTime(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeResponse, error)
}

type timeServerClient struct {
	cc *grpc.ClientConn
}

func NewTimeServerClient(cc *grpc.ClientConn) TimeServerClient {
	return &timeServerClient{cc}
}

func (c *timeServerClient) GetTime(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeResponse, error) {
	out := new(TimeResponse)
	err := c.cc.Invoke(ctx, "/time.TimeServer/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServerServer is the server API for TimeServer service.
type TimeServerServer interface {
	GetTime(context.Context, *TimeRequest) (*TimeResponse, error)
}

func RegisterTimeServerServer(s *grpc.Server, srv TimeServerServer) {
	s.RegisterService(&_TimeServer_serviceDesc, srv)
}

func _TimeServer_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServerServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.TimeServer/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServerServer).GetTime(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimeServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "time.TimeServer",
	HandlerType: (*TimeServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _TimeServer_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time.proto",
}
