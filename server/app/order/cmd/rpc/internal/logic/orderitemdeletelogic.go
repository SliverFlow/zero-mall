package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemDeleteLogic {
	return &OrderItemDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderItemDelete
// Author [SliverFlow]
// @desc 删除订单详细
func (l *OrderItemDeleteLogic) OrderItemDelete(in *pb.OrderItemIdReq) (*pb.OrderNil, error) {

	err := l.svcCtx.OrderModel.OrderItemDeleteByID(l.ctx, in.GetOrderItemID())
	if err != nil {
		return nil, status.Errorf(100001, "订单详细删除失败")
	}

	return &pb.OrderNil{}, nil
}
