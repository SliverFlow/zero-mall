package private

import (
	"context"
	cartpb "server/app/cart/cmd/rpc/pb"
	productpb "server/app/product/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/cart/cmd/api/internal/svc"
	"server/app/cart/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartCreateLogic {
	return &CartCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartCreate
// Author [SliverFlow]
// @desc 用户端添加购物车
func (l *CartCreateLogic) CartCreate(req *types.CartCreateReq) (resp *types.Nil, err error) {
	// 验证商品
	pbfind, err := l.svcCtx.ProductRpc.ProductFind(l.ctx, &productpb.ProductFindReq{ProductID: req.ProductID})
	if err != nil {
		return nil, err
	}

	// 验证库存
	if pbfind.Stock-req.Quantity < 0 {
		return nil, xerr.NewErrMsg("当前商品库存不足")
	}

	// 获取用户信息
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户信息失败")
	}

	// 添加购物车
	_, err = l.svcCtx.CartRpc.CartCreate(l.ctx, &cartpb.CartCreateReq{
		UserID:      uuid,
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		Checked:     req.Checked,
		ProductName: pbfind.Name,
	})
	if err != nil {
		return nil, err
	}

	return
}
