package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminChangeRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminChangeRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminChangeRoleLogic {
	return &AdminChangeRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminChangeRoleLogic) AdminChangeRole(in *pb.AdminChangeRoleReq) (*pb.Nil, error) {
	err := l.svcCtx.UserModel.AdminChangeRole(l.ctx, in.GetUsername(), in.GetRole())
	if err != nil {
		return nil, status.Errorf(100001, err.Error())
	}
	return &pb.Nil{}, nil
}
