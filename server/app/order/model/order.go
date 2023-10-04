package model

import (
	"gorm.io/gorm"
	"server/common"
)

type (
	orderModel interface {
	}

	defaultOrderModel struct {
		db *gorm.DB
	}

	// Order
	// Author [SliverFlow]
	// @desc 业务订单表
	Order struct {
		common.GloModel
		OrderID     string  `json:"orderId" gorm:"not null;default:'';comment:订单业务id"`
		UserID      string  `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
		ShoppingID  string  `json:"shoppingId" gorm:"not null;default:'';comment:收获信息业务id"`
		Payment     float64 `json:"payment" gorm:"not null;default:0;comment:实际付款金额"`
		PaymentType int64   `json:"paymentType" gorm:"not null;default:1;comment:支付类型,1-在线支付"`
		Postage     int64   `json:"postage" gorm:"not null;default:0;comment:运费 单位是 元"`
		Status      int64   `json:"status" gorm:"not null;default:10;comment:订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭"`
		PaymentTime int64   `json:"paymentTime" gorm:"not null;default:0;comment:支付时间"`
		SendTime    int64   `json:"sendTime" gorm:"not null;default:0;comment:发货时间"`
		EndTime     int64   `json:"endTime" gorm:"not null;default:0;comment:交易完成时间"`
		CloseTime   int64   `json:"closeTime" gorm:"not null;default:0;comment:交易关闭时间"`
	}
)

func (o *Order) TableName() string {
	return "order"
}

func newOrderModel(db *gorm.DB) *defaultOrderModel {
	return &defaultOrderModel{db: db}
}
