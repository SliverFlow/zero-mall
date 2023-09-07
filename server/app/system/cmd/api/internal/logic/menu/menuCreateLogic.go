package menu

import (
	"context"
	"github.com/jinzhu/copier"
	systempb "server/app/system/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuCreateLogic {
	return &MenuCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuCreateLogic) MenuCreate(req *types.MenuCreateReq) (resp *types.NilReply, err error) {
	uuid, _ := xjwt.GetUserUUID(l.ctx)

	user, err := l.svcCtx.UserRpc.UserFindByUUID(l.ctx, &userpb.UUIDReq{UUID: uuid})
	if err != nil {
		logx.Error("创建订单失败", err.Error())
		return nil, xerr.NewErrMsg("创建菜单失败")
	}
	if user.User.Role != 3 {
		logx.Error("该用户不能创建菜单 username ", user.User.Username)
		return nil, xerr.NewErrMsg("该用户不能创建菜单")
	}

	var menupb systempb.MenuCreateReq
	_ = copier.Copy(&menupb, req)
	menupb.Title = req.Meta.Title
	menupb.Icon = req.Meta.Icon

	_, err = l.svcCtx.SystemRpc.MenuCreate(l.ctx, &menupb)
	if err != nil {
		return nil, err
	}
	return
}
