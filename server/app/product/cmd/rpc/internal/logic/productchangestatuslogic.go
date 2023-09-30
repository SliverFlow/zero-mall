package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductChangeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductChangeStatusLogic {
	return &ProductChangeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductChangeStatus
// Author [SliverFlow]
// @desc  更新商品状态
func (l *ProductChangeStatusLogic) ProductChangeStatus(in *pb.ProductChangeStatusReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.ProductChangeStatus(l.ctx, in.GetProductID(), in.GetStatus())
	if err != nil {
		logx.Error("update product status failed", err.Error())
		return nil, status.Errorf(100001, "更新商品状态失败")
	}

	return &pb.ProductNil{}, nil
}
