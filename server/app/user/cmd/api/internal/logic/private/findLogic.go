package private

import (
	"context"
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
	// todo: add your logic here and delete this line
	reply, err := l.svcCtx.UserRpc.Find(l.ctx, &pb.IDReq{UserID: req.Id})
	if err != nil {
		return nil, err
	}
	u := reply.GetUser()
	return &types.UserReply{User: types.User{
		ID:        u.ID,
		UserID:    u.UserID,
		Username:  u.Username,
		UUID:      u.UUID,
		Nickname:  u.Nickname,
		Email:     u.Email,
		Avatar:    u.Avatar,
		Type:      u.Type,
		CreatedAt: u.CreatedAt,
	}}, nil
}
