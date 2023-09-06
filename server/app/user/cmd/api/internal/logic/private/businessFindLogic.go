package private

import (
	"context"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessFindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessFindLogic {
	return &BusinessFindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessFindLogic) BusinessFind(req *types.Nil) (resp *types.Nil, err error) {
	// todo: add your logic here and delete this line

	return
}
