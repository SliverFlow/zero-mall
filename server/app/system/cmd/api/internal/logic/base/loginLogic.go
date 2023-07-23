package base

import (
	"context"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login
// Author [SliverFlow]
// @desc 系统登录
// @param req *types.LoginReq
// @return resp *types.LoginReply, err error
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	return
}
