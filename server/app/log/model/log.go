package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
	"time"
)

type (
	logModel interface {
		Create(ctx context.Context, log *Log) error
		FindPageList(ctx context.Context, req *LogPageReq) ([]*Log, int64, error)
	}

	defaultLogModel struct {
		db *gorm.DB
	}

	Log struct {
		common.GloModel
		UserID         string `json:"userId" gorm:"not null;default:0;comment:用户id"`
		Username       string `json:"username" gorm:"not null;default:'';comment:用户名"`
		Ip             string `json:"ip" gorm:"not null;default:'';comment:ip"`
		Path           string `json:"path" gorm:"not null;default:'';comment:请求路径"`
		Method         string `json:"method" gorm:"not null;default:'';comment:请求方法"`
		RequestParams  string `json:"requestParams" gorm:"not null;default:'';comment:请求体"`
		Status         int64  `json:"status" gorm:"not null;default:0;comment:请求状态"`
		ResponseParams string `json:"responseParams" gorm:"not null;default:'';comment:响应体"`
		Latency        int64  `json:"latency" gorm:"not null;default:0;comment:延迟时间"`
	}
)

func (l *Log) TableName() string {
	return "sys_log"
}

func newLogModel(db *gorm.DB) *defaultLogModel {
	return &defaultLogModel{db: db}
}

func (d *defaultLogModel) Create(ctx context.Context, log *Log) error {
	span, _ := common.Tracer(ctx, "log-create")
	defer span.End()

	tx := d.db.WithContext(ctx)

	return tx.Create(log).Error
}

func (d *defaultLogModel) FindPageList(ctx context.Context, req *LogPageReq) ([]*Log, int64, error) {
	span, _ := common.Tracer(ctx, "log-findPageList")
	defer span.End()

	tx := d.db.WithContext(ctx)

	var list []*Log
	var total int64

	if req.Username != "" {
		tx = tx.Where("username like ?", "%"+req.Username+"%")
	}
	if req.Method != "" {
		tx = tx.Where("method = ?", req.Method)
	}
	if req.Path != "" {
		tx = tx.Where("path like ?", "%"+req.Path+"%")
	}

	if req.StartTime != 0 {
		tx = tx.Where("created_at >= ?", time.Unix(req.StartTime, 0))
	}
	if req.EndTime != 0 {
		tx = tx.Where("created_at <= ?", time.Unix(req.EndTime, 0))
	}

	err := tx.Model(&Log{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = tx.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
