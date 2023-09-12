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

func (l *UserUpdateLogic) UserUpdate(req *types.UpdateReq) (resp *types.NilReply, err error) {
	var u pb.UpdateReq
	_ = copier.Copy(&u, req)
	_, err = l.svcCtx.UserRpc.UserUpdate(l.ctx, &u)
	if err != nil {
		return nil, err
	}
	return
}
