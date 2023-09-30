package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/app/product/model"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductUpdate
// Author [SliverFlow]
// @desc 下架某商家的所有商品
func (l *ProductUpdateLogic) ProductUpdate(in *pb.ProductUpdateReq) (*pb.ProductNil, error) {
	var product model.Product
	_ = copier.Copy(&product, in)
	if in.GetCategories() != nil && len(in.GetCategories()) > 0 {
		var clist []model.Category
		for _, v := range in.GetCategories() {
			var c model.Category
			_ = copier.Copy(&c, v)
			clist = append(clist, c)
		}
		product.Categories = clist
	}
	err := l.svcCtx.ProductModel.ProductUpdate(l.ctx, &product)
	if err != nil {
		logx.Error("update product error", err.Error())
		return nil, status.Errorf(100001, "更新商品信息失败")
	}

	return &pb.ProductNil{}, nil
}
