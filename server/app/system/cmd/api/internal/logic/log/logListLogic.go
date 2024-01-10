package log

import (
	"context"
	"server/app/log/cmd/rpc/pb"
	"server/common"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogListLogic {
	return &LogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogListLogic) LogList(req *types.LogListReq) (resp *types.LogListReply, err error) {

	var endTime, startTime int64
	if req.StartTime != "" {
		t, _ := common.FormatTime(req.StartTime)
		startTime = t.Unix()
	}
	if req.EndTime != "" {
		t, _ := common.FormatTime(req.EndTime)
		endTime = t.Unix()
	}

	logPageList, err := l.svcCtx.LogRpc.LogPageList(l.ctx, &pb.LogPageListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Username:  req.Username,
		Method:    req.Method,
		Path:      req.Path,
		StartTime: startTime,
		EndTime:   endTime,
	})

	if err != nil {
		return
	}

	var list []types.LoggerReply
	for _, item := range logPageList.List {
		list = append(list, types.LoggerReply{
			ID:       item.ID,
			Username: item.Username,
			IP:       item.IP,
			Method:   item.Method,
			Path:     item.Path,
			Latency:  item.Latency,
			Time:     item.Time,
		})
	}

	return &types.LogListReply{
		List:  list,
		Total: logPageList.Total,
	}, nil
}
