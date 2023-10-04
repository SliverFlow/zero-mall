package public

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"

	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryList
// Author [SliverFlow]
// @desc 用户端分类接口
func (l *CategoryListLogic) CategoryList(req *types.Nil) (resp *types.CategoryListReply, err error) {
	pbreply, err := l.svcCtx.ProductRpc.CategoryListAll(l.ctx, &pb.ProductNil{})
	if err != nil {
		return nil, err
	}

	list := make([]types.Category, 0)
	if len(pbreply.List) != 0 {
		activeList := make([]*pb.Category, 0)
		for _, category := range pbreply.List {
			if category.Status == 1 {
				activeList = append(activeList, category)
			}
		}
		list = categoryTree("0", activeList)
	}

	return &types.CategoryListReply{List: list}, nil
}

func categoryTree(root string, enter []*pb.Category) (rep []types.Category) {
	list := make([]types.Category, 0)
	for _, pbcategory := range enter {
		if pbcategory.ParentId == root {
			var category types.Category
			_ = copier.Copy(&category, pbcategory)

			category.Children = categoryTree(pbcategory.CategoryID, enter)
			list = append(list, category)
		}
	}
	return list
}
