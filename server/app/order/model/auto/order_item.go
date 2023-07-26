package auto

import "server/common"

// OrderItem
// Author [SliverFlow]
// @desc 业务订单子表
type OrderItem struct {
	common.GloModel
	OrderItemID      string  `json:"orderItemId" gorm:"not null;default:'';comment:订单子表业务id"`
	OrderID          string  `json:"orderId" gorm:"not null;default:'';comment:订单业务id"`
	UserID           string  `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
	ProductID        string  `json:"productId" gorm:"not null;default:'';comment:商品业务 id"`
	ProdName         string  `json:"prodName" gorm:"not null;default:'';comment:商品名称"`
	ProdImage        string  `json:"prodImage" gorm:"not null;default:'';comment:商品图片地址"`
	CurrentunitPrice float64 `json:"currentunitPrice" gorm:"not null;default:0;comment:生成订单时的商品单价，单位是元,保留两位小数"`
	Quantity         int64   `json:"quantity" gorm:"not null;default:0;comment:商品数量"`
	TotalPrice       float64 `json:"price" gorm:"not null;default:0;comment:商品总价,单位是元,保留两位小数"`
}
