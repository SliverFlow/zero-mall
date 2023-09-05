package public

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"server/app/product/cmd/api/internal/svc"
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

func (l *ProductListLogic) ProductList() error {
	// todo: add your logic here and delete this line

	return nil
}
