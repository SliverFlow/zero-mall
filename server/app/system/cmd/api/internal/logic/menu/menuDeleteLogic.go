package menu

import (
	"context"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req *types.IDReq) (resp *types.NilReply, err error) {
	var ids []int64
	ids = append(ids, req.ID)
	pbmenu, err := l.svcCtx.SystemRpc.MenuFind(l.ctx, &pb.MenuIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if pbmenu.ParentId == 0 { // 父菜单
		pbreply, err := l.svcCtx.SystemRpc.MenuFindChildrenID(l.ctx, &pb.MenuIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		ids = append(ids, pbreply.IDs...)
	}

	_, err = l.svcCtx.SystemRpc.MenuDelete(l.ctx, &pb.MenuIDsReq{IDs: ids})
	if err != nil {
		return nil, err
	}

	return
}
