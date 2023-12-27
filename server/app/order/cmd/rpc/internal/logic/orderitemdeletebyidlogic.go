package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemDeleteByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemDeleteByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemDeleteByIDLogic {
	return &OrderItemDeleteByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemDeleteByIDLogic) OrderItemDeleteByID(in *pb.OrderItemDeleteByIDReq) (*pb.OrderNil, error) {
	err := l.svcCtx.OrderModel.OrderItemDeleteByID(l.ctx, in.GetID())
	if err != nil {
		return nil, status.Error(100001, "订单详细删除失败")
	}

	return &pb.OrderNil{}, nil
}
