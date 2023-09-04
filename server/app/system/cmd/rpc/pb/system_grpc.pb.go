// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: system.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SystemClient is the client API for System service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemClient interface {
	// 系统登录
	Login(ctx context.Context, in *SystemLoginReq, opts ...grpc.CallOption) (*SystemLoginReply, error)
	// 创建角色
	RoleCreate(ctx context.Context, in *CreateRole, opts ...grpc.CallOption) (*NilReply, error)
}

type systemClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemClient(cc grpc.ClientConnInterface) SystemClient {
	return &systemClient{cc}
}

func (c *systemClient) Login(ctx context.Context, in *SystemLoginReq, opts ...grpc.CallOption) (*SystemLoginReply, error) {
	out := new(SystemLoginReply)
	err := c.cc.Invoke(ctx, "/pb.system/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) RoleCreate(ctx context.Context, in *CreateRole, opts ...grpc.CallOption) (*NilReply, error) {
	out := new(NilReply)
	err := c.cc.Invoke(ctx, "/pb.system/RoleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServer is the server API for System service.
// All implementations must embed UnimplementedSystemServer
// for forward compatibility
type SystemServer interface {
	// 系统登录
	Login(context.Context, *SystemLoginReq) (*SystemLoginReply, error)
	// 创建角色
	RoleCreate(context.Context, *CreateRole) (*NilReply, error)
	mustEmbedUnimplementedSystemServer()
}

// UnimplementedSystemServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServer struct {
}

func (UnimplementedSystemServer) Login(context.Context, *SystemLoginReq) (*SystemLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSystemServer) RoleCreate(context.Context, *CreateRole) (*NilReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleCreate not implemented")
}
func (UnimplementedSystemServer) mustEmbedUnimplementedSystemServer() {}

// UnsafeSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServer will
// result in compilation errors.
type UnsafeSystemServer interface {
	mustEmbedUnimplementedSystemServer()
}

func RegisterSystemServer(s grpc.ServiceRegistrar, srv SystemServer) {
	s.RegisterService(&System_ServiceDesc, srv)
}

func _System_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.system/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).Login(ctx, req.(*SystemLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_RoleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).RoleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.system/RoleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).RoleCreate(ctx, req.(*CreateRole))
	}
	return interceptor(ctx, in, info, handler)
}

// System_ServiceDesc is the grpc.ServiceDesc for System service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var System_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.system",
	HandlerType: (*SystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _System_Login_Handler,
		},
		{
			MethodName: "RoleCreate",
			Handler:    _System_RoleCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system.proto",
}
