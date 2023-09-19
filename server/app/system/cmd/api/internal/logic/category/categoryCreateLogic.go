package category

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryCreateLogic {
	return &CategoryCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryCreateLogic) CategoryCreate(req *types.CategoryCreateReq) (resp *types.NilReply, err error) {

	_, err = l.svcCtx.ProductRpc.CategoryCreate(l.ctx, &pb.CategoryCreateReq{
		ParentId: req.ParentId,
		Name:     req.Name,
		Status:   req.Status,
		Type:     req.Type,
		Sorted:   req.Sorted,
	})
	if err != nil {
		return nil, err
	}

	return
}
