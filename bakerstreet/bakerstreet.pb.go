// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bakerstreet/bakerstreet.proto

/*
Package com_appknox_bakerstreet is a generated protocol buffer package.

It is generated from these files:
	bakerstreet/bakerstreet.proto

It has these top-level messages:
	Message
	Apps
	App
	Device
	Finding
	InstallReq
	ConfigProxyReq
	Empty
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
	Data string `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Apps struct {
	App []*App `protobuf:"bytes,1,rep,name=App" json:"App,omitempty"`
}

func (m *Apps) Reset()                    { *m = Apps{} }
func (m *Apps) String() string            { return proto.CompactTextString(m) }
func (*Apps) ProtoMessage()               {}
func (*Apps) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Apps) GetApp() []*App {
	if m != nil {
		return m.App
	}
	return nil
}

type App struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Device struct {
	Uuid            string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	IsTablet        bool   `protobuf:"varint,2,opt,name=IsTablet" json:"IsTablet,omitempty"`
	Platform        int32  `protobuf:"varint,3,opt,name=Platform" json:"Platform,omitempty"`
	PlatformVersion string `protobuf:"bytes,4,opt,name=PlatformVersion" json:"PlatformVersion,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

func (m *Device) GetPlatformVersion() string {
	if m != nil {
		return m.PlatformVersion
	}
	return ""
}

type Finding struct {
	Title       string `protobuf:"bytes,1,opt,name=Title" json:"Title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
}

func (m *Finding) Reset()                    { *m = Finding{} }
func (m *Finding) String() string            { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()               {}
func (*Finding) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Finding) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Finding) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type InstallReq struct {
	URL string `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
}

func (m *InstallReq) Reset()                    { *m = InstallReq{} }
func (m *InstallReq) String() string            { return proto.CompactTextString(m) }
func (*InstallReq) ProtoMessage()               {}
func (*InstallReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InstallReq) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type ConfigProxyReq struct {
	IP   string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	PORT string `protobuf:"bytes,2,opt,name=PORT" json:"PORT,omitempty"`
}

func (m *ConfigProxyReq) Reset()                    { *m = ConfigProxyReq{} }
func (m *ConfigProxyReq) String() string            { return proto.CompactTextString(m) }
func (*ConfigProxyReq) ProtoMessage()               {}
func (*ConfigProxyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ConfigProxyReq) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ConfigProxyReq) GetPORT() string {
	if m != nil {
		return m.PORT
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Message)(nil), "com.appknox.bakerstreet.Message")
	proto.RegisterType((*Apps)(nil), "com.appknox.bakerstreet.Apps")
	proto.RegisterType((*App)(nil), "com.appknox.bakerstreet.App")
	proto.RegisterType((*Device)(nil), "com.appknox.bakerstreet.Device")
	proto.RegisterType((*Finding)(nil), "com.appknox.bakerstreet.Finding")
	proto.RegisterType((*InstallReq)(nil), "com.appknox.bakerstreet.InstallReq")
	proto.RegisterType((*ConfigProxyReq)(nil), "com.appknox.bakerstreet.ConfigProxyReq")
	proto.RegisterType((*Empty)(nil), "com.appknox.bakerstreet.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Moriarty service

type MoriartyClient interface {
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	LaunchApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error)
	ClearProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	RemovePackage(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error)
	InstallPackage(ctx context.Context, in *InstallReq, opts ...grpc.CallOption) (*Message, error)
	ConfigureProxy(ctx context.Context, in *ConfigProxyReq, opts ...grpc.CallOption) (*Message, error)
	ConfigureGadget(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Device, error)
	ListPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Apps, error)
}

type moriartyClient struct {
	cc *grpc.ClientConn
}

func NewMoriartyClient(cc *grpc.ClientConn) MoriartyClient {
	return &moriartyClient{cc}
}

func (c *moriartyClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) LaunchApp(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/LaunchApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ClearProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ClearProxy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) RemovePackage(ctx context.Context, in *App, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/RemovePackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) InstallPackage(ctx context.Context, in *InstallReq, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/InstallPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ConfigureProxy(ctx context.Context, in *ConfigProxyReq, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ConfigureProxy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ConfigureGadget(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ConfigureGadget", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moriartyClient) ListPackages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Apps, error) {
	out := new(Apps)
	err := grpc.Invoke(ctx, "/com.appknox.bakerstreet.Moriarty/ListPackages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Moriarty service

type MoriartyServer interface {
	Echo(context.Context, *Message) (*Message, error)
	LaunchApp(context.Context, *App) (*Message, error)
	ClearProxy(context.Context, *Empty) (*Message, error)
	HealthCheck(context.Context, *Empty) (*Message, error)
	RemovePackage(context.Context, *App) (*Message, error)
	InstallPackage(context.Context, *InstallReq) (*Message, error)
	ConfigureProxy(context.Context, *ConfigProxyReq) (*Message, error)
	ConfigureGadget(context.Context, *Message) (*Message, error)
	Info(context.Context, *Empty) (*Device, error)
	ListPackages(context.Context, *Empty) (*Apps, error)
}

func RegisterMoriartyServer(s *grpc.Server, srv MoriartyServer) {
	s.RegisterService(&_Moriarty_serviceDesc, srv)
}

func _Moriarty_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_LaunchApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).LaunchApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/LaunchApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).LaunchApp(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ClearProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ClearProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ClearProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ClearProxy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_RemovePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(App)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).RemovePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/RemovePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).RemovePackage(ctx, req.(*App))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_InstallPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).InstallPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/InstallPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).InstallPackage(ctx, req.(*InstallReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ConfigureProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigProxyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ConfigureProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ConfigureProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ConfigureProxy(ctx, req.(*ConfigProxyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ConfigureGadget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ConfigureGadget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ConfigureGadget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ConfigureGadget(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moriarty_ListPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoriartyServer).ListPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.appknox.bakerstreet.Moriarty/ListPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoriartyServer).ListPackages(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Moriarty_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.appknox.bakerstreet.Moriarty",
	HandlerType: (*MoriartyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Moriarty_Echo_Handler,
		},
		{
			MethodName: "LaunchApp",
			Handler:    _Moriarty_LaunchApp_Handler,
		},
		{
			MethodName: "ClearProxy",
			Handler:    _Moriarty_ClearProxy_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _Moriarty_HealthCheck_Handler,
		},
		{
			MethodName: "RemovePackage",
			Handler:    _Moriarty_RemovePackage_Handler,
		},
		{
			MethodName: "InstallPackage",
			Handler:    _Moriarty_InstallPackage_Handler,
		},
		{
			MethodName: "ConfigureProxy",
			Handler:    _Moriarty_ConfigureProxy_Handler,
		},
		{
			MethodName: "ConfigureGadget",
			Handler:    _Moriarty_ConfigureGadget_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Moriarty_Info_Handler,
		},
		{
			MethodName: "ListPackages",
			Handler:    _Moriarty_ListPackages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bakerstreet/bakerstreet.proto",
}

func init() { proto.RegisterFile("bakerstreet/bakerstreet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x95, 0x36, 0xfd, 0xb3, 0xd3, 0xf7, 0xed, 0x90, 0x85, 0x44, 0x98, 0xe8, 0x88, 0xc2,
	0x05, 0xb9, 0x0a, 0xd2, 0x40, 0xdc, 0x57, 0xed, 0x18, 0x81, 0x76, 0x8b, 0xa2, 0x16, 0x2e, 0xb8,
	0x72, 0x53, 0xb7, 0xb5, 0x9a, 0xc4, 0xc6, 0x76, 0xa7, 0xf5, 0x86, 0x0f, 0xc6, 0xa7, 0x43, 0xf6,
	0x92, 0x52, 0x90, 0xb2, 0x56, 0x62, 0x77, 0x8f, 0x7d, 0x1e, 0xff, 0x92, 0xe3, 0xe7, 0xc8, 0xd0,
	0x9b, 0xe1, 0x35, 0x11, 0x52, 0x09, 0x42, 0xd4, 0x9b, 0x3d, 0x1d, 0x70, 0xc1, 0x14, 0x43, 0xcf,
	0x12, 0x96, 0x05, 0x98, 0xf3, 0x75, 0xce, 0xee, 0x82, 0xbd, 0xb2, 0xd7, 0x83, 0xd6, 0x98, 0x48,
	0x89, 0x97, 0x04, 0x21, 0xb0, 0x87, 0x58, 0x61, 0xc7, 0x72, 0x2d, 0xff, 0x24, 0x36, 0xda, 0x7b,
	0x0f, 0x76, 0x9f, 0x73, 0x89, 0x02, 0xa8, 0xf7, 0x39, 0x77, 0x2c, 0xb7, 0xee, 0x77, 0x2e, 0x5e,
	0x04, 0x15, 0xb4, 0xa0, 0xcf, 0x79, 0xac, 0x8d, 0xde, 0x73, 0xe3, 0xd7, 0xc8, 0x6b, 0x9c, 0x91,
	0x12, 0xa9, 0xb5, 0xf7, 0x03, 0x9a, 0x43, 0x72, 0x4b, 0x13, 0xf3, 0xc1, 0xe9, 0x86, 0xce, 0xcb,
	0xaa, 0xd6, 0xe8, 0x0c, 0xda, 0xa1, 0x9c, 0xe0, 0x59, 0x4a, 0x94, 0x53, 0x73, 0x2d, 0xbf, 0x1d,
	0xef, 0xd6, 0xba, 0x16, 0xa5, 0x58, 0x2d, 0x98, 0xc8, 0x9c, 0xba, 0x6b, 0xf9, 0x8d, 0x78, 0xb7,
	0x46, 0x3e, 0x9c, 0x96, 0xfa, 0x0b, 0x11, 0x92, 0xb2, 0xdc, 0xb1, 0x0d, 0xf6, 0xef, 0x6d, 0xaf,
	0x0f, 0xad, 0x0f, 0x34, 0x9f, 0xd3, 0x7c, 0x89, 0x9e, 0x42, 0x63, 0x42, 0x55, 0x5a, 0xfe, 0xdf,
	0xfd, 0x02, 0xb9, 0xd0, 0x19, 0x12, 0x99, 0x08, 0xca, 0x95, 0xc6, 0xd4, 0x4c, 0x6d, 0x7f, 0xcb,
	0x3b, 0x07, 0x08, 0x73, 0xa9, 0x70, 0x9a, 0xc6, 0xe4, 0x3b, 0x7a, 0x02, 0xf5, 0x69, 0x3c, 0x2a,
	0x18, 0x5a, 0x7a, 0xef, 0xa0, 0x3b, 0x60, 0xf9, 0x82, 0x2e, 0x23, 0xc1, 0xee, 0xb6, 0xda, 0xd3,
	0x85, 0x5a, 0x18, 0x15, 0x96, 0x5a, 0x18, 0xe9, 0xd6, 0xa3, 0x9b, 0x78, 0x52, 0xc0, 0x8d, 0xf6,
	0x5a, 0xd0, 0xb8, 0xcc, 0xb8, 0xda, 0x5e, 0xfc, 0x6c, 0x42, 0x7b, 0xcc, 0x04, 0xc5, 0x42, 0x6d,
	0xd1, 0x27, 0xb0, 0x2f, 0x93, 0x15, 0x43, 0x6e, 0xe5, 0xa5, 0x17, 0xf9, 0x9d, 0x1d, 0x74, 0xa0,
	0xcf, 0x70, 0x32, 0xc2, 0x9b, 0x3c, 0x59, 0xe9, 0x6c, 0x1e, 0x4c, 0xf1, 0x08, 0xd8, 0x35, 0xc0,
	0x20, 0x25, 0x58, 0x98, 0x1e, 0xd1, 0x79, 0xa5, 0xdf, 0xf4, 0x74, 0x04, 0xef, 0x06, 0x3a, 0x1f,
	0x09, 0x4e, 0xd5, 0x6a, 0xb0, 0x22, 0xc9, 0xfa, 0x51, 0x80, 0xff, 0xc7, 0x24, 0x63, 0xb7, 0x24,
	0xc2, 0xc9, 0x5a, 0x6f, 0xfc, 0x6b, 0xc7, 0x5f, 0xa1, 0x5b, 0xc4, 0x5e, 0x12, 0x5f, 0x55, 0x9e,
	0xf9, 0x3d, 0x1f, 0x47, 0x80, 0xbf, 0x95, 0xf3, 0xb2, 0x11, 0xe4, 0xfe, 0x3a, 0x5f, 0x57, 0x9e,
	0xf9, 0x73, 0xb0, 0x8e, 0x80, 0x4f, 0xe1, 0x74, 0x07, 0xbf, 0xc2, 0xf3, 0x25, 0x51, 0x8f, 0x32,
	0x4b, 0x57, 0x60, 0x87, 0xf9, 0x82, 0x1d, 0xcc, 0xe9, 0x65, 0x65, 0xbd, 0x78, 0x05, 0xc6, 0xf0,
	0xdf, 0x88, 0x4a, 0x55, 0x5c, 0xa9, 0x3c, 0x08, 0xec, 0x3d, 0x94, 0xa2, 0x9c, 0x35, 0xcd, 0x83,
	0xf7, 0xf6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x04, 0x53, 0x88, 0x11, 0x05, 0x00, 0x00,
}
