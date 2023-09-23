package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"
	"server/app/product/model"
	"server/common"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryCreateLogic {
	return &CategoryCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CategoryCreate
// Author [SliverFlow]
// @desc 创建分类
func (l *CategoryCreateLogic) CategoryCreate(in *pb.CategoryCreateReq) (*pb.ProductNil, error) {
	id := strconv.FormatInt(common.SnowWorker.NextId(), 10)
	err := l.svcCtx.ProductModel.CategoryCreate(l.ctx, &model.Category{
		CategoryID: id,
		ParentId:   in.GetParentId(),
		Name:       in.GetName(),
		Status:     in.GetStatus(),
		Type:       in.GetType(),
		BusinessID: in.GetBusinessID(),
		Sorted:     in.GetSorted(),
	})
	if err != nil {
		return nil, status.Errorf(100001, "创建商品分类失败")
	}

	return &pb.ProductNil{}, nil
}
