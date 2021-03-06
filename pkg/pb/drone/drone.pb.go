// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/drone/drone.proto

/*
Package pbdrone is a generated protocol buffer package.

It is generated from these files:
	openpitrix/drone/drone.proto

It has these top-level messages:
*/
package pbdrone

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import openpitrix_types "openpitrix.io/openpitrix/pkg/pb/types"
import openpitrix_types1 "openpitrix.io/openpitrix/pkg/pb/types"
import openpitrix_types2 "openpitrix.io/openpitrix/pkg/pb/types"
import openpitrix_types4 "openpitrix.io/openpitrix/pkg/pb/types"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DroneService service

type DroneServiceClient interface {
	GetDroneConfig(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types2.DroneConfig, error)
	GetConfdConfig(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types1.ConfdConfig, error)
	SetConfdConfig(ctx context.Context, in *openpitrix_types1.ConfdConfig, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
	GetFrontgateConfig(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types4.FrontgateConfig, error)
	SetFrontgateConfig(ctx context.Context, in *openpitrix_types4.FrontgateConfig, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
	IsConfdRunning(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Bool, error)
	StartConfd(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
	StopConfd(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
	GetTemplateFiles(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.StringList, error)
	GetValues(ctx context.Context, in *openpitrix_types.StringList, opts ...grpc.CallOption) (*openpitrix_types.StringMap, error)
	PingPilot(ctx context.Context, in *openpitrix_types4.FrontgateEndpoint, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
	PingFrontgate(ctx context.Context, in *openpitrix_types4.FrontgateEndpoint, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
	PingDrone(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Empty, error)
}

type droneServiceClient struct {
	cc *grpc.ClientConn
}

func NewDroneServiceClient(cc *grpc.ClientConn) DroneServiceClient {
	return &droneServiceClient{cc}
}

func (c *droneServiceClient) GetDroneConfig(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types2.DroneConfig, error) {
	out := new(openpitrix_types2.DroneConfig)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetDroneConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetConfdConfig(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types1.ConfdConfig, error) {
	out := new(openpitrix_types1.ConfdConfig)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetConfdConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetConfdConfig(ctx context.Context, in *openpitrix_types1.ConfdConfig, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/SetConfdConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetFrontgateConfig(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types4.FrontgateConfig, error) {
	out := new(openpitrix_types4.FrontgateConfig)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetFrontgateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) SetFrontgateConfig(ctx context.Context, in *openpitrix_types4.FrontgateConfig, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/SetFrontgateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) IsConfdRunning(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Bool, error) {
	out := new(openpitrix_types.Bool)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/IsConfdRunning", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StartConfd(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/StartConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) StopConfd(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/StopConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetTemplateFiles(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.StringList, error) {
	out := new(openpitrix_types.StringList)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetTemplateFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) GetValues(ctx context.Context, in *openpitrix_types.StringList, opts ...grpc.CallOption) (*openpitrix_types.StringMap, error) {
	out := new(openpitrix_types.StringMap)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/GetValues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingPilot(ctx context.Context, in *openpitrix_types4.FrontgateEndpoint, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/PingPilot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingFrontgate(ctx context.Context, in *openpitrix_types4.FrontgateEndpoint, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/PingFrontgate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *droneServiceClient) PingDrone(ctx context.Context, in *openpitrix_types.Empty, opts ...grpc.CallOption) (*openpitrix_types.Empty, error) {
	out := new(openpitrix_types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.drone.DroneService/PingDrone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DroneService service

type DroneServiceServer interface {
	GetDroneConfig(context.Context, *openpitrix_types.Empty) (*openpitrix_types2.DroneConfig, error)
	GetConfdConfig(context.Context, *openpitrix_types.Empty) (*openpitrix_types1.ConfdConfig, error)
	SetConfdConfig(context.Context, *openpitrix_types1.ConfdConfig) (*openpitrix_types.Empty, error)
	GetFrontgateConfig(context.Context, *openpitrix_types.Empty) (*openpitrix_types4.FrontgateConfig, error)
	SetFrontgateConfig(context.Context, *openpitrix_types4.FrontgateConfig) (*openpitrix_types.Empty, error)
	IsConfdRunning(context.Context, *openpitrix_types.Empty) (*openpitrix_types.Bool, error)
	StartConfd(context.Context, *openpitrix_types.Empty) (*openpitrix_types.Empty, error)
	StopConfd(context.Context, *openpitrix_types.Empty) (*openpitrix_types.Empty, error)
	GetTemplateFiles(context.Context, *openpitrix_types.Empty) (*openpitrix_types.StringList, error)
	GetValues(context.Context, *openpitrix_types.StringList) (*openpitrix_types.StringMap, error)
	PingPilot(context.Context, *openpitrix_types4.FrontgateEndpoint) (*openpitrix_types.Empty, error)
	PingFrontgate(context.Context, *openpitrix_types4.FrontgateEndpoint) (*openpitrix_types.Empty, error)
	PingDrone(context.Context, *openpitrix_types.Empty) (*openpitrix_types.Empty, error)
}

func RegisterDroneServiceServer(s *grpc.Server, srv DroneServiceServer) {
	s.RegisterService(&_DroneService_serviceDesc, srv)
}

func _DroneService_GetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetDroneConfig(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetConfdConfig(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetConfdConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types1.ConfdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetConfdConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/SetConfdConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetConfdConfig(ctx, req.(*openpitrix_types1.ConfdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetFrontgateConfig(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_SetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types4.FrontgateConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).SetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/SetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).SetFrontgateConfig(ctx, req.(*openpitrix_types4.FrontgateConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_IsConfdRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).IsConfdRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/IsConfdRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).IsConfdRunning(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StartConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StartConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/StartConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StartConfd(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_StopConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).StopConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/StopConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).StopConfd(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetTemplateFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetTemplateFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetTemplateFiles(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.StringList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).GetValues(ctx, req.(*openpitrix_types.StringList))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingPilot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types4.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingPilot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/PingPilot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingPilot(ctx, req.(*openpitrix_types4.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingFrontgate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types4.FrontgateEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingFrontgate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/PingFrontgate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingFrontgate(ctx, req.(*openpitrix_types4.FrontgateEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _DroneService_PingDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(openpitrix_types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DroneServiceServer).PingDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.drone.DroneService/PingDrone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DroneServiceServer).PingDrone(ctx, req.(*openpitrix_types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DroneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.drone.DroneService",
	HandlerType: (*DroneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDroneConfig",
			Handler:    _DroneService_GetDroneConfig_Handler,
		},
		{
			MethodName: "GetConfdConfig",
			Handler:    _DroneService_GetConfdConfig_Handler,
		},
		{
			MethodName: "SetConfdConfig",
			Handler:    _DroneService_SetConfdConfig_Handler,
		},
		{
			MethodName: "GetFrontgateConfig",
			Handler:    _DroneService_GetFrontgateConfig_Handler,
		},
		{
			MethodName: "SetFrontgateConfig",
			Handler:    _DroneService_SetFrontgateConfig_Handler,
		},
		{
			MethodName: "IsConfdRunning",
			Handler:    _DroneService_IsConfdRunning_Handler,
		},
		{
			MethodName: "StartConfd",
			Handler:    _DroneService_StartConfd_Handler,
		},
		{
			MethodName: "StopConfd",
			Handler:    _DroneService_StopConfd_Handler,
		},
		{
			MethodName: "GetTemplateFiles",
			Handler:    _DroneService_GetTemplateFiles_Handler,
		},
		{
			MethodName: "GetValues",
			Handler:    _DroneService_GetValues_Handler,
		},
		{
			MethodName: "PingPilot",
			Handler:    _DroneService_PingPilot_Handler,
		},
		{
			MethodName: "PingFrontgate",
			Handler:    _DroneService_PingFrontgate_Handler,
		},
		{
			MethodName: "PingDrone",
			Handler:    _DroneService_PingDrone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openpitrix/drone/drone.proto",
}

func init() { proto.RegisterFile("openpitrix/drone/drone.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x6f, 0x4a, 0x07, 0x2d, 0x65, 0x0f, 0x0a, 0xb1, 0x82, 0xe2, 0xd9, 0x06, 0xf4, 0x28,
	0x15, 0xac, 0xb6, 0x69, 0xc1, 0x42, 0x30, 0xe2, 0xc1, 0x5b, 0xda, 0x4c, 0xc3, 0x62, 0xba, 0xb3,
	0x24, 0x53, 0xb1, 0x5f, 0xce, 0xcf, 0x26, 0xcd, 0xfa, 0x67, 0x35, 0x6e, 0x8a, 0xf5, 0xb2, 0x21,
	0xfb, 0x9b, 0xf7, 0xf6, 0xed, 0x84, 0x09, 0xb4, 0x49, 0xa3, 0xd2, 0x92, 0x73, 0xf9, 0xe2, 0x27,
	0x39, 0x29, 0x34, 0x6b, 0x47, 0xe7, 0xc4, 0x24, 0x5a, 0x5f, 0xb4, 0x53, 0xee, 0x7b, 0x76, 0x3d,
	0x2f, 0x35, 0x16, 0x66, 0x35, 0xf5, 0xbf, 0xd0, 0x29, 0xa9, 0x59, 0xe2, 0xa4, 0xd6, 0x59, 0xde,
	0x51, 0x85, 0xce, 0x72, 0x52, 0x9c, 0xc6, 0xfc, 0x5e, 0x71, 0xf6, 0xba, 0x0d, 0x3b, 0x37, 0x2b,
	0x45, 0x84, 0xf9, 0xb3, 0x9c, 0xa2, 0x18, 0x42, 0x33, 0x40, 0x2e, 0xb7, 0xae, 0x49, 0xcd, 0x64,
	0x2a, 0xf6, 0x3b, 0x56, 0x62, 0x93, 0xac, 0x3f, 0xd7, 0xbc, 0xf4, 0x0e, 0xab, 0xc0, 0xd6, 0x19,
	0xa7, 0xd5, 0x4b, 0xb2, 0x81, 0x93, 0xad, 0x1b, 0x42, 0x33, 0xfa, 0xee, 0x54, 0x2f, 0xf0, 0x5c,
	0x07, 0x89, 0x10, 0x44, 0x80, 0x3c, 0xf8, 0x68, 0xc2, 0xba, 0x5c, 0xc7, 0x55, 0xf0, 0x53, 0x1b,
	0x82, 0x88, 0xaa, 0x8e, 0xeb, 0x85, 0xee, 0x8c, 0x57, 0xd0, 0x1c, 0x15, 0xe5, 0x6d, 0xee, 0x16,
	0x4a, 0x49, 0x55, 0x93, 0x6f, 0xaf, 0x0a, 0x7a, 0x44, 0x99, 0xb8, 0x04, 0x88, 0x38, 0xce, 0x4d,
	0xcb, 0xdc, 0x72, 0x67, 0x84, 0x2e, 0x34, 0x22, 0x26, 0xbd, 0xa9, 0x7c, 0x04, 0xad, 0x00, 0xf9,
	0x1e, 0xe7, 0x3a, 0x8b, 0x19, 0x07, 0x32, 0xc3, 0xc2, 0xed, 0xd2, 0xae, 0x82, 0x88, 0x73, 0xa9,
	0xd2, 0x5b, 0x59, 0xb0, 0x18, 0x40, 0x23, 0x40, 0x7e, 0x88, 0xb3, 0x05, 0x16, 0xa2, 0xb6, 0xd4,
	0x3b, 0x70, 0xd1, 0x71, 0xac, 0xc5, 0x08, 0x1a, 0xa1, 0x54, 0x69, 0x28, 0x33, 0x62, 0x71, 0x52,
	0xf3, 0x75, 0xfa, 0x2a, 0xd1, 0x24, 0x15, 0xbb, 0x6f, 0x37, 0x86, 0xdd, 0x95, 0xd5, 0xa7, 0xe2,
	0x9f, 0x76, 0x5d, 0x93, 0xac, 0x9c, 0x9c, 0xbf, 0xf7, 0xba, 0xe7, 0x3f, 0x9e, 0x5a, 0x44, 0x92,
	0x6f, 0x8d, 0xbc, 0x7e, 0x4a, 0x7d, 0x3d, 0x31, 0x7f, 0x84, 0x0b, 0x3d, 0x29, 0x9f, 0x93, 0xad,
	0x72, 0xf0, 0xcf, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x1f, 0xa2, 0x98, 0xa6, 0x04, 0x00,
	0x00,
}
