package category

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryFindLogic {
	return &CategoryFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryFind
// Author [SliverFlow]
// @desc  分类查询
func (l *CategoryFindLogic) CategoryFind(req *types.CategoryIDReq) (resp *types.Category, err error) {
	pbreply, err := l.svcCtx.ProductRpc.CategoryFind(l.ctx, &pb.CategoryIDReq{CategoryID: req.CategoryID})
	if err != nil {
		return nil, err
	}
	var category types.Category
	_ = copier.Copy(&category, pbreply)
	return &category, nil
}
