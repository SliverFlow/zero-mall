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

type UserFindByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFindByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindByUUIDLogic {
	return &UserFindByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserFindByUUID
// Author [SliverFlow]
// @desc 根据 uuid 查找用户
func (l *UserFindByUUIDLogic) UserFindByUUID(in *pb.UUIDReq) (*pb.UserInfoReply, error) {
	user, err := l.svcCtx.UserModel.UserFindByUUID(l.ctx, in.UUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100101, "user not exist by uuid = %s", in.GetUUID())
		}
		return nil, err
	}

	u := copyUserFoReply(user)
	return &pb.UserInfoReply{User: u}, nil
}
