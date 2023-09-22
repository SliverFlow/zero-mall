package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"
	"server/app/system/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleCreateLogic {
	return &RoleCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleCreate 创建角色
func (l *RoleCreateLogic) RoleCreate(in *pb.CreateRole) (*pb.SystemNil, error) {
	err := l.svcCtx.SystemModel.RoleCreate(l.ctx, &model.Role{
		Name: in.GetName(),
	})
	if err != nil {
		logx.Error("create role err", err.Error())
		return nil, status.Errorf(100001, "创建角色失败")
	}
	return &pb.SystemNil{}, nil
}
