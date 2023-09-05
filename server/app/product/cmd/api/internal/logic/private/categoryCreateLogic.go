package private

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"server/app/product/cmd/api/internal/svc"
)

type CategoryCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryCreateLogic {
	return &CategoryCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryCreateLogic) CategoryCreate() error {
	// todo: add your logic here and delete this line

	return nil
}
