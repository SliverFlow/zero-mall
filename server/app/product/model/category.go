package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	categoryModel interface {
		CategoryCreate(ctx context.Context, category *Category) (err error)
		CategoryList(ctx context.Context, page *common.PageInfo) (enter *[]Category, total int64, err error)
		CategoryFind(ctx context.Context, categoryId string) (enter *Category, err error)
		CategoryUpdate(ctx context.Context, category *Category) (err error)
		CategoryDelete(ctx context.Context, categoryId string) (err error)

		CategoryListAll(ctx context.Context) (enter *[]Category, err error)
		CategoryChangeStatus(ctx context.Context, categoryId string, status int64) (err error)
		CategoryFindChildrenID(ctx context.Context, categoryId string) (enter *[]string, err error)
		CategoryBatchDelete(ctx context.Context, ids []string) (err error)
	}

	defaultCategoryModel struct {
		db *gorm.DB
	}

	// Category
	// Author [SliverFlow]
	// @desc 业务分类表
	Category struct {
		common.GloModel
		CategoryID string    `json:"categoryId" gorm:"not null;default:'';comment:商品分类id"`
		ParentId   string    `json:"parentId" gorm:"not null;default:'';comment:父分类id 当id=0时说明是根节点,一级类别 "`
		Name       string    `json:"name" gorm:"not null;default:'';comment:分类名"`
		Status     int64     `json:"status" gorm:"not null;default:0;comment:类别状态 0-暂存 1-正常,2-已废弃"`
		Sorted     int64     `json:"sorted" gorm:"not null;default:0;comment:排序标记"`
		Type       int64     `json:"type" gorm:"not null;default:0;comment:分类类别 0 系统类别 1 商家添加的类别"`
		BusinessID string    `json:"businessId" gorm:"not null;default:'';comment:商户 id 只有为商家类别的时候才拥有值"`
		Products   []Product `json:"products" gorm:"many2many:category_product;"`
	}
)

func (c *Category) TableName() string {
	return "category"
}

func newDefaultCategoryModel(db *gorm.DB) *defaultCategoryModel {
	return &defaultCategoryModel{
		db: db,
	}
}

// CategoryCreate
// Author [SliverFlow]
// @desc 添加分类
func (d *defaultCategoryModel) CategoryCreate(ctx context.Context, category *Category) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&category).Create(category).Error
}

// CategoryList
// Author [SliverFlow]
// @desc 分页查询分类
func (d *defaultCategoryModel) CategoryList(ctx context.Context, page *common.PageInfo) (enter *[]Category, total int64, err error) {
	tx := d.db.WithContext(ctx)

	limit, offset, keyWord := page.LimitAndOffsetAndKeyWord()

	tx = tx.Model(&Category{}).Where("name like ?", keyWord)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []Category
	if err = tx.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}

// CategoryFind
// Author [SliverFlow]
// @desc 根据分类 id 查询分类
func (d *defaultCategoryModel) CategoryFind(ctx context.Context, categoryId string) (enter *Category, err error) {
	tx := d.db.WithContext(ctx)

	var c Category
	if err = tx.Model(&Category{}).Where("category_id = ?", categoryId).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// CategoryUpdate
// Author [SliverFlow]
// @desc 更新分类信息
func (d *defaultCategoryModel) CategoryUpdate(ctx context.Context, category *Category) (err error) {
	tx := d.db.WithContext(ctx)

	cm := make(map[string]any)
	cm["parent_id"] = category.ParentId
	cm["name"] = category.Name
	cm["status"] = category.Status
	cm["type"] = category.Type
	cm["sorted"] = category.Sorted

	return tx.Model(&Category{}).Where("category_id = ?", category.CategoryID).Updates(cm).Error
}

// CategoryDelete
// Author [SliverFlow]
// @desc 删除分类信息
func (d *defaultCategoryModel) CategoryDelete(ctx context.Context, categoryId string) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Where("category_id = ?", categoryId).Delete(&Category{}).Error
}

// CategoryListAll
// Author [SliverFlow]
// @desc 查询所有分类信息 按照排序标记排序
func (d *defaultCategoryModel) CategoryListAll(ctx context.Context) (enter *[]Category, err error) {
	tx := d.db.WithContext(ctx)

	var list []Category
	if err = tx.Model(&Category{}).Order("sorted asc").Find(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

// CategoryChangeStatus
// Author [SliverFlow]
// @desc 更新分类状态
func (d *defaultCategoryModel) CategoryChangeStatus(ctx context.Context, categoryId string, status int64) (err error) {
	tx := d.db.WithContext(ctx)
	var ids []string
	ids = append(ids, categoryId)

	enter, err := d.CategoryFind(ctx, categoryId)
	if err != nil {
		return err
	}
	if enter.ParentId == "0" && status == 0 {
		var list []Category
		if err = tx.Where(&Category{}).Where("parent_id = ?", enter.CategoryID).Find(&list).Error; err != nil {
			return err
		}
		for _, c := range list {
			ids = append(ids, c.CategoryID)
		}
	}

	return tx.Model(&Category{}).Where("category_id in ?", ids).Update("status", status).Error
}

// CategoryFindChildrenID
// Author [SliverFlow]
// @desc 查询某个分类ID下的所有子分类ID
func (d *defaultCategoryModel) CategoryFindChildrenID(ctx context.Context, categoryId string) (enter *[]string, err error) {
	tx := d.db.WithContext(ctx)

	var categorys []Category
	if err = tx.Model(&Category{}).Where("parent_id = ?", categoryId).Find(&categorys).Error; err != nil {
		return nil, err
	}
	ids := make([]string, 0)
	for _, c := range categorys {
		ids = append(ids, c.CategoryID)
	}
	return &ids, nil
}

// CategoryBatchDelete
// Author [SliverFlow]
// @desc 批量删除
func (d *defaultCategoryModel) CategoryBatchDelete(ctx context.Context, ids []string) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Where("category_id in ?", ids).Delete(&Category{}).Error
}
