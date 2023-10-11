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

type UserFindByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFindByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindByPhoneLogic {
	return &UserFindByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserFindByPhone
// Author [SliverFlow]
// @desc 通过手机号查找用户
func (l *UserFindByPhoneLogic) UserFindByPhone(in *pb.PhoneReq) (*pb.UserInfoReply, error) {
	user, err := l.svcCtx.UserModel.UserFindByPhone(l.ctx, in.GetPhone())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "当前手机号码未注册")
		}
		return nil, status.Errorf(100001, "查询失败")
	}

	return &pb.UserInfoReply{
		User: copyUserFoReply(user),
	}, nil
}
