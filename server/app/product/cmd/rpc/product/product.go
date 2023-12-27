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
	Category                       = pb.Category
	CategoryChangeStatusReq        = pb.CategoryChangeStatusReq
	CategoryCreateReq              = pb.CategoryCreateReq
	CategoryIDByProductListReply   = pb.CategoryIDByProductListReply
	CategoryIDByProductListReq     = pb.CategoryIDByProductListReq
	CategoryIDReq                  = pb.CategoryIDReq
	CategoryIDSReply               = pb.CategoryIDSReply
	CategoryIDSReq                 = pb.CategoryIDSReq
	CategoryListAllReply           = pb.CategoryListAllReply
	CategoryUpdateReq              = pb.CategoryUpdateReq
	ProductAddStockReq             = pb.ProductAddStockReq
	ProductChangeStatusReq         = pb.ProductChangeStatusReq
	ProductCreateReq               = pb.ProductCreateReq
	ProductDeductionStockReq       = pb.ProductDeductionStockReq
	ProductDeleteReq               = pb.ProductDeleteReq
	ProductFindListByBusinessIDReq = pb.ProductFindListByBusinessIDReq
	ProductFindReq                 = pb.ProductFindReq
	ProductIDsReq                  = pb.ProductIDsReq
	ProductListReply               = pb.ProductListReply
	ProductListReq                 = pb.ProductListReq
	ProductNil                     = pb.ProductNil
	ProductReply                   = pb.ProductReply
	ProductStagingReq              = pb.ProductStagingReq
	ProductUpdateReq               = pb.ProductUpdateReq

	Product interface {
		// @desc 所有的分类列表
		CategoryListAll(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*CategoryListAllReply, error)
		// @desc 分类创建
		CategoryCreate(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*ProductNil, error)
		// @desc 修改分类状态
		CategoryChangeStatus(ctx context.Context, in *CategoryChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error)
		// @desc 根据 分类 id 查询分类
		CategoryFind(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*Category, error)
		// @desc 更新分类
		CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*ProductNil, error)
		// @desc 批量删除分类
		CategoryBatchDelete(ctx context.Context, in *CategoryIDSReq, opts ...grpc.CallOption) (*ProductNil, error)
		// @desc 分局父分类 id 查询子分类 id 列表
		CategoryFindChildrenID(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*CategoryIDSReply, error)
		// @desc 单个删除分类
		CategoryDelete(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*ProductNil, error)
		// 商品列表 分页
		ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListReply, error)
		ProductCreate(ctx context.Context, in *ProductCreateReq, opts ...grpc.CallOption) (*ProductNil, error)
		ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductNil, error)
		// @desc 下架某商家的所有商品
		ProductStaging(ctx context.Context, in *ProductStagingReq, opts ...grpc.CallOption) (*ProductNil, error)
		ProductChangeStatus(ctx context.Context, in *ProductChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error)
		ProductFind(ctx context.Context, in *ProductFindReq, opts ...grpc.CallOption) (*ProductReply, error)
		ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductNil, error)
		ProductDeductionStock(ctx context.Context, in *ProductDeductionStockReq, opts ...grpc.CallOption) (*ProductNil, error)
		ProductAddStock(ctx context.Context, in *ProductAddStockReq, opts ...grpc.CallOption) (*ProductNil, error)
		ProductFindListByIDs(ctx context.Context, in *ProductIDsReq, opts ...grpc.CallOption) (*ProductListReply, error)
		ProductFindListByBusinessID(ctx context.Context, in *ProductFindListByBusinessIDReq, opts ...grpc.CallOption) (*ProductListReply, error)
		CategoryIDByProductList(ctx context.Context, in *CategoryIDByProductListReq, opts ...grpc.CallOption) (*CategoryIDByProductListReply, error)
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

// @desc 所有的分类列表
func (m *defaultProduct) CategoryListAll(ctx context.Context, in *ProductNil, opts ...grpc.CallOption) (*CategoryListAllReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryListAll(ctx, in, opts...)
}

// @desc 分类创建
func (m *defaultProduct) CategoryCreate(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryCreate(ctx, in, opts...)
}

// @desc 修改分类状态
func (m *defaultProduct) CategoryChangeStatus(ctx context.Context, in *CategoryChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryChangeStatus(ctx, in, opts...)
}

// @desc 根据 分类 id 查询分类
func (m *defaultProduct) CategoryFind(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*Category, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryFind(ctx, in, opts...)
}

// @desc 更新分类
func (m *defaultProduct) CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryUpdate(ctx, in, opts...)
}

// @desc 批量删除分类
func (m *defaultProduct) CategoryBatchDelete(ctx context.Context, in *CategoryIDSReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryBatchDelete(ctx, in, opts...)
}

// @desc 分局父分类 id 查询子分类 id 列表
func (m *defaultProduct) CategoryFindChildrenID(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*CategoryIDSReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryFindChildrenID(ctx, in, opts...)
}

// @desc 单个删除分类
func (m *defaultProduct) CategoryDelete(ctx context.Context, in *CategoryIDReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryDelete(ctx, in, opts...)
}

// 商品列表 分页
func (m *defaultProduct) ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}

func (m *defaultProduct) ProductCreate(ctx context.Context, in *ProductCreateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductCreate(ctx, in, opts...)
}

func (m *defaultProduct) ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductDelete(ctx, in, opts...)
}

// @desc 下架某商家的所有商品
func (m *defaultProduct) ProductStaging(ctx context.Context, in *ProductStagingReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductStaging(ctx, in, opts...)
}

func (m *defaultProduct) ProductChangeStatus(ctx context.Context, in *ProductChangeStatusReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductChangeStatus(ctx, in, opts...)
}

func (m *defaultProduct) ProductFind(ctx context.Context, in *ProductFindReq, opts ...grpc.CallOption) (*ProductReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductFind(ctx, in, opts...)
}

func (m *defaultProduct) ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductUpdate(ctx, in, opts...)
}

func (m *defaultProduct) ProductDeductionStock(ctx context.Context, in *ProductDeductionStockReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductDeductionStock(ctx, in, opts...)
}

func (m *defaultProduct) ProductAddStock(ctx context.Context, in *ProductAddStockReq, opts ...grpc.CallOption) (*ProductNil, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductAddStock(ctx, in, opts...)
}

func (m *defaultProduct) ProductFindListByIDs(ctx context.Context, in *ProductIDsReq, opts ...grpc.CallOption) (*ProductListReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductFindListByIDs(ctx, in, opts...)
}

func (m *defaultProduct) ProductFindListByBusinessID(ctx context.Context, in *ProductFindListByBusinessIDReq, opts ...grpc.CallOption) (*ProductListReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductFindListByBusinessID(ctx, in, opts...)
}

func (m *defaultProduct) CategoryIDByProductList(ctx context.Context, in *CategoryIDByProductListReq, opts ...grpc.CallOption) (*CategoryIDByProductListReply, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryIDByProductList(ctx, in, opts...)
}
