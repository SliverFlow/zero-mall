package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	productModel interface {
		ProductCreate(ctx context.Context, product *Product) (err error)
		ProductDelete(ctx context.Context, ids []string) (err error)
	}

	defaultProductModel struct {
		db *gorm.DB
	}

	// Product
	// Author [SliverFlow]
	// @desc 业务商品表
	Product struct {
		common.GloModel
		ProductID  string  `json:"productId" gorm:"not null;default:'';comment:商品业务 id"`
		Name       string  `json:"name" gorm:"not null;default:'';comment:商品名"`
		CategoryID string  `json:"categoryId" gorm:"not null;default:'';comment:商品分类id"`
		Subtitle   string  `json:"subtitle" gorm:"not null;default:'';comment:商品副标题"`
		Image      string  `json:"image" gorm:"type:text;not null;comment:图片地址 json 格式"`
		Detail     string  `json:"detail" gorm:"type:text;not null;comment:商品详情"`
		Price      float64 `json:"price" gorm:"not null;default:0;comment:商品价格"`
		Stock      int64   `json:"stock" gorm:"not null;default:0;comment:商品数量"`
		Status     int64   `json:"status" gorm:"not null;default:0;comment:商品状态.0-暂存 1-在售 2-下架 3-删除"`
	}
)

func (p *Product) TableName() string {
	return "product"
}

func newDefaultProductModel(db *gorm.DB) *defaultProductModel {
	return &defaultProductModel{db: db}
}

func (d *defaultProductModel) ProductCreate(ctx context.Context, product *Product) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&Product{}).Create(product).Error
}

func (d *defaultProductModel) ProductDelete(ctx context.Context, ids []string) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Where("product_id in ?", ids).Delete(&Product{}).Error
}
