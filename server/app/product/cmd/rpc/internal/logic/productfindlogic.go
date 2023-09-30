package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFindLogic {
	return &ProductFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductFind
// Author [SliverFlow]
// @desc  查询商品信息
func (l *ProductFindLogic) ProductFind(in *pb.ProductFindReq) (*pb.ProductReply, error) {
	product, err := l.svcCtx.ProductModel.ProductFind(l.ctx, in.GetProductID())
	if err != nil {
		logx.Error("find product err", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "不存在编号为%s的商品", in.GetProductID())
		}
		return nil, status.Errorf(100001, "商品信息查询失败")
	}

	var productReply pb.ProductReply
	_ = copier.Copy(&productReply, product)
	productReply.CreatedAt = product.CreatedAt.Unix()
	productReply.UpdatedAt = product.UpdatedAt.Unix()

	if product.Categories != nil && len(product.Categories) > 0 {
		var pbcategoryList []*pb.Category
		for _, c := range product.Categories {
			var pbc pb.Category
			_ = copier.Copy(&pbc, c)
			pbcategoryList = append(pbcategoryList, &pbc)
		}
		productReply.Categories = pbcategoryList
	}

	return &productReply, nil
}
