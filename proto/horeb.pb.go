// Code generated by protoc-gen-go.
// source: horeb.proto
// DO NOT EDIT!

/*
Package horeb is a generated protocol buffer package.

It is generated from these files:
	horeb.proto

It has these top-level messages:
	RuneRequest
	Rune
*/
package horeb

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

type RuneRequest struct {
	Num   int32  `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Block string `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *RuneRequest) Reset()                    { *m = RuneRequest{} }
func (m *RuneRequest) String() string            { return proto.CompactTextString(m) }
func (*RuneRequest) ProtoMessage()               {}
func (*RuneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RuneRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *RuneRequest) GetBlock() string {
	if m != nil {
		return m.Block
	}
	return ""
}

type Rune struct {
	R string `protobuf:"bytes,1,opt,name=r" json:"r,omitempty"`
}

func (m *Rune) Reset()                    { *m = Rune{} }
func (m *Rune) String() string            { return proto.CompactTextString(m) }
func (*Rune) ProtoMessage()               {}
func (*Rune) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rune) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func init() {
	proto.RegisterType((*RuneRequest)(nil), "horeb.RuneRequest")
	proto.RegisterType((*Rune)(nil), "horeb.Rune")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Horeb service

type HorebClient interface {
	GetStream(ctx context.Context, in *RuneRequest, opts ...grpc.CallOption) (Horeb_GetStreamClient, error)
}

type horebClient struct {
	cc *grpc.ClientConn
}

func NewHorebClient(cc *grpc.ClientConn) HorebClient {
	return &horebClient{cc}
}

func (c *horebClient) GetStream(ctx context.Context, in *RuneRequest, opts ...grpc.CallOption) (Horeb_GetStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Horeb_serviceDesc.Streams[0], c.cc, "/horeb.Horeb/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &horebGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Horeb_GetStreamClient interface {
	Recv() (*Rune, error)
	grpc.ClientStream
}

type horebGetStreamClient struct {
	grpc.ClientStream
}

func (x *horebGetStreamClient) Recv() (*Rune, error) {
	m := new(Rune)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Horeb service

type HorebServer interface {
	GetStream(*RuneRequest, Horeb_GetStreamServer) error
}

func RegisterHorebServer(s *grpc.Server, srv HorebServer) {
	s.RegisterService(&_Horeb_serviceDesc, srv)
}

func _Horeb_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RuneRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HorebServer).GetStream(m, &horebGetStreamServer{stream})
}

type Horeb_GetStreamServer interface {
	Send(*Rune) error
	grpc.ServerStream
}

type horebGetStreamServer struct {
	grpc.ServerStream
}

func (x *horebGetStreamServer) Send(m *Rune) error {
	return x.ServerStream.SendMsg(m)
}

var _Horeb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "horeb.Horeb",
	HandlerType: (*HorebServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Horeb_GetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "horeb.proto",
}

func init() { proto.RegisterFile("horeb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc8, 0x2f, 0x4a,
	0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x4c, 0xb9, 0xb8, 0x83,
	0x4a, 0xf3, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0x4a,
	0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xa4, 0x9c,
	0xfc, 0xe4, 0x6c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x49, 0x84, 0x8b, 0x05,
	0xa4, 0x4d, 0x88, 0x87, 0x8b, 0xb1, 0x08, 0xac, 0x9a, 0x33, 0x88, 0xb1, 0xc8, 0xc8, 0x92, 0x8b,
	0xd5, 0x03, 0x64, 0xaa, 0x90, 0x01, 0x17, 0xa7, 0x7b, 0x6a, 0x49, 0x70, 0x49, 0x51, 0x6a, 0x62,
	0xae, 0x90, 0x90, 0x1e, 0xc4, 0x5e, 0x24, 0x7b, 0xa4, 0xb8, 0x91, 0xc4, 0x94, 0x18, 0x0c, 0x18,
	0x93, 0xd8, 0xc0, 0xae, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xe6, 0xf5, 0x16, 0xa4,
	0x00, 0x00, 0x00,
}
