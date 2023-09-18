package logic

import (
	"context"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBatchDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBatchDeleteLogic {
	return &UserBatchDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserBatchDeleteLogic) UserBatchDelete(in *pb.IDsReq) (*pb.UserNil, error) {
	// todo: add your logic here and delete this line

	return &pb.UserNil{}, nil
}
