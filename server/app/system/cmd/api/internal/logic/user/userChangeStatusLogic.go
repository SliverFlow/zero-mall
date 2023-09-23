package user

import (
	"context"
	"server/app/user/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChangeStatusLogic {
	return &UserChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChangeStatusLogic) UserChangeStatus(req *types.UserChangeStatusReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.UserRpc.UserChangeStatus(l.ctx, &pb.UserChangeStatusReq{
		ID:     req.ID,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return
}
