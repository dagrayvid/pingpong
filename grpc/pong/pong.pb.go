// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pong.proto

/*
Package pong is a generated protocol buffer package.

It is generated from these files:
	pong.proto

It has these top-level messages:
	PongData
*/
package pong

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

type PongData struct {
	Msg  string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Ball int32  `protobuf:"varint,2,opt,name=ball" json:"ball,omitempty"`
}

func (m *PongData) Reset()                    { *m = PongData{} }
func (m *PongData) String() string            { return proto.CompactTextString(m) }
func (*PongData) ProtoMessage()               {}
func (*PongData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PongData) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *PongData) GetBall() int32 {
	if m != nil {
		return m.Ball
	}
	return 0
}

func init() {
	proto.RegisterType((*PongData)(nil), "pong.PongData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PongService service

type PongServiceClient interface {
	PingPongRPC(ctx context.Context, opts ...grpc.CallOption) (PongService_PingPongRPCClient, error)
}

type pongServiceClient struct {
	cc *grpc.ClientConn
}

func NewPongServiceClient(cc *grpc.ClientConn) PongServiceClient {
	return &pongServiceClient{cc}
}

func (c *pongServiceClient) PingPongRPC(ctx context.Context, opts ...grpc.CallOption) (PongService_PingPongRPCClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PongService_serviceDesc.Streams[0], c.cc, "/pong.PongService/PingPongRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &pongServicePingPongRPCClient{stream}
	return x, nil
}

type PongService_PingPongRPCClient interface {
	Send(*PongData) error
	Recv() (*PongData, error)
	grpc.ClientStream
}

type pongServicePingPongRPCClient struct {
	grpc.ClientStream
}

func (x *pongServicePingPongRPCClient) Send(m *PongData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pongServicePingPongRPCClient) Recv() (*PongData, error) {
	m := new(PongData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PongService service

type PongServiceServer interface {
	PingPongRPC(PongService_PingPongRPCServer) error
}

func RegisterPongServiceServer(s *grpc.Server, srv PongServiceServer) {
	s.RegisterService(&_PongService_serviceDesc, srv)
}

func _PongService_PingPongRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PongServiceServer).PingPongRPC(&pongServicePingPongRPCServer{stream})
}

type PongService_PingPongRPCServer interface {
	Send(*PongData) error
	Recv() (*PongData, error)
	grpc.ServerStream
}

type pongServicePingPongRPCServer struct {
	grpc.ServerStream
}

func (x *pongServicePingPongRPCServer) Send(m *PongData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pongServicePingPongRPCServer) Recv() (*PongData, error) {
	m := new(PongData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PongService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pong.PongService",
	HandlerType: (*PongServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingPongRPC",
			Handler:       _PongService_PingPongRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pong.proto",
}

func init() { proto.RegisterFile("pong.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcf, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x0c, 0xb8, 0x38, 0x02, 0xf2,
	0xf3, 0xd2, 0x5d, 0x12, 0x4b, 0x12, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xc4, 0x9c, 0x1c, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0xdb, 0xc8, 0x89, 0x8b, 0x1b, 0xa4, 0x23, 0x38, 0xb5, 0xa8,
	0x2c, 0x33, 0x39, 0x55, 0xc8, 0x98, 0x8b, 0x3b, 0x20, 0x33, 0x2f, 0x1d, 0x24, 0x14, 0x14, 0xe0,
	0x2c, 0xc4, 0xa7, 0x07, 0xb6, 0x02, 0x66, 0xa6, 0x14, 0x1a, 0x5f, 0x89, 0x41, 0x83, 0xd1, 0x80,
	0x31, 0x89, 0x0d, 0xec, 0x04, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x0d, 0xd2, 0x55,
	0x90, 0x00, 0x00, 0x00,
}
