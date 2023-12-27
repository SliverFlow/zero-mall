package private

import (
	"context"
	"server/app/cart/cmd/rpc/pb"

	"server/app/cart/cmd/api/internal/svc"
	"server/app/cart/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartDeleteLogic {
	return &CartDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartDelete
// Author [SliverFlow]
// @desc 用户端删除购物车
func (l *CartDeleteLogic) CartDelete(req *types.CartDeleteReq) (resp *types.Nil, err error) {
	_, err = l.svcCtx.CartRpc.CartDelete(l.ctx, &pb.CartDeleteReq{CartID: req.CartID})
	if err != nil {
		return nil, err
	}

	return
}
