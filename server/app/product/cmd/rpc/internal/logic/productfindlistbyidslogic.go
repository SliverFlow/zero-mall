package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFindListByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFindListByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFindListByIDsLogic {
	return &ProductFindListByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFindListByIDsLogic) ProductFindListByIDs(in *pb.ProductIDsReq) (*pb.ProductListReply, error) {

	products, err := l.svcCtx.ProductModel.ProductFindListByIDs(l.ctx, in.Ids)
	if err != nil {
		return nil, status.Errorf(100001, "商品列表查询失败")
	}

	var productListReply []*pb.ProductReply
	for i := range products {
		productReply := &pb.ProductReply{
			ProductID:  products[i].ProductID,
			BusinessID: products[i].BusinessID,
			Name:       products[i].Name,
			Subtitle:   products[i].Subtitle,
			Image:      products[i].Image,
			Detail:     products[i].Detail,
			Price:      products[i].Price,
			Stock:      products[i].Stock,
			Status:     products[i].Status,
			CreatedAt:  products[i].CreatedAt.Unix(),
			UpdatedAt:  products[i].UpdatedAt.Unix(),
		}
		productListReply = append(productListReply, productReply)
	}

	return &pb.ProductListReply{
		List: productListReply,
	}, nil
}
