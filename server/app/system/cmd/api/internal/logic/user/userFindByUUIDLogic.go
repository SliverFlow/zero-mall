package user

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFindByUUIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFindByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindByUUIDLogic {
	return &UserFindByUUIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFindByUUIDLogic) UserFindByUUID(req *types.NilReq) (resp *types.UserInfoReply, err error) {
	uuid, _ := xjwt.GetUserUUID(l.ctx)
	reply, err := l.svcCtx.UserRpc.UserFindByUUID(l.ctx, &pb.UUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}
	var u types.User
	_ = copier.Copy(&u, reply.User)

	return &types.UserInfoReply{User: u}, nil
}
