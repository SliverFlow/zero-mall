package common

import (
	"gorm.io/gorm"
	"time"
)

// GloModel
// 全局公用的 model
type GloModel struct {
	ID        uint           `json:"ID" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}

type (
	PageInfo struct {
		Page     int
		PageSize int
		KeyWord  string
	}
)

// LimitAndOffset 获取分页以及偏移量
func (pi *PageInfo) LimitAndOffset() (limit, offset int) {
	return pi.Page, (pi.Page - 1) * pi.PageSize
}
