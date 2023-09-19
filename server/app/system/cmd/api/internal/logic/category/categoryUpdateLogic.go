package category

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryUpdateLogic) CategoryUpdate(req *types.CategoryUpdateReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.ProductRpc.CategoryUpdate(l.ctx, &pb.CategoryUpdateReq{
		CategoryID: req.CategoryID,
		ParentId:   req.ParentId,
		Name:       req.Name,
		Status:     req.Status,
		Sorted:     req.Sorted,
		Type:       req.Type,
	})
	if err != nil {
		return nil, err
	}

	return
}
