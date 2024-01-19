package model

import (
	"context"

	"gorm.io/gorm"
)

// table orders
type Order struct {
	BaseModel
	OrderNumber string  `json:"order_number" gorm:"column:order_number;not null"` // 订单号
	UserId      int     `json:"user_id" gorm:"column:user_id;not null"`           // 客户ID
	Amount      float64 `json:"amount" gorm:"column:amount;not null"`             // 订单金额
	Status      string  `json:"status" gorm:"column:status;default:PENDING"`             // 订单状态
}

type orderModel DB

func NewOrderModel(tx *gorm.DB) *orderModel {
	db := tx.Table("orders").Model(&Order{})
	return &orderModel{db: db}
}

func (m *orderModel) SetIds(ids ...int64) *orderModel {
	m.db = m.db.Where("order_id IN (?)", ids)
	return m
}

func (m *orderModel) Order(expr string) *orderModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *orderModel) Select(fields ...string) *orderModel {
	m.db = m.db.Select(fields)
	return m
}

func (m *orderModel) WithContext(ctx context.Context) *orderModel {
	m.db = m.db.WithContext(ctx)
	return m
}

func (m *orderModel) Create(order Order) error {
	return m.db.Create(&order).Error
}

func (m *orderModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *orderModel) FirstOne() (data Order, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *orderModel) LastOne() (data Order, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *orderModel) DeleteByPrimKey(key interface{}) error {
	return m.db.Where("order_id IN (?)", key).Delete(&Order{}).Error
}

func (m *orderModel) List() (data []Order) {
	m.db.Find(&data)
	return
}

func (m *orderModel) PageList(limit, offset int) (data []Order, total int64, err error) {
	query := m.db.Model(&Order{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}
