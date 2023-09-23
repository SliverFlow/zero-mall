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

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserDelete
// Author [SliverFlow]
// @desc 用户删除
func (l *UserDeleteLogic) UserDelete(in *pb.IDReq) (*pb.UserNil, error) {
	_, err := l.svcCtx.UserModel.UserFind(l.ctx, in.GetID())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logx.Error("user not exist in in = ", in.ID)
			return nil, status.Errorf(100001, "不存在 id 为 ", in.GetID(), "的用户")
		}
		logx.Error("find user info err", err)
		return nil, status.Errorf(100001, "查询用户信息失败")
	}

	err = l.svcCtx.UserModel.UserDelete(l.ctx, in.GetID())
	if err != nil {
		logx.Error("delete user err", err)
		return nil, status.Errorf(100001, "删除用户失败")
	}

	return &pb.UserNil{}, nil
}
