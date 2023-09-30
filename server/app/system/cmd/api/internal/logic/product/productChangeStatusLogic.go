package product

import (
	"context"
	productpb "server/app/product/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductChangeStatusLogic {
	return &ProductChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductChangeStatus
// Author [SliverFlow]
// @desc  创建商品
func (l *ProductChangeStatusLogic) ProductChangeStatus(req *types.ProductChangeStatusReq) (resp *types.NilReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户 uuid 失败")
	}

	if req.Status == 1 { // 将商品上架
		pbbusiness, err := l.svcCtx.UserRpc.BusinessFindByUUID(l.ctx, &userpb.BusinessUUIDReq{UUID: uuid})
		if err != nil {
			return nil, err
		}
		if pbbusiness.Status == 0 {
			return nil, xerr.NewErrMsg("当前商户处于非营业状态，商品不能上架")
		}
		if pbbusiness.Status == 2 {
			return nil, xerr.NewErrMsg("当前商户处于禁用状态，商品不能上架")
		}
	}

	_, err = l.svcCtx.ProductRpc.ProductChangeStatus(l.ctx, &productpb.ProductChangeStatusReq{
		ProductID: req.ProductID,
		Status:    req.Status,
	})
	if err != nil {
		return nil, err
	}

	return
}
