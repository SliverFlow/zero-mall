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
