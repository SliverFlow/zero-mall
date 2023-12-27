package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"
	"server/app/order/model"
	"server/common"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderCreate
// Author [SliverFlow]
// @desc 创建订单
func (l *OrderCreateLogic) OrderCreate(in *pb.OrderCreateReq) (*pb.OrderIDReply, error) {

	orderId := strconv.FormatInt(common.SnowWorker.NextId(), 10)

	err := l.svcCtx.OrderModel.OrderCreate(l.ctx, &model.Order{
		OrderID:    orderId,
		UserID:     in.GetUserID(),
		ShoppingID: in.GetShoppingID(),
		Postage:    in.GetPostage(),
		BusinessID: in.GetBusinessID(),
		Status:     10,
	})
	if err != nil {
		return nil, status.Errorf(100001, "订单创建失败")
	}

	orderItemId := strconv.FormatInt(common.SnowWorker.NextId(), 10)
	err = l.svcCtx.OrderModel.OrderItemCreate(l.ctx, &model.OrderItem{
		OrderItemID:      strconv.FormatInt(common.SnowWorker.NextId(), 10),
		OrderID:          orderId,
		UserID:           in.GetUserID(),
		ProductID:        in.GetProductID(),
		ProdName:         in.GetProdName(),
		ProdImage:        in.GetProdImage(),
		CurrentunitPrice: in.GetCurrentunitPrice(),
		Quantity:         in.GetQuantity(),
		TotalPrice:       in.GetTotalPrice(),
	})
	if err != nil {
		_ = l.svcCtx.OrderModel.OrderDeleteByID(l.ctx, orderId)
		return nil, status.Errorf(100001, "订单创建失败")
	}

	return &pb.OrderIDReply{OrderID: orderId, OrderItemID: orderItemId}, nil
}
