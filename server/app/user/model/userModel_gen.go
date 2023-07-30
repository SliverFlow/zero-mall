// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheApUserUserIdPrefix = "cache:apUser:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *User) error

		FindOne(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, tx *gorm.DB, data *User) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error

		// FindByUserID 自定义的查询方法
		FindByUserID(ctx context.Context, userId string) (*User, error)
		List(ctx context.Context, limit, offset int, keyWord string) (total int64, list *[]User, err error)
		UpdateByUserID(ctx context.Context, userId string, data *User) error
	}

	defaultUserModel struct {
		gormc.CachedConn
		table string
	}

	User struct {
		Id        int64          `gorm:"column:id"`
		CreatedAt sql.NullTime   `gorm:"column:created_at"`
		UpdatedAt sql.NullTime   `gorm:"column:updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
		UserId    string         `gorm:"column:user_id"`  // 用户业务id
		Uuid      string         `gorm:"column:uuid"`     // 用户uuid
		Username  string         `gorm:"column:username"` // 用户登录名
		Email     string         `gorm:"column:email"`    // 用户邮箱
		Nickname  string         `gorm:"column:nickname"` // 用户显示昵称
		Password  string         `gorm:"column:password"` // 用户登录密码
		Avatar    string         `gorm:"column:avatar"`   // 用户头像
		Type      int64          `gorm:"column:type"`     // 用户类型 0 普通用户 1 商家 2 系统管理员
		Status    int64          `gorm:"column:status"`   // 用户状态 0 开启 1 禁用
	}
)

func (User) TableName() string {
	return "`user`"
}

func newUserModel(conn *gorm.DB, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, tx *gorm.DB, data *User) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	apUserUserIdKey := fmt.Sprintf("%s%v", cacheApUserUserIdPrefix, id)
	var resp User
	err := m.QueryCtx(ctx, &resp, apUserUserIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Update 更新用户信息 （修改过）
func (m *defaultUserModel) Update(ctx context.Context, tx *gorm.DB, data *User) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultUserModel) getCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	apUserUserIdKey := fmt.Sprintf("%s%v", cacheApUserUserIdPrefix, data.Id)
	cacheKeys := []string{
		apUserUserIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultUserModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&User{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultUserModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheApUserUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&User{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultUserModel) tableName() string {
	return m.table
}

func (m *defaultUserModel) FindByUserID(ctx context.Context, userId string) (*User, error) {
	apUserUserIdKey := fmt.Sprintf("%s%v", cacheApUserUserIdPrefix, userId)
	var resp User
	err := m.QueryCtx(ctx, &resp, apUserUserIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("`user_id` = ?", userId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) List(ctx context.Context, limit, offset int, keyWord string) (total int64, list *[]User, err error) {
	var resp []User

	err = m.QueryNoCacheCtx(ctx, &total, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("nickname like ?", keyWord).Count(&total).Error
	})

	err = m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("nickname like ?", keyWord).Find(&resp).Limit(limit).Offset(offset).Error
	})
	switch err {
	case nil:
		return total, &resp, nil
	case gormc.ErrNotFound:
		return 0, nil, ErrNotFound
	default:
		return 0, nil, err
	}
}

// UpdateByUserID 通过 userId  更新用户信息
func (m *defaultUserModel) UpdateByUserID(ctx context.Context, userId string, data *User) error {
	_, err := m.FindByUserID(ctx, userId)
	if err != nil {
		return err
	}
	return m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Model(&User{}).Where("user_id = ?", userId).Updates(data).Error
	}, fmt.Sprintf("%s%v", cacheApUserUserIdPrefix, userId))
}