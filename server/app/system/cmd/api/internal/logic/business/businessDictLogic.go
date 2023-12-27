package business

import (
	"context"
	"server/app/user/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessDictLogic {
	return &BusinessDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessDictLogic) BusinessDict(req *types.NilReq) (resp *types.BusinessDictReply, err error) {
	dict, err := l.svcCtx.UserRpc.BusinessDict(l.ctx, &pb.UserNil{})
	if err != nil {
		return nil, err
	}
	var list []types.BusinessDict
	for _, v := range dict.List {
		list = append(list, types.BusinessDict{
			Label: v.BusinessName,
			Value: v.BusinessID,
		})
	}

	return &types.BusinessDictReply{
		List: list,
	}, nil
}
