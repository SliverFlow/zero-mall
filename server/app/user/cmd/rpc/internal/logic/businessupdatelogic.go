package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"
	"server/app/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessUpdateLogic {
	return &BusinessUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessUpdate
// Author [SliverFlow]
// @desc 商家更新信息
func (l *BusinessUpdateLogic) BusinessUpdate(in *pb.Business) (*pb.UserNil, error) {
	err := l.svcCtx.UserModel.BusinessUpdate(l.ctx, &model.Business{
		BusinessID: in.GetBusinessID(),
		UUID:       in.GetUUID(),
		Name:       in.GetName(),
		Detail:     in.GetDetail(),
		Score:      in.GetScore(),
		Image:      in.GetImage(),
		Status:     in.GetStatus(),
	})
	if err != nil {
		logx.Error("update business info err ", err.Error())
		return nil, status.Errorf(100001, "更新商户信息失败")
	}

	return &pb.UserNil{}, nil
}
