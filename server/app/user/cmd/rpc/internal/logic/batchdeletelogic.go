package logic

import (
	"context"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteLogic {
	return &BatchDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchDeleteLogic) BatchDelete(in *pb.IDsReq) (*pb.Nil, error) {
	// todo: add your logic here and delete this line

	return &pb.Nil{}, nil
}
