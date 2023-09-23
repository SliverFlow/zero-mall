package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChangeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChangeStatusLogic {
	return &UserChangeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserChangeStatus
// Author [SliverFlow]
// @desc 管理员账户切换角色
func (l *UserChangeStatusLogic) UserChangeStatus(in *pb.UserChangeStatusReq) (*pb.UserNil, error) {
	err := l.svcCtx.UserModel.UserChangeStatus(l.ctx, in.GetID(), in.GetStatus())
	if err != nil {
		return nil, status.Errorf(100001, "更新用户状态失败")
	}

	return &pb.UserNil{}, nil
}
