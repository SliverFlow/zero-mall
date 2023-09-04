package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	// 数据库操作方法接口
	menuModel interface {
	}
	// 默认数据库连接对象
	defaultMenuModel struct {
		db *gorm.DB
	}
	// Menu menu orm 映射
	Menu struct {
		common.GloModel
		ParentId  uint                                       `json:"parentId" gorm:"not null;default:0;comment:父菜单 id"`
		Name      string                                     `json:"name" gorm:"not null;default:'';comment:路由name"` // 菜单显示名字
		Path      string                                     `json:"path" gorm:"not null;default'';comment:路由path"`
		Component string                                     `json:"component" gorm:"not null;default:'';comment:组件位置对应的前端路径"`
		Sorted    int64                                      `json:"sorted" gorm:"not null;default:0;comment:排序标记"`
		Meta      `json:"meta" gorm:"embedded;comment:附加属性"` // embedded gorm tag 嵌套字段
	}
	Meta struct {
		Title string `json:"title" gorm:"not null;default:'';comment:菜单名"` // 菜单名
		Icon  string `json:"icon" gorm:"not null;default:'';comment:菜单图标"` // 菜单图标
	}
)

func (m *Menu) TableName() string {
	return "sys_menu"
}

func newMenuModel(db *gorm.DB) *defaultMenuModel {
	return &defaultMenuModel{db}
}
