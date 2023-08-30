package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"strconv"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pb.PageReq) (*pb.PageReply, error) {
	page, err := strconv.Atoi(strconv.FormatInt(in.Page, 10))
	pageSize, err := strconv.Atoi(strconv.FormatInt(in.PageSize, 10))
	if err != nil {
		return nil, status.Errorf(100001, "数据转换错误")
	}
	limit := pageSize
	offset := (page - 1) * pageSize
	keyWork := fmt.Sprintf("%%%s%%", in.KeyWork)
	total, list, err := l.svcCtx.UserModel.List(l.ctx, limit, offset, keyWork)
	if err != nil {
		return nil, status.Errorf(100001, "查询失败")
	}
	var enter []*pb.UserReply
	for _, user := range *list {
		enter = append(enter, &pb.UserReply{
			ID:        user.Id,
			UUID:      user.Uuid,
			Username:  user.Username,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Password:  user.Password,
			Avatar:    user.Avatar,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Time.Unix(),
			UpdatedAt: user.CreatedAt.Time.Unix(),
		})
	}

	return &pb.PageReply{
		List:  enter,
		Total: total,
	}, nil
}
