package private

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"server/app/product/cmd/api/internal/svc"
)

type ProductCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCreateLogic) ProductCreate() error {
	// todo: add your logic here and delete this line

	return nil
}
