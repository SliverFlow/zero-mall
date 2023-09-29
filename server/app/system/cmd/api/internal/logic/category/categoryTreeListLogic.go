package category

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryTreeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryTreeListLogic {
	return &CategoryTreeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryTreeList
// Author [SliverFlow]
// @desc  添加商品时的分类列表
func (l *CategoryTreeListLogic) CategoryTreeList(req *types.NilReq) (resp *types.CategoryListAllReply, err error) {
	reply, err := l.svcCtx.ProductRpc.CategoryListAll(l.ctx, &pb.ProductNil{})
	if err != nil {
		return nil, err
	}

	list := make([]types.Category, 0)
	if len(reply.List) != 0 {
		activeList := make([]*pb.Category, 0)
		for _, category := range reply.List {
			if category.Status == 1 {
				activeList = append(activeList, category)
			}
		}
		list = categoryTree("0", activeList)
	}

	return &types.CategoryListAllReply{List: list}, nil
}
