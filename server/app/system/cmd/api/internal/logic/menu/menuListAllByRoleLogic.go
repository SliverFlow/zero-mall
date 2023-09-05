package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/system/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListAllByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListAllByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListAllByRoleLogic {
	return &MenuListAllByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListAllByRoleLogic) MenuListAllByRole(req *types.RoleIDReq) (resp *types.MenuListByRoleReply, err error) {
	reply, err := l.svcCtx.SystemRpc.MenuListAllByRole(l.ctx, &pb.RoleIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	list := findChildren(reply.List, 0)
	return &types.MenuListByRoleReply{List: list}, nil
}

func findChildren(list []*pb.Menu, root int64) []types.Menu {
	var rootList []types.Menu
	for _, reply := range list {
		var rep types.Menu
		if reply.ParentId == root {
			_ = copier.Copy(&rep, reply)
			rep.Meta.Title = reply.Title
			rep.Meta.Icon = reply.Icon
			rep.Children = findChildren(list, rep.ID)
			if len(rep.Children) == 0 {
				rep.Children = make([]types.Menu, 0)
			}
			rootList = append(rootList, rep)
		}
	}
	return rootList
}
