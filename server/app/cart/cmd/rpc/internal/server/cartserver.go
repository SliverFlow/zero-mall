// Code generated by goctl. DO NOT EDIT.
// Source: cart.proto

package server

import (
	"context"

	"server/app/cart/cmd/rpc/internal/logic"
	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"
)

type CartServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCartServer
}

func NewCartServer(svcCtx *svc.ServiceContext) *CartServer {
	return &CartServer{
		svcCtx: svcCtx,
	}
}

func (s *CartServer) CartCreate(ctx context.Context, in *pb.CartNil) (*pb.CartNil, error) {
	l := logic.NewCartCreateLogic(ctx, s.svcCtx)
	return l.CartCreate(in)
}