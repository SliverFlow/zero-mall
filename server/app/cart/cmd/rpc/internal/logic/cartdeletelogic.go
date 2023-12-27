package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartDeleteLogic {
	return &CartDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartDelete
// Author [SliverFlow]
// @desc 删除购物车信息
func (l *CartDeleteLogic) CartDelete(in *pb.CartDeleteReq) (*pb.CartNil, error) {
	err := l.svcCtx.CartModel.CartDelete(l.ctx, in.GetCartID())
	if err != nil {
		return nil, status.Errorf(100001, "删除购物车信息失败")
	}

	return &pb.CartNil{}, nil
}
