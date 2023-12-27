package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/order/cmd/rpc/internal/svc"
	"server/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteByIDLogic {
	return &OrderDeleteByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDeleteByIDLogic) OrderDeleteByID(in *pb.OrderDeleteByIDReq) (*pb.OrderNil, error) {
	err := l.svcCtx.OrderModel.OrderDeleteByID(l.ctx, in.GetID())
	if err != nil {
		return nil, status.Error(100001, "订单删除失败")
	}

	return &pb.OrderNil{}, nil
}
