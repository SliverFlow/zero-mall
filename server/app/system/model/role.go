package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	roleModel interface {
	}
	defaultRoleModel struct {
		db *gorm.DB
	}

	Role struct {
		common.GloModel
	}
)

func (r *Role) TableName() string {
	return "sys_role"
}

func newRoleModel(db *gorm.DB) *defaultRoleModel {
	return &defaultRoleModel{db}
}
