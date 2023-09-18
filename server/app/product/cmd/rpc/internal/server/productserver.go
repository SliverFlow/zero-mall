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

func (s *ProductServer) ProductList(ctx context.Context, in *pb.ProductNil) (*pb.ProductNil, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

func (s *ProductServer) CategoryListAll(ctx context.Context, in *pb.ProductNil) (*pb.CategoryListAllReply, error) {
	l := logic.NewCategoryListAllLogic(ctx, s.svcCtx)
	return l.CategoryListAll(in)
}

func (s *ProductServer) CategoryCreate(ctx context.Context, in *pb.CategoryCreateReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryCreateLogic(ctx, s.svcCtx)
	return l.CategoryCreate(in)
}

func (s *ProductServer) CategoryChangeStatus(ctx context.Context, in *pb.CategoryChangeStatusReq) (*pb.ProductNil, error) {
	l := logic.NewCategoryChangeStatusLogic(ctx, s.svcCtx)
	return l.CategoryChangeStatus(in)
}
