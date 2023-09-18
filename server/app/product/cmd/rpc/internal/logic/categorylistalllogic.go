package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListAllLogic {
	return &CategoryListAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryListAllLogic) CategoryListAll(in *pb.ProductNil) (*pb.CategoryListAllReply, error) {
	enter, err := l.svcCtx.ProductModel.CategoryListAll(l.ctx)
	if err != nil {
		return nil, status.Errorf(100001, "商品分类查询失败")
	}

	pblist := make([]*pb.Category, 0)
	i := len(*enter)
	if i != 0 {
		for _, category := range *enter {
			var pbcategory pb.Category
			_ = copier.Copy(&pbcategory, &category)
			pbcategory.CreatedAt = category.CreatedAt.Unix()
			pbcategory.UpdatedAt = category.UpdatedAt.Unix()
			pblist = append(pblist, &pbcategory)
		}
	}

	return &pb.CategoryListAllReply{
		List: pblist,
	}, nil
}
