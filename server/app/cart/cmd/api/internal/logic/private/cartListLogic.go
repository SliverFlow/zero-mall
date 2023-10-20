package private

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/cart/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/cart/cmd/api/internal/svc"
	"server/app/cart/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartList
// Author [SliverFlow]
// @desc 用户端获取购物车列表
func (l *CartListLogic) CartList(req *types.CartListReq) (resp *types.CartListReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户信息失败")
	}

	pbreply, err := l.svcCtx.CartRpc.CartList(l.ctx, &pb.CartListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		KeyWord:  req.KeyWork,
		UserID:   uuid,
	})
	if err != nil {
		return nil, err
	}

	var list []types.Cart
	for _, pbcart := range pbreply.List {
		var tycart types.Cart
		_ = copier.Copy(&tycart, pbcart)
		list = append(list, tycart)
	}

	return &types.CartListReply{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    pbreply.GetTotal(),
		List:     list,
	}, nil
}
