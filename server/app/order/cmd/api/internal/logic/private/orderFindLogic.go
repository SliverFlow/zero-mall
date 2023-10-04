package private

import (
	"context"

	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderFindLogic {
	return &OrderFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderFindLogic) OrderFind(req *types.Nil) (resp *types.Nil, err error) {
	// todo: add your logic here and delete this line

	return
}
