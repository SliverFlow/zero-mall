package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/app/product/model"
	"server/common"
	"strconv"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCreate
// Author [SliverFlow]
// @desc  创建商品
func (l *ProductCreateLogic) ProductCreate(in *pb.ProductCreateReq) (*pb.ProductNil, error) {
	var product model.Product
	_ = copier.Copy(&product, in)
	product.ProductID = strconv.FormatInt(common.SnowWorker.NextId(), 10)

	err := l.svcCtx.ProductModel.ProductCreate(l.ctx, &product)
	if err != nil {
		logx.Error("create product err", err.Error())
		return nil, status.Errorf(100001, "创建商品失败")
	}

	return &pb.ProductNil{}, nil
}
