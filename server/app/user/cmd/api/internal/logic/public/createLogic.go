package public

import (
	"context"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateReq) (resp *types.NilReply, err error) {
	var createIn pb.UserCreateReq
	_ = copier.Copy(&createIn, req)
	_, err = l.svcCtx.UserRpc.UserCreate(l.ctx, &createIn)
	if err != nil {
		return nil, err
	}
	return
}
