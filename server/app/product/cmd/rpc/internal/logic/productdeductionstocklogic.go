package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeductionStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDeductionStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeductionStockLogic {
	return &ProductDeductionStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductDeductionStock
// Author [SliverFlow]
// @desc  商品扣减库存
func (l *ProductDeductionStockLogic) ProductDeductionStock(in *pb.ProductDeductionStockReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.ProductDeductionStock(l.ctx, in.GetProductID(), in.GetQuantity())
	if err != nil {
		return nil, status.Errorf(10001, "商品扣减库存失败")
	}

	return &pb.ProductNil{}, nil
}
