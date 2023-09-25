package business

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

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

// BusinessFind
// Author [SliverFlow]
// @desc 查询商户信息
func (l *BusinessFindLogic) BusinessFind(req *types.NilReq) (resp *types.Business, err error) {
	uuid, err := xjwt.GetUserUUID(l.ctx)
	if err != nil {
		logx.Error("获取 uuid 失败")
		return nil, xerr.NewErrMsg("获取 uuid 失败")
	}

	pbreply, err := l.svcCtx.UserRpc.BusinessFindByUUID(l.ctx, &pb.BusinessUUIDReq{UUID: uuid})
	if err != nil {
		return nil, err
	}

	var tybusiness types.Business
	_ = copier.Copy(&tybusiness, pbreply)
	err = json.Unmarshal([]byte(pbreply.Image), &tybusiness.Image)
	if err != nil {
		return nil, xerr.NewErrMsg("序列化失败")
	}
	return &tybusiness, nil
}
