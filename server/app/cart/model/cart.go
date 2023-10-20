package model

import (
	"context"
	"gorm.io/gorm"
	"server/common"
)

type (
	cartModel interface {
		CartCreate(ctx context.Context, cart *Cart) (err error)
		CartUpdate(ctx context.Context, cart *Cart) (err error)
		CartDelete(ctx context.Context, cartId string) (err error)
		CartFind(ctx context.Context, cartId string) (enter *Cart, err error)
		CartCountByUserID(ctx context.Context, userId string) (total int64, err error)
		CartListByUserID(ctx context.Context, userId string, page *common.PageInfo) (enter *[]Cart, total int64, err error)
	}

	defaultCartModel struct {
		db *gorm.DB
	}

	// Cart 购物车
	Cart struct {
		common.GloModel
		CartID      string `json:"cartId" gorm:"not null;default:'';comment:购物车业务id"`
		UserID      string `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
		ProductID   string `json:"productId" gorm:"not null;default:'';comment:商品业务 id "`
		ProductName string `json:"productName" gorm:"not null;default:'';comment:冗余商品名称"`
		Quantity    int64  `json:"quantity" gorm:"not null;default:0;comment:数量"`
		Checked     int64  `json:"checked" gorm:"not null;default:0;comment:是否选择,1=已勾选,0=未勾选"`
	}
)

func (c *Cart) TableName() string {
	return "cart"
}

func newCartModel(db *gorm.DB) *defaultCartModel {
	return &defaultCartModel{db: db}
}

// CartCreate 创建一条购物车信息
func (d *defaultCartModel) CartCreate(ctx context.Context, cart *Cart) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "cart-create")
	defer span.End()

	return tx.Model(&Cart{}).Create(cart).Error
}

// CartUpdate 更新购物车信息
func (d *defaultCartModel) CartUpdate(ctx context.Context, cart *Cart) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "cart-update")
	defer span.End()

	cMap := make(map[string]interface{})
	cMap["checked"] = cart.Checked
	cMap["quantity"] = cart.Quantity

	return tx.Model(&Cart{}).Where("cart_id = ?", cart.CartID).Updates(cMap).Error
}

// CartDelete 删除一条购物车信息
func (d *defaultCartModel) CartDelete(ctx context.Context, cartId string) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "cart-delete")
	defer span.End()

	return tx.Where("cart_id = ?", cartId).Delete(&Cart{}).Error
}

// CartFind 查询一条购物车信息
func (d *defaultCartModel) CartFind(ctx context.Context, cartId string) (enter *Cart, err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "cart-delete")
	defer span.End()

	var c Cart
	if err = tx.Model(&Cart{}).Where("cart_id = ?", cartId).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

// CartCountByUserID 根据用户ID统计购物车信息
func (d *defaultCartModel) CartCountByUserID(ctx context.Context, userId string) (total int64, err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "cart-count-by-userId")
	defer span.End()

	if err = tx.Model(&Cart{}).Where("user_id = ?", userId).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// CartListByUserID 查询用户购物车列表
func (d *defaultCartModel) CartListByUserID(ctx context.Context, userId string, page *common.PageInfo) (enter *[]Cart, total int64, err error) {
	tx := d.db.WithContext(ctx)

	limit, offset, keyWord := page.LimitAndOffsetAndKeyWord()

	tx = tx.Model(&Cart{}).Where("user_id = ? and product_name like ?", userId, keyWord)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []Cart
	if err = tx.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}
