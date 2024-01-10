package model

import "gorm.io/gorm"

type LogModel interface {
	logModel
}

type customLogModel struct {
	*defaultLogModel
}

func NewLogModel(db *gorm.DB) *customLogModel {
	return &customLogModel{
		defaultLogModel: newLogModel(db),
	}
}
