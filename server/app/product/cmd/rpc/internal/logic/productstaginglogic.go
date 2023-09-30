package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductStagingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductStagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductStagingLogic {
	return &ProductStagingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductStaging
// Author [SliverFlow]
// @desc 下架某商家的所有商品
func (l *ProductStagingLogic) ProductStaging(in *pb.ProductStagingReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.ProductStagingByBusinessID(l.ctx, in.GetBusinessID())
	if err != nil {
		logx.Error("product staging err", err.Error())
		return nil, status.Errorf(100001, "下架商户商品失败")
	}

	return &pb.ProductNil{}, nil
}
