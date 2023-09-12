package private

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"

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

func (l *UserFindLogic) UserFind(req *types.IdReq) (resp *types.UserReply, err error) {
	reply, err := l.svcCtx.UserRpc.UserFind(l.ctx, &pb.IDReq{ID: req.Id})
	if err != nil {
		return nil, err
	}
	user := reply.GetUser()
	var u types.User
	_ = copier.Copy(&u, user)
	return &types.UserReply{User: u}, nil
}
