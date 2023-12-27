package private

import (
	"context"
	cartpb "server/app/cart/cmd/rpc/pb"
	productpb "server/app/product/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/cart/cmd/api/internal/svc"
	"server/app/cart/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateLogic {
	return &CartUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartUpdate
// Author [SliverFlow]
// @desc 用户端更新购物车
func (l *CartUpdateLogic) CartUpdate(req *types.CartUpdateReq) (resp *types.Nil, err error) {
	find, err := l.svcCtx.CartRpc.CartFind(l.ctx, &cartpb.CartFindReq{CartID: req.CartID})
	if err != nil {
		return nil, err
	}

	// 验证商品
	pbfind, err := l.svcCtx.ProductRpc.ProductFind(l.ctx, &productpb.ProductFindReq{ProductID: find.ProductID})
	if err != nil {
		return nil, err
	}

	// 验证库存
	if pbfind.Stock-req.Quantity < 0 {
		return nil, xerr.NewErrMsg("当前商品库存不足")
	}

	// 更新信息
	_, err = l.svcCtx.CartRpc.CartUpdate(l.ctx, &cartpb.CartUpdateReq{
		Quantity: req.Quantity,
		Checked:  req.Checked,
		CartID:   req.CartID,
	})
	if err != nil {
		return nil, err
	}

	return
}
