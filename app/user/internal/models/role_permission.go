package models

import "gorm.io/gorm"

// role_permission
type RolePermission struct {
	BaseModel
	RoleId       int `json:"role_id" gorm:"column:role_id;not null"`
	PermissionId int `json:"permission_id" gorm:"column:permission_id;not null"`
}

type rolePermissionModel DB

func NewRolePermissionModel(tx *gorm.DB) *rolePermissionModel {
	db := tx.Table("role_permission").Model(&RolePermission{})
	return &rolePermissionModel{db: db}
}

func (m *rolePermissionModel) SetId(id int) *rolePermissionModel {
	m.db = m.db.Where("id = ?", id)
	return m
}

func (m *rolePermissionModel) SetRoleId(roleId uint) *rolePermissionModel {
	m.db = m.db.Where("role_id = ?", roleId)
	return m
}

func (m *rolePermissionModel) SetRoleIds(roleIds ...int) *rolePermissionModel {
	m.db = m.db.Where("role_id in (?)", roleIds)
	return m
}

func (m *rolePermissionModel) SetPermissionId(permissionId uint) *rolePermissionModel {
	m.db = m.db.Where("permission_id = ?", permissionId)
	return m
}

func (m *rolePermissionModel) Order(expr string) *rolePermissionModel {
	m.db = m.db.Order(expr)
	return m
}

func (m *rolePermissionModel) Create(rolePermission RolePermission) error {
	return m.db.Create(&rolePermission).Error
}

func (m *rolePermissionModel) Updates(values interface{}) error {
	return m.db.Updates(values).Error
}

func (m *rolePermissionModel) FirstOne() (data RolePermission, err error) {
	err = m.db.Take(&data).Error
	return
}

func (m *rolePermissionModel) LastOne() (data RolePermission, err error) {
	err = m.db.Last(&data).Error
	return
}

func (m *rolePermissionModel) DeleteById(id int) error {
	return m.db.Where("id = ?", id).Delete(&RolePermission{}).Error
}

func (m *rolePermissionModel) PageList(limit, offset int) (data []RolePermission, total int64, err error) {
	query := m.db.Model(&RolePermission{})
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&data).Error
	return
}

func (m *rolePermissionModel) List() (data []RolePermission) {
	m.db.Find(&data)
	return
}

func (m *rolePermissionModel) PluckPermissionID() (permissionIDs []uint) {
	m.db.Pluck("permission_id", &permissionIDs)
	return
}

func (m *rolePermissionModel) Delete() error {
	return m.db.Delete(&RolePermission{}).Error
}
