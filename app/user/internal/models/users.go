package models

import (
	"context"

	"gorm.io/gorm"
)

// Users  用户表
type User struct {
	BaseModel
	Username string  `gorm:"column:username" json:"username"` //  用户名
	Password string  `gorm:"column:password" json:"password"` //  密码
	Email    string  `gorm:"column:email" json:"email"`       //  邮箱
	Phone    string  `gorm:"column:phone" json:"phone"`       //  电话
	Amount   float64 `gorm:"column:amount" json:"amount"`     // 用户余额
}

type userModel DB

func NewUserModel(tx ...*gorm.DB) *userModel {
	db := getDB(tx...).Table("users").Model(&User{})
	return &userModel{db: db}
}

func (u *userModel) WithContext(ctx context.Context) *userModel {
	u.db = u.db.WithContext(ctx)
	return u
}

func (u *userModel) SetId(id uint) *userModel {
	u.db = u.db.Where("id = ?", id)
	return u
}

func (u *userModel) SetIds(ids []int) *userModel {
	u.db = u.db.Where("id in (?)", ids)
	return u
}

func (u *userModel) SetUsername(username string) *userModel {
	u.db = u.db.Where("username = ?", username)
	return u
}

func (u *userModel) SetPhoneNumber(phoneNumber string) *userModel {
	u.db = u.db.Where("phone = ?", phoneNumber)
	return u
}

func (u *userModel) PhoneNumberLike(phoneNumber string) *userModel {
	u.db = u.db.Where("phone like ?", "%"+phoneNumber+"%")
	return u
}

func (u *userModel) Set(opts ...Options) *userModel {
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func (u *userModel) Create(user User) (User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userModel) FirstOne() (data User, err error) {
	err = u.db.First(&data).Error
	return
}

func (u *userModel) LastOne() (data User, err error) {
	err = u.db.Last(&data).Error
	return
}

func (u *userModel) Update(values interface{}) (err error) {
	err = u.db.Updates(values).Error
	return
}

func (u *userModel) DeleteById(id int) error {
	return u.db.Where("id = ?", id).Delete(&User{}).Error
}

func (u *userModel) List(limit, offset int) (users []User, total int64, err error) {
	query := u.db
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&users).Error
	return
}

func (u *userModel) ListAll() (users []User, err error) {
	err = u.db.Find(&users).Error
	return
}

type Options func(*userModel)

func WithUserPhone(phoneNumber string) Options {
	return func(um *userModel) {
		um.SetPhoneNumber(phoneNumber)
	}
}

func WithUserName(userName string) Options {
	return func(um *userModel) {
		um.SetUsername(userName)
	}
}
