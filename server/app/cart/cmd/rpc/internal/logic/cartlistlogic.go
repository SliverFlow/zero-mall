package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/common"

	"server/app/cart/cmd/rpc/internal/svc"
	"server/app/cart/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CartList
// Author [SliverFlow]
// @desc 用户购物车列表
func (l *CartListLogic) CartList(in *pb.CartListReq) (*pb.CartListReply, error) {
	list, total, err := l.svcCtx.CartModel.CartListByUserID(l.ctx, in.GetUserID(), &common.PageInfo{
		Page:     int(in.GetPage()),
		PageSize: int(in.GetPageSize()),
		KeyWord:  in.GetKeyWord(),
	})
	if err != nil {
		return nil, status.Errorf(100001, "购物车列表获取失败")
	}

	var pbcartList []*pb.CartInfo
	for _, cart := range *list {
		var pbcart pb.CartInfo
		_ = copier.Copy(&pbcart, &cart)
		pbcartList = append(pbcartList, &pbcart)
	}

	return &pb.CartListReply{
		Total: total,
		List:  pbcartList,
	}, nil
}
