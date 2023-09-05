package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListAllByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListAllByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListAllByRoleLogic {
	return &MenuListAllByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListAllByRoleLogic) MenuListAllByRole(in *pb.RoleIDReq) (*pb.MenuListReply, error) {
	list, err := l.svcCtx.SystemModel.MenuTreeListAllByRole(l.ctx, in.ID)
	if err != nil {
		logx.Error("find menu list by role err ", err.Error())
		return nil, status.Errorf(100001, "查询菜单列表失败")
	}
	var enter []*pb.Menu
	for _, menu := range *list {
		var m pb.Menu
		_ = copier.Copy(&m, &menu)
		enter = append(enter, &m)
	}
	return &pb.MenuListReply{List: enter}, nil
}
