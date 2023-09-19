// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package product

import (
	"context"

	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Category                = pb.Category
	CategoryChangeStatusReq = pb.CategoryChangeStatusReq
	CategoryCreateReq       = pb.CategoryCreateReq
	CategoryIDReq           = pb.CategoryIDReq
	CategoryIDSReply        = pb.CategoryIDSReply
	CategoryIDSReq          = pb.CategoryIDSReq
	CategoryListAllReply    = pb.CategoryListAllReply
	CategoryUpdateReq       = pb.CategoryUpdateReq
	ProductNil              = pb.ProductNil

	Product interface {
		ProductList(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*ProductNil, error)
		CategoryListAll(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*CategoryListAllReply, error)
		CategoryCreate(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*ProductNil, error)
		CategoryChangeStatus(ctx context.Context, in *CategoryChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error)
		CategoryFind(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*Category, error)
		CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*ProductNil, error)
		CategoryBatchDelete(ctx context.Context, in *CategoryIDSReq, opts ...grpc.CallOption) (*ProductNil, error)
		CategoryFindChildrenID(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*CategoryIDSReply, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) ProductList(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}

func (m *defaultProduct) CategoryListAll(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*CategoryListAllReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryListAll(ctx, in, opts...)
}

func (m *defaultProduct) CategoryCreate(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryCreate(ctx, in, opts...)
}

func (m *defaultProduct) CategoryChangeStatus(ctx context.Context, in *CategoryChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryChangeStatus(ctx, in, opts...)
}

func (m *defaultProduct) CategoryFind(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*Category, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryFind(ctx, in, opts...)
}

func (m *defaultProduct) CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryUpdate(ctx, in, opts...)
}

func (m *defaultProduct) CategoryBatchDelete(ctx context.Context, in *CategoryIDSReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryBatchDelete(ctx, in, opts...)
}

func (m *defaultProduct) CategoryFindChildrenID(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*CategoryIDSReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryFindChildrenID(ctx, in, opts...)
}
