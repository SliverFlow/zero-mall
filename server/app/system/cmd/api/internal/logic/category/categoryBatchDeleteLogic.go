package category

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryBatchDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryBatchDeleteLogic {
	return &CategoryBatchDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryBatchDeleteLogic) CategoryBatchDelete(req *types.CategoryBatchDelteReq) (resp *types.NilReply, err error) {
	var ids []string
	for _, kv := range req.KVS {
		if kv.ParentId == "0" {
			pbids, err := l.svcCtx.ProductRpc.CategoryFindChildrenID(l.ctx, &pb.CategoryIDReq{CategoryID: kv.CategoryID})
			if err != nil {
				return nil, err
			}
			ids = append(ids, pbids.IDS...)
		}
		ids = append(ids, kv.CategoryID)
	}

	_, err = l.svcCtx.ProductRpc.CategoryBatchDelete(l.ctx, &pb.CategoryIDSReq{IDS: ids})
	if err != nil {
		return nil, err
	}

	return
}
