package model

import (
	"context"

	"gorm.io/gorm"
)

// table products
type Product struct {
	BaseModel
	ProductId   int     `json:"product_id" gorm:"column:product_id;primaryKey;not null"` // 商品ID
	Name        string  `json:"name" gorm:"column:name;default:NULL"`                    // 商品名称
	Description string  `json:"description" gorm:"column:description;default:NULL"`      // 商品描述
	Price       float64 `json:"price" gorm:"column:price;default:NULL"`                  // 商品价格
}

type productModel DB

func NewProductModel(tx *gorm.DB) *productModel {
	db := tx.Table("products").Model(&Product{})
	return &productModel{db: db}
}

func (m *productModel) SetIds(ids ...int64) *productModel {
	m.db = m.db.Where("product_id IN (?)", ids)
	return m
}

func (m *productModel) Order(expr string) *productModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *productModel) Select(fields ...string) *productModel {
	m.db = m.db.Select(fields)
	return m
}

func (m *productModel) WithContext(ctx context.Context) *productModel {
	m.db = m.db.WithContext(ctx)
	return m
}

func (m *productModel) Create(product Product) error {
	return m.db.Create(&product).Error
}

func (m *productModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *productModel) FirstOne() (data Product, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *productModel) LastOne() (data Product, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *productModel) DeleteByPrimKey(key interface{}) error {
	return m.db.Where("product_id IN (?)", key).Delete(&Product{}).Error
}

func (m *productModel) List() (data []Product) {
	m.db.Find(&data)
	return
}

func (m *productModel) PageList(limit, offset int) (data []Product, total int64, err error) {
	query := m.db.Model(&Product{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}
