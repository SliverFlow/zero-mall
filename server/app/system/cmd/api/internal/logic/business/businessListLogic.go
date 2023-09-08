package business

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.PageReq) (resp *types.BusinessListReply, err error) {
	reply, err := l.svcCtx.UserRpc.BusinessList(l.ctx, &pb.PageReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		KeyWork:  req.KeyWord,
	})
	if err != nil {
		return nil, err
	}
	var list []types.Business
	for _, business := range reply.Business {
		var tyBusiness types.Business
		_ = copier.Copy(&tyBusiness, &business)

		err := json.Unmarshal([]byte(business.Image), &tyBusiness.Image)
		if err != nil {
			return nil, xerr.NewErrMsg("序列化失败")
		}
		list = append(list, tyBusiness)
	}

	return &types.BusinessListReply{
		List:     list,
		Total:    reply.Total,
		Page:     reply.Page,
		PageSize: reply.PageSize,
	}, nil
}
