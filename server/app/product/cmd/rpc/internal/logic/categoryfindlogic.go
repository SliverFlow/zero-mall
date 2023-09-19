package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryFindLogic {
	return &CategoryFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryFindLogic) CategoryFind(in *pb.CategoryIDReq) (*pb.Category, error) {
	enter, err := l.svcCtx.ProductModel.CategoryFind(l.ctx, in.GetCategoryID())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "不存在 id 为"+in.CategoryID+"的分类")
		}
		return nil, status.Errorf(100001, "查询失败")
	}

	var pbcategory pb.Category
	_ = copier.Copy(&pbcategory, enter)
	pbcategory.UpdatedAt = enter.UpdatedAt.Unix()
	pbcategory.CreatedAt = enter.CreatedAt.Unix()
	return &pbcategory, nil
}
