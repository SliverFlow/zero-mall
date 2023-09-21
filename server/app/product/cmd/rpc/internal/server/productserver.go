// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"server/app/product/cmd/rpc/internal/logic"
	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

// 所有的分类列表
func (s *ProductServer) CategoryListAll(ctx context.Context, in *pb.ProductNil) (*pb.CategoryListAllReply, error) {
	l := logic.NewCategoryListAllLogic(ctx, s.svcCtx)
	return l.CategoryListAll(in)
}

// 分类创建
func (s *ProductServer) CategoryCreate(ctx context.Context, in *pb.CategoryCreateReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryCreateLogic(ctx, s.svcCtx)
	return l.CategoryCreate(in)
}

// 修改分类状态
func (s *ProductServer) CategoryChangeStatus(ctx context.Context, in *pb.CategoryChangeStatusReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryChangeStatusLogic(ctx, s.svcCtx)
	return l.CategoryChangeStatus(in)
}

// 根据 分类 id 查询分类
func (s *ProductServer) CategoryFind(ctx context.Context, in *pb.CategoryIDReq) (*pb.Category, error) {
	l := logic.NewCategoryFindLogic(ctx, s.svcCtx)
	return l.CategoryFind(in)
}

// 更新分类
func (s *ProductServer) CategoryUpdate(ctx context.Context, in *pb.CategoryUpdateReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryUpdateLogic(ctx, s.svcCtx)
	return l.CategoryUpdate(in)
}

// 批量删除分类
func (s *ProductServer) CategoryBatchDelete(ctx context.Context, in *pb.CategoryIDSReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryBatchDeleteLogic(ctx, s.svcCtx)
	return l.CategoryBatchDelete(in)
}

// 分局父分类 id 查询子分类 id 列表
func (s *ProductServer) CategoryFindChildrenID(ctx context.Context, in *pb.CategoryIDReq) (*pb.CategoryIDSReply, error) {
	l := logic.NewCategoryFindChildrenIDLogic(ctx, s.svcCtx)
	return l.CategoryFindChildrenID(in)
}

// @desc  单个删除分类
func (s *ProductServer) CategoryDelete(ctx context.Context, in *pb.CategoryIDReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryDeleteLogic(ctx, s.svcCtx)
	return l.CategoryDelete(in)
}

// 商品列表 分页
func (s *ProductServer) ProductList(ctx context.Context, in *pb.ProductNil) (*pb.ProductNil, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}
