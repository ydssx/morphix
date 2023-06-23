package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.UserRepo = (*userRepo)(nil)

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser implements biz.UserRepo.
func (*userRepo) CreateUser(context.Context, *biz.User) error {
	panic("unimplemented")
}

// ListUser implements biz.UserRepo.
func (*userRepo) ListUser(context.Context) ([]biz.User, error) {
	return []biz.User{{
		ID:       1,
		Username: "ydssx",
		Password: "1234556",
		Email:    "456@qq.com",
		Phone:    "1562659746",
	}}, nil
}
