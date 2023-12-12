package models

import "gorm.io/gorm"

// roles 用户角色表
type Role struct {
	BaseModel
	Name     string `json:"name" gorm:"column:name;not null"`                     // 角色名
	ParentId int    `json:"parent_id" gorm:"column:parent_id;not null;default:0"` // 父角色ID
}

type roleModel DB

func NewRoleModel(tx *gorm.DB) *roleModel {
	db := tx.Table("roles").Model(&Role{})
	return &roleModel{db: db}
}

func (m *roleModel) SetIds(id ...int) *roleModel {
	m.db = m.db.Where("id in (?)", id)
	return m
}

func (m *roleModel) SetParentId(parentId int) *roleModel {
	m.db = m.db.Where("parent_id = ?", parentId)
	return m
}

func (m *roleModel) Clone() *roleModel {
	m.db = m.db.Session(&gorm.Session{Initialized: true}).Session(&gorm.Session{})
	return m
}

func (m *roleModel) Order(expr string) *roleModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *roleModel) Create(role Role) error {
	return m.db.Create(&role).Error
}

func (m *roleModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *roleModel) FirstOne() (data Role, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *roleModel) LastOne() (data Role, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *roleModel) DeleteById(id int) error {
	return m.db.Where("id = ?", id).Delete(&Role{}).Error
}

func (m *roleModel) List(limit, offset int) (data []Role, total int64, err error) {
	query := m.db.Model(&Role{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}
