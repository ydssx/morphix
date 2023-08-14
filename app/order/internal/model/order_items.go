package model

import (
	"context"

	"gorm.io/gorm"
)

// table order_items
type OrderItem struct {
	BaseModel
	ItemId    int     `json:"item_id" gorm:"column:item_id;primaryKey;not null"` // 订单商品ID
	OrderId   int     `json:"order_id" gorm:"column:order_id;default:NULL"`      // 订单ID
	ProductId int     `json:"product_id" gorm:"column:product_id;default:NULL"`  // 商品ID
	Quantity  int     `json:"quantity" gorm:"column:quantity;default:NULL"`      // 商品数量
	Price     float64 `json:"price" gorm:"column:price;default:NULL"`            // 商品单价
}

type orderItemModel DB

func NewOrderItemModel(tx *gorm.DB) *orderItemModel {
	db := tx.Table("order_items").Model(&OrderItem{})
	return &orderItemModel{db: db}
}

func (m *orderItemModel) SetIds(ids ...int64) *orderItemModel {
	m.db = m.db.Where("item_id IN (?)", ids)
	return m
}

func (m *orderItemModel) Order(expr string) *orderItemModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *orderItemModel) Select(fields ...string) *orderItemModel {
	m.db = m.db.Select(fields)
	return m
}

func (m *orderItemModel) WithContext(ctx context.Context) *orderItemModel {
	m.db = m.db.WithContext(ctx)
	return m
}

func (m *orderItemModel) Create(orderItem OrderItem) error {
	return m.db.Create(&orderItem).Error
}

func (m *orderItemModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *orderItemModel) FirstOne() (data OrderItem, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *orderItemModel) LastOne() (data OrderItem, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *orderItemModel) DeleteByPrimKey(key interface{}) error {
	return m.db.Where("item_id IN (?)", key).Delete(&OrderItem{}).Error
}

func (m *orderItemModel) List() (data []OrderItem) {
	m.db.Find(&data)
	return
}

func (m *orderItemModel) PageList(limit, offset int) (data []OrderItem, total int64, err error) {
	query := m.db.Model(&OrderItem{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}
