package private

import (
	"context"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindLogic {
	return &UserFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserFind
// Author [SliverFlow]
// @desc 用户端查询用户信息
func (l *UserFindLogic) UserFind(req *types.Nil) (resp *types.UserReply, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户 uuid 失败")
	}

	pbreply, err := l.svcCtx.UserRpc.UserFindByUUID(l.ctx, &pb.UUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}

	u := pbreply.GetUser()

	return &types.UserReply{
		UUID:     uuid,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Phone:    u.Phone,
		Email:    u.Email,
	}, nil
}
