package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	// 数据库操作方法接口
	menuModel interface {
		MenuCreate(ctx context.Context, menu *Menu) (err error)
		MenuListByRole(ctx context.Context, role int64) (list *[]Menu, err error)
		MenuChangeStatus(ctx context.Context, id int64, pid int64, status int64) (err error)
		MenuTreeListAllByRole(ctx context.Context, role int64) (list *[]Menu, err error)
		MenuUpdate(ctx context.Context, menu *Menu) (err error)
		MenuFind(ctx context.Context, id int64) (enter *Menu, err error)
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
		Role      int64                                      `json:"role" gorm:"not null;default:0;comment:角色 1 普通用户 2 系统商家 3 系统管理员"`
		Meta      `json:"meta" gorm:"embedded;comment:附加属性"` // embedded gorm tag 嵌套字段
		Status    int64                                      `json:"status" gorm:"not null;default:0;comment:菜单状态"`
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

func (d *defaultMenuModel) MenuCreate(ctx context.Context, menu *Menu) (err error) {
	tx := d.db.WithContext(ctx)
	return tx.Model(&menu).Create(menu).Error
}

// MenuListByRole
// Author [SliverFlow]
// @desc  根据菜单角色获取菜单（状态为激活）
func (d *defaultMenuModel) MenuListByRole(ctx context.Context, role int64) (list *[]Menu, err error) {
	tx := d.db.WithContext(ctx)

	var enter []Menu
	if role == 10 {
		err = tx.Model(&Menu{}).Where("status = ?", 1).Order("sorted asc").Find(&enter).Error
	} else {
		err = tx.Model(&Menu{}).Where("role = ? and status = ?", role, 1).Order("sorted asc").Find(&enter).Error
	}

	if err != nil {
		return nil, err
	}
	return &enter, nil
}

// MenuChangeStatus 修改菜单状态
func (d *defaultMenuModel) MenuChangeStatus(ctx context.Context, id int64, pid int64, status int64) (err error) {
	tx := d.db.WithContext(ctx)

	var ids []int64
	ids = append(ids, id)
	if pid == 0 && status == 0 { // 根菜单 将所有的子菜单全部修改
		var list []Menu
		if err = tx.Model(&Menu{}).Where("parent_id = ?", id).Find(&list).Error; err != nil {
			return err
		}
		for _, menu := range list {
			ids = append(ids, int64(menu.ID))
		}
	}
	if pid != 0 && status == 1 {
		ids = append(ids, pid)
	}
	return tx.Model(&Menu{}).Where("id in ?", ids).Update("status", status).Error
}

// MenuTreeListAllByRole
// Author [SliverFlow]
// @desc  查询某角色下的所有菜单
func (d *defaultMenuModel) MenuTreeListAllByRole(ctx context.Context, role int64) (list *[]Menu, err error) {
	tx := d.db.WithContext(ctx)

	var enter []Menu
	if role == 10 {
		err = tx.Model(&Menu{}).Order("sorted asc").Find(&enter).Error
	} else {
		err = tx.Model(&Menu{}).Where("role = ?", role).Order("sorted asc").Find(&enter).Error
	}

	if err != nil {
		return nil, err
	}
	return &enter, nil
}

// MenuUpdate
// Author [SliverFlow]
// @desc 更新菜单
func (d *defaultMenuModel) MenuUpdate(ctx context.Context, menu *Menu) (err error) {
	tx := d.db.WithContext(ctx)

	mMap := make(map[string]any)
	mMap["name"] = menu.Name
	mMap["title"] = menu.Title
	mMap["component"] = menu.Component
	mMap["icon"] = menu.Icon
	mMap["status"] = menu.Status
	mMap["role"] = menu.Role
	mMap["sorted"] = menu.Sorted
	mMap["path"] = menu.Path
	mMap["parent_id"] = menu.ParentId
	return tx.Model(&Menu{}).Where("id = ?", menu.ID).Updates(mMap).Error
}

// MenuFind
// Author [SliverFlow]
// @desc 根据 id 查询菜单
func (d *defaultMenuModel) MenuFind(ctx context.Context, id int64) (enter *Menu, err error) {
	tx := d.db.WithContext(ctx)

	var m Menu
	if err = tx.Model(&Menu{}).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
