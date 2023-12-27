package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartFindLogic {
	return &CartFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartFindLogic) CartFind(in *pb.CartFindReq) (*pb.CartInfo, error) {
	cart, err := l.svcCtx.CartModel.CartFind(l.ctx, in.GetCartID())
	if err != nil {
		return nil, status.Errorf(100001, "查询失败")
	}

	return &pb.CartInfo{
		UserID:      cart.UserID,
		ProductID:   cart.ProductID,
		Quantity:    cart.Quantity,
		Checked:     cart.Checked,
		CartID:      cart.CartID,
		ProductName: cart.ProductName,
	}, nil
}
