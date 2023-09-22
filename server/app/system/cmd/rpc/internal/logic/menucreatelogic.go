package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/app/system/model"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuCreateLogic {
	return &MenuCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuCreateLogic) MenuCreate(in *pb.MenuCreateReq) (*pb.SystemNil, error) {
	var menu model.Menu
	_ = copier.Copy(&menu, in)
	menu.Title = in.GetTitle()
	menu.Icon = in.GetIcon()

	err := l.svcCtx.SystemModel.MenuCreate(l.ctx, &menu)
	if err != nil {
		logx.Error("create menu err", err.Error())
		return nil, status.Errorf(100001, "菜单创建失败")
	}

	return &pb.SystemNil{}, nil
}
