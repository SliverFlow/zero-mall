package public

import (
	"context"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

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

func (l *BusinessCreateLogic) BusinessCreate(req *types.CreateBusinessReq) (resp *types.Nil, err error) {

	_, err = l.svcCtx.UserRpc.UserCreate(l.ctx, &pb.CreateReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("创建商家管理员失败")
	}

	user, _ := l.svcCtx.UserRpc.UserFindByUsername(l.ctx, &pb.UsernameReq{Username: req.Username})

	_, err = l.svcCtx.UserRpc.BusinessCreate(l.ctx, &pb.BusinessCreateReq{
		Business: &pb.Business{
			Name:   req.Name,
			Detail: req.Detail,
			Image:  req.Image,
		},
		UUID: user.User.UUID,
	})
	if err != nil {
		return nil, xerr.NewErrMsg("创建商家失败")
	}
	return
}
