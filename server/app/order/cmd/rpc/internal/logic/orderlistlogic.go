package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/common"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderList
// Author [SliverFlow]
// @desc 订单列表查询
func (l *OrderListLogic) OrderList(in *pb.OrderListReq) (*pb.OrderListReply, error) {
	enter, total, err := l.svcCtx.OrderModel.OrderList(l.ctx, &common.PageInfo{
		Page:     int(in.GetPage()),
		PageSize: int(in.GetPageSize()),
		KeyWord:  in.GetKeyWord(),
	}, in.GetUserID())
	if err != nil {
		return nil, status.Errorf(100001, "订单列表查询失败")
	}

	var pblist []*pb.OrderReply

	for _, order := range *enter {
		item, _ := l.svcCtx.OrderModel.OrderItemFindByOrderID(l.ctx, order.OrderID)
		pblist = append(pblist, &pb.OrderReply{
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
				ProductName:      item.ProdName,
				ProductImage:     item.ProdImage,
				CurrentunitPrice: item.CurrentunitPrice,
				Quantity:         item.Quantity,
				TotalPrice:       item.TotalPrice,
			},
		})
	}

	return &pb.OrderListReply{
		List:  pblist,
		Total: total,
	}, nil
}
