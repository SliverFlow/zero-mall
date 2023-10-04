package public

import (
	"context"

	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.Nil) (resp *types.CategoryListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
