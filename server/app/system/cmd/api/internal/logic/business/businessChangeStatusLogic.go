package business

import (
	"context"
	"server/app/user/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessChangeStatusLogic {
	return &BusinessChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BusinessChangeStatus
// Author [SliverFlow]
// @desc 修改商户状态
func (l *BusinessChangeStatusLogic) BusinessChangeStatus(req *types.BusinessChangeStatusReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.UserRpc.BusinessChangeStatus(l.ctx, &pb.BusinessChangeStatusReq{
		BusinessID: req.BusinessID,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}

	return
}
