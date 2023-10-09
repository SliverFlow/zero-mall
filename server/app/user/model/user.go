package model

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"server/common"
	"server/common/xerr"
)

type (
	userModel interface { // 数据库操作接口
		UserCreate(ctx context.Context, u *User) error
		UserDelete(ctx context.Context, id int64) (err error)
		UserUpdate(ctx context.Context, u *User) (err error)
		UserFind(ctx context.Context, id int64) (enter *User, err error)
		UserList(ctx context.Context, page *common.PageInfo) (enter *[]User, total int64, err error)

		UserFindByUsername(ctx context.Context, username string) (enter *User, err error)
		UserFindByUUID(ctx context.Context, uuid string) (enter *User, err error)
		AdminChangeRole(ctx context.Context, username string, role int64) (err error)
		CheckUsername(ctx context.Context, username string) (ok bool, err error)
		UserChangeStatus(ctx context.Context, id int64, status int64) (err error)
		UserFindByPhoneOrUsername(ctx context.Context, phone string, username string) (enter *User, err error)
		CheckPhone(ctx context.Context, phone string) (ok bool, err error)
	}

	defaultUserModel struct {
		db *gorm.DB
	}

	User struct {
		common.GloModel
		UUID     string `json:"uuid" gorm:"not null;default:'';comment:用户uuid"`
		Username string `json:"username" gorm:"not null;default:'';comment:用户登录名"`
		Email    string `json:"email" gorm:"not null;default:'';comment:用户邮箱"`
		Nickname string `json:"nickname" gorm:"not null;default:'';comment:用户显示昵称"`
		Password string `json:"password" gorm:"not null;default:'';comment:用户登录密码"`
		Avatar   string `json:"avatar" gorm:"not null;default:'';comment:用户头像"`
		Role     int64  `json:"role" gorm:"not null;default:1;comment:用户类型 1 普通用户 2 商家 3 系统管理员"`
		Status   int64  `json:"status" gorm:"not null;default:1;comment:用户状态 1 开启 0 禁用 "`
		Phone    string `json:"phone"  gorm:"not null;default:'';comment:手机号码"`
	}
)

func (u *User) TableName() string {
	return "user"
}

func newUserModel(db *gorm.DB) *defaultUserModel {
	return &defaultUserModel{db}
}

func (d *defaultUserModel) UserCreate(ctx context.Context, u *User) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&User{}).Create(u).Error
}

func (d *defaultUserModel) UserDelete(ctx context.Context, id int64) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Where("id = ?", id).Delete(&User{}).Error
}

func (d *defaultUserModel) UserUpdate(ctx context.Context, u *User) (err error) {
	tx := d.db.WithContext(ctx)

	um := make(map[string]any)
	um["username"] = u.Username
	um["email"] = u.Email
	um["role"] = u.Role
	um["avatar"] = u.Avatar
	um["status"] = u.Status
	um["nickname"] = u.Nickname
	um["password"] = common.BcryptHash(u.Password)

	return tx.Model(&User{}).Where("id = ?", u.ID).Updates(&um).Error
}

func (d *defaultUserModel) UserFind(ctx context.Context, id int64) (enter *User, err error) {
	tx := d.db.WithContext(ctx)
	span, carrier := common.Tracer(ctx, "find-user")
	carrier.Set("name", "gorm")
	defer span.End()

	var u User
	if err = tx.Model(&User{}).Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *defaultUserModel) UserList(ctx context.Context, page *common.PageInfo) (enter *[]User, total int64, err error) {
	tx := d.db.WithContext(ctx)

	limit, offset, keyWord := page.LimitAndOffsetAndKeyWord()

	tx = tx.Model(&User{}).Where("nickname like ?", keyWord)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []User
	if err = tx.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}

func (d *defaultUserModel) UserFindByUsername(ctx context.Context, username string) (enter *User, err error) {
	tx := d.db.WithContext(ctx)
	var u User
	if err = tx.Model(&User{}).Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *defaultUserModel) UserFindByUUID(ctx context.Context, uuid string) (enter *User, err error) {
	tx := d.db.WithContext(ctx)
	var u User
	if err = tx.Model(&User{}).Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// AdminChangeRole 管理员方便测试 修改角色
func (d *defaultUserModel) AdminChangeRole(ctx context.Context, username string, role int64) (err error) {
	tx := d.db.WithContext(ctx)

	if username != "admin" {
		return xerr.NewErrMsg("此用户不能修改角色")
	}
	return tx.Model(&User{}).Where("username = ?", username).Update("role", role).Error
}

// CheckUsername 验证用户名是否已经存在
func (d *defaultUserModel) CheckUsername(ctx context.Context, username string) (ok bool, err error) {
	tx := d.db.WithContext(ctx)
	var user User
	if err = tx.Model(&User{}).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

// UserChangeStatus 更新用户状态
func (d *defaultUserModel) UserChangeStatus(ctx context.Context, id int64, status int64) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&User{}).Where("id = ?", id).Update("status", status).Error
}

// UserFindByPhoneOrUsername 通过用户电话号码或者用户名查找用户
func (d *defaultUserModel) UserFindByPhoneOrUsername(ctx context.Context, phone string, username string) (enter *User, err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "find-user-by-phone-or-username")
	defer span.End()

	var u User
	if err = tx.Model(&User{}).Where("username = ? or phone = ?", username, phone).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

// CheckPhone 验证手机号是否已经存在
func (d *defaultUserModel) CheckPhone(ctx context.Context, phone string) (ok bool, err error) {
	tx := d.db.WithContext(ctx)
	var user User
	if err = tx.Model(&User{}).Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
