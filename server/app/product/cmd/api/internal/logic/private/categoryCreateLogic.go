package private

import (
	"context"

	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CategoryCreateLogic) CategoryCreate(req *types.Nil) (resp *types.Nil, err error) {
	// todo: add your logic here and delete this line

	return
}
