package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListByRoleLogic {
	return &MenuListByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuListByRole 查询某个角色的菜单
func (l *MenuListByRoleLogic) MenuListByRole(in *pb.RoleIDReq) (*pb.MenuListReply, error) {
	list, err := l.svcCtx.SystemModel.MenuListByRole(l.ctx, in.ID)
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
	fmt.Println("enter", enter)
	return &pb.MenuListReply{List: enter}, nil
}
