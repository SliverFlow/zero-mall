package model

import "gorm.io/gorm"

type (
	CartModel interface {
		cartModel
	}

	customCartModel struct {
		*defaultCartModel
	}
)

func NewCartModel(db *gorm.DB) CartModel {
	return customCartModel{
		defaultCartModel: newCartModel(db),
	}
}
