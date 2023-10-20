package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"
	"server/app/cart/model"
	"server/common"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartCreateLogic {
	return &CartCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartCreate
// Author [SliverFlow]
// @desc 添加购物车
func (l *CartCreateLogic) CartCreate(in *pb.CartCreateReq) (*pb.CartNil, error) {
	err := l.svcCtx.CartModel.CartCreate(l.ctx, &model.Cart{
		CartID:      strconv.FormatInt(common.SnowWorker.NextId(), 10),
		UserID:      in.GetUserID(),
		ProductID:   in.GetProductID(),
		Quantity:    in.GetQuantity(),
		Checked:     in.GetChecked(),
		ProductName: in.GetProductName(),
	})
	if err != nil {
		return nil, status.Errorf(100001, "商品添加购物车失败")
	}

	return &pb.CartNil{}, nil
}
