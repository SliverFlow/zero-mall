package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	categoryModel interface {
	}

	defaultCategoryModel struct {
		db *gorm.DB
	}

	// Category
	// Author [SliverFlow]
	// @desc 业务分类表
	Category struct {
		common.GloModel
		CategoryID string `json:"categoryId" gorm:"not null;default:'';comment:商品分类id"`
		ParentId   string `json:"parentId" gorm:"not null;default:'';comment:父分类id 当id=0时说明是根节点,一级类别 "`
		Name       string `json:"name" gorm:"not null;default:'';comment:分类名"`
		Status     int64  `json:"status" gorm:"not null;default:0;comment:类别状态 0-暂存 1-正常,2-已废弃"`
		Type       int64  `json:"type" gorm:"not null;default:0;comment:分类类别 0 系统类别 1 商家添加的类别"`
		BusinessID string `json:"businessId" gorm:"not null;default:'';comment:商户 id 只有为商家类别的时候才拥有值"`
	}
)

func (c *Category) TableName() string {
	return "category"
}

func newDefaultCategoryModel(db *gorm.DB) *defaultCategoryModel {
	return &defaultCategoryModel{
		db: db,
	}
}