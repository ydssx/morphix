package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
	"gorm.io/gorm"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// GetUserByName implements biz.UserRepo.
func (*userRepo) GetUserByName(ctx context.Context, username string) (*models.User, error) {
	user, err := models.NewUserModel().SetUsername(username).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func NewUserRepo(data *Data, log log.Logger) biz.UserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) GetUserByID(ctx context.Context, id uint, tx ...*gorm.DB) (*models.User, error) {
	user, err := models.NewUserModel(tx...).SetId(id).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return models.NewUserModel().SetId(user.ID).Update(user)
}

func (r *userRepo) DeleteUser(ctx context.Context, id uint) error {
	return models.NewUserModel().DeleteById(int(id))
}

func (r *userRepo) GetUsersByRole(ctx context.Context, roleID int) (result []models.User, err error) {
	return
}

// CreateUser implements biz.UserRepo.
func (*userRepo) CreateUser(ctx context.Context, user *models.User, tx ...*gorm.DB) (userId int, err error) {
	userInfo, err := models.NewUserModel(tx...).Create(*user)
	if err != nil {
		return 0, err
	}
	return int(userInfo.ID), nil
}

func (*userRepo) ListUser(ctx context.Context) []models.User {
	users, _, _ := models.NewUserModel().List(10, 0)
	return users
}
