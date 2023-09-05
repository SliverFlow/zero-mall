// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdminChangeRoleReq = pb.AdminChangeRoleReq
	CreateReq          = pb.CreateReq
	IDReq              = pb.IDReq
	IDsReq             = pb.IDsReq
	Nil                = pb.Nil
	PageReply          = pb.PageReply
	PageReq            = pb.PageReq
	UUIDReq            = pb.UUIDReq
	UpdateReq          = pb.UpdateReq
	UserInfoReply      = pb.UserInfoReply
	UserLoginReply     = pb.UserLoginReply
	UserLoginReq       = pb.UserLoginReq
	UserReply          = pb.UserReply

	User interface {
		Find(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserInfoReply, error)
		List(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*PageReply, error)
		Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Nil, error)
		Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Nil, error)
		Delete(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*Nil, error)
		BatchDelete(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*Nil, error)
		Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginReply, error)
		FindByUUID(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfoReply, error)
		AdminChangeRole(ctx context.Context, in *AdminChangeRoleReq, opts ...grpc.CallOption) (*Nil, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Find(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Find(ctx, in, opts...)
}

func (m *defaultUser) List(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*PageReply, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultUser) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Nil, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultUser) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Nil, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultUser) Delete(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*Nil, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

func (m *defaultUser) BatchDelete(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*Nil, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.BatchDelete(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginReply, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) FindByUUID(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.FindByUUID(ctx, in, opts...)
}

func (m *defaultUser) AdminChangeRole(ctx context.Context, in *AdminChangeRoleReq, opts ...grpc.CallOption) (*Nil, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.AdminChangeRole(ctx, in, opts...)
}
