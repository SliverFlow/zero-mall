package common

import (
	"fmt"
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

type PageInfo struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	KeyWord   string `json:"keyWord"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Status    int    `json:"status"`
	Role      int    `json:"role"`
}

// LimitAndOffsetAndKeyWord 获取分页以及偏移量
func (pi *PageInfo) LimitAndOffsetAndKeyWord() (limit, offset int, keyWord string) {
	return pi.PageSize, (pi.Page - 1) * pi.PageSize, fmt.Sprintf("%%%s%%", pi.KeyWord)
}
