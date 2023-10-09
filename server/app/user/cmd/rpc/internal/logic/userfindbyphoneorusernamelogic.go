package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFindByPhoneOrUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFindByPhoneOrUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindByPhoneOrUsernameLogic {
	return &UserFindByPhoneOrUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserFindByPhoneOrUsername
// Author [SliverFlow]
// @desc 通过手机号或用户名查找用户
func (l *UserFindByPhoneOrUsernameLogic) UserFindByPhoneOrUsername(in *pb.UserFindByPhoneOrUsernameReq) (*pb.UserInfoReply, error) {
	user, err := l.svcCtx.UserModel.UserFindByPhoneOrUsername(l.ctx, in.GetPhone(), in.GetUsername())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "用户不存在")
		}
		return nil, status.Errorf(100001, "查询用户信息失败")
	}

	return &pb.UserInfoReply{
		User: copyUserFoReply(user),
	}, nil
}
