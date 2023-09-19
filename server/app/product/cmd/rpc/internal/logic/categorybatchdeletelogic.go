package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryBatchDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryBatchDeleteLogic {
	return &CategoryBatchDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryBatchDeleteLogic) CategoryBatchDelete(in *pb.CategoryIDSReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.CategoryBatchDelete(l.ctx, in.IDS)
	if err != nil {
		return nil, status.Errorf(100001, "批量删除分类失败")
	}
	return &pb.ProductNil{}, nil
}
