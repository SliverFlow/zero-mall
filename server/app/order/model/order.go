package model

import (
	"context"
	"gorm.io/gorm"
	"server/app/order/model/request"
	"server/common"
	"time"
)

type (
	orderModel interface {
		OrderCreate(ctx context.Context, Order *Order) (err error)
		OrderList(ctx context.Context, page *common.PageInfo, userId string) (enter *[]Order, total int64, err error)
		OrderFind(ctx context.Context, OrderId string) (enter *Order, err error)
		OrderUpdate(ctx context.Context, Order *Order) (err error)
		OrderDelete(ctx context.Context, orderIds []string) (err error)
		OrderListPage(ctx context.Context, req *request.PageListReq) (enter []Order, total int64, err error)
		OrderDisable(ctx context.Context, OrderId string) (err error)
		OrderFindListByIDs(ctx context.Context, orderIds []string) (orders []Order, err error)
		OrderDeleteByID(ctx context.Context, orderId string) (err error)
		OrderListPageForUser(ctx context.Context, req *request.PageListReq) (enter []Order, total int64, err error)
	}

	defaultOrderModel struct {
		db *gorm.DB
	}

	// Order
	// Author [SliverFlow]
	// @desc 业务订单表
	Order struct {
		common.GloModel
		OrderID     string  `json:"orderId" gorm:"not null;default:'';comment:订单业务id"`
		UserID      string  `json:"userId" gorm:"not null;default:'';comment:用户业务id"`
		ShoppingID  string  `json:"shoppingId" gorm:"not null;default:'';comment:收获信息业务id"`
		Payment     float64 `json:"payment" gorm:"type:decimal(10,2);not null;default:0;comment:实际付款金额"`
		PaymentType int64   `json:"paymentType" gorm:"not null;default:1;comment:支付类型,1-在线支付"`
		Postage     int64   `json:"postage" gorm:"not null;default:0;comment:运费 单位是 元"`
		Status      int64   `json:"status" gorm:"not null;default:10;comment:订单状态:70-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭"`
		PaymentTime int64   `json:"paymentTime" gorm:"not null;default:0;comment:支付时间"`
		SendTime    int64   `json:"sendTime" gorm:"not null;default:0;comment:发货时间"`
		EndTime     int64   `json:"endTime" gorm:"not null;default:0;comment:交易完成时间"`
		CloseTime   int64   `json:"closeTime" gorm:"not null;default:0;comment:交易关闭时间"`
		BusinessID  string  `json:"businessId" gorm:"not null;default:'';comment:冗余字段 商家业务id"`
	}
)

func (o *Order) TableName() string {
	return "order"
}

func newOrderModel(db *gorm.DB) *defaultOrderModel {
	return &defaultOrderModel{db: db}
}

// OrderCreate
// Author [SliverFlow]
// @desc 添加分类
func (d *defaultOrderModel) OrderCreate(ctx context.Context, order *Order) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Model(&Order{}).Create(order).Error
}

// OrderList
// Author [SliverFlow]
// @desc 分页查询分类
func (d *defaultOrderModel) OrderList(ctx context.Context, page *common.PageInfo, userId string) (enter *[]Order, total int64, err error) {
	tx := d.db.WithContext(ctx)

	limit, offset, _ := page.LimitAndOffsetAndKeyWord()

	tx = tx.Model(&Order{}).Where("user_id = ?", userId)
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []Order
	if err = tx.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return &list, total, nil
}

// OrderFind
// Author [SliverFlow]
// @desc 根据分类 id 查询分类
func (d *defaultOrderModel) OrderFind(ctx context.Context, orderId string) (enter *Order, err error) {
	tx := d.db.WithContext(ctx)

	var c Order
	if err = tx.Model(&Order{}).Where("order_id = ?", orderId).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// OrderUpdate
// Author [SliverFlow]
// @desc 更新分类信息
func (d *defaultOrderModel) OrderUpdate(ctx context.Context, order *Order) (err error) {
	tx := d.db.WithContext(ctx)

	cm := make(map[string]any)

	return tx.Model(&Order{}).Where("order_id = ?", order.OrderID).Updates(cm).Error
}

// OrderDelete
// Author [SliverFlow]
// @desc 删除分类信息
func (d *defaultOrderModel) OrderDelete(ctx context.Context, orderIds []string) (err error) {
	tx := d.db.WithContext(ctx)

	return tx.Where("order_id in ?", orderIds).Delete(&Order{}).Error
}

// OrderListByRole
// Author [SliverFlow]
// @desc 分页查询分类
func (d *defaultOrderModel) OrderListByRole(ctx context.Context) (enter *[]Order, total int64, err error) {

	return nil, 0, err
}

func (d *defaultOrderModel) OrderListPage(ctx context.Context, req *request.PageListReq) (enter []Order, total int64, err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "orderListPage")
	defer span.End()

	query := tx.Model(&Order{}).Unscoped()

	if req.BusinessID != "" {
		query = query.Where("business_id = ?", req.BusinessID)
	}

	if req.UserID != "" {
		query = query.Where("user_id = ?", req.UserID)
	}

	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}

	if req.StartTime != 0 {
		query = query.Where("created_at >= ?", time.Unix(req.StartTime, 0))
	}

	if req.EndTime != 0 {
		query = query.Where("created_at <= ?", time.Unix(req.EndTime, 0))
	}

	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	limit, offset := req.LimitAndOffset()
	err = query.Limit(limit).Offset(offset).Find(&enter).Error
	if err != nil {
		return nil, 0, err
	}

	return enter, total, nil
}

func (d *defaultOrderModel) OrderDisable(ctx context.Context, OrderId string) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "order-disable")
	defer span.End()

	return tx.Model(&Order{}).Where("order_id = ?", OrderId).Update("status", 70).Error
}

func (d *defaultOrderModel) OrderFindListByIDs(ctx context.Context, orderIds []string) (orders []Order, err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "order-find-list-by-ids")
	defer span.End()

	var enter []Order
	if err = tx.Model(&Order{}).Unscoped().Where("order_id in ?", orderIds).Find(&enter).Error; err != nil {
		return nil, err
	}
	return enter, nil
}

func (d *defaultOrderModel) OrderDeleteByID(ctx context.Context, orderId string) (err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "order-delete-by-id")
	defer span.End()

	return tx.Where("order_id = ?", orderId).Delete(&Order{}).Error
}

func (d *defaultOrderModel) OrderListPageForUser(ctx context.Context, req *request.PageListReq) (enter []Order, total int64, err error) {
	tx := d.db.WithContext(ctx)
	span, _ := common.Tracer(ctx, "orderListPage")
	defer span.End()

	query := tx.Model(&Order{})

	if req.BusinessID != "" {
		query = query.Where("business_id = ?", req.BusinessID)
	}

	if req.UserID != "" {
		query = query.Where("user_id = ?", req.UserID)
	}

	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}

	if req.StartTime != 0 {
		query = query.Where("created_at >= ?", time.Unix(req.StartTime, 0))
	}

	if req.EndTime != 0 {
		query = query.Where("created_at <= ?", time.Unix(req.EndTime, 0))
	}

	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	limit, offset := req.LimitAndOffset()
	err = query.Limit(limit).Offset(offset).Find(&enter).Error
	if err != nil {
		return nil, 0, err
	}

	return enter, total, nil
}
