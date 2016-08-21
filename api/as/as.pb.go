// Code generated by protoc-gen-go.
// source: as.proto
// DO NOT EDIT!

/*
Package as is a generated protocol buffer package.

It is generated from these files:
	as.proto

It has these top-level messages:
	JoinRequestRequest
	JoinRequestResponse
*/
package as

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}
var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}
func (RXWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JoinRequestRequest struct {
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	DevAddr    []byte `protobuf:"bytes,2,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
}

func (m *JoinRequestRequest) Reset()                    { *m = JoinRequestRequest{} }
func (m *JoinRequestRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestRequest) ProtoMessage()               {}
func (*JoinRequestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JoinRequestResponse struct {
	PhyPayload  []byte   `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	NwkSKey     []byte   `protobuf:"bytes,2,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	AppEUI      []byte   `protobuf:"bytes,3,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	DevEUI      []byte   `protobuf:"bytes,4,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	RxDelay     uint32   `protobuf:"varint,5,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset uint32   `protobuf:"varint,6,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList      []uint32 `protobuf:"varint,7,rep,packed,name=cFList" json:"cFList,omitempty"`
	RxWindow    RXWindow `protobuf:"varint,8,opt,name=rxWindow,enum=as.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR       uint32   `protobuf:"varint,9,opt,name=rx2DR" json:"rx2DR,omitempty"`
}

func (m *JoinRequestResponse) Reset()                    { *m = JoinRequestResponse{} }
func (m *JoinRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestResponse) ProtoMessage()               {}
func (*JoinRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*JoinRequestRequest)(nil), "as.JoinRequestRequest")
	proto.RegisterType((*JoinRequestResponse)(nil), "as.JoinRequestResponse")
	proto.RegisterEnum("as.RXWindow", RXWindow_name, RXWindow_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ApplicationServer service

type ApplicationServerClient interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error)
}

type applicationServerClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServerClient(cc *grpc.ClientConn) ApplicationServerClient {
	return &applicationServerClient{cc}
}

func (c *applicationServerClient) JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error) {
	out := new(JoinRequestResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/JoinRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationServer service

type ApplicationServerServer interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(context.Context, *JoinRequestRequest) (*JoinRequestResponse, error)
}

func RegisterApplicationServerServer(s *grpc.Server, srv ApplicationServerServer) {
	s.RegisterService(&_ApplicationServer_serviceDesc, srv)
}

func _ApplicationServer_JoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).JoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/JoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).JoinRequest(ctx, req.(*JoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "as.ApplicationServer",
	HandlerType: (*ApplicationServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinRequest",
			Handler:    _ApplicationServer_JoinRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("as.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x4a, 0xf3, 0x30,
	0x18, 0xc6, 0xd7, 0xed, 0xdb, 0xd6, 0xef, 0xdd, 0x26, 0x33, 0xca, 0x0c, 0x22, 0x52, 0x7a, 0x54,
	0x3c, 0x28, 0xac, 0xde, 0x80, 0x83, 0x2a, 0xf8, 0x07, 0x95, 0x8c, 0xe1, 0x4e, 0xe3, 0x92, 0x61,
	0xb1, 0x34, 0x31, 0xa9, 0x5d, 0x7b, 0x1d, 0xde, 0xb0, 0xb4, 0x69, 0xa5, 0xb2, 0x13, 0x8f, 0x92,
	0xdf, 0xef, 0x09, 0x4f, 0xe0, 0x7d, 0xc1, 0xa6, 0xda, 0x97, 0x4a, 0xa4, 0x02, 0x75, 0xa9, 0x76,
	0x1f, 0x01, 0xdd, 0x89, 0x28, 0x21, 0xfc, 0xe3, 0x93, 0xeb, 0xb4, 0x3e, 0xd0, 0x39, 0x80, 0x7c,
	0x2b, 0x9e, 0x69, 0x11, 0x0b, 0xca, 0xb0, 0xe5, 0x58, 0xde, 0x98, 0xb4, 0x0c, 0xc2, 0x30, 0x64,
	0x3c, 0x5b, 0x30, 0xa6, 0x70, 0xb7, 0x0a, 0x1b, 0x74, 0xbf, 0xba, 0x70, 0xf4, 0xab, 0x50, 0x4b,
	0x91, 0x68, 0xfe, 0x97, 0xc6, 0x64, 0xf7, 0xbe, 0xbc, 0xe7, 0x45, 0xd3, 0x58, 0x23, 0x9a, 0xc1,
	0x80, 0x4a, 0x79, 0xbd, 0xba, 0xc5, 0xbd, 0x2a, 0xa8, 0xa9, 0xf4, 0x8c, 0x67, 0xa5, 0xff, 0x67,
	0xbc, 0xa1, 0xb2, 0x49, 0xe5, 0x21, 0x8f, 0x69, 0x81, 0xfb, 0x8e, 0xe5, 0x4d, 0x48, 0x83, 0xc8,
	0x81, 0x91, 0xca, 0xe7, 0x21, 0x79, 0xda, 0x6e, 0x35, 0x4f, 0xf1, 0xa0, 0x4a, 0xdb, 0xaa, 0xec,
	0xdc, 0xdc, 0x3c, 0x44, 0x3a, 0xc5, 0x43, 0xa7, 0xe7, 0x4d, 0x48, 0x4d, 0xc8, 0x03, 0x5b, 0xe5,
	0x2f, 0x51, 0xc2, 0xc4, 0x0e, 0xdb, 0x8e, 0xe5, 0x1d, 0x04, 0x63, 0x9f, 0x6a, 0x9f, 0xac, 0x8d,
	0x23, 0x3f, 0x29, 0x3a, 0x86, 0xbe, 0xca, 0x83, 0x90, 0xe0, 0xff, 0x55, 0xbb, 0x81, 0x8b, 0x33,
	0xb0, 0x9b, 0xb7, 0x68, 0x08, 0x3d, 0xb2, 0x9e, 0x4f, 0x3b, 0xe6, 0x12, 0x4c, 0xad, 0x60, 0x05,
	0x87, 0x0b, 0x29, 0xe3, 0x68, 0x43, 0xd3, 0x48, 0x24, 0x4b, 0xae, 0x32, 0xae, 0xd0, 0x15, 0x8c,
	0x5a, 0x73, 0x44, 0xb3, 0xf2, 0xbf, 0xfd, 0x4d, 0x9d, 0x9e, 0xec, 0x79, 0x33, 0x70, 0xb7, 0xf3,
	0x3a, 0xa8, 0xb6, 0x7c, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x69, 0xff, 0xab, 0xed, 0xf1, 0x01,
	0x00, 0x00,
}
