package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/common"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessList
// Author [SliverFlow]
// @desc 商家分页列表
func (l *BusinessListLogic) BusinessList(in *pb.PageReq) (*pb.BusinessPageReply, error) {
	reply, total, err := l.svcCtx.UserModel.BusinessList(l.ctx, &common.PageInfo{
		Page:     int(in.GetPage()),
		PageSize: int(in.GetPageSize()),
		KeyWord:  in.KeyWork,
	})
	if err != nil {
		logx.Error("find business page list err", err.Error())
		return nil, status.Errorf(100001, "商家列表查询失败")
	}
	var list []*pb.Business
	for _, business := range *reply {
		var pbBusiness pb.Business
		_ = copier.Copy(&pbBusiness, &business)
		pbBusiness.UpdatedAt = business.UpdatedAt.Unix()
		pbBusiness.CreatedAt = business.CreatedAt.Unix()
		list = append(list, &pbBusiness)
	}

	return &pb.BusinessPageReply{
		Business: list,
		Total:    total,
		PageSize: in.GetPageSize(),
		Page:     in.GetPage(),
	}, nil
}
