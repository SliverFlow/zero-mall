package category

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteLogic {
	return &CategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryDelete
// Author [SliverFlow]
// @desc 单个删除分类
func (l *CategoryDeleteLogic) CategoryDelete(req *types.CategoryIDReq) (resp *types.NilReply, err error) {
	pbreply, err := l.svcCtx.ProductRpc.CategoryFind(l.ctx, &pb.CategoryIDReq{CategoryID: req.CategoryID})
	if err != nil {
		return nil, err
	}

	// 主分类 删除及相关所有分类
	if pbreply.GetParentId() == "0" {
		res, _ := l.svcCtx.ProductRpc.CategoryFindChildrenID(l.ctx, &pb.CategoryIDReq{CategoryID: req.CategoryID})
		var ids []string
		ids = append(ids, res.IDS...)
		ids = append(ids, req.CategoryID)
		_, err = l.svcCtx.ProductRpc.CategoryBatchDelete(l.ctx, &pb.CategoryIDSReq{IDS: ids})
		if err != nil {
			return nil, err
		}
		return
	}

	// 非主分类 单个删除
	_, err = l.svcCtx.ProductRpc.CategoryDelete(l.ctx, &pb.CategoryIDReq{CategoryID: req.CategoryID})
	if err != nil {
		return nil, err
	}

	return
}
