package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFindListByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFindListByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFindListByIDsLogic {
	return &UserFindListByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFindListByIDsLogic) UserFindListByIDs(in *pb.UserPageListByIDsReq) (*pb.UserListReply, error) {
	users, err := l.svcCtx.UserModel.FindListByIDs(l.ctx, in.IDs)
	if err != nil {
		return nil, status.Errorf(100001, "用户列表查询失败")
	}

	var list []*pb.UserReply
	for _, v := range users {
		list = append(list, &pb.UserReply{
			UUID:     v.UUID,
			Username: v.Username,
			Phone:    v.Phone,
			Role:     v.Role,
			Status:   v.Status,
			Nickname: v.Nickname,
		})
	}

	return &pb.UserListReply{
		List: list,
	}, nil
}
