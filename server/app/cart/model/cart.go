package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	cartModel interface {
	}

	defaultCartModel struct {
		db *gorm.DB
	}

	// Cart 购物车
	Cart struct {
		common.GloModel
		CartID    string `json:"cartId" gorm:"not null;default:'';comment:购物车业务id"`
		UserID    string `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
		ProductID string `json:"productId" gorm:"not null;default:'';comment:商品业务 id "`
		Quantity  int64  `json:"quantity" gorm:"not null;default:0;comment:数量"`
		Checked   int64  `json:"checked" gorm:"not null;default:0;comment:是否选择,1=已勾选,0=未勾选"`
	}
)

func (c *Cart) TableName() string {
	return "cart"
}

func newCartModel(db *gorm.DB) *defaultCartModel {
	return &defaultCartModel{db: db}
}
