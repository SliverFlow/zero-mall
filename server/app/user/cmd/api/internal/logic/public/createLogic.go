package public

import (
	"context"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateReq) (resp *types.NilReply, err error) {
	_, err = l.svcCtx.UserRpc.Create(l.ctx, &pb.CreateReq{
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   req.Avatar,
		Type:     req.Type,
	})
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	return
}
