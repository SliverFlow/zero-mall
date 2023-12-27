package private

import (
	"context"
	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"
	orderpb "server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderDelete
// Author [SliverFlow]
// @desc 用户查询订单列表
func (l *OrderDeleteLogic) OrderDelete(req *types.OrderDeleteReq) (resp *types.Nil, err error) {

	_, err = l.svcCtx.OrderRpc.OrderDelete(l.ctx, &orderpb.OrderDeleteReq{IDs: req.IDs})
	if err != nil {
		return nil, err
	}

	return
}
