package private

import (
	"context"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Update 更新用户信息
func (l *UpdateLogic) Update(req *types.UpdateReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.UserRpc.Update(l.ctx, &pb.UpdateReq{
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   req.Avatar,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	return
}
