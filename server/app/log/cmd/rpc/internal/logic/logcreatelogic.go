package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/log/cmd/rpc/internal/svc"
	"server/app/log/cmd/rpc/pb"
	"server/app/log/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogCreateLogic {
	return &LogCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogCreateLogic) LogCreate(in *pb.LogCreateReq) (*pb.LogNil, error) {
	err := l.svcCtx.LogModel.Create(l.ctx, &model.Log{
		UserID:         in.UserID,
		Username:       in.Username,
		Ip:             in.IP,
		Path:           in.Path,
		Method:         in.Method,
		RequestParams:  in.RequestParams,
		Status:         in.StatusCode,
		ResponseParams: in.ResponseParams,
		Latency:        in.Latency,
	})
	if err != nil {
		return nil, status.Errorf(100001, "创建日志失败")
	}

	return &pb.LogNil{}, nil
}
