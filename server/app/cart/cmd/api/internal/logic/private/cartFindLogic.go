package private

import (
	"context"

	"server/app/cart/cmd/api/internal/svc"
	"server/app/cart/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartFindLogic {
	return &CartFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartFindLogic) CartFind(req *types.Nil) (resp *types.Nil, err error) {
	// todo: add your logic here and delete this line

	return
}
