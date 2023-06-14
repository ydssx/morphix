package service

import (
	user "github.com/ydssx/morphix/app/user/api"
	"github.com/ydssx/morphix/app/user/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	user.UnimplementedUserServiceServer

	uc *biz.GreeterUsecase
}

// NewUserService new a greeter service.
func NewUserService() *UserService {
	return &UserService{}
}
