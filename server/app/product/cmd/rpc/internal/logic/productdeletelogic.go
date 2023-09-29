package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductDelete
// Author [SliverFlow]
// @desc  删除商品
func (l *ProductDeleteLogic) ProductDelete(in *pb.ProductDeleteReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.ProductDelete(l.ctx, in.IDs)
	if err != nil {
		return nil, status.Errorf(100001, "商品删除失败")
	}

	return &pb.ProductNil{}, nil
}
