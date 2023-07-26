package auto

import "server/common"

// Shopping
// Author [SliverFlow]
// @desc 业务收货信息表
type Shopping struct {
	common.GloModel
	ShoppingID       string `json:"shoppingId" gorm:"not null;default:'';comment:收货信息业务id"`
	OrderID          string `json:"orderId" gorm:"not null;default:'';comment:订单业务id"`
	UserID           string `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
	ReceiverName     string `json:"receiverName" gorm:"not null;default:'';comment:收货姓名"`
	ReceiverPhone    string `json:"receiverPhone" gorm:"not null;default:'';comment:收货固定电话"`
	ReceiverMobile   string `json:"receiverMobile" gorm:"not null;default:'';comment:收货移动电话"`
	ReceiverProvince string `json:"receiverProvince" gorm:"not null;default:'';comment:省份"`
	ReceiverCity     string `json:"receiverCity" gorm:"not null;default:'';comment:城市"`
	ReceiverDistrict string `json:"receiverDistrict" gorm:"not null;default:'';comment:区/县"`
	ReceiverAddress  string `json:"receiverAddress" gorm:"not null;default:'';comment:详细地址"`
}
