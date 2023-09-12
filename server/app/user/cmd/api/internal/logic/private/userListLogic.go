package private

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

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

func (l *UserListLogic) UserList(req *types.PageReq) (resp *types.ListReply, err error) {
	reply, err := l.svcCtx.UserRpc.UserList(l.ctx, &pb.PageReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		KeyWork:  req.KeyWord,
	})
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}

	var list []types.User
	if len(list) == 0 {
		list = make([]types.User, 0)
	}
	for _, u := range reply.List {
		var user types.User
		_ = copier.Copy(&user, &u)
		list = append(list, user)
	}

	return &types.ListReply{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    reply.Total,
		List:     list,
	}, nil

	return
}
