package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"
	"server/app/system/model"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *pb.MenuUpdateReq) (*pb.SystemNil, error) {
	var m model.Menu
	_ = copier.Copy(&m, in)
	m.ID = uint(in.GetID())
	m.Meta.Icon = in.Icon
	m.Meta.Title = in.Title

	if err := l.svcCtx.SystemModel.MenuUpdate(l.ctx, &m); err != nil {
		return nil, status.Errorf(100001, "菜单更新失败")
	}

	return &pb.SystemNil{}, nil
}
