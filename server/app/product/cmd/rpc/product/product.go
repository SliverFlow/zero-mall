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
		// 所有的分类列表
		CategoryListAll(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*CategoryListAllReply, error)
		// 分类创建
		CategoryCreate(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*ProductNil, error)
		// 修改分类状态
		CategoryChangeStatus(ctx context.Context, in *CategoryChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error)
		// 根据 分类 id 查询分类
		CategoryFind(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*Category, error)
		// 更新分类
		CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*ProductNil, error)
		// 批量删除分类
		CategoryBatchDelete(ctx context.Context, in *CategoryIDSReq, opts ...grpc.CallOption) (*ProductNil, error)
		// 分局父分类 id 查询子分类 id 列表
		CategoryFindChildrenID(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*CategoryIDSReply, error)
		// @desc  单个删除分类
		CategoryDelete(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*ProductNil, error)
		// 商品列表 分页
		ProductList(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*ProductNil, error)
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

// 所有的分类列表
func (m *defaultProduct) CategoryListAll(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*CategoryListAllReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryListAll(ctx, in, opts...)
}

// 分类创建
func (m *defaultProduct) CategoryCreate(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryCreate(ctx, in, opts...)
}

// 修改分类状态
func (m *defaultProduct) CategoryChangeStatus(ctx context.Context, in *CategoryChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryChangeStatus(ctx, in, opts...)
}

// 根据 分类 id 查询分类
func (m *defaultProduct) CategoryFind(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*Category, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryFind(ctx, in, opts...)
}

// 更新分类
func (m *defaultProduct) CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryUpdate(ctx, in, opts...)
}

// 批量删除分类
func (m *defaultProduct) CategoryBatchDelete(ctx context.Context, in *CategoryIDSReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryBatchDelete(ctx, in, opts...)
}

// 分局父分类 id 查询子分类 id 列表
func (m *defaultProduct) CategoryFindChildrenID(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*CategoryIDSReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryFindChildrenID(ctx, in, opts...)
}

// @desc  单个删除分类
func (m *defaultProduct) CategoryDelete(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryDelete(ctx, in, opts...)
}

// 商品列表 分页
func (m *defaultProduct) ProductList(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}
