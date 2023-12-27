package private

import (
	"context"
	"encoding/json"
	"server/app/order/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderList
// Author [SliverFlow]
// @desc 用户查询订单列表
func (l *OrderListLogic) OrderList(req *types.OrderListReq) (resp *types.OrderListReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("用户 uuid 获取失败")
	}

	pbreply, err := l.svcCtx.OrderRpc.OrderPageList(l.ctx, &pb.OrderPageListReq{
		UserID:   uuid,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list []types.OrderListInfo
	for _, v := range pbreply.List {
		o := types.OrderListInfo{
			OrderID:          v.OrderID,
			OrderItemID:      v.OrderItemID,
			ProductName:      v.ProductName,
			BusinessName:     v.BusinessName,
			BusinessID:       v.BusinessID,
			Price:            v.Price,
			OrderStatus:      v.OrderStatus,
			Quantity:         v.ProductNum,
			Payment:          v.Payment,
			PaymentType:      v.PaymentType,
			PaymentTime:      v.PaymentTime,
			SendTime:         v.SendTime,
			PayEndTime:       v.PayEndTime,
			PayCloseTime:     v.PayCloseTime,
			CurrentunitPrice: v.CurrentunitPrice,
			TotalPrice:       v.TotalPrice,
			CreateTime:       v.CreateTime,
			EndTime:          v.EndTime,
		}
		_ = json.Unmarshal([]byte(v.ProductImage), &o.ProductImage)
		list = append(list, o)
	}

	return &types.OrderListReply{
		List:     list,
		Total:    pbreply.GetTotal(),
		PageSize: req.PageSize,
		Page:     req.Page,
	}, nil
}
