package models

import "gorm.io/gorm"

// permission
type Permission struct {
	BaseModel
	Name        string `json:"name" gorm:"column:name;not null"`
	Description string `json:"description" gorm:"column:description;default:NULL"`
}

type permissionModel DB

func NewPermissionModel(tx *gorm.DB) *permissionModel {
	db := tx.Table("permission").Model(&Permission{})
	return &permissionModel{db: db}
}

func (m *permissionModel) SetId(id ...uint) *permissionModel {
	m.db = m.db.Where("id IN (?)", id)
	return m
}

func (m *permissionModel) Order(expr string) *permissionModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *permissionModel) Create(permission Permission) error {
	return m.db.Create(&permission).Error
}

func (m *permissionModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *permissionModel) FirstOne() (data Permission, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *permissionModel) LastOne() (data Permission, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *permissionModel) DeleteById(id int) error {
	return m.db.Where("id = ?", id).Delete(&Permission{}).Error
}

func (m *permissionModel) PageList(limit, offset int) (data []Permission, total int64, err error) {
	query := m.db.Model(&Permission{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}

func (m *permissionModel) List() (data []Permission) {
	m.db.Find(&data)
	return
}
