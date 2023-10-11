// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"server/app/order/cmd/rpc/internal/logic"
	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) OrderCreate(ctx context.Context, in *pb.OrderCreateReq) (*pb.OrderIDReply, error) {
	l := logic.NewOrderCreateLogic(ctx, s.svcCtx)
	return l.OrderCreate(in)
}

func (s *OrderServer) OrderDelete(ctx context.Context, in *pb.OrderIdReq) (*pb.OrderNil, error) {
	l := logic.NewOrderDeleteLogic(ctx, s.svcCtx)
	return l.OrderDelete(in)
}

func (s *OrderServer) OrderItemDelete(ctx context.Context, in *pb.OrderItemIdReq) (*pb.OrderNil, error) {
	l := logic.NewOrderItemDeleteLogic(ctx, s.svcCtx)
	return l.OrderItemDelete(in)
}
