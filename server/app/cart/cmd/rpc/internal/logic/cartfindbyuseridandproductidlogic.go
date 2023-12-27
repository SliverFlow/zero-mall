package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartFindByUserIDAndProductIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartFindByUserIDAndProductIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartFindByUserIDAndProductIDLogic {
	return &CartFindByUserIDAndProductIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartFindByUserIDAndProductIDLogic) CartFindByUserIDAndProductID(in *pb.CartFindByUserIDAndProductIDReq) (*pb.CartInfo, error) {
	cart, err := l.svcCtx.CartModel.CartFindByUserIDAndProductID(l.ctx, in.GetUserID(), in.GetProductID())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "当前商品不存在购物车中")
		}
		return nil, status.Errorf(100001, "查询购物车失败")
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
