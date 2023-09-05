package menu

import (
	"context"
	"fmt"
	"server/app/system/cmd/rpc/pb"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuChangeStatusLogic {
	return &MenuChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuChangeStatusLogic) MenuChangeStatus(req *types.MenuChangeStatusReq) (resp *types.NilReply, err error) {
	fmt.Println("req", req)
	_, err = l.svcCtx.SystemRpc.MenuChangeStatus(l.ctx, &pb.MenuChangeStatusReq{
		ID:     req.ID,
		PID:    req.PID,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return
}
