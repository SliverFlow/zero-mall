package business

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessUpdateLogic {
	return &BusinessUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BusinessUpdate
// Author [SliverFlow]
// @desc 修改商户信息
func (l *BusinessUpdateLogic) BusinessUpdate(req *types.BusinessUpdateReq) (resp *types.NilReply, err error) {
	var pbbusiness pb.Business
	_ = copier.Copy(&pbbusiness, req)

	bt, err := json.Marshal(req.Image)
	if err != nil {
		return nil, xerr.NewErrMsg("json 序列化失败")
	}

	pbbusiness.Image = string(bt)

	_, err = l.svcCtx.UserRpc.BusinessUpdate(l.ctx, &pbbusiness)
	if err != nil {
		return nil, err
	}

	return
}
