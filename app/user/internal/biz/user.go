package biz

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User, tx ...*gorm.DB) (userId int, err error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	ListUser(ctx context.Context, cond *ListUserCond) []models.User
	GetUsersByRole(ctx context.Context, roleID int) ([]models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByName(ctx context.Context, username string) (*models.User, error)
	GetUserByPhone(ctx context.Context, phoneNumber string) (*models.User, error)
}

type ListUserCond struct {
	Page  int64
	Limit int64
	Phone string
}

type UserRepoWithCache interface {
	UserRepo
	cache.Cache
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

func (uc *UserUsecase) ListUser(ctx context.Context, req *userv1.UserListRequest) (*userv1.UserListResponse, error) {
	users := uc.repo.ListUser(ctx, &ListUserCond{Page: req.Page, Limit: req.Limit})

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
	if req.Username == "" || req.PhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "用户名和手机号不能同时为空")
	}

	var userInfo *models.User
	var err error
	if req.Username != "" {
		userInfo, err = uc.repo.GetUserByName(ctx, req.Username)
	} else {
		userInfo, err = uc.repo.GetUserByPhone(ctx, req.Username)
	}
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

func (uc *UserUsecase) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	user, err := uc.repo.GetUserByID(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return &userv1.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Email,
	}, nil
}

func (uc *UserUsecase) ResetPassword(ctx context.Context, req *userv1.ResetPasswordRequest) error {
	user, err := uc.repo.GetUserByName(ctx, req.Username)
	if err != nil {
		return err
	}

	return uc.repo.UpdateUser(ctx, &models.User{BaseModel: models.BaseModel{ID: user.ID}, Password: util.MD5(req.NewPassword)})
}
