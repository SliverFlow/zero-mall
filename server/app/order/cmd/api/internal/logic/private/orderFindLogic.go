package private

import (
	"context"
	"encoding/json"
	orderpb "server/app/order/cmd/rpc/pb"
	productpb "server/app/product/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"

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

// OrderFind
// Author [SliverFlow]
// @desc 用户查询订单详细
func (l *OrderFindLogic) OrderFind(req *types.OrderFindReq) (resp *types.OrderFindReply, err error) {
	pborder, err := l.svcCtx.OrderRpc.OrderFind(l.ctx, &orderpb.OrderFindReq{OrderID: req.OrderID})
	if err != nil {
		return nil, err
	}

	pbproduct, err := l.svcCtx.ProductRpc.ProductFind(l.ctx, &productpb.ProductFindReq{ProductID: pborder.OrderItem.ProductID})
	if err != nil {
		return nil, err
	}

	pbbusiness, err := l.svcCtx.UserRpc.BusinessFind(l.ctx, &userpb.BusinessIDReq{BusinessID: pbproduct.BusinessID})

	reply := &types.OrderFindReply{
		OrderID:          pborder.OrderID,
		OrderItemID:      pborder.OrderItem.OrderItemID,
		ProductName:      pborder.OrderItem.ProductName,
		BusinessName:     pbbusiness.Name,
		Price:            pborder.Payment,
		OrderStatus:      pborder.Status,
		Quantity:         pborder.OrderItem.Quantity,
		Payment:          pborder.Payment,
		PaymentType:      pborder.PaymentType,
		PaymentTime:      pborder.PaymentTime,
		SendTime:         pborder.SendTime,
		EndTime:          pborder.EndTime,
		PayCloseTime:     pborder.CloseTime,
		CurrentunitPrice: pborder.OrderItem.CurrentunitPrice,
		TotalPrice:       pborder.OrderItem.TotalPrice,
	}
	_ = json.Unmarshal([]byte(pborder.OrderItem.ProductImage), &reply.ProductImage)
	return reply, nil
}
