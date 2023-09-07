package user

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.PageReq) (resp *types.UserListReply, err error) {

	rep, err := l.svcCtx.UserRpc.UserList(l.ctx, &pb.PageReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		KeyWork:  req.KeyWord,
	})
	if err != nil {
		return nil, err
	}

	enter := make([]types.User, 0)
	for _, userInfo := range rep.List {
		var user types.User
		_ = copier.Copy(&user, userInfo)
		enter = append(enter, user)
	}

	return &types.UserListReply{
		User:     enter,
		Total:    rep.Total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
