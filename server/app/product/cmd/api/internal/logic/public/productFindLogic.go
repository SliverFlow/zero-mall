package public

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/product/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"

	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductFindLogic {
	return &ProductFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductFind
// Author [SliverFlow]
// @desc 用户端商品详细
func (l *ProductFindLogic) ProductFind(req *types.ProductFindReq) (resp *types.ProductFindReqly, err error) {
	pbreply, err := l.svcCtx.ProductRpc.ProductFind(l.ctx, &pb.ProductFindReq{ProductID: req.ProductID})
	if err != nil {
		return nil, err
	}

	var typroduct types.ProductFindReqly
	_ = copier.Copy(&typroduct, pbreply)
	_ = json.Unmarshal([]byte(pbreply.Image), &typroduct.Image)

	pbnusiness, err := l.svcCtx.UserRpc.BusinessFind(l.ctx, &userpb.BusinessIDReq{BusinessID: pbreply.BusinessID})
	if err != nil {
		return nil, err
	}
	typroduct.BudinessName = pbnusiness.Name

	return &typroduct, nil
}
