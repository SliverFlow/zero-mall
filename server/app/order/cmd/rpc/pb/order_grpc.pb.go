// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: order.proto

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	OrderCreate(ctx context.Context, in *OrderCreateReq, opts ...grpc.CallOption) (*OrderIDReply, error)
	OrderDelete(ctx context.Context, in *OrderDeleteReq, opts ...grpc.CallOption) (*OrderNil, error)
	OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListReply, error)
	OrderFind(ctx context.Context, in *OrderFindReq, opts ...grpc.CallOption) (*OrderReply, error)
	OrderDeleteByID(ctx context.Context, in *OrderDeleteByIDReq, opts ...grpc.CallOption) (*OrderNil, error)
	// 订单子表相关
	OrderItemDelete(ctx context.Context, in *OrderItemIdReq, opts ...grpc.CallOption) (*OrderNil, error)
	OrderPageList(ctx context.Context, in *OrderPageListReq, opts ...grpc.CallOption) (*OrderPageListReply, error)
	OrderDisable(ctx context.Context, in *OrderDisableReq, opts ...grpc.CallOption) (*OrderNil, error)
	OrderItemDeleteByID(ctx context.Context, in *OrderItemDeleteByIDReq, opts ...grpc.CallOption) (*OrderNil, error)
	OrderPageListForUser(ctx context.Context, in *OrderPageListReq, opts ...grpc.CallOption) (*OrderPageListReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) OrderCreate(ctx context.Context, in *OrderCreateReq, opts ...grpc.CallOption) (*OrderIDReply, error) {
	out := new(OrderIDReply)
	err := c.cc.Invoke(ctx, "/pb.order/orderCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDelete(ctx context.Context, in *OrderDeleteReq, opts ...grpc.CallOption) (*OrderNil, error) {
	out := new(OrderNil)
	err := c.cc.Invoke(ctx, "/pb.order/orderDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListReply, error) {
	out := new(OrderListReply)
	err := c.cc.Invoke(ctx, "/pb.order/orderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderFind(ctx context.Context, in *OrderFindReq, opts ...grpc.CallOption) (*OrderReply, error) {
	out := new(OrderReply)
	err := c.cc.Invoke(ctx, "/pb.order/orderFind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDeleteByID(ctx context.Context, in *OrderDeleteByIDReq, opts ...grpc.CallOption) (*OrderNil, error) {
	out := new(OrderNil)
	err := c.cc.Invoke(ctx, "/pb.order/orderDeleteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderItemDelete(ctx context.Context, in *OrderItemIdReq, opts ...grpc.CallOption) (*OrderNil, error) {
	out := new(OrderNil)
	err := c.cc.Invoke(ctx, "/pb.order/orderItemDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderPageList(ctx context.Context, in *OrderPageListReq, opts ...grpc.CallOption) (*OrderPageListReply, error) {
	out := new(OrderPageListReply)
	err := c.cc.Invoke(ctx, "/pb.order/orderPageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDisable(ctx context.Context, in *OrderDisableReq, opts ...grpc.CallOption) (*OrderNil, error) {
	out := new(OrderNil)
	err := c.cc.Invoke(ctx, "/pb.order/orderDisable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderItemDeleteByID(ctx context.Context, in *OrderItemDeleteByIDReq, opts ...grpc.CallOption) (*OrderNil, error) {
	out := new(OrderNil)
	err := c.cc.Invoke(ctx, "/pb.order/orderItemDeleteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderPageListForUser(ctx context.Context, in *OrderPageListReq, opts ...grpc.CallOption) (*OrderPageListReply, error) {
	out := new(OrderPageListReply)
	err := c.cc.Invoke(ctx, "/pb.order/orderPageListForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	OrderCreate(context.Context, *OrderCreateReq) (*OrderIDReply, error)
	OrderDelete(context.Context, *OrderDeleteReq) (*OrderNil, error)
	OrderList(context.Context, *OrderListReq) (*OrderListReply, error)
	OrderFind(context.Context, *OrderFindReq) (*OrderReply, error)
	OrderDeleteByID(context.Context, *OrderDeleteByIDReq) (*OrderNil, error)
	// 订单子表相关
	OrderItemDelete(context.Context, *OrderItemIdReq) (*OrderNil, error)
	OrderPageList(context.Context, *OrderPageListReq) (*OrderPageListReply, error)
	OrderDisable(context.Context, *OrderDisableReq) (*OrderNil, error)
	OrderItemDeleteByID(context.Context, *OrderItemDeleteByIDReq) (*OrderNil, error)
	OrderPageListForUser(context.Context, *OrderPageListReq) (*OrderPageListReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) OrderCreate(context.Context, *OrderCreateReq) (*OrderIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCreate not implemented")
}
func (UnimplementedOrderServer) OrderDelete(context.Context, *OrderDeleteReq) (*OrderNil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDelete not implemented")
}
func (UnimplementedOrderServer) OrderList(context.Context, *OrderListReq) (*OrderListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServer) OrderFind(context.Context, *OrderFindReq) (*OrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderFind not implemented")
}
func (UnimplementedOrderServer) OrderDeleteByID(context.Context, *OrderDeleteByIDReq) (*OrderNil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDeleteByID not implemented")
}
func (UnimplementedOrderServer) OrderItemDelete(context.Context, *OrderItemIdReq) (*OrderNil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderItemDelete not implemented")
}
func (UnimplementedOrderServer) OrderPageList(context.Context, *OrderPageListReq) (*OrderPageListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPageList not implemented")
}
func (UnimplementedOrderServer) OrderDisable(context.Context, *OrderDisableReq) (*OrderNil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDisable not implemented")
}
func (UnimplementedOrderServer) OrderItemDeleteByID(context.Context, *OrderItemDeleteByIDReq) (*OrderNil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderItemDeleteByID not implemented")
}
func (UnimplementedOrderServer) OrderPageListForUser(context.Context, *OrderPageListReq) (*OrderPageListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPageListForUser not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_OrderCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderCreate(ctx, req.(*OrderCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDelete(ctx, req.(*OrderDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderList(ctx, req.(*OrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderFind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderFind(ctx, req.(*OrderFindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDeleteByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderDeleteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDeleteByID(ctx, req.(*OrderDeleteByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderItemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderItemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderItemDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderItemDelete(ctx, req.(*OrderItemIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderPageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderPageList(ctx, req.(*OrderPageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDisableReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDisable(ctx, req.(*OrderDisableReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderItemDeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemDeleteByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderItemDeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderItemDeleteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderItemDeleteByID(ctx, req.(*OrderItemDeleteByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderPageListForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderPageListForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.order/orderPageListForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderPageListForUser(ctx, req.(*OrderPageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "orderCreate",
			Handler:    _Order_OrderCreate_Handler,
		},
		{
			MethodName: "orderDelete",
			Handler:    _Order_OrderDelete_Handler,
		},
		{
			MethodName: "orderList",
			Handler:    _Order_OrderList_Handler,
		},
		{
			MethodName: "orderFind",
			Handler:    _Order_OrderFind_Handler,
		},
		{
			MethodName: "orderDeleteByID",
			Handler:    _Order_OrderDeleteByID_Handler,
		},
		{
			MethodName: "orderItemDelete",
			Handler:    _Order_OrderItemDelete_Handler,
		},
		{
			MethodName: "orderPageList",
			Handler:    _Order_OrderPageList_Handler,
		},
		{
			MethodName: "orderDisable",
			Handler:    _Order_OrderDisable_Handler,
		},
		{
			MethodName: "orderItemDeleteByID",
			Handler:    _Order_OrderItemDeleteByID_Handler,
		},
		{
			MethodName: "orderPageListForUser",
			Handler:    _Order_OrderPageListForUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
