package auto

import "server/common"

type User struct {
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

func (su *User) TableName() string {
	return "sys_user"
}
