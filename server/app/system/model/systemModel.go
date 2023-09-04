package model

import "gorm.io/gorm"

// MenuModel 菜单模型 所有的数据库操作将挂载在 MenuModel 上
type (
	SystemModel interface {
		menuModel
		roleModel
	}
	customMenuModel struct {
		*defaultMenuModel
		*defaultRoleModel
	}
)

// NewSystemModel 实例化 system model
func NewSystemModel(db *gorm.DB) SystemModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(db),
		defaultRoleModel: newRoleModel(db),
	}
}
