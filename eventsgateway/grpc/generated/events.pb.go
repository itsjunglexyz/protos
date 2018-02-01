// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

/*
Package eventforwarder is a generated protocol buffer package.

It is generated from these files:
	events.proto

It has these top-level messages:
	Event
	Response
*/
package eventforwarder

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

type Event struct {
	Id    string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Topic string            `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	Props map[string]string `protobuf:"bytes,4,rep,name=props" json:"props,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Event) GetProps() map[string]string {
	if m != nil {
		return m.Props
	}
	return nil
}

type Response struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "eventforwarder.Event")
	proto.RegisterType((*Response)(nil), "eventforwarder.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GRPCForwarder service

type GRPCForwarderClient interface {
	SendEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
}

type gRPCForwarderClient struct {
	cc *grpc.ClientConn
}

func NewGRPCForwarderClient(cc *grpc.ClientConn) GRPCForwarderClient {
	return &gRPCForwarderClient{cc}
}

func (c *gRPCForwarderClient) SendEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GRPCForwarder service

type GRPCForwarderServer interface {
	SendEvent(context.Context, *Event) (*Response, error)
}

func RegisterGRPCForwarderServer(s *grpc.Server, srv GRPCForwarderServer) {
	s.RegisterService(&_GRPCForwarder_serviceDesc, srv)
}

func _GRPCForwarder_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCForwarder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventforwarder.GRPCForwarder",
	HandlerType: (*GRPCForwarderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _GRPCForwarder_SendEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}

func init() { proto.RegisterFile("events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0xd2, 0xa8, 0x1d, 0xb5, 0xc8, 0xa0, 0xb0, 0xf4, 0x14, 0x72, 0xea, 0x29, 0x87,
	0x0a, 0x12, 0x3c, 0x78, 0x91, 0xea, 0xd1, 0xb2, 0xfe, 0x82, 0xd8, 0x1d, 0x25, 0x68, 0x77, 0x97,
	0xdd, 0x58, 0xc9, 0xef, 0xf2, 0x0f, 0xca, 0xce, 0xba, 0x88, 0xd2, 0xdb, 0x7b, 0x8f, 0x99, 0x79,
	0x1f, 0x03, 0xa7, 0xb4, 0x23, 0x3d, 0xf8, 0xc6, 0x3a, 0x33, 0x18, 0x9c, 0xb1, 0x7b, 0x31, 0xee,
	0xb3, 0x73, 0x8a, 0x5c, 0xfd, 0x95, 0x41, 0xb9, 0x0a, 0x11, 0xce, 0x20, 0xef, 0x95, 0xc8, 0xaa,
	0x6c, 0x31, 0x95, 0x79, 0xaf, 0x10, 0x61, 0xa2, 0xbb, 0x2d, 0x89, 0x9c, 0x13, 0xd6, 0x78, 0x01,
	0xe5, 0x60, 0x6c, 0xbf, 0x11, 0x05, 0x87, 0xd1, 0xe0, 0x35, 0x94, 0xd6, 0x19, 0xeb, 0xc5, 0xa4,
	0x2a, 0x16, 0x27, 0xcb, 0xaa, 0xf9, 0xdb, 0xd1, 0xf0, 0xfd, 0x66, 0x1d, 0x46, 0x56, 0x7a, 0x70,
	0xa3, 0x8c, 0xe3, 0xf3, 0x16, 0xe0, 0x37, 0xc4, 0x73, 0x28, 0xde, 0x68, 0xfc, 0x01, 0x08, 0x32,
	0xb4, 0xed, 0xba, 0xf7, 0x8f, 0x84, 0x10, 0xcd, 0x4d, 0xde, 0x66, 0x75, 0x0b, 0xc7, 0x92, 0xbc,
	0x35, 0xda, 0x53, 0xe0, 0xdc, 0x18, 0x45, 0xbc, 0x58, 0x4a, 0xd6, 0x28, 0xe0, 0x68, 0x4b, 0xde,
	0x77, 0xaf, 0x69, 0x37, 0xd9, 0xe5, 0x23, 0x9c, 0x3d, 0xc8, 0xf5, 0xdd, 0x7d, 0x82, 0xc3, 0x5b,
	0x98, 0x3e, 0x91, 0x56, 0xf1, 0x07, 0x97, 0x7b, 0xd1, 0xe7, 0xe2, 0x7f, 0x9c, 0xca, 0xeb, 0x83,
	0xe7, 0x43, 0xfe, 0xeb, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xd2, 0xc0, 0xd5, 0x67,
	0x01, 0x00, 0x00,
}
