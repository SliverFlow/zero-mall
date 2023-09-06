package public

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.CreateReq) (resp *types.NilReply, err error) {
	var createIn pb.CreateReq
	_ = copier.Copy(&createIn, req)
	_, err = l.svcCtx.UserRpc.UserCreate(l.ctx, &createIn)
	if err != nil {
		return nil, err
	}

	return
}
