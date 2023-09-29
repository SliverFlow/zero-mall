package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	businessModel interface {
		BusinessCreate(ctx context.Context, business *Business) (err error)
		BusinessList(ctx context.Context, page *common.PageInfo) (enter *[]Business, total int64, err error)
		BusinessChangeStatus(ctx context.Context, id string, status int64) (err error)
		BusinessFind(ctx context.Context, businessId string) (enter *Business, err error)
		BusinessFindByUUID(ctx context.Context, uuid string) (enter *Business, err error)
		BusinessUpdate(ctx context.Context, business *Business) (err error)
	}

	defaultBusinessModel struct {
		db *gorm.DB
	}

	// Business
	// Author [SliverFlow]
	// @desc 商家表
	Business struct {
		common.GloModel
		BusinessID string `json:"businessId" gorm:"not null;default:'';comment:商户 id"`
		UUID       string `json:"uuid" gorm:"not null;default:'';comment:关联用户已存在的 user "`
		Name       string `json:"name" gorm:"not null;default:'';comment:商家名称"`
		Detail     string `json:"detail" gorm:"type:text;not null;comment:商家详情"`
		Score      int64  `json:"score" gorm:"not null;default:0;comment:商家评分"`
		Image      string `json:"image" gorm:"type:text;not null;comment:图片地址 json 格式"`
		Status     int64  `json:"status" gorm:"not null;default:0;comment:商品状态.0-暂存 1-正常 2-禁用"`
	}
)

func (b *Business) TableName() string {
	return "business"
}

func newBusinessModel(db *gorm.DB) *defaultBusinessModel {
	return &defaultBusinessModel{db: db}
}

// BusinessCreate 创建商家信息
func (d *defaultBusinessModel) BusinessCreate(ctx context.Context, business *Business) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&Business{}).Create(business).Error
}

func (d *defaultBusinessModel) BusinessList(ctx context.Context, page *common.PageInfo) (enter *[]Business, total int64, err error) {
	tx := d.db.WithContext(ctx)

	limit, offset, keyWord := page.LimitAndOffsetAndKeyWord()
	tx = tx.Model(&Business{}).Where("name like ?", keyWord)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []Business
	if err = tx.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}

// BusinessChangeStatus 修改商户状态
func (d *defaultBusinessModel) BusinessChangeStatus(ctx context.Context, id string, status int64) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&Business{}).Where("business_id = ?", id).Update("status", status).Error
}

// BusinessFind 查询商户信息
func (d *defaultBusinessModel) BusinessFind(ctx context.Context, businessId string) (enter *Business, err error) {
	span, _ := common.Tracer(ctx, "business-find")
	defer span.End()

	var b Business
	if err = d.db.Model(&Business{}).Where("business_id = ?", businessId).First(&b).Error; err != nil {
		return nil, err
	}

	return &b, nil
}

// BusinessFindByUUID 根据用户 uuid 查询商户信息
func (d *defaultBusinessModel) BusinessFindByUUID(ctx context.Context, uuid string) (enter *Business, err error) {
	span, _ := common.Tracer(ctx, "business-find-uuid")
	defer span.End()

	var b Business
	if err = d.db.Model(&Business{}).Where("uuid = ?", uuid).First(&b).Error; err != nil {
		return nil, err
	}

	return &b, nil
}

// BusinessUpdate 更新商户信息
func (d *defaultBusinessModel) BusinessUpdate(ctx context.Context, business *Business) (err error) {
	span, _ := common.Tracer(ctx, "business-update")
	defer span.End()

	bm := make(map[string]any)
	bm["name"] = business.Name
	bm["image"] = business.Image
	bm["status"] = business.Status
	bm["detail"] = business.Detail
	bm["score"] = business.Score

	return d.db.Model(&Business{}).Where("business_id = ?", business.BusinessID).Updates(&bm).Error
}
