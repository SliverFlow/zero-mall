package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/system/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuFindLogic {
	return &MenuFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuFindLogic) MenuFind(req *types.IDReq) (resp *types.Menu, err error) {
	pbmenu, err := l.svcCtx.SystemRpc.MenuFind(l.ctx, &pb.MenuIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}

	var tymenu types.Menu
	_ = copier.Copy(&tymenu, pbmenu)
	tymenu.Meta.Icon = pbmenu.Icon
	tymenu.Meta.Title = pbmenu.Title

	return &tymenu, nil
}
