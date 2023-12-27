package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/user/model"
	"server/common"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateByUUIDLogic {
	return &UserUpdateByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateByUUIDLogic) UserUpdateByUUID(in *pb.UserUpdateByUUIDReq) (*pb.UserNil, error) {

	u := &model.User{
		UUID:     in.UUID,
		Username: in.Username,
		Email:    in.Email,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Role:     in.Role,
		Status:   in.Status,
		Phone:    in.Phone,
	}
	if in.Password != "" {
		u.Password = common.BcryptHash(in.Password)
	}

	err := l.svcCtx.UserModel.UserUpdateByUUID(l.ctx, u.UUID, u)
	if err != nil {
		return nil, status.Errorf(100001, "更新用户失败")
	}

	return &pb.UserNil{}, nil
}
