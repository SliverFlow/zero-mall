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
func (s *SystemServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}