// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pkg/proto/agent.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error)
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	DestroyInstance(ctx context.Context, in *DestroyInstanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartInstance(ctx context.Context, in *StartInstanceRequest, opts ...grpc.CallOption) (*StartInstanceResponse, error)
	StopInstance(ctx context.Context, in *StopInstanceRequest, opts ...grpc.CallOption) (*StopInstanceResponse, error)
	InstanceExec(ctx context.Context, in *InstanceExecRequest, opts ...grpc.CallOption) (*InstanceExecResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentService/CreateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error) {
	out := new(GetInstanceResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentService/GetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentService/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) DestroyInstance(ctx context.Context, in *DestroyInstanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.AgentService/DestroyInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StartInstance(ctx context.Context, in *StartInstanceRequest, opts ...grpc.CallOption) (*StartInstanceResponse, error) {
	out := new(StartInstanceResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentService/StartInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) StopInstance(ctx context.Context, in *StopInstanceRequest, opts ...grpc.CallOption) (*StopInstanceResponse, error) {
	out := new(StopInstanceResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentService/StopInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) InstanceExec(ctx context.Context, in *InstanceExecRequest, opts ...grpc.CallOption) (*InstanceExecResponse, error) {
	out := new(InstanceExecResponse)
	err := c.cc.Invoke(ctx, "/proto.AgentService/InstanceExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations should embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error)
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	DestroyInstance(context.Context, *DestroyInstanceRequest) (*emptypb.Empty, error)
	StartInstance(context.Context, *StartInstanceRequest) (*StartInstanceResponse, error)
	StopInstance(context.Context, *StopInstanceRequest) (*StopInstanceResponse, error)
	InstanceExec(context.Context, *InstanceExecRequest) (*InstanceExecResponse, error)
}

// UnimplementedAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedAgentServiceServer) GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedAgentServiceServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedAgentServiceServer) DestroyInstance(context.Context, *DestroyInstanceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyInstance not implemented")
}
func (UnimplementedAgentServiceServer) StartInstance(context.Context, *StartInstanceRequest) (*StartInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartInstance not implemented")
}
func (UnimplementedAgentServiceServer) StopInstance(context.Context, *StopInstanceRequest) (*StopInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopInstance not implemented")
}
func (UnimplementedAgentServiceServer) InstanceExec(context.Context, *InstanceExecRequest) (*InstanceExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstanceExec not implemented")
}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/GetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_DestroyInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).DestroyInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/DestroyInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).DestroyInstance(ctx, req.(*DestroyInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StartInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).StartInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/StartInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).StartInstance(ctx, req.(*StartInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_StopInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).StopInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/StopInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).StopInstance(ctx, req.(*StopInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_InstanceExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).InstanceExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/InstanceExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).InstanceExec(ctx, req.(*InstanceExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstance",
			Handler:    _AgentService_CreateInstance_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _AgentService_GetInstance_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _AgentService_ListInstances_Handler,
		},
		{
			MethodName: "DestroyInstance",
			Handler:    _AgentService_DestroyInstance_Handler,
		},
		{
			MethodName: "StartInstance",
			Handler:    _AgentService_StartInstance_Handler,
		},
		{
			MethodName: "StopInstance",
			Handler:    _AgentService_StopInstance_Handler,
		},
		{
			MethodName: "InstanceExec",
			Handler:    _AgentService_InstanceExec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/agent.proto",
}
