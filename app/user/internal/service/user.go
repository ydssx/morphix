package service

import (
	"context"

	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GreeterService is a greeter service.
type UserService struct {
	userv1.UnimplementedUserServiceServer

	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

// 用户注册
func (s *UserService) Register(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	return s.uc.Register(ctx, req)
}

// 用户登录
func (s *UserService) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.AuthenticationResponse, error) {
	return s.uc.Login(ctx, req)
}

func (s *UserService) UpdateProfile(ctx context.Context, req *userv1.UpdateProfileRequest) (*userv1.User, error) {
	return s.uc.UpdateProfile(ctx, req)
}

// 用户重置密码
func (s *UserService) ResetPassword(ctx context.Context, req *userv1.ResetPasswordRequest) (*emptypb.Empty, error) {
	if err := s.uc.ResetPassword(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserService) GetUserList(ctx context.Context, req *userv1.UserListRequest) (*userv1.UserListResponse, error) {
	return s.uc.ListUser(ctx, req)
}

func (s *UserService) ManageUserPermission(ctx context.Context, req *userv1.ManageUserPermissionRequest) (*userv1.User, error) {
	return s.uc.ManageUserPermission(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	return s.uc.GetUser(ctx, req)
}

func (s *UserService) Logout(ctx context.Context, req *userv1.LogoutRequest) (res *emptypb.Empty, err error) {
	return s.uc.Logout(ctx, req)
}

func (s *UserService) Authenticate(ctx context.Context, req *emptypb.Empty) (res *userv1.AuthenticationResponse, err error) {
	return s.uc.Authenticate(ctx, req)
}

func (s *UserService) Authorize(ctx context.Context, req *userv1.AuthorizationRequest) (res *emptypb.Empty, err error) {
	return s.uc.Authorize(ctx, req)
}

func (s *UserService) LogActivity(ctx context.Context, req *userv1.LogEntry) (res *emptypb.Empty, err error) {
	return s.uc.LogActivity(ctx, req)
}

func (s *UserService) GetUserActivity(ctx context.Context, req *userv1.GetUserActivityRequest) (res *userv1.UserActivityListResponse, err error) {
	return s.uc.GetUserActivity(ctx, req)
}

