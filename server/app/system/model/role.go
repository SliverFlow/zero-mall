package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	roleModel interface {
		RoleCreate(ctx context.Context, role *Role) (err error)
	}
	defaultRoleModel struct {
		db *gorm.DB
	}

	Role struct {
		common.GloModel
		Name   string `json:"name" gorm:"not null;default:'';comment:角色名称"`
		Status int64  `json:"status" gorm:"not null;default:0;comment:角色状态"`
	}
)

func (r *Role) TableName() string {
	return "sys_role"
}

func newRoleModel(db *gorm.DB) *defaultRoleModel {
	return &defaultRoleModel{db}
}

func (d *defaultRoleModel) RoleCreate(ctx context.Context, role *Role) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&Role{}).Create(role).Error
}
