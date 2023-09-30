package product

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFindLogic {
	return &ProductFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductFind
// Author [SliverFlow]
// @desc  创建商品
func (l *ProductFindLogic) ProductFind(req *types.ProductFindReq) (resp *types.Product, err error) {
	pbreply, err := l.svcCtx.ProductRpc.ProductFind(l.ctx, &pb.ProductFindReq{ProductID: req.ProductID})
	if err != nil {
		return nil, err
	}
	var typroduct types.Product
	_ = copier.Copy(&typroduct, pbreply)
	_ = json.Unmarshal([]byte(pbreply.Image), &typroduct.Image)

	if pbreply.Categories != nil && len(pbreply.Categories) > 0 {
		var tycategoryList []types.Category
		for _, c := range pbreply.GetCategories() {
			var tyc types.Category
			_ = copier.Copy(&tyc, c)
			tyc.Children = make([]types.Category, 0)
			tycategoryList = append(tycategoryList, tyc)
		}
		typroduct.Categories = tycategoryList
	}

	return &typroduct, nil
}
