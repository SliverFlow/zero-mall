package private

import (
	"context"
	"encoding/json"
	orderpb "server/app/cart/cmd/rpc/pb"
	productpb "server/app/product/cmd/rpc/pb"
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

	pbreply, err := l.svcCtx.CartRpc.CartList(l.ctx, &orderpb.CartListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		KeyWord:  req.KeyWord,
		UserID:   uuid,
	})
	if err != nil {
		return nil, err
	}

	var productIDs []string
	for _, v := range pbreply.List {
		productIDs = append(productIDs, v.ProductID)
	}

	productListReply, err := l.svcCtx.ProductRpc.ProductFindListByIDs(l.ctx, &productpb.ProductIDsReq{Ids: productIDs})
	if err != nil {
		return nil, err
	}

	var list []types.Cart
	for _, v := range pbreply.List {
		var c types.Cart
		c.UserID = v.UserID
		c.ProductID = v.ProductID
		c.CartID = v.CartID
		c.Quantity = v.Quantity
		c.ProductName = v.ProductName

		for _, vv := range productListReply.List {
			if v.ProductID == vv.ProductID {
				c.ProductPrice = vv.Price
				_ = json.Unmarshal([]byte(vv.Image), &c.ProductImage)
			}
		}
		c.Price = c.ProductPrice * float64(c.Quantity)
		list = append(list, c)
	}

	return &types.CartListReply{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    pbreply.GetTotal(),
		List:     list,
	}, nil
}
