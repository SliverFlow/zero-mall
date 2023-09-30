package product

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductUpdate
// Author [SliverFlow]
// @desc  商品更新信息
func (l *ProductUpdateLogic) ProductUpdate(req *types.ProductUpdateReq) (resp *types.NilReply, err error) {
	var pbproduct pb.ProductUpdateReq
	_ = copier.Copy(&pbproduct, req)
	bt, err := json.Marshal(req.Image)
	if err != nil {
		return nil, xerr.NewErrMsg("json序列化失败")
	}
	pbproduct.Image = string(bt)

	if req.Categories != nil && len(req.Categories) > 0 {
		var pbclist []*pb.Category
		for _, c := range req.Categories {
			var pbc pb.Category
			_ = copier.Copy(&pbc, c)
			pbclist = append(pbclist, &pbc)
		}
		pbproduct.Categories = pbclist
	}

	_, err = l.svcCtx.ProductRpc.ProductUpdate(l.ctx, &pbproduct)
	if err != nil {
		return nil, err
	}

	return
}
