package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/common"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *pb.PageReq) (*pb.PageReply, error) {
	list, total, err := l.svcCtx.UserModel.UserList(l.ctx, &common.PageInfo{
		Page:     int(in.Page),
		PageSize: int(in.PageSize),
		KeyWord:  in.KeyWork,
	})
	if err != nil {
		return nil, status.Errorf(100001, "查询失败")
	}

	var enter []*pb.UserReply
	for _, user := range *list {
		u := copyUserFoReply(&user)
		enter = append(enter, u)
	}

	return &pb.PageReply{
		List:  enter,
		Total: total,
	}, nil
}
