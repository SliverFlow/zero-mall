package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"server/app/user/model"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserUpdate
// Author [SliverFlow]
// @desc 用户更新
func (l *UserUpdateLogic) UserUpdate(in *pb.UserUpdateReq) (*pb.UserNil, error) {
	var u model.User
	_ = copier.Copy(&u, in)
	u.ID = uint(int(in.ID))

	err := l.svcCtx.UserModel.UserUpdate(l.ctx, &u)
	// 返回nil表示成功
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "user not exist in userId = %s", in.GetID())
		}
		return nil, err
	}
	return &pb.UserNil{}, nil
}
