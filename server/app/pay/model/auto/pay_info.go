package auto

import "server/common"

// PayInfo
// Author [SliverFlow]
// @desc 业务支付信息表
type PayInfo struct {
	common.GloModel
	PayInfoID  string `json:"payInfoId" gorm:"not null;default:'';comment:支付信息业务id"`
	OrderID    string `json:"orderId" gorm:"not null;default:'';comment:订单业务id"`
	UserID     string `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
	PlatForm   int64  `json:"platForm" gorm:"not null;default:0;comment:支付平台:1-支付宝,2-微信"`
	PlatNumber string `json:"platNumber" gorm:"not null;default:0;comment:支付流水号"`
	PlatStatus string `json:"platStatus" gorm:"not null;default:0;comment:支付状态"`
}
