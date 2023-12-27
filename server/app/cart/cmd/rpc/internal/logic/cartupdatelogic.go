package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"
	"server/app/cart/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateLogic {
	return &CartUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartUpdate
// Author [SliverFlow]
// @desc 更新购物车信息
func (l *CartUpdateLogic) CartUpdate(in *pb.CartUpdateReq) (*pb.CartNil, error) {
	err := l.svcCtx.CartModel.CartUpdate(l.ctx, &model.Cart{
		CartID:   in.GetCartID(),
		Quantity: in.GetQuantity(),
		Checked:  in.GetChecked(),
	})
	if err != nil {
		logx.Error("update cart info error:", err.Error())
		return nil, status.Errorf(100001, "更新购物车信息失败")
	}

	return &pb.CartNil{}, nil
}
