package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/system/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListByRoleLogic {
	return &MenuListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListByRoleLogic) MenuListByRole(req *types.RoleIDReq) (resp *types.MenuListByRoleReply, err error) {
	reply, err := l.svcCtx.SystemRpc.MenuListByRole(l.ctx, &pb.RoleIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	list := l.findChildren(reply.List, 0)
	return &types.MenuListByRoleReply{List: list}, nil
}

func (l *MenuListByRoleLogic) findChildren(list []*pb.Menu, root int64) []types.Menu {
	var rootList []types.Menu
	for _, reply := range list {
		var rep types.Menu
		if reply.ParentId == root {
			_ = copier.Copy(&rep, reply)
			rep.Meta.Title = reply.Title
			rep.Meta.Icon = reply.Icon
			rep.Children = l.findChildren(list, rep.ID)
			if len(rep.Children) == 0 {
				rep.Children = make([]types.Menu, 0)
			}
			rootList = append(rootList, rep)
		}
	}
	return rootList
}
