package private

import (
	"context"
	"go.etcd.io/etcd/client/v3/concurrency"
	"google.golang.org/grpc/status"
	orderpb "server/app/order/cmd/rpc/pb"
	productpb "server/app/product/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"
	"strconv"

	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// OrderCreate
// Author [SliverFlow]
// @desc 用户下单
func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreateReq) (resp *types.OrderCreateReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("用户 uuid 获取失败")
	}

	// 查询用户信息
	pbuser, err := l.svcCtx.UserRpc.UserFindByUUID(l.ctx, &userpb.UUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}

	if pbuser.GetUser().Role != 1 {
		return nil, status.Errorf(100001, "当前用户不能购买商品，请注册普通用户进行购买")
	}

	pbproduct, err := l.svcCtx.ProductRpc.ProductFind(l.ctx, &productpb.ProductFindReq{ProductID: req.ProductID})
	if err != nil {
		return nil, err
	}

	if pbproduct.Stock-req.Quantity < 0 {
		return nil, xerr.NewErrMsg("当前商品库存不足")
	}

	totalS := strconv.FormatInt(req.Quantity, 10)
	total, _ := strconv.ParseFloat(totalS, 64)

	// 创建订单
	pbreply, err := l.svcCtx.OrderRpc.OrderCreate(l.ctx, &orderpb.OrderCreateReq{
		UserID:           uuid,
		ShoppingID:       "",
		ProductID:        pbproduct.ProductID,
		Postage:          0,
		ProdName:         pbproduct.Name,
		ProdImage:        pbproduct.Image,
		CurrentunitPrice: pbproduct.Price,
		Quantity:         req.Quantity,
		TotalPrice:       total * pbproduct.Price,
	})
	if err != nil {
		return nil, err
	}

	session, err := concurrency.NewSession(l.svcCtx.EtcdCli)
	locker := concurrency.NewLocker(session, req.ProductID)
	locker.Lock()
	defer locker.Unlock()

	// 扣减库存
	_, err = l.svcCtx.ProductRpc.ProductDeductionStock(l.ctx, &productpb.ProductDeductionStockReq{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	})
	if err != nil {
		// 删除创建的订单
		_, _ = l.svcCtx.OrderRpc.OrderDelete(l.ctx, &orderpb.OrderIdReq{OrderID: pbreply.OrderID})
		_, _ = l.svcCtx.OrderRpc.OrderItemDelete(l.ctx, &orderpb.OrderItemIdReq{OrderItemID: pbreply.OrderItemID})
		return nil, err
	}

	return &types.OrderCreateReply{
		OrderID:     pbreply.OrderID,
		OrderItemID: pbreply.GetOrderItemID(),
	}, nil
}
