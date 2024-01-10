package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"server/app/log/model"

	"server/app/log/cmd/rpc/internal/svc"
	"server/app/log/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogPageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogPageListLogic {
	return &LogPageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogPageListLogic) LogPageList(in *pb.LogPageListReq) (*pb.LogPageListReply, error) {
	list, total, err := l.svcCtx.LogModel.FindPageList(l.ctx, &model.LogPageReq{
		Page:      in.Page,
		PageSize:  in.GetPageSize(),
		Username:  in.Username,
		Method:    in.Method,
		Path:      in.Path,
		StartTime: in.StartTime,
		EndTime:   in.EndTime,
	})
	if err != nil {
		return nil, status.Errorf(100001, "日志列表查询失败")
	}

	var listReply []*pb.LogReply
	for _, v := range list {
		listReply = append(listReply, &pb.LogReply{
			ID:             int64(v.ID),
			Username:       v.Username,
			IP:             v.Ip,
			Method:         v.Method,
			Path:           v.Path,
			Latency:        v.Latency,
			ResponseParams: v.ResponseParams,
			Time:           v.CreatedAt.Unix(),
		})
	}

	return &pb.LogPageListReply{
		List:  listReply,
		Total: total,
	}, nil
}
