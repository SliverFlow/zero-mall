package public

import (
	"context"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhealthLogic {
	return &PhealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhealthLogic) Phealth(req *types.NilReq) (resp *types.HealthReply, err error) {
	// todo: add your logic here and delete this line

	return
}
