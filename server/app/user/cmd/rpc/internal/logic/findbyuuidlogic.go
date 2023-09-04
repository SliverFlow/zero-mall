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

type FindByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByUUIDLogic {
	return &FindByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByUUIDLogic) FindByUUID(in *pb.UUIDReq) (*pb.UserInfoReply, error) {
	user, err := l.svcCtx.UserModel.FindByUUID(l.ctx, in.UUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100101, "user not exist by uuid = %s", in.GetUUID())
		}
		return nil, err
	}

	u := copyUserFoReply(user)
	return &pb.UserInfoReply{
		User: u,
	}, nil
}
