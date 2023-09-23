package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryChangeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryChangeStatusLogic {
	return &CategoryChangeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CategoryChangeStatus
// Author [SliverFlow]
// @desc 修改分类状态
func (l *CategoryChangeStatusLogic) CategoryChangeStatus(in *pb.CategoryChangeStatusReq) (*pb.ProductNil, error) {
	err := l.svcCtx.ProductModel.CategoryChangeStatus(l.ctx, in.GetCategoryID(), in.Status)
	if err != nil {
		return nil, status.Errorf(100001, "更新菜单状态失败")
	}

	return &pb.ProductNil{}, nil
}
