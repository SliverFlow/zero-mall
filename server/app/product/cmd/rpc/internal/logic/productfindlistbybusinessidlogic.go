package logic

import (
	"context"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFindListByBusinessIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductFindListByBusinessIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFindListByBusinessIDLogic {
	return &ProductFindListByBusinessIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductFindListByBusinessIDLogic) ProductFindListByBusinessID(in *pb.ProductFindListByBusinessIDReq) (*pb.ProductListReply, error) {
	// todo: add your logic here and delete this line

	return &pb.ProductListReply{}, nil
}
