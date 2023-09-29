package category

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListAllLogic {
	return &CategoryListAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryListAll
// Author [SliverFlow]
// @desc 分类列表管理
func (l *CategoryListAllLogic) CategoryListAll(req *types.NilReq) (resp *types.CategoryListAllReply, err error) {
	reply, err := l.svcCtx.ProductRpc.CategoryListAll(l.ctx, &pb.ProductNil{})
	if err != nil {
		return nil, err
	}

	list := make([]types.Category, 0)
	if len(reply.List) != 0 {
		list = categoryTree("0", reply.List)
	}
	return &types.CategoryListAllReply{List: list}, nil
}

func categoryTree(root string, enter []*pb.Category) (rep []types.Category) {
	list := make([]types.Category, 0)
	for _, pbcategory := range enter {
		if pbcategory.ParentId == root {
			var category types.Category
			_ = copier.Copy(&category, pbcategory)
			category.Children = categoryTree(category.CategoryID, enter)
			list = append(list, category)
		}
	}
	return list
}
