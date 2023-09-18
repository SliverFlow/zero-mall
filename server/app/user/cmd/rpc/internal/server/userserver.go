// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"server/app/user/cmd/rpc/internal/logic"
	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 用户相关
func (s *UserServer) UserFind(ctx context.Context, in *pb.IDReq) (*pb.UserInfoReply, error) {
	l := logic.NewUserFindLogic(ctx, s.svcCtx)
	return l.UserFind(in)
}

func (s *UserServer) UserFindByUsername(ctx context.Context, in *pb.UsernameReq) (*pb.UserInfoReply, error) {
	l := logic.NewUserFindByUsernameLogic(ctx, s.svcCtx)
	return l.UserFindByUsername(in)
}

func (s *UserServer) UserList(ctx context.Context, in *pb.PageReq) (*pb.PageReply, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

func (s *UserServer) UserCreate(ctx context.Context, in *pb.CreateReq) (*pb.UserNil, error) {
	l := logic.NewUserCreateLogic(ctx, s.svcCtx)
	return l.UserCreate(in)
}

func (s *UserServer) UserUpdate(ctx context.Context, in *pb.UpdateReq) (*pb.UserNil, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *UserServer) UserDelete(ctx context.Context, in *pb.IDReq) (*pb.UserNil, error) {
	l := logic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

func (s *UserServer) UserBatchDelete(ctx context.Context, in *pb.IDsReq) (*pb.UserNil, error) {
	l := logic.NewUserBatchDeleteLogic(ctx, s.svcCtx)
	return l.UserBatchDelete(in)
}

func (s *UserServer) Login(ctx context.Context, in *pb.UserLoginReq) (*pb.UserLoginReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) UserFindByUUID(ctx context.Context, in *pb.UUIDReq) (*pb.UserInfoReply, error) {
	l := logic.NewUserFindByUUIDLogic(ctx, s.svcCtx)
	return l.UserFindByUUID(in)
}

func (s *UserServer) AdminChangeRole(ctx context.Context, in *pb.AdminChangeRoleReq) (*pb.UserNil, error) {
	l := logic.NewAdminChangeRoleLogic(ctx, s.svcCtx)
	return l.AdminChangeRole(in)
}

// 商家相关
func (s *UserServer) BusinessCreate(ctx context.Context, in *pb.BusinessCreateReq) (*pb.UserNil, error) {
	l := logic.NewBusinessCreateLogic(ctx, s.svcCtx)
	return l.BusinessCreate(in)
}

func (s *UserServer) BusinessList(ctx context.Context, in *pb.PageReq) (*pb.BusinessPageReply, error) {
	l := logic.NewBusinessListLogic(ctx, s.svcCtx)
	return l.BusinessList(in)
}
