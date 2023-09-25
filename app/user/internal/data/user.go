package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// GetUserByPhone implements biz.UserRepo.
func (r *userRepo) GetUserByPhone(ctx context.Context, phoneNumber string) (*models.User, error) {
	user, err := models.NewUserModel(r.data.DB(ctx)).SetPhoneNumber(phoneNumber).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// GetUserByName implements biz.UserRepo.
func (r *userRepo) GetUserByName(ctx context.Context, username string) (*models.User, error) {
	user, err := models.NewUserModel(r.data.DB(ctx)).SetUsername(username).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func NewUserRepo(data *Data, logger log.Logger) *userRepo {
	return &userRepo{data: data, log: log.NewHelper(logger)}
}

func (r *userRepo) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := models.NewUserModel(r.data.DB(ctx)).SetId(id).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return models.NewUserModel(r.data.DB(ctx)).SetId(user.ID).Update(user)
}

func (r *userRepo) DeleteUser(ctx context.Context, id uint) error {
	return models.NewUserModel(r.data.DB(ctx)).DeleteById(int(id))
}

func (r *userRepo) GetUsersByRole(ctx context.Context, roleID int) (result []models.User, err error) {
	return
}

// CreateUser implements biz.UserRepo.
func (r *userRepo) CreateUser(ctx context.Context, user *models.User) (userId int, err error) {
	userInfo, err := models.NewUserModel(r.data.DB(ctx)).Create(*user)
	if err != nil {
		return 0, err
	}
	return int(userInfo.ID), nil
}

func (r *userRepo) ListUser(ctx context.Context, cond *biz.ListUserCond) []models.User {
	model := models.NewUserModel(r.data.DB(ctx)).WithContext(ctx)
	if cond.Phone != "" {
		model.PhoneNumberLike(cond.Phone)
	}
	if cond.Page == 0 {
		cond.Page = 1
	}
	if cond.Limit == 0 {
		cond.Limit = 10
	}

	users, _, _ := model.List(int(cond.Limit), (int(cond.Page)-1)*int(cond.Limit))

	return users
}
