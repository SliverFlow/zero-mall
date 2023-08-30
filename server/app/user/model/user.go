package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	userModelGorm interface { // 数据库操作接口
		Create(u *UserGorm) error
	}

	defaultUserModelGorm struct {
		db *gorm.DB
	}

	UserGorm struct {
		common.GloModel
		UUID     string `json:"uuid" gorm:"not null;default:'';comment:用户uuid"`
		Username string `json:"username" gorm:"not null;default:'';comment:用户登录名"`
		Email    string `json:"email" gorm:"not null;default:'';comment:用户邮箱"`
		Nickname string `json:"nickname" gorm:"not null;default:'';comment:用户显示昵称"`
		Password string `json:"password" gorm:"not null;default:'';comment:用户登录密码"`
		Avatar   string `json:"avatar" gorm:"not null;default:'';comment:用户头像"`
		Role     int64  `json:"role" gorm:"not null;default:0;comment:用户类型 0 普通用户 2 商家 1 系统管理员"`
		Status   int64  `json:"status" gorm:"not null;default:0;comment:用户状态 0 开启 1 禁用 "`
	}
)

func (u *UserGorm) TableName() string {
	return "user"
}

func newUserModelGorm(db *gorm.DB) *defaultUserModelGorm {
	return &defaultUserModelGorm{db}
}

func (d *defaultUserModelGorm) Create(u *UserGorm) error {
	if err := d.db.Model(&UserGorm{}).Create(u).Error; err != nil {
		return err
	}
	return nil
}
