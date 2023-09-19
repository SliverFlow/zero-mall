package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"
	"server/app/product/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryUpdateLogic) CategoryUpdate(in *pb.CategoryUpdateReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.CategoryUpdate(l.ctx, &model.Category{
		CategoryID: in.GetCategoryID(),
		ParentId:   in.GetParentId(),
		Name:       in.GetName(),
		Status:     in.GetStatus(),
		Sorted:     in.GetSorted(),
		Type:       in.GetType(),
		BusinessID: in.GetBusinessID(),
	})
	if err != nil {
		return nil, status.Errorf(100001, "更新分类失败")
	}

	return &pb.ProductNil{}, nil
}
