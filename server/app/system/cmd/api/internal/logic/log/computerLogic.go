package log

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"server/common/xerr"
	"time"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComputerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComputerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComputerLogic {
	return &ComputerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComputerLogic) Computer(req *types.NilReq) (resp *types.ComputerReply, err error) {

	info, err := cpu.Info()
	if err != nil {
		return nil, xerr.NewErrMsg("获取计算机cpu信息失败")
	}
	percent, _ := cpu.Percent(time.Second, false)
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, xerr.NewErrMsg("获取计算机内存信息失败")
	}

	resp = &types.ComputerReply{
		Cores:       int64(info[0].Cores),
		ModelName:   info[0].ModelName,
		VendorID:    info[0].VendorID,
		Family:      info[0].Family,
		Stepping:    fmt.Sprintf("%d", info[0].Stepping),
		CacheSize:   fmt.Sprintf("%.2fMB", float64(info[0].CacheSize)/1024/1024),
		CpuUsage:    fmt.Sprintf("%.2f", percent[0]),
		MemSize:     fmt.Sprintf("%.2fGB", float64(memory.Total)/1024/1024/1024),
		Available:   fmt.Sprintf("%.2fGB", float64(memory.Available/1024/1024/1024)),
		Used:        fmt.Sprintf("%.2fGB", float64(memory.Used/1024/1024/1024)),
		UsedPercent: fmt.Sprintf("%.2f", memory.UsedPercent),
	}

	return resp, nil
}
