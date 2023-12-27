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
func (l *OrderDeleteLogic) OrderDelete(in *pb.OrderDeleteReq) (*pb.OrderNil, error) {

	orders, err := l.svcCtx.OrderModel.OrderFindListByIDs(l.ctx, in.GetIDs())
	if err != nil {
		return nil, status.Errorf(100001, "订单删除失败")
	}

	for _, order := range orders {
		if order.Status != 70 && order.Status != 80 {
			return nil, status.Errorf(100001, "选中了不可删除的订单")
		}
	}

	// 删除订单表
	err = l.svcCtx.OrderModel.OrderDelete(l.ctx, in.GetIDs())
	if err != nil {
		return nil, status.Errorf(100001, "订单删除失败")
	}

	// 删除子表
	_ = l.svcCtx.OrderModel.OrderItemDelete(l.ctx, in.GetIDs())

	return &pb.OrderNil{}, nil
}
