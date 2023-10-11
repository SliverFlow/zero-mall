package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	orderItemModel interface {
		OrderItemCreate(ctx context.Context, orderItem *OrderItem) (err error)
		OrderItemDelete(ctx context.Context, orderItemID string) (err error)
	}

	defaultOrderItemModel struct {
		db *gorm.DB
	}

	// OrderItem
	// Author [SliverFlow]
	// @desc 业务订单子表
	OrderItem struct {
		common.GloModel
		OrderItemID      string  `json:"orderItemId" gorm:"not null;default:'';comment:订单子表业务id"`
		OrderID          string  `json:"orderId" gorm:"not null;default:'';comment:订单业务id"`
		UserID           string  `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
		ProductID        string  `json:"productId" gorm:"not null;default:'';comment:商品业务 id"`
		ProdName         string  `json:"prodName" gorm:"not null;default:'';comment:商品名称"`
		ProdImage        string  `json:"prodImage" gorm:"not null;default:'';comment:商品图片地址"`
		CurrentunitPrice float64 `json:"currentunitPrice" gorm:"type:decimal(10,2);not null;default:0;comment:生成订单时的商品单价，单位是元,保留两位小数"`
		Quantity         int64   `json:"quantity" gorm:"not null;default:0;comment:商品数量"`
		TotalPrice       float64 `json:"price" gorm:"not null;type:decimal(10,2);default:0;comment:商品总价,单位是元,保留两位小数"`
	}
)

func (ot *OrderItem) TableName() string {
	return "order_item"
}

func newOrderItemModel(db *gorm.DB) *defaultOrderItemModel {
	return &defaultOrderItemModel{db}
}

// OrderItemCreate 创建订单字数据
func (d *defaultOrderItemModel) OrderItemCreate(ctx context.Context, orderItem *OrderItem) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "order_item-create")
	defer span.End()

	return tx.Model(&OrderItem{}).Create(orderItem).Error
}

// OrderItemDelete 删除订单详情
func (d *defaultOrderItemModel) OrderItemDelete(ctx context.Context, orderItemId string) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "order_item-create")
	defer span.End()

	return tx.Where("order_item_id = ?", orderItemId).Delete(&OrderItem{}).Error
}
