package product

import (
	"context"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductDelete
// Author [SliverFlow]
// @desc  删除商品
func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.ProductRpc.ProductDelete(l.ctx, &pb.ProductDeleteReq{IDs: []string{req.ProductID}})
	if err != nil {
		return nil, err
	}

	return
}
