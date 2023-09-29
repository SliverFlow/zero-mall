package category

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryChangeStatusLogic {
	return &CategoryChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryChangeStatus
// Author [SliverFlow]
// @desc  分类修改状态
func (l *CategoryChangeStatusLogic) CategoryChangeStatus(req *types.CategoryChangeStatusReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.ProductRpc.CategoryChangeStatus(l.ctx, &pb.CategoryChangeStatusReq{
		CategoryID: req.CategoryID,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
