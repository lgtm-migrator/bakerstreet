// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bakerstreet.proto

/*
Package com_appknox_bakerstreet is a generated protocol buffer package.

It is generated from these files:
	bakerstreet.proto

It has these top-level messages:
	Message
	Package
	Device
*/
package com_appknox_bakerstreet

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

type Message struct {
	Title       string `protobuf:"bytes,1,opt,name=Title,json=title" json:"Title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,json=description" json:"Description,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Message) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Package struct {
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Device struct {
	Url             string `protobuf:"bytes,1,opt,name=Url,json=url" json:"Url,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	Uuid            string `protobuf:"bytes,3,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`
	IsTablet        bool   `protobuf:"varint,4,opt,name=IsTablet,json=isTablet" json:"IsTablet,omitempty"`
	Platform        int32  `protobuf:"varint,5,opt,name=Platform,json=platform" json:"Platform,omitempty"`
	NotifyUrl       string `protobuf:"bytes,6,opt,name=NotifyUrl,json=notifyUrl" json:"NotifyUrl,omitempty"`
	SerialNumber    string `protobuf:"bytes,8,opt,name=SerialNumber,json=serialNumber" json:"SerialNumber,omitempty"`
	IsProxyEnabled  bool   `protobuf:"varint,7,opt,name=IsProxyEnabled,json=isProxyEnabled" json:"IsProxyEnabled,omitempty"`
	PlatformVersion string `protobuf:"bytes,9,opt,name=PlatformVersion,json=platformVersion" json:"PlatformVersion,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Device) GetIsTablet() bool {
	if m != nil {
		return m.IsTablet
	}
	return false
}

func (m *Device) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *Device) GetNotifyUrl() string {
	if m != nil {
		return m.NotifyUrl
	}
	return ""
}

func (m *Device) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Device) GetIsProxyEnabled() bool {
	if m != nil {
		return m.IsProxyEnabled
	}
	return false
}

func (m *Device) GetPlatformVersion() string {
	if m != nil {
		return m.PlatformVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "com.appknox.bakerstreet.Message")
	proto.RegisterType((*Package)(nil), "com.appknox.bakerstreet.Package")
	proto.RegisterType((*Device)(nil), "com.appknox.bakerstreet.Device")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agent service

type AgentClient interface {
	Info(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Device, error)
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	LaunchApp(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	HealthCheck(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	RemovePackage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	InstallPackage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ListPackages(ctx context.Context, in *Message, opts ...grpc.CallOption) (Agent_ListPackagesClient, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Info(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Agent/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Agent/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) LaunchApp(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Agent/LaunchApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) HealthCheck(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Agent/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RemovePackage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Agent/RemovePackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) InstallPackage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Agent/InstallPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ListPackages(ctx context.Context, in *Message, opts ...grpc.CallOption) (Agent_ListPackagesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Agent_serviceDesc.Streams[0], c.cc, "/com.appknox.bakerstreet.Agent/ListPackages", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentListPackagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agent_ListPackagesClient interface {
	Recv() (*Package, error)
	grpc.ClientStream
}

type agentListPackagesClient struct {
	grpc.ClientStream
}

func (x *agentListPackagesClient) Recv() (*Package, error) {
	m := new(Package)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Agent service

type AgentServer interface {
	Info(context.Context, *Message) (*Device, error)
	Echo(context.Context, *Message) (*Message, error)
	LaunchApp(context.Context, *Message) (*Message, error)
	HealthCheck(context.Context, *Message) (*Message, error)
	RemovePackage(context.Context, *Message) (*Message, error)
	InstallPackage(context.Context, *Message) (*Message, error)
	ListPackages(*Message, Agent_ListPackagesServer) error
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Agent/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Info(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Agent/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_LaunchApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).LaunchApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Agent/LaunchApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).LaunchApp(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Agent/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).HealthCheck(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RemovePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RemovePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Agent/RemovePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RemovePackage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_InstallPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).InstallPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Agent/InstallPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).InstallPackage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ListPackages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServer).ListPackages(m, &agentListPackagesServer{stream})
}

type Agent_ListPackagesServer interface {
	Send(*Package) error
	grpc.ServerStream
}

type agentListPackagesServer struct {
	grpc.ServerStream
}

func (x *agentListPackagesServer) Send(m *Package) error {
	return x.ServerStream.SendMsg(m)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.appknox.bakerstreet.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Agent_Info_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _Agent_Echo_Handler,
		},
		{
			MethodName: "LaunchApp",
			Handler:    _Agent_LaunchApp_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _Agent_HealthCheck_Handler,
		},
		{
			MethodName: "RemovePackage",
			Handler:    _Agent_RemovePackage_Handler,
		},
		{
			MethodName: "InstallPackage",
			Handler:    _Agent_InstallPackage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPackages",
			Handler:       _Agent_ListPackages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bakerstreet.proto",
}

func init() { proto.RegisterFile("bakerstreet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x36, 0x69, 0x93, 0xd3, 0xd2, 0x81, 0x85, 0x44, 0x34, 0x81, 0x88, 0x7a, 0x81,
	0x7a, 0x15, 0x21, 0x78, 0x82, 0x8a, 0x4d, 0x22, 0x68, 0x94, 0xd2, 0x75, 0xdc, 0x3b, 0xc9, 0xe9,
	0x62, 0xc5, 0xb1, 0x2d, 0xdb, 0x99, 0xb6, 0x97, 0xe0, 0xa5, 0x78, 0x31, 0x94, 0x90, 0xb0, 0xed,
	0xa2, 0x02, 0x69, 0xb9, 0x3c, 0xdf, 0x7f, 0xfc, 0x9d, 0xc8, 0x27, 0x86, 0x17, 0x29, 0x2d, 0x51,
	0x1b, 0xab, 0x11, 0x6d, 0xac, 0xb4, 0xb4, 0x92, 0xbc, 0xca, 0x64, 0x15, 0x53, 0xa5, 0x4a, 0x21,
	0x6f, 0xe3, 0x07, 0xf1, 0x72, 0x0d, 0xd3, 0xaf, 0x68, 0x0c, 0xbd, 0x46, 0xf2, 0x12, 0xbc, 0x3d,
	0xb3, 0x1c, 0x43, 0x27, 0x72, 0x56, 0xc1, 0xce, 0xb3, 0x4d, 0x41, 0x22, 0x98, 0x9d, 0xa1, 0xc9,
	0x34, 0x53, 0x96, 0x49, 0x11, 0x8e, 0xda, 0x6c, 0x96, 0xdf, 0xa3, 0xe5, 0x1b, 0x98, 0x6e, 0x69,
	0x56, 0x36, 0x0a, 0x02, 0xee, 0x86, 0x56, 0xbd, 0xc1, 0x15, 0xb4, 0xc2, 0xe5, 0xcf, 0x11, 0x4c,
	0xce, 0xf0, 0x86, 0x65, 0x48, 0x9e, 0xc3, 0xf8, 0x4a, 0xf3, 0x2e, 0x1d, 0xd7, 0x9a, 0xff, 0x3d,
	0x30, 0xba, 0x3f, 0xd0, 0xb0, 0xab, 0x9a, 0xe5, 0xe1, 0xf8, 0x0f, 0xab, 0x6b, 0x96, 0x93, 0x53,
	0xf0, 0x13, 0xb3, 0xa7, 0x29, 0x47, 0x1b, 0xba, 0x91, 0xb3, 0xf2, 0x77, 0x3e, 0xeb, 0xea, 0x26,
	0xdb, 0x72, 0x6a, 0x0f, 0x52, 0x57, 0xa1, 0x17, 0x39, 0x2b, 0x6f, 0xe7, 0xab, 0xae, 0x26, 0xaf,
	0x21, 0xd8, 0x48, 0xcb, 0x0e, 0x77, 0xcd, 0xdc, 0x49, 0x2b, 0x0c, 0x44, 0x0f, 0xc8, 0x12, 0xe6,
	0x97, 0xa8, 0x19, 0xe5, 0x9b, 0xba, 0x4a, 0x51, 0x87, 0x7e, 0xdb, 0x30, 0x37, 0x0f, 0x18, 0x79,
	0x07, 0x8b, 0xc4, 0x6c, 0xb5, 0xbc, 0xbd, 0x3b, 0x17, 0xcd, 0xbc, 0x3c, 0x9c, 0xb6, 0xf3, 0x17,
	0xec, 0x11, 0x25, 0x2b, 0x38, 0xe9, 0xbf, 0xe2, 0x07, 0x6a, 0xd3, 0xdc, 0x55, 0xd0, 0xea, 0x4e,
	0xd4, 0x63, 0xfc, 0xe1, 0x97, 0x0b, 0xde, 0xfa, 0x1a, 0x85, 0x25, 0x09, 0xb8, 0x89, 0x38, 0x48,
	0x12, 0xc5, 0x47, 0xd6, 0x13, 0x77, 0xbb, 0x39, 0x7d, 0x7b, 0xb4, 0xa3, 0xbb, 0xda, 0x2f, 0xe0,
	0x9e, 0x67, 0xc5, 0xff, 0xa8, 0xfe, 0xd9, 0x41, 0xbe, 0x41, 0x70, 0x41, 0x6b, 0x91, 0x15, 0x6b,
	0xa5, 0x06, 0x11, 0x7e, 0x87, 0xd9, 0x67, 0xa4, 0xdc, 0x16, 0x9f, 0x0a, 0xcc, 0xca, 0x41, 0x94,
	0x97, 0xf0, 0x6c, 0x87, 0x95, 0xbc, 0xc1, 0xfe, 0xd7, 0x1b, 0x42, 0xba, 0x87, 0x45, 0x22, 0x8c,
	0xa5, 0x9c, 0x0f, 0x6b, 0x9d, 0x5f, 0x30, 0x63, 0x3b, 0xa5, 0x79, 0x92, 0xb3, 0x93, 0xbc, 0x77,
	0xd2, 0x49, 0xfb, 0xb0, 0x3f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb9, 0xb1, 0xdb, 0xed,
	0x03, 0x00, 0x00,
}
