package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderDelete
// Author [SliverFlow]
// @desc 删除订单
func (l *OrderDeleteLogic) OrderDelete(in *pb.OrderIdReq) (*pb.OrderNil, error) {
	err := l.svcCtx.OrderModel.OrderDelete(l.ctx, in.GetOrderID())
	if err != nil {
		return nil, status.Errorf(100001, "订单删除失败")
	}

	return &pb.OrderNil{}, nil
}
