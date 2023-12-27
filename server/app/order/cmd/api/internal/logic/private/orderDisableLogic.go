package private

import (
	"context"
	"server/app/order/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDisableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDisableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDisableLogic {
	return &OrderDisableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDisableLogic) OrderDisable(req *types.OrderDisableReq) (resp *types.Nil, err error) {
	orderpb, err := l.svcCtx.OrderRpc.OrderFind(l.ctx, &pb.OrderFindReq{OrderID: req.OrderID})
	if err != nil {
		return nil, err
	}

	if orderpb.Status != 10 {
		return nil, xerr.NewErrMsg("订单状态不是未付款，无法取消")
	}

	_, err = l.svcCtx.OrderRpc.OrderDisable(l.ctx, &pb.OrderDisableReq{OrderID: req.OrderID})
	if err != nil {
		return nil, err
	}

	return
}
