package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
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

// 用户注册
func (uc *UserService) Register(_ context.Context, _ *userv1.RegistrationRequest) (*userv1.User, error) {
	panic("not implemented") // TODO: Implement
}

// 用户登录
func (uc *UserService) Login(_ context.Context, _ *userv1.LoginRequest) (*userv1.AuthenticationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) Logout(_ context.Context, _ *userv1.LogoutRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) UpdateProfile(_ context.Context, _ *userv1.UpdateProfileRequest) (*userv1.User, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) ResetPassword(_ context.Context, _ *userv1.ResetPasswordRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) Authenticate(_ context.Context, _ *emptypb.Empty) (*userv1.AuthenticationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) Authorize(_ context.Context, _ *userv1.AuthorizationRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) GetUserList(ctx context.Context, _ *emptypb.Empty) (*userv1.UserListResponse, error) {
	res, _ := uc.sms.SendSMS(ctx, &smsv1.SendSMSRequest{
		MobileNumber:       "15623562713",
		Message:            "测试短信",
		SenderId:           "",
		TemplateId:         "",
		TemplateParameters: "",
	})
	if !res.Success {
		return nil, errors.New(401, "发送短信失败", "发送失败")
	}

	uc.uc.ListUser(ctx, nil)
	return &userv1.UserListResponse{Users: []*userv1.User{{
		Id:       "1",
		Username: "wangxin",
		Password: "123456",
		Email:    "",
		Phone:    "",
	}}}, nil
}

func (uc *UserService) ManageUserPermission(_ context.Context, _ *userv1.ManageUserPermissionRequest) (*userv1.User, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *UserService) LogActivity(_ context.Context, _ *userv1.LogEntry) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}
