package model

import "gorm.io/gorm"

// MenuModel 菜单模型 所有的数据库操作将挂载在 MenuModel 上
type (
	MenuModel interface {
		menuModel
	}
	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel 实例化 menu model
func NewMenuModel(db *gorm.DB) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(db),
	}
}
