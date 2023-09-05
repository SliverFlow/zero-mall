package public

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"server/app/product/cmd/api/internal/svc"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList() error {
	// todo: add your logic here and delete this line

	return nil
}
