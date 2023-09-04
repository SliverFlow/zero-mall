package logic

import (
	"context"
	"time"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 系统登录
func (l *LoginLogic) Login(in *pb.SystemLoginReq) (*pb.SystemLoginReply, error) {
	// todo: add your logic here and delete this line

	return &pb.SystemLoginReply{User: &pb.User{
		UserID:    "21111",
		UUID:      "111",
		Username:  "1111",
		Email:     "1111",
		Nickname:  "111",
		Avatar:    "11111",
		Type:      0,
		Status:    0,
		ID:        0,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}}, nil
}
