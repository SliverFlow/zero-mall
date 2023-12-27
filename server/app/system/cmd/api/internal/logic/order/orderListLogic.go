package order

import (
	"context"
	"encoding/json"
	"server/app/order/cmd/rpc/pb"
	"server/common"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

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

func (l *OrderListLogic) OrderList(req *types.OrderListReq) (resp *types.OrderListReply, err error) {
	var startTime int64
	if req.StartTime != "" {
		t, _ := common.FormatTime(req.StartTime)
		startTime = t.Unix()
	}

	var endTime int64
	if req.EndTime != "" {
		t, _ := common.FormatTime(req.EndTime)
		endTime = t.Unix()
	}

	orderPageList, err := l.svcCtx.OrderRpc.OrderPageList(l.ctx, &pb.OrderPageListReq{
		BusinessID: req.BusinessId,
		Page:       req.Page,
		PageSize:   req.PageSize,
		StartTime:  startTime,
		EndTime:    endTime,
		Status:     req.Status,
	})
	if err != nil {
		return
	}
	var list []types.OrderReply
	for _, v := range orderPageList.List {
		orderReply := types.OrderReply{
			OrderID:      v.OrderID,
			Username:     v.Username,
			BusinessName: v.BusinessName,
			ProductName:  v.ProductName,
			ProductNum:   v.ProductNum,
			Payment:      v.Payment,
			OrderStatus:  v.OrderStatus,
			CreateTime:   v.CreateTime,
			EndTime:      v.EndTime,
			Quantity:     v.ProductNum,
		}
		_ = json.Unmarshal([]byte(v.ProductImage), &orderReply.ProductImage)
		list = append(list, orderReply)
	}

	return &types.OrderListReply{
		Total:    orderPageList.Total,
		List:     list,
		PageSize: req.PageSize,
		Page:     req.Page,
	}, nil
}
