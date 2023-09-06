package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	businessModel interface {
	}

	defaultBusinessModel struct {
		db *gorm.DB
	}

	// Business
	// Author [SliverFlow]
	// @desc 商家表
	Business struct {
		common.GloModel
		BusinessID string `json:"businessId" gorm:"not null;default:'';comment:商户 id"`
		UUID       string `json:"uuid" gorm:"not null;default:'';comment:关联用户已存在的 user "`
		Name       string `json:"name" gorm:"not null;default:'';comment:商家名称"`
		Detail     string `json:"detail" gorm:"type:text;not null;comment:商家详情"`
		Score      int64  `json:"score" gorm:"not null;default:0;comment:商家评分"`
		Image      string `json:"image" gorm:"type:text;not null;comment:图片地址 json 格式"`
		Status     int64  `json:"status" gorm:"not null;default:0;comment:商品状态.0-暂存 1-正常 2-禁用"`
	}
)

func (b *Business) TableName() string {
	return "business"
}

func newDefaultBusinessModel(db *gorm.DB) *defaultBusinessModel {
	return &defaultBusinessModel{db: db}
}
