package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/common"

	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *pb.ProductListReq) (*pb.ProductListReply, error) {
	list, total, err := l.svcCtx.ProductModel.ProductList(l.ctx, &common.PageInfo{
		Page:     int(in.GetPage()),
		PageSize: int(in.GetPageSize()),
		KeyWord:  in.GetKeyWord(),
	})
	if err != nil {
		logx.Error("productList find error", err.Error())
		return nil, status.Errorf(100001, "商品列表查询失败")
	}

	var pblist []*pb.ProductReply
	for _, v := range *list {
		var pbv pb.ProductReply

		_ = copier.Copy(&pbv, v)
		pbv.UpdatedAt = v.UpdatedAt.Unix()
		pbv.CreatedAt = v.CreatedAt.Unix()
		if v.Categories != nil { // 组装分类
			var pbc []*pb.Category
			for _, c := range v.Categories {
				var pbcategory pb.Category
				_ = copier.Copy(&pbcategory, c)
				pbc = append(pbc, &pbcategory)
			}
			pbv.Categories = pbc
		}
		pblist = append(pblist, &pbv)
	}

	return &pb.ProductListReply{
		Page:     in.GetPage(),
		PageSize: in.GetPageSize(),
		Total:    total,
		List:     pblist,
	}, nil
}
