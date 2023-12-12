package biz

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserRepo 是用户仓库接口
type UserRepo interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, user *models.User) (userId int, err error)
	// UpdateUser 更新用户
	UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, id uint) error
	// ListUser 获取用户列表
	ListUser(ctx context.Context, cond *ListUserCond) []models.User
	// GetUsersByRole 根据角色ID获取用户列表
	GetUsersByRole(ctx context.Context, roleID int) ([]models.User, error)
	// GetUserByID 根据用户ID获取用户
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	// GetUserByName 根据用户名获取用户
	GetUserByName(ctx context.Context, username string) (*models.User, error)
	// GetUserByPhone 根据手机号码获取用户
	GetUserByPhone(ctx context.Context, phoneNumber string) (*models.User, error)

	AddUserPermission(ctx context.Context, userID int, permissionID int) error
	// GetUserPermission 根据用户ID获取用户权限列表
	GetUserPermission(ctx context.Context, userID int) ([]models.Permission, error)
	// DeleteUserPermission 删除用户权限
	DeleteUserPermission(ctx context.Context, userID int, permissionID int) error
	// GetUserPermissionByRole 根据角色ID获取用户权限列表
	GetUserPermissionByRole(ctx context.Context, roleID int) ([]models.Permission, error)
	// AddRolePermission 添加角色权限
	AddRolePermission(ctx context.Context, roleID int, permissionID int) error
	// DeleteRolePermission 删除角色权限
	DeleteRolePermission(ctx context.Context, roleID int, permissionID int) error
	// GetRolePermission 根据角色ID获取角色权限列表
	GetRolePermission(ctx context.Context, roleID int) ([]models.Permission, error)

	// AddUserRole 添加用户角色
	AddUserRole(ctx context.Context, userID int, roleID int) error
	// GetUserRole 根据用户ID获取用户角色列表
	GetUserRole(ctx context.Context, userID int) ([]models.Role, error)
	// DeleteUserRole 删除用户角色
	DeleteUserRole(ctx context.Context, userID int, roleID int) error
}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type ListUserCond struct {
	Page  int64
	Limit int64
	Phone string
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
	sms  smsv1.SMSServiceClient
	tm   Transaction
}

func NewUserUseCase(userRepo UserRepo, logger log.Logger, smsClient smsv1.SMSServiceClient, transaction Transaction) *UserUseCase {
	return &UserUseCase{
		repo: userRepo,
		log:  log.NewHelper(logger),
		sms:  smsClient,
		tm:   transaction,
	}
}

// Register is a function that handles user registration.
// It takes a context and a registration request as input and returns a user object and an error.
func (uc *UserUseCase) Register(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	// Check SMS status
	checkResult, err := uc.sms.CheckSMSStatus(ctx, &smsv1.QuerySMSStatusRequest{MobileNumber: req.Phone, SmsCode: req.SmsCode, Scene: smsv1.SmsScene_USER_REGISTER})
	if err != nil {
		return nil, err
	}
	if !checkResult.Status {
		return nil, errors.New("failed to check SMS verification code")
	}

	// Validate username and password
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username and password cannot be empty")
	}

	// Validate phone number
	if !util.IsPhoneNumber(req.Phone) {
		return nil, errors.New("incorrect phone number format")
	}

	// Create user object
	user := &models.User{
		Username: req.Username,
		Password: util.MD5(req.Password),
		Email:    req.Email,
		Phone:    req.Phone,
	}

	// Create user in the database
	userId, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = uint(userId)

	// Create response object
	response := &userv1.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return response, nil
}

// ListUser retrieves a paginated list of users from the repository.
// It takes a request with pagination parameters and returns a list response containing
// user objects and any error.
func (uc *UserUseCase) ListUser(ctx context.Context, req *userv1.UserListRequest) (*userv1.UserListResponse, error) {
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

// Login authenticates a user and generates a token.
func (uc *UserUseCase) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.AuthenticationResponse, error) {
	// Check if either the username or phone number is empty
	if req.Username == "" || req.PhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "用户名和手机号不能同时为空")
	}

	var userInfo *models.User
	var err error
	if req.Username != "" {
		// Get user information by username
		userInfo, err = uc.repo.GetUserByName(ctx, req.Username)
	} else {
		// Get user information by phone number
		userInfo, err = uc.repo.GetUserByPhone(ctx, req.Username)
	}
	if err != nil {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	// Check if the provided password matches the stored password
	if util.MD5(req.Password) != userInfo.Password {
		return nil, status.Error(codes.InvalidArgument, "密码错误")
	}

	// Generate a token for the user
	token, err := jwt.GenerateToken(int64(userInfo.ID), userInfo.Username, "")
	if err != nil {
		return nil, status.Error(codes.Internal, "token生成失败")
	}

	// Return the authentication response with the token and user ID
	return &userv1.AuthenticationResponse{Token: token, UserId: strconv.Itoa(int(userInfo.ID))}, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	user, err := uc.repo.GetUserByID(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	return &userv1.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}

// ResetPassword resets the password for a user after verifying the SMS verification code.
// It checks the verification code status via SMS service, gets the user by username,
// and updates the user's password in the repository.
func (uc *UserUseCase) ResetPassword(ctx context.Context, req *userv1.ResetPasswordRequest) error {
	// Check the SMS status for verification code
	checkResult, err := uc.sms.CheckSMSStatus(ctx, &smsv1.QuerySMSStatusRequest{
		MobileNumber: "",
		SmsCode:      req.VerificationCode,
		Scene:        smsv1.SmsScene_USER_RESET_PASSWORD,
	})
	if err != nil {
		return err
	}

	// If SMS verification fails, return an error
	if !checkResult.Status {
		return errors.New("failed to verify SMS code")
	}

	// Get the user by username
	user, err := uc.repo.GetUserByName(ctx, req.Username)
	if err != nil {
		return err
	}

	// Update the user's password
	return uc.repo.UpdateUser(ctx, &models.User{
		BaseModel: models.BaseModel{
			ID: user.ID,
		},
		Password: util.MD5(req.NewPassword),
	})
}

// UpdateProfile updates the email and phone number for the authenticated user.
// It gets the user ID from the authentication claims in the context, updates
// the user in the repository, and returns the updated user object.
func (uc *UserUseCase) UpdateProfile(ctx context.Context, req *userv1.UpdateProfileRequest) (*userv1.User, error) {
	// Get the claims from the context
	claims, _ := interceptors.AuthFromContext(ctx)

	// Update the user in the repository
	err := uc.repo.UpdateUser(ctx, &models.User{
		BaseModel: models.BaseModel{
			ID: uint(claims.Uid),
		},
		Email: req.Email,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	// Get the updated user from the database
	return uc.GetUser(ctx, &userv1.GetUserRequest{
		UserId: claims.Uid,
	})
}

func (uc *UserUseCase) Logout(ctx context.Context, req *userv1.LogoutRequest) (res *emptypb.Empty, err error) {
	res = new(emptypb.Empty)

	// TODO:ADD logic here and delete this line.

	return
}

func (uc *UserUseCase) Authenticate(ctx context.Context, req *emptypb.Empty) (res *userv1.AuthenticationResponse, err error) {
	res = new(userv1.AuthenticationResponse)

	// TODO:ADD logic here and delete this line.

	return
}

func (uc *UserUseCase) Authorize(ctx context.Context, req *userv1.AuthorizationRequest) (res *emptypb.Empty, err error) {
	res = new(emptypb.Empty)

	return
}

func (uc *UserUseCase) GetUserList(ctx context.Context, req *userv1.UserListRequest) (res *userv1.UserListResponse, err error) {
	res = new(userv1.UserListResponse)

	// TODO:ADD logic here and delete this line.

	return
}

func (uc *UserUseCase) ManageUserPermission(ctx context.Context, req *userv1.ManageUserPermissionRequest) (res *userv1.User, err error) {
	res = new(userv1.User)

	// TODO:ADD logic here and delete this line.

	return
}

func (uc *UserUseCase) LogActivity(ctx context.Context, req *userv1.LogEntry) (res *emptypb.Empty, err error) {
	res = new(emptypb.Empty)
	return
}
