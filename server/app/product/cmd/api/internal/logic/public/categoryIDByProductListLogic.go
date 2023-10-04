package public

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"

	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryIDByProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryIDByProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryIDByProductListLogic {
	return &CategoryIDByProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategoryIDByProductList
// Author [SliverFlow]
// @desc 用户端按照分类 ID 查询商品列表
func (l *CategoryIDByProductListLogic) CategoryIDByProductList(req *types.CategoryIDByProductListReq) (resp *types.CategoryIDByProductListReply, err error) {
	pbreply, err := l.svcCtx.ProductRpc.CategoryIDByProductList(l.ctx, &pb.CategoryIDByProductListReq{
		CategoryID: req.CategoryID,
		KeyWord:    req.KeyWord,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var tycategory types.CategoryIDByProductListReply
	_ = copier.Copy(&tycategory, pbreply)

	tycategory.ProductList = make([]types.Product, 0)
	for _, product := range pbreply.Products {
		var typroduct types.Product
		_ = copier.Copy(&typroduct, product)
		_ = json.Unmarshal([]byte(product.Image), &typroduct.Image)
		tycategory.ProductList = append(tycategory.ProductList, typroduct)
	}

	return &tycategory, nil
}
