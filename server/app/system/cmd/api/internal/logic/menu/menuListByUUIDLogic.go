package menu

import (
	"context"
	systempb "server/app/system/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListByUUIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListByUUIDLogic {
	return &MenuListByUUIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListByUUIDLogic) MenuListByUUID(req *types.NilReq) (resp *types.MenuListByRoleReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	user, err := l.svcCtx.UserRpc.UserFindByUUID(l.ctx, &userpb.UUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}

	reply, err := l.svcCtx.SystemRpc.MenuListByRole(l.ctx, &systempb.RoleIDReq{ID: user.User.Role})
	if err != nil {
		return nil, err
	}
	list := findChildren(reply.List, 0)
	return &types.MenuListByRoleReply{List: list}, nil
}
