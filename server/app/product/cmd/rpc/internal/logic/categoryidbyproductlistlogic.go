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

type CategoryIDByProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryIDByProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryIDByProductListLogic {
	return &CategoryIDByProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryIDByProductListLogic) CategoryIDByProductList(in *pb.CategoryIDByProductListReq) (*pb.CategoryIDByProductListReply, error) {
	enter, total, err := l.svcCtx.ProductModel.CategoryIDByProductList(l.ctx, in.GetCategoryID(), &common.PageInfo{
		Page:     int(in.GetPage()),
		PageSize: int(in.GetPageSize()),
		KeyWord:  in.GetKeyWord(),
	})
	if err != nil {
		return nil, status.Errorf(100001, "查询失败")
	}

	var reply pb.CategoryIDByProductListReply
	_ = copier.Copy(&reply, enter)
	reply.Total = total
	reply.Page = in.GetPage()
	reply.PageSize = in.GetPageSize()

	//for _, product := range enter.Products {
	//	var pbproduct pb.ProductReply
	//	_ = copier.Copy(&pbproduct, product)
	//	reply.Products = append(reply.Products, &pbproduct)
	//}

	return &reply, nil
}
