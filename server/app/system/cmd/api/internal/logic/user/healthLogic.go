package user

import (
	"context"
	"fmt"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health(req *types.NilReq) (resp *types.HealthReply, err error) {
	// todo: add your logic here and delete this line
	id, err := xjwt.GetUserID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户 id 失败")
	}
	fmt.Println(id)
	return
}
