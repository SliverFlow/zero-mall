package user

import (
	"context"
	"server/app/user/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserDelete
// Author [SliverFlow]
// @desc 删除用户
func (l *UserDeleteLogic) UserDelete(req *types.UserIDReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.UserRpc.UserDelete(l.ctx, &pb.IDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}

	return
}
