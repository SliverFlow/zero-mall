package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"server/app/product/model/vo_fk"
	"server/common"
)

type (
	fkProductModel interface {
		Create(ctx context.Context, fkProduct *FkProduct) error
	}
	defaultFkProductModel struct {
		db    *gorm.DB
		cache *defaultFkProductCustomModel
	}
	FkProduct struct {
		common.GloModel
		FkProductID  string  `gorm:"column:fk_product_id;primaryKey;not null"`
		BusinessID   string  `gorm:"column:business_id;not null"`          // 商家ID\
		Title        string  `gorm:"column:title;not null"`                // 商品标题
		SubTitle     string  `gorm:"column:sub_title;not null;default:''"` // 商品副标题
		Img          string  `gorm:"column:img;not null;default:''"`       // 商品图片
		Detail       string  `gorm:"column:detail;not null;default:''"`    // 商品详情
		SeckillPrice float64 `gorm:"column:seckill_price;"`                // 秒杀价格
		Price        float64 `gorm:"column:price;"`                        // 原有 价格
		Stock        int64   `gorm:"column:stock;"`                        // 库存
		SoldCount    int64   `gorm:"column:sold_count;"`                   // 已售数量
		Status       int64   `gorm:"column:status;not null;default:0"`     // 状态 0-暂存 1-在售 2-下架 3-删除
		CategoryID   string  `gorm:"column:category_id;not null"`          // 分类ID
		StartDay     int64   `gorm:"column:start_day;not null;default:0"`  // 开始日期
		EndDay       int64   `gorm:"column:end_day;not null;default:0"`    // 结束日期
	}
)

func (fk *FkProduct) TableName() string {
	return "fk_product"
}

func newDefaultFkProductModel(db *gorm.DB, rdb *redis.Client) *defaultFkProductModel {
	return &defaultFkProductModel{
		db:    db,
		cache: &defaultFkProductCustomModel{rdb: rdb},
	}
}

func (d *defaultFkProductModel) Create(ctx context.Context, fkProduct *FkProduct) error {
	span, _ := common.Tracer(ctx, "fk_product-create")
	defer span.End()

	tx := d.db.WithContext(ctx)

	return tx.Create(fkProduct).Error
}

func (d *defaultFkProductModel) FindByID(ctx context.Context, id string) (*FkProduct, error) {
	span, _ := common.Tracer(ctx, "fk_product-find_by_id")
	defer span.End()

	tx := d.db.WithContext(ctx)

	fp, ok := d.cache.CacheGetByID(ctx, id)
	if ok {
		return fp, nil
	}

	var fkProduct FkProduct
	err := tx.Where("fk_product_id = ?", id).First(&fkProduct).Error
	if err != nil {
		return nil, err
	}

	d.cache.CacheInsert(ctx, &fkProduct)

	return &fkProduct, nil
}

func (d *defaultFkProductModel) UpdateByID(ctx context.Context, fkProduct *FkProduct) error {
	span, _ := common.Tracer(ctx, "fk_product-update_by_id")
	defer span.End()

	tx := d.db.WithContext(ctx)

	err := tx.Model(&FkProduct{}).Where("fk_product_id = ?", fkProduct).Updates(fkProduct).Error
	if err != nil {
		return err
	}

	d.cache.CacheDelete(ctx, fkProduct.FkProductID)

	return nil
}

func (d *defaultFkProductModel) DeleteByID(ctx context.Context, id string) error {
	span, _ := common.Tracer(ctx, "fk_product-delete_by_id")
	defer span.End()

	tx := d.db.WithContext(ctx)

	err := tx.Where("fk_product_id = ?", id).Delete(&FkProduct{}).Error
	if err != nil {
		return err
	}

	d.cache.CacheDelete(ctx, id)
	return nil
}

func (d *defaultFkProductModel) FindListPage(ctx context.Context, req *vo_fk.FkProductListReq) (fps []*FkProduct, total int64, err error) {
	span, _ := common.Tracer(ctx, "fk_product-find_list_page")
	defer span.End()

	// tx := d.db.WithContext(ctx)

	//cacheKey := d.cache.getCacheKey(ctx, req)
	//list, ids, ok := d.cache.CacheGetList(ctx, cacheKey)
	//if ok && len(ids) == 0 {
	//	return list, d.cache.getTotal(ctx, cacheKey+":"+"count"), nil
	//}
	//
	//if ok && len(ids) > 0 {
	//	ds, err := d.findListByIDs(ctx, ids)
	//	if err != nil {
	//		return nil, 0, err
	//	}
	//	list = append(list, ds...)
	//	return list, d.cache.getTotal(ctx, cacheKey+":"+"count"), nil
	//}
	//
	//if !ok && len(ids) == 0 {
	//	err = tx.Where("start_time >= ? and end_time <= ?", req.StartTime, req.EndTime).Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&fps).Error
	//	if err != nil {
	//		return nil, 0, err
	//	}
	//
	//	return fps, total, nil
	//}

	return fps, total, nil
}
