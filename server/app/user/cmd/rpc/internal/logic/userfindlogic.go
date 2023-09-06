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

type UserFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindLogic {
	return &UserFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户相关
func (l *UserFindLogic) UserFind(in *pb.IDReq) (*pb.UserInfoReply, error) {
	user, err := l.svcCtx.UserModel.UserFind(l.ctx, in.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100101, "user not exist in id = %d", in.GetID())
		}
		return nil, err
	}

	u := copyUserFoReply(user)

	return &pb.UserInfoReply{User: u}, nil
}

func copyUserFoReply(user *model.User) *pb.UserReply {
	var u pb.UserReply
	_ = copier.Copy(&u, &user)
	u.ID = int64(user.ID)
	u.CreatedAt = user.CreatedAt.Unix()
	u.UpdatedAt = user.UpdatedAt.Unix()
	return &u
}
