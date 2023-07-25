package auto

import "server/common"

// User
// Author [SliverFlow]
// @desc 业务用户表
type User struct {
	common.GloModel
	UserID   string `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
	UUID     string `json:"uuid" gorm:"not null;default:'';comment:用户uuid"`
	Username string `json:"username" gorm:"not null;default:'';comment:用户登录名"`
	Email    string `json:"email" gorm:"not null;default:'';comment:用户邮箱"`
	Nickname string `json:"nickname" gorm:"not null;default:'';comment:用户显示昵称"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户登录密码"`
	Avatar   string `json:"avatar" gorm:"not null;default:'';comment:用户头像"`
	Type     int64  `json:"type" gorm:"not null;default:0;comment:用户类型 0 普通用户 1 商家 2 系统管理员"`
	Status   int64  `json:"status" gorm:"not null;default:0;comment:用户状态 0 开启 1 禁用 "`
}

func (u *User) TableName() string {
	return "user"
}
