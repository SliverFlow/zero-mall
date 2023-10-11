package model

import "gorm.io/gorm"

type (
	OrderModel interface {
		orderModel
		orderItemModel
	}

	customOrderModel struct {
		*defaultOrderModel
		*defaultOrderItemModel
	}
)

func NewOrderModel(db *gorm.DB) OrderModel {
	return &customOrderModel{
		defaultOrderModel:     newOrderModel(db),
		defaultOrderItemModel: newOrderItemModel(db),
	}
}
