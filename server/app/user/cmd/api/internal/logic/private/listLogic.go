package private

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.PageReq) (resp *types.ListReply, err error) {
	// todo: add your logic here and delete this line
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
}
