package auto

import "server/common"

// Category
// Author [SliverFlow]
// @desc 业务分类表
type Category struct {
	common.GloModel
	CategoryID string `json:"categoryId" gorm:"not null;default:'';comment:商品分类id"`
	ParentId   string `json:"parentId" gorm:"not null;default:'';comment:父分类id 当id=0时说明是根节点,一级类别 "`
	Name       string `json:"name" gorm:"not null;default:'';comment:分类名"`
	Status     int64  `json:"status" gorm:"not null;default:0;comment:类别状态 0-暂存 1-正常,2-已废弃"`
}
