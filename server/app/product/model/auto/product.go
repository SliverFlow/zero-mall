package auto

import "server/common"

// Product
// Author [SliverFlow]
// @desc 业务商品表
type Product struct {
	common.GloModel
	ProductID  string  `json:"productId" gorm:"not null;default:'';comment:商品业务 id"`
	Name       string  `json:"name" gorm:"not null;default:'';comment:商品名"`
	CategoryID string  `json:"categoryId" gorm:"not null;default:'';comment:商品分类id"`
	Subtitle   string  `json:"subtitle" gorm:"not null;default:'';comment:商品副标题"`
	Image      string  `json:"image" gorm:"type:text;not null;default:'';comment:图片地址 json 格式"`
	Detail     string  `json:"detail" gorm:"type:text;not null;default:'';comment:商品详情"`
	Price      float64 `json:"price" gorm:"not null;default:0;comment:商品价格"`
	Stock      int64   `json:"stock" gorm:"not null;default:0;comment:商品数量"`
	Status     int64   `json:"status" gorm:"not null;default:0;comment:商品状态.0-暂存 1-在售 2-下架 3-删除"`
}
