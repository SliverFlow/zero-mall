package logic

import (
	"context"
	"google.golang.org/grpc/status"
	orderpb "server/app/order/cmd/rpc/pb"
	"server/app/order/model/request"
	productpb "server/app/product/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPageListForUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPageListForUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPageListForUserLogic {
	return &OrderPageListForUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPageListForUserLogic) OrderPageListForUser(in *pb.OrderPageListReq) (*pb.OrderPageListReply, error) {

	orderList, total, err := l.svcCtx.OrderModel.OrderListPageForUser(l.ctx, &request.PageListReq{
		BusinessID: in.GetBusinessID(),
		Page:       in.GetPage(),
		PageSize:   in.GetPageSize(),
		StartTime:  in.GetStartTime(),
		EndTime:    in.GetEndTime(),
		Status:     in.GetStatus(),
		UserID:     in.GetUserID(),
	})
	if err != nil {
		return nil, status.Errorf(100001, "订单列表查询失败")
	}

	var orderIDs []string
	var userIds []string
	var productIds []string
	for _, order := range orderList {
		orderIDs = append(orderIDs, order.OrderID)
		userIds = append(userIds, order.UserID)
	}

	orderItems, err := l.svcCtx.OrderModel.OrderItemFindByOrderIDs(l.ctx, orderIDs)
	if err != nil {
		return nil, status.Errorf(100001, "订单列表查询失败")
	}

	for i := range orderItems {
		productIds = append(productIds, orderItems[i].ProductID)
	}

	products, err := l.svcCtx.ProductRpc.ProductFindListByIDs(l.ctx, &productpb.ProductIDsReq{Ids: productIds})
	if err != nil {
		return nil, err
	}
	var businessIds []string
	for _, v := range products.GetList() {
		businessIds = append(businessIds, v.BusinessID)
	}

	businessList, err := l.svcCtx.UserRpc.BusinessFindListByIDs(l.ctx, &userpb.BusinessFindListReq{IDs: businessIds})
	if err != nil {
		return nil, err
	}

	users, err := l.svcCtx.UserRpc.UserFindListByIDs(l.ctx, &userpb.UserPageListByIDsReq{IDs: userIds})
	if err != nil {
		return nil, err
	}

	var orderListReply []*orderpb.OrderPageReply
	for i := range orderList {
		orderReply := &orderpb.OrderPageReply{
			PaymentType:  orderList[i].PaymentType,
			PaymentTime:  orderList[i].PaymentTime,
			SendTime:     orderList[i].SendTime,
			PayEndTime:   orderList[i].EndTime,
			PayCloseTime: orderList[i].CloseTime,
			OrderID:      orderList[i].OrderID,
			Payment:      orderList[i].Payment,
			OrderStatus:  orderList[i].Status,
			CreateTime:   orderList[i].CreatedAt.Unix(),
			EndTime:      orderList[i].EndTime,
		}
		for _, v := range orderItems {
			if v.OrderID == orderList[i].OrderID {
				orderReply.OrderItemID = v.OrderItemID
				orderReply.ProductImage = v.ProdImage
				orderReply.ProductNum = v.Quantity
				orderReply.ProductName = v.ProdName
				orderReply.CurrentunitPrice = v.CurrentunitPrice
				for _, vv := range products.GetList() {
					if vv.ProductID == v.ProductID {
						orderReply.Price = vv.Price
						for _, vvv := range businessList.GetList() {
							if vvv.BusinessID == vv.BusinessID {
								orderReply.BusinessName = vvv.Name
								orderReply.BusinessID = vvv.BusinessID
							}
						}
					}
				}
				userId := v.UserID
				for _, v := range users.GetList() {
					if v.UUID == userId {
						orderReply.Username = v.Nickname
					}
				}
				if orderReply.Username == "" {
					orderReply.Username = "未知"
				}
				if orderReply.BusinessName == "" {
					orderReply.BusinessName = "未知"
					orderReply.BusinessID = ""
				}
			}
		}
		orderReply.TotalPrice = orderReply.Price * float64(orderReply.ProductNum)
		orderListReply = append(orderListReply, orderReply)
	}

	return &orderpb.OrderPageListReply{
		Total: total,
		List:  orderListReply,
	}, nil
}
