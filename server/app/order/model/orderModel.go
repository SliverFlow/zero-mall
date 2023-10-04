package model

import "gorm.io/gorm"

type (
	OrderModel interface {
		orderModel
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

func NewOrderModel(db *gorm.DB) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(db),
	}
}
