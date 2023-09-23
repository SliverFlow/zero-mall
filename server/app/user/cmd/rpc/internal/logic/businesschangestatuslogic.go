package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessChangeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessChangeStatusLogic {
	return &BusinessChangeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessChangeStatus
// Author [SliverFlow]
// @desc 管理员账户切换角色
func (l *BusinessChangeStatusLogic) BusinessChangeStatus(in *pb.BusinessChangeStatusReq) (*pb.UserNil, error) {
	err := l.svcCtx.UserModel.BusinessChangeStatus(l.ctx, in.GetBusinessID(), in.GetStatus())
	if err != nil {
		return nil, status.Errorf(100001, "更新商家状态失败")
	}

	return &pb.UserNil{}, nil
}
