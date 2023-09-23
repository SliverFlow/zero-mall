package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFindByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFindByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindByUsernameLogic {
	return &UserFindByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserFindByUsername
// Author [SliverFlow]
// @desc 根据用户名查找用户
func (l *UserFindByUsernameLogic) UserFindByUsername(in *pb.UsernameReq) (*pb.UserInfoReply, error) {
	user, err := l.svcCtx.UserModel.UserFindByUsername(l.ctx, in.GetUsername())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "not find user by username is %s", in.Username)
		}
		return nil, err
	}

	return &pb.UserInfoReply{
		User: copyUserFoReply(user),
	}, nil
}
