// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package server

import (
	"context"

	"server/app/system/cmd/rpc/internal/logic"
	"server/app/system/cmd/rpc/internal/svc"
	"server/app/system/cmd/rpc/pb"
)

type SystemServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSystemServer
}

func NewSystemServer(svcCtx *svc.ServiceContext) *SystemServer {
	return &SystemServer{
		svcCtx: svcCtx,
	}
}

// 系统登录
func (s *SystemServer) Login(ctx context.Context, in *pb.SystemLoginReq) (*pb.SystemLoginReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 创建角色
func (s *SystemServer) RoleCreate(ctx context.Context, in *pb.CreateRole) (*pb.SystemNil, error) {
	l := logic.NewRoleCreateLogic(ctx, s.svcCtx)
	return l.RoleCreate(in)
}

// 查询某个角色的菜单
func (s *SystemServer) MenuListByRole(ctx context.Context, in *pb.RoleIDReq) (*pb.MenuListReply, error) {
	l := logic.NewMenuListByRoleLogic(ctx, s.svcCtx)
	return l.MenuListByRole(in)
}

func (s *SystemServer) MenuListAllByRole(ctx context.Context, in *pb.RoleIDReq) (*pb.MenuListReply, error) {
	l := logic.NewMenuListAllByRoleLogic(ctx, s.svcCtx)
	return l.MenuListAllByRole(in)
}

func (s *SystemServer) MenuChangeStatus(ctx context.Context, in *pb.MenuChangeStatusReq) (*pb.SystemNil, error) {
	l := logic.NewMenuChangeStatusLogic(ctx, s.svcCtx)
	return l.MenuChangeStatus(in)
}

func (s *SystemServer) MenuCreate(ctx context.Context, in *pb.MenuCreateReq) (*pb.SystemNil, error) {
	l := logic.NewMenuCreateLogic(ctx, s.svcCtx)
	return l.MenuCreate(in)
}

func (s *SystemServer) MenuUpdate(ctx context.Context, in *pb.MenuUpdateReq) (*pb.SystemNil, error) {
	l := logic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *SystemServer) MenuFind(ctx context.Context, in *pb.MenuIDReq) (*pb.Menu, error) {
	l := logic.NewMenuFindLogic(ctx, s.svcCtx)
	return l.MenuFind(in)
}

func (s *SystemServer) MenuDelete(ctx context.Context, in *pb.MenuIDsReq) (*pb.SystemNil, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}

func (s *SystemServer) MenuFindChildrenID(ctx context.Context, in *pb.MenuIDReq) (*pb.MenuIDsReply, error) {
	l := logic.NewMenuFindChildrenIDLogic(ctx, s.svcCtx)
	return l.MenuFindChildrenID(in)
}
