package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryFindChildrenIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryFindChildrenIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryFindChildrenIDLogic {
	return &CategoryFindChildrenIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryFindChildrenIDLogic) CategoryFindChildrenID(in *pb.CategoryIDReq) (*pb.CategoryIDSReply, error) {
	enter, err := l.svcCtx.ProductModel.CategoryFindChildrenID(l.ctx, in.CategoryID)
	if err != nil {
		return nil, status.Errorf(10001, "插叙子分类ID失败")
	}

	return &pb.CategoryIDSReply{
		IDS: *enter,
	}, nil
}
