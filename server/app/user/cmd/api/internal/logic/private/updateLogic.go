package private

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Update 更新用户信息
func (l *UpdateLogic) Update(req *types.UpdateReq) (resp *types.NilReply, err error) {
	var u pb.UpdateReq
	_ = copier.Copy(&u, req)
	_, err = l.svcCtx.UserRpc.Update(l.ctx, &u)
	if err != nil {
		return nil, err
	}
	return
}
