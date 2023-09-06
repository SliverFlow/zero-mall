package private

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLogic {
	return &FindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLogic) Find(req *types.IdReq) (resp *types.UserReply, err error) {
	reply, err := l.svcCtx.UserRpc.UserFind(l.ctx, &pb.IDReq{ID: req.Id})
	if err != nil {
		return nil, err
	}
	user := reply.GetUser()
	var u types.User
	_ = copier.Copy(&u, user)
	return &types.UserReply{User: u}, nil
}
