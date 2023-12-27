package user

import (
	"context"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateByUUIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateByUUIDLogic {
	return &UserUpdateByUUIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateByUUIDLogic) UserUpdateByUUID(req *types.UserUpdateByUUIDReq) (resp *types.NilReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, err
	}
	uuidRep, err := l.svcCtx.UserRpc.UserFindByUUID(l.ctx, &userpb.UUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}
	if uuidRep.GetUser().ID == 0 {
		return nil, xerr.NewErrMsg("用户不存在")
	}
	_, err = l.svcCtx.UserRpc.UserUpdateByUUID(l.ctx, &userpb.UserUpdateByUUIDReq{
		UUID:     uuid,
		Phone:    req.Phone,
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   uuidRep.GetUser().Avatar,
		Role:     uuidRep.GetUser().Role,
		Status:   uuidRep.GetUser().Status,
	})
	if err != nil {
		return nil, err
	}

	return
}
