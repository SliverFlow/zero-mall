package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuChangeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuChangeStatusLogic {
	return &MenuChangeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuChangeStatusLogic) MenuChangeStatus(in *pb.MenuChangeStatusReq) (*pb.SystemNil, error) {
	err := l.svcCtx.SystemModel.MenuChangeStatus(l.ctx, in.GetID(), in.GetPID(), in.GetStatus())
	if err != nil {
		logx.Error("update menu status err", err.Error())
		return nil, status.Errorf(100001, "更新菜单状态失败")
	}

	return &pb.SystemNil{}, nil
}
