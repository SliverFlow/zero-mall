// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package system

import (
	"context"

	"server/app/system/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRole       = pb.CreateRole
	Menu             = pb.Menu
	MenuListReply    = pb.MenuListReply
	NilReply         = pb.NilReply
	NilReq           = pb.NilReq
	RoleIDReq        = pb.RoleIDReq
	SystemLoginReply = pb.SystemLoginReply
	SystemLoginReq   = pb.SystemLoginReq
	User             = pb.User

	System interface {
		// 系统登录
		Login(ctx context.Context, in *SystemLoginReq, opts ...grpc.CallOption) (*SystemLoginReply, error)
		// 创建角色
		RoleCreate(ctx context.Context, in *CreateRole, opts ...grpc.CallOption) (*NilReply, error)
		// 查询某个角色的菜单
		MenuListByRole(ctx context.Context, in *RoleIDReq, opts ...grpc.CallOption) (*MenuListReply, error)
	}

	defaultSystem struct {
		cli zrpc.Client
	}
)

func NewSystem(cli zrpc.Client) System {
	return &defaultSystem{
		cli: cli,
	}
}

// 系统登录
func (m *defaultSystem) Login(ctx context.Context, in *SystemLoginReq, opts ...grpc.CallOption) (*SystemLoginReply, error) {
	client := pb.NewSystemClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 创建角色
func (m *defaultSystem) RoleCreate(ctx context.Context, in *CreateRole, opts ...grpc.CallOption) (*NilReply, error) {
	client := pb.NewSystemClient(m.cli.Conn())
	return client.RoleCreate(ctx, in, opts...)
}

// 查询某个角色的菜单
func (m *defaultSystem) MenuListByRole(ctx context.Context, in *RoleIDReq, opts ...grpc.CallOption) (*MenuListReply, error) {
	client := pb.NewSystemClient(m.cli.Conn())
	return client.MenuListByRole(ctx, in, opts...)
}
