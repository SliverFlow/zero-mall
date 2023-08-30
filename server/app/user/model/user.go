package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	userModelGorm interface { // 数据库操作接口
		Create(ctx context.Context, u *UserGorm) error
		Delete(ctx context.Context, id int64) (err error)
		Update(ctx context.Context, u *UserGorm) (err error)
		Find(ctx context.Context, id int64) (enter *UserGorm, err error)
		List(ctx context.Context, page common.PageInfo) (enter *[]UserGorm, total int64, err error)
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

func (d *defaultUserModelGorm) Create(ctx context.Context, u *UserGorm) (err error) {
	tx := d.db.WithContext(ctx)

	if err = tx.Model(&UserGorm{}).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (d *defaultUserModelGorm) Delete(ctx context.Context, id int64) (err error) {
	tx := d.db.WithContext(ctx)

	if err = tx.Where("id = ?", id).Delete(&UserGorm{}).Error; err != nil {
		return err
	}

	return nil
}

func (d *defaultUserModelGorm) Update(ctx context.Context, u *UserGorm) (err error) {
	tx := d.db.WithContext(ctx)

	var um map[string]any
	um["username"] = u.Username
	um["email"] = u.Email
	um["role"] = u.Role
	um["avatar"] = u.Avatar
	um["status"] = u.Status
	um["nickname"] = u.Nickname
	um["password"] = common.BcryptHash(u.Password)

	return tx.Model(&UserGorm{}).Where("id = ?", u.ID).Updates(&um).Error
}

func (d *defaultUserModelGorm) Find(ctx context.Context, id int64) (enter *UserGorm, err error) {
	tx := d.db.WithContext(ctx)

	var u UserGorm
	if err = tx.Model(&UserGorm{}).Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *defaultUserModelGorm) List(ctx context.Context, page common.PageInfo) (enter *[]UserGorm, total int64, err error) {
	tx := d.db.WithContext(ctx)

	tx = tx.Model(&UserGorm{}).Where("username = ?", page.KeyWord)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []UserGorm
	limit, offset := page.LimitAndOffset()
	if err = tx.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}
