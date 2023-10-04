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

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductList
// Author [SliverFlow]
// @desc   商品分页列表
func (l *ProductListLogic) ProductList(req *types.ProductListReq) (resp *types.ProductListReply, err error) {
	pbreply, err := l.svcCtx.ProductRpc.ProductList(l.ctx, &pb.ProductListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		KeyWord:  req.KeyWord,
	})
	if err != nil {
		return nil, err
	}

	var productList []types.Product
	for _, v := range pbreply.List {
		var p types.Product
		_ = copier.Copy(&p, v)
		_ = json.Unmarshal([]byte(v.Image), &p.Image)
		if v.Categories != nil { // 组装分类
			var list []types.Category
			for _, category := range v.Categories { // 只取子分类
				if category.ParentId != "0" {
					var c types.Category
					_ = copier.Copy(&c, category)
					list = append(list, c)
				}
			}
			p.Categories = list
		}
		productList = append(productList, p)
	}

	return &types.ProductListReply{
		Page:     pbreply.Page,
		PageSize: pbreply.PageSize,
		Total:    pbreply.Total,
		List:     productList,
	}, nil
}
