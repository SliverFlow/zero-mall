package private

import (
	"context"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessCreateLogic {
	return &BusinessCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessCreateLogic) BusinessCreate(req *types.Nil) (resp *types.Nil, err error) {

	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		return nil, xerr.NewErrMsg("获取用户UUID失败")
	}

	_, err = l.svcCtx.UserRpc.BusinessCreate(l.ctx, &pb.BusinessCreateReq{
		UUID: uuid,
	})
	if err != nil {
		return nil, err
	}
	return
}
