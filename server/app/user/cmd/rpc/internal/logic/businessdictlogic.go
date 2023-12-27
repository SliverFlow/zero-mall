package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessDictLogic {
	return &BusinessDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BusinessDictLogic) BusinessDict(in *pb.UserNil) (*pb.BusinessDictReply, error) {
	businesses, err := l.svcCtx.UserModel.BusinessListAll(l.ctx)
	if err != nil {
		return nil, status.Errorf(100001, "业务字典查询失败")
	}
	var list []*pb.BusinessDict
	for _, v := range businesses {
		list = append(list, &pb.BusinessDict{
			BusinessID:   v.BusinessID,
			BusinessName: v.Name,
		})
	}

	return &pb.BusinessDictReply{
		List: list,
	}, nil
}
