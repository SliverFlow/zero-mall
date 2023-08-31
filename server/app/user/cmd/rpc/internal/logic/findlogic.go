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

type FindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLogic {
	return &FindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLogic) Find(in *pb.IDReq) (*pb.UserInfoReply, error) {
	u, err := l.svcCtx.UserModel.Find(l.ctx, in.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100101, "user not exist in id = %d", in.GetID())
		}
		return nil, err
	}
	return &pb.UserInfoReply{User: &pb.UserReply{
		ID:        int64(u.ID),
		UUID:      u.UUID,
		Username:  u.Username,
		Email:     u.Email,
		Nickname:  u.Nickname,
		Password:  u.Password,
		Avatar:    u.Avatar,
		Status:    u.Status,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}}, nil
}
