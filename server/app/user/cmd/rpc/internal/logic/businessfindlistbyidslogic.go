package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessFindListByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessFindListByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessFindListByIDsLogic {
	return &BusinessFindListByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BusinessFindListByIDsLogic) BusinessFindListByIDs(in *pb.BusinessFindListReq) (*pb.BusinessListReply, error) {
	enter, err := l.svcCtx.UserModel.BusinessFindListByIDs(l.ctx, in.IDs)
	if err != nil {
		return nil, status.Errorf(100001, "商家列表查询失败")
	}

	var list []*pb.Business
	for _, v := range enter {
		list = append(list, &pb.Business{
			BusinessID: v.BusinessID,
			Name:       v.Name,
			Detail:     v.Detail,
			Score:      v.Score,
			Image:      v.Image,
			Status:     v.Status,
		})
	}
	return &pb.BusinessListReply{
		List: list,
	}, nil
}
