package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteLogic {
	return &CategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CategoryDelete
// Author [SliverFlow]
// @desc 单个删除分类
func (l *CategoryDeleteLogic) CategoryDelete(in *pb.CategoryIDReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.CategoryDelete(l.ctx, in.GetCategoryID())
	if err != nil {
		logx.Error("category delete err", err.Error())
		return nil, status.Errorf(100001, "分类删除失败")
	}

	return &pb.ProductNil{}, nil
}
