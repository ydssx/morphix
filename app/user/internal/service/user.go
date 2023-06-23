package service

import (
	"context"

	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GreeterService is a greeter service.
type UserService struct {
	userv1.UnimplementedUserServiceServer

	uc  *biz.UserUsecase
	sms smsv1.SMSServiceClient
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, sms smsv1.SMSServiceClient) *UserService {
	return &UserService{uc: uc, sms: sms}
}

// Register 实现用户注册接口
func (s *UserService) Register(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	// 调用业务逻辑层的方法进行用户注册处理
	user, err := s.uc.RegisterUser(ctx, req.Username, req.Password, req.Email, req.Phone)
	if err != nil {
		// 注册失败，返回错误信息
		return nil, err
	}

	// 将业务层返回的用户对象转换为接口定义的 User 对象，并返回
	response := &userv1.User{
		Id:       "123",
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		// 其他个人信息字段...
	}

	return response, nil
}

// Login 实现用户登录接口
func (s *UserService) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.AuthenticationResponse, error) {
	// 在这里实现用户登录逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// Logout 实现用户登出接口
func (s *UserService) Logout(ctx context.Context, req *userv1.LogoutRequest) (*emptypb.Empty, error) {
	// 在这里实现用户登出逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// UpdateProfile 实现更新用户信息接口
func (s *UserService) UpdateProfile(ctx context.Context, req *userv1.UpdateProfileRequest) (*userv1.User, error) {
	// 在这里实现更新用户信息逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// ResetPassword 实现重置密码接口
func (s *UserService) ResetPassword(ctx context.Context, req *userv1.ResetPasswordRequest) (*emptypb.Empty, error) {
	// 在这里实现重置密码逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// Authenticate 实现用户身份认证接口
func (s *UserService) Authenticate(ctx context.Context, req *emptypb.Empty) (*userv1.AuthenticationResponse, error) {
	// 在这里实现用户身份认证逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// Authorize 实现用户授权接口
func (s *UserService) Authorize(ctx context.Context, req *userv1.AuthorizationRequest) (*emptypb.Empty, error) {
	// 在这里实现用户授权逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// GetUserList 实现获取用户列表接口
func (s *UserService) GetUserList(ctx context.Context, req *emptypb.Empty) (*userv1.UserListResponse, error) {
	
	return s.uc.GetUserList(ctx, req) 
}

// ManageUserPermission 实现管理用户权限接口
func (s *UserService) ManageUserPermission(ctx context.Context, req *userv1.ManageUserPermissionRequest) (*userv1.User, error) {
	// 在这里实现管理用户权限逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}

// LogActivity 实现记录用户活动接口
func (s *UserService) LogActivity(ctx context.Context, req *userv1.LogEntry) (*emptypb.Empty, error) {
	// 在这里实现记录用户活动逻辑
	// 使用 s.userRepository 进行数据库操作
	return nil, nil
}
