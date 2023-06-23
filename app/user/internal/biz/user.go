package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
	Phone    string
}

type UserRepo interface {
	CreateUser(context.Context, *User) error
	ListUser(context.Context) ([]User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// RegisterUser 用户注册逻辑
func (uc *UserUsecase) RegisterUser(ctx context.Context, username, password, email, phone string) (*User, error) {
	// TODO: 执行用户注册的业务逻辑，例如验证参数、生成用户ID、存储用户信息等
	// 这里只是一个示例，你需要根据实际需求进行实现

	// 验证参数
	if username == "" || password == "" {
		return nil, nil
	}

	// 生成用户ID
	userID := 1

	// 创建用户对象
	user := &User{
		ID:       int64(userID),
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
		// 其他个人信息字段...
	}

	// 存储用户信息到数据仓库
	err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		// 存储失败，返回错误信息
		return nil, err
	}

	return user, nil
}

func (uc *UserUsecase) GetUserList(ctx context.Context, req *emptypb.Empty) (*userv1.UserListResponse, error) {
	users, err := uc.repo.ListUser(ctx)
	if err != nil {
		return nil, err
	}
	resp := new(userv1.UserListResponse)
	for _, v := range users {
		resp.Users = append(resp.Users, &userv1.User{
			Id:       "1",
			Username: v.Username,
			Password: v.Password,
			Email:    v.Email,
			Phone:    v.Phone,
		})
	}
	return resp, nil
}
