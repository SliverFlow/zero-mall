package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderFindLogic {
	return &OrderFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderFind
// Author [SliverFlow]
// @desc 订单查询
func (l *OrderFindLogic) OrderFind(in *pb.OrderFindReq) (*pb.OrderReply, error) {
	item, err := l.svcCtx.OrderModel.OrderItemFindByOrderID(l.ctx, in.OrderID)
	if err != nil {
		return nil, status.Errorf(100001, "查询失败")
	}

	order, err := l.svcCtx.OrderModel.OrderFind(l.ctx, in.GetOrderID())
	if err != nil {
		return nil, status.Errorf(100001, "查询失败")
	}

	return &pb.OrderReply{
		OrderID:     order.OrderID,
		UserID:      order.UserID,
		ShoppingID:  order.ShoppingID,
		Payment:     order.Payment,
		PaymentType: order.PaymentType,
		Postage:     order.Postage,
		Status:      order.Status,
		PaymentTime: order.PaymentTime,
		SendTime:    order.SendTime,
		EndTime:     order.EndTime,
		CloseTime:   order.CloseTime,
		OrderItem: &pb.OrderItemReply{
			OrderItemID:      item.OrderItemID,
			OrderID:          item.OrderID,
			UserID:           item.UserID,
			ProductID:        item.ProductID,
			ProductName:      item.ProductID,
			ProductImage:     item.ProductID,
			CurrentunitPrice: item.CurrentunitPrice,
			Quantity:         item.Quantity,
			TotalPrice:       item.TotalPrice,
		},
	}, nil
}
