package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuFindChildrenIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuFindChildrenIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuFindChildrenIDLogic {
	return &MenuFindChildrenIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuFindChildrenIDLogic) MenuFindChildrenID(in *pb.MenuIDReq) (*pb.MenuIDsReply, error) {
	ids, err := l.svcCtx.SystemModel.MenuFindChildrenID(l.ctx, in.GetID())
	if err != nil {
		return nil, status.Errorf(100001, "查询子菜单 id 失败")
	}

	return &pb.MenuIDsReply{
		IDs: ids,
	}, nil
}
