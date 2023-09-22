package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuDeleteLogic) MenuDelete(in *pb.MenuIDsReq) (*pb.SystemNil, error) {
	err := l.svcCtx.SystemModel.MenuBatchDelete(l.ctx, in.GetIDs())
	if err != nil {
		return nil, status.Errorf(100001, "删除菜单失败")
	}

	return &pb.SystemNil{}, nil
}
