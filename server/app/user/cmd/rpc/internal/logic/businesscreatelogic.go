package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"server/app/user/model"
	"server/common"
	"strconv"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessCreateLogic {
	return &BusinessCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessCreate
// Author [SliverFlow]
// @desc 创建商家
func (l *BusinessCreateLogic) BusinessCreate(in *pb.BusinessCreateReq) (*pb.UserNil, error) {
	user, err := l.svcCtx.UserModel.UserFindByUUID(l.ctx, in.GetUUID())
	if err != nil {
		logx.Error("find business admin user err", err.Error())
		return nil, status.Errorf(100001, "未查找到商家管理员")
	}

	var bu model.Business
	_ = copier.Copy(&bu, in.GetBusiness())
	bu.UUID = user.UUID
	bu.BusinessID = strconv.FormatInt(common.SnowWorker.NextId(), 10)

	err = l.svcCtx.UserModel.BusinessCreate(l.ctx, &bu)
	if err != nil {
		logx.Error("create business err", err.Error())
		return nil, status.Errorf(100001, "创建商家失败")
	}

	return &pb.UserNil{}, nil
}
