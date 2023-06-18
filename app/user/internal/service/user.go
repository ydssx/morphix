package service

import (
	"context"

	user "github.com/ydssx/morphix/app/user/api"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GreeterService is a greeter service.
type UserService struct {
	user.UnimplementedUserServiceServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// 用户注册
func (uc *UserService) Register(_ context.Context, _ *user.RegistrationRequest) (*user.User, error) {
	panic("not implemented") // TODO: Implement
}

// 用户登录
func (uc *UserService) Login(_ context.Context, _ *user.LoginRequest) (*user.AuthenticationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) Logout(_ context.Context, _ *user.LogoutRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) UpdateProfile(_ context.Context, _ *user.UpdateProfileRequest) (*user.User, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) ResetPassword(_ context.Context, _ *user.ResetPasswordRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) Authenticate(_ context.Context, _ *emptypb.Empty) (*user.AuthenticationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) Authorize(_ context.Context, _ *user.AuthorizationRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) GetUserList(ctx context.Context, _ *emptypb.Empty) (*user.UserListResponse, error) {
	uc.uc.ListUser(ctx, nil)
	return &user.UserListResponse{Users: []*user.User{{
		Id:       "1",
		Username: "wangxin",
		Password: "123456",
		Email:    "",
		Phone:    "",
	}}}, nil
}

func (uc *UserService) ManageUserPermission(_ context.Context, _ *user.ManageUserPermissionRequest) (*user.User, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) LogActivity(_ context.Context, _ *user.LogEntry) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}
