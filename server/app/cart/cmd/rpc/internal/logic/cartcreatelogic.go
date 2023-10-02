package logic

import (
	"context"

	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartCreateLogic {
	return &CartCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartCreateLogic) CartCreate(in *pb.CartNil) (*pb.CartNil, error) {
	// todo: add your logic here and delete this line

	return &pb.CartNil{}, nil
}
