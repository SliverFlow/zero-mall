package private

import (
	"context"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"strconv"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line
	if l.svcCtx.UserRpc == nil {
		return nil, xerr.NewCodeError(100001)
	}
	info, err := l.svcCtx.UserRpc.GetUserInfoById(l.ctx, &pb.UserInfoReq{ID: 1})
	if err != nil {
		return nil, xerr.NewCodeError(100001)
	}
	return &types.UserInfoReply{UserID: strconv.Itoa(int(info.ID))}, nil
}
