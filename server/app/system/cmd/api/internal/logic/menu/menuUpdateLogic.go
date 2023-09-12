package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/system/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.MenuUpdateReq) (resp *types.NilReply, err error) {
	var pbmenu pb.MenuUpdateReq
	_ = copier.Copy(&pbmenu, req)
	pbmenu.Title = req.Meta.Title
	pbmenu.Icon = req.Meta.Icon

	_, err = l.svcCtx.SystemRpc.MenuUpdate(l.ctx, &pbmenu)
	if err != nil {
		return nil, err
	}

	return
}
