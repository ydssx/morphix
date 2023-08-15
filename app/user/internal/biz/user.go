package biz

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByID(ctx context.Context, id uint, tx ...*gorm.DB) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User, tx ...*gorm.DB) (userId int, err error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	GetUsersByRole(ctx context.Context, roleID int) ([]models.User, error)
	ListUser(ctx context.Context) []models.User
	GetUserByName(ctx context.Context, username string) (*models.User, error)
}

type UserRepoWithCache interface {
	UserRepo
}

type UserUsecase struct {
	repo UserRepoWithCache
	log  *log.Helper
}

func NewUserUsecase(repo UserRepoWithCache, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// RegisterUser 用户注册逻辑
func (uc *UserUsecase) RegisterUser(ctx context.Context, username, password, email, phone string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	if !util.IsPhoneNumber(phone) {
		return nil, errors.New("手机号格式不正确")
	}

	// 创建用户对象
	user := &models.User{
		Username: username,
		Password: util.MD5(password),
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

func (uc *UserUsecase) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.AuthenticationResponse, error) {
	userInfo, err := uc.repo.GetUserByName(ctx, req.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	if util.MD5(req.Password) != userInfo.Password {
		return nil, status.Error(codes.InvalidArgument, "密码错误")
	}

	token, err := jwt.GenerateToken(int64(userInfo.ID), userInfo.Username, "")
	if err != nil {
		return nil, status.Error(codes.Internal, "token生成失败")
	}

	return &userv1.AuthenticationResponse{Token: token, UserId: strconv.Itoa(int(userInfo.ID))}, nil
}
