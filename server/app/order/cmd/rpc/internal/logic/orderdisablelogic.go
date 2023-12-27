package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDisableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDisableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDisableLogic {
	return &OrderDisableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDisableLogic) OrderDisable(in *pb.OrderDisableReq) (*pb.OrderNil, error) {
	err := l.svcCtx.OrderModel.OrderDisable(l.ctx, in.GetOrderID())
	if err != nil {
		return nil, status.Error(10001, "订单取消失败")
	}

	return &pb.OrderNil{}, nil
}
