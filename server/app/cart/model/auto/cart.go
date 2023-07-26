package auto

import "server/common"

// Cart
// Author [SliverFlow]
// @desc 业务购物车表
type Cart struct {
	common.GloModel
	CartID    string `json:"cartId" gorm:"not null;default:'';comment:购物车业务id"`
	UserID    string `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
	ProductID string `json:"productId" gorm:"not null;default:'';comment:商品业务 id "`
	Quantity  int64  `json:"quantity" gorm:"not null;default:0;comment:数量"`
	Checked   int64  `json:"checked" gorm:"not null;default:0;comment:是否选择,1=已勾选,0=未勾选"`
}
