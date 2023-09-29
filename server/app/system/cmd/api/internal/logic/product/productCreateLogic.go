package product

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	productpb "server/app/product/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCreate
// Author [SliverFlow]
// @desc  创建商品
func (l *ProductCreateLogic) ProductCreate(req *types.ProductCreateReq) (resp *types.NilReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户 uuid 失败")
	}

	pbbusiness, err := l.svcCtx.UserRpc.BusinessFindByUUID(l.ctx, &userpb.BusinessUUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}

	var pbproduct productpb.ProductCreateReq
	_ = copier.Copy(&pbproduct, req)
	pbproduct.BusinessID = pbbusiness.BusinessID
	bt, err := json.Marshal(req.Image)
	if err != nil {
		return nil, xerr.NewErrMsg("json 序列化错误")
	}

	pbproduct.Image = string(bt)

	if _, err = l.svcCtx.ProductRpc.ProductCreate(l.ctx, &pbproduct); err != nil {
		return nil, err
	}

	return
}
