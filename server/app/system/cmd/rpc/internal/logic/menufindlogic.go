package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuFindLogic {
	return &MenuFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuFindLogic) MenuFind(in *pb.MenuIDReq) (*pb.Menu, error) {
	menu, err := l.svcCtx.SystemModel.MenuFind(l.ctx, in.GetID())
	if err != nil {
		return nil, status.Errorf(100001, "分类查询失败")
	}
	var pbmenu pb.Menu
	_ = copier.Copy(&pbmenu, menu)
	pbmenu.Icon = menu.Meta.Icon
	pbmenu.Title = menu.Meta.Title
	pbmenu.CreatedAt = menu.CreatedAt.Unix()
	pbmenu.UpdatedAt = menu.UpdatedAt.Unix()

	return &pbmenu, nil
}
