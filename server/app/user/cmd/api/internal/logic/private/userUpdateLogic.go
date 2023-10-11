package private

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserUpdate
// Author [SliverFlow]
// @desc 用户端更新信息
func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.Nil, err error) {
	var u pb.UserUpdateReq
	_ = copier.Copy(&u, req)
	_, err = l.svcCtx.UserRpc.UserUpdate(l.ctx, &u)
	if err != nil {
		return nil, err
	}
	return
}
