package models

import "gorm.io/gorm"

// user_role
type UserRole struct {
	BaseModel
	UserId int `json:"user_id" gorm:"column:user_id;not null"`
	RoleId int `json:"role_id" gorm:"column:role_id;not null"`
}

type userRoleModel DB

func NewUserRoleModel(tx *gorm.DB) *userRoleModel {
	db := tx.Table("user_role").Model(&UserRole{})
	return &userRoleModel{db: db}
}

func (m *userRoleModel) SetId(id int) *userRoleModel {
	m.db = m.db.Where("id = ?", id)
	return m
}

func (m *userRoleModel) Order(expr string) *userRoleModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *userRoleModel) Create(userRole UserRole) error {
	return m.db.Create(&userRole).Error
}

func (m *userRoleModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *userRoleModel) FirstOne() (data UserRole, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *userRoleModel) LastOne() (data UserRole, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *userRoleModel) DeleteById(id int) error {
	return m.db.Where("id = ?", id).Delete(&UserRole{}).Error
}

func (m *userRoleModel) List(limit, offset int) (data []UserRole, total int64, err error) {
	query := m.db.Model(&UserRole{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}
