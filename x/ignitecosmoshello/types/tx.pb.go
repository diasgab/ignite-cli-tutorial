// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ignitecosmoshello/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("ignitecosmoshello/tx.proto", fileDescriptor_c95ac3e116495519) }

var fileDescriptor_c95ac3e116495519 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4c, 0xcf, 0xcb,
	0x2c, 0x49, 0x4d, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xce, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xc6, 0x90, 0xd3, 0xc3, 0x10, 0x31, 0x62, 0xe5,
	0x62, 0xf6, 0x2d, 0x4e, 0x77, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0x5d, 0x88, 0x1e, 0x5d, 0x88, 0x26, 0x5d, 0x88, 0x1d, 0x15, 0xfa, 0x58, 0xec, 0xad, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0xdb, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x5c, 0x38,
	0x0f, 0x99, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ignitecosmoshello.ignitecosmoshello.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "ignitecosmoshello/tx.proto",
}
