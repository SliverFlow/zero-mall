package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAddStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddStockLogic {
	return &ProductAddStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAddStock
// Author [SliverFlow]
// @desc 商品返还库存
func (l *ProductAddStockLogic) ProductAddStock(in *pb.ProductAddStockReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.ProductAddStock(l.ctx, in.GetProductID(), in.GetQuantity())
	if err != nil {
		return nil, status.Errorf(100001, "商品添加库存失败")
	}

	return &pb.ProductNil{}, nil
}
