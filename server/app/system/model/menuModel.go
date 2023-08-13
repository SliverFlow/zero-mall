package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
		customMenuLogicModel
	}

	customMenuModel struct {
		*defaultMenuModel
	}

	customMenuLogicModel interface {
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(conn *gorm.DB, c cache.CacheConf) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(conn, c),
	}
}

func (m *defaultMenuModel) customCacheKeys(data *Menu) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
