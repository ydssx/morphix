package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/logger"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByID(ctx context.Context, id uint, tx ...*gorm.DB) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User, tx ...*gorm.DB) (userId int, err error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	GetUsersByRole(ctx context.Context, roleID int) ([]models.User, error)
	ListUser(ctx context.Context) []models.User
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// RegisterUser 用户注册逻辑
func (uc *UserUsecase) RegisterUser(ctx context.Context, username, password, email, phone string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	// 创建用户对象
	user := &models.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	userId, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = uint(userId)
	return user, nil
}

func (uc *UserUsecase) ListUser(ctx context.Context) (*userv1.UserListResponse, error) {
	userInfo, _ := interceptors.AuthFromContext(ctx)
	logger.Infof(ctx, "userInfo: %#+v", userInfo)

	users := uc.repo.ListUser(ctx)

	resp := new(userv1.UserListResponse)
	for _, user := range users {
		resp.Users = append(resp.Users, &userv1.User{
			Id:       int64(user.ID),
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}
	return resp, nil
}
