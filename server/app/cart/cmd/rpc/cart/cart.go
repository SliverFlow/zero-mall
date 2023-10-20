// Code generated by goctl. DO NOT EDIT.
// Source: cart.proto

package cart

import (
	"context"

	"server/app/cart/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CartCreateReq = pb.CartCreateReq
	CartInfo      = pb.CartInfo
	CartListReply = pb.CartListReply
	CartListReq   = pb.CartListReq
	CartNil       = pb.CartNil

	Cart interface {
		CartCreate(ctx context.Context, in *CartCreateReq, opts ...grpc.CallOption) (*CartNil, error)
		CartList(ctx context.Context, in *CartListReq, opts ...grpc.CallOption) (*CartListReply, error)
	}

	defaultCart struct {
		cli zrpc.Client
	}
)

func NewCart(cli zrpc.Client) Cart {
	return &defaultCart{
		cli: cli,
	}
}

func (m *defaultCart) CartCreate(ctx context.Context, in *CartCreateReq, opts ...grpc.CallOption) (*CartNil, error) {
	client := pb.NewCartClient(m.cli.Conn())
	return client.CartCreate(ctx, in, opts...)
}

func (m *defaultCart) CartList(ctx context.Context, in *CartListReq, opts ...grpc.CallOption) (*CartListReply, error) {
	client := pb.NewCartClient(m.cli.Conn())
	return client.CartList(ctx, in, opts...)
}
