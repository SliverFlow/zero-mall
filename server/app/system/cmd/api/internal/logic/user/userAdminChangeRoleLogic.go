package user

import (
	"context"
	"server/app/user/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAdminChangeRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAdminChangeRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAdminChangeRoleLogic {
	return &UserAdminChangeRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAdminChangeRoleLogic) UserAdminChangeRole(req *types.AdminChangeRoleReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.UserRpc.AdminChangeRole(l.ctx, &pb.AdminChangeRoleReq{
		Username: req.Username,
		Role:     req.Role,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
