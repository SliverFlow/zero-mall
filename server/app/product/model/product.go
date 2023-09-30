package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"server/common"
)

type (
	productModel interface {
		ProductCreate(ctx context.Context, product *Product) (err error)
		ProductDelete(ctx context.Context, ids []string) (err error)
		ProductList(ctx context.Context, page *common.PageInfo) (enter *[]Product, total int64, err error)
		ProductFind(ctx context.Context, productId string) (enter *Product, err error)
		ProductUpdate(ctx context.Context, product *Product) (err error)

		ProductStagingByBusinessID(ctx context.Context, businessId string) (err error)
		ProductChangeStatus(ctx context.Context, productId string, status int64) (err error)
	}

	defaultProductModel struct {
		db *gorm.DB
		*defaultProductCustomModel
	}

	// Product
	// Author [SliverFlow]
	// @desc 业务商品表
	Product struct {
		common.GloModel
		ProductID  string  `json:"productId" gorm:"not null;default:'';comment:商品业务 id"`
		BusinessID string  `json:"businessId" gorm:"not null;default:'';comment:所属商户 id "`
		Name       string  `json:"name" gorm:"not null;default:'';comment:商品名"`
		Subtitle   string  `json:"subtitle" gorm:"not null;default:'';comment:商品副标题"`
		Image      string  `json:"image" gorm:"type:text;not null;comment:图片地址 json 格式"`
		Detail     string  `json:"detail" gorm:"type:text;not null;comment:商品详情"`
		Price      float64 `json:"price" gorm:"type:decimal(10,2);column:price;not null;default:0;comment:商品价格"`
		Stock      int64   `json:"stock" gorm:"not null;default:0;comment:商品数量"`
		Status     int64   `json:"status" gorm:"not null;default:0;comment:商品状态.0-暂存 1-在售 2-下架 3-删除"`

		Categories []Category `json:"categories" gorm:"many2many:category_product;"`
	}
)

func (p *Product) TableName() string {
	return "product"
}

func newDefaultProductModel(db *gorm.DB) *defaultProductModel {
	return &defaultProductModel{
		db: db,
		defaultProductCustomModel: &defaultProductCustomModel{
			cache: "name",
		},
	}
}

// ProductCreate
// Author [SliverFlow]
// @desc 商品创建
func (d *defaultProductModel) ProductCreate(ctx context.Context, product *Product) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&Product{}).Create(product).Error
}

// ProductDelete
// Author [SliverFlow]
// @desc 商品删除
func (d *defaultProductModel) ProductDelete(ctx context.Context, ids []string) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := d.tracer(ctx, "product-delete")
	defer span.End()

	var p Product
	// 更新关系
	err = tx.Model(&Product{}).Preload("Categories").Where("product_id = ?", ids[0]).First(&p).Error
	err = tx.Model(&p).Association("Categories").Delete(p.Categories)
	if err != nil {
		return err
	}

	return tx.Where("product_id in ?", ids).Delete(&Product{}).Error
}

// ProductList
// Author [SliverFlow]
// @desc 分页查询分类
func (d *defaultProductModel) ProductList(ctx context.Context, page *common.PageInfo) (enter *[]Product, total int64, err error) {
	tx := d.db.WithContext(ctx)

	limit, offset, keyWord := page.LimitAndOffsetAndKeyWord()

	tx = tx.Model(&Product{}).Where("name like ?", keyWord)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []Product
	if err = tx.Preload("Categories").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}

// ProductFind
// Author [SliverFlow]
// @desc 根据分类 id 查询分类
func (d *defaultProductModel) ProductFind(ctx context.Context, productId string) (enter *Product, err error) {
	tx := d.db.WithContext(ctx)

	var c Product
	if err = tx.Model(&Product{}).Preload("Categories").Where("product_id = ?", productId).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// ProductUpdate
// Author [SliverFlow]
// @desc 更新分类信息
func (d *defaultProductModel) ProductUpdate(ctx context.Context, product *Product) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := d.tracer(ctx, "product-update")
	defer span.End()

	fmt.Println(product.Categories)

	var p Product
	// 更新关系
	err = tx.Model(&Product{}).Preload("Categories").Where("product_id = ?", product.ProductID).First(&p).Error
	err = tx.Model(&p).Association("Categories").Replace(product.Categories)
	if err != nil {
		return err
	}

	return tx.Model(&Product{}).Where("product_id = ?", product.ProductID).Updates(product).Error
}

// ProductStagingByBusinessID
// Author [SliverFlow]
// @desc 根据商户 id 下架所有商品
func (d *defaultProductModel) ProductStagingByBusinessID(ctx context.Context, businessId string) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := d.tracer(ctx, "product-update-status-staging")
	defer span.End()

	return tx.Model(&Product{}).Where("business_id = ?", businessId).Update("status", 0).Error
}

// ProductChangeStatus
// Author [SliverFlow]
// @desc 更新商品状态
func (d *defaultProductModel) ProductChangeStatus(ctx context.Context, productId string, status int64) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := d.tracer(ctx, "product-update-status")
	defer span.End()

	return tx.Model(&Product{}).Where("product_id = ?", productId).Update("status", status).Error
}
