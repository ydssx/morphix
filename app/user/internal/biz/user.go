package biz

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
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

	// AddUserPermission 添加用户权限
	AddUserPermission(ctx context.Context, userID int64, permissionID ...int64) error
	// GetUserPermission 根据用户ID获取用户权限列表
	GetUserPermission(ctx context.Context, userID int) ([]models.Permission, error)
	// DeleteUserPermission 删除用户权限
	DeleteUserPermission(ctx context.Context, userID int64, permissionID ...int64) error
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
	// ListRole 获取角色列表
	ListRole(ctx context.Context, cond *ListRoleCond) []models.Role

	// CreateUserActivity 创建用户活动
	CreateUserActivity(ctx context.Context, userActivity *models.UserActivity) error
	// GetUserActivity 根据用户ID获取用户活动列表
	GetUserActivity(ctx context.Context, userID int, page int64, limit int64) ([]models.UserActivity, error)
}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type ListUserCond struct {
	Page  int64
	Limit int64
	Phone string
}

type ListRoleCond struct {
	Page  int64
	Limit int64
	Name  string
}

type UserUseCase struct {
	repo   UserRepo
	log    *log.Helper
	sms    smsv1.SMSServiceClient
	tm     Transaction
	jobCli jobv1.JobServiceClient
}

func NewUserUseCase(userRepo UserRepo, logger log.Logger, smsClient smsv1.SMSServiceClient, transaction Transaction, jobClient jobv1.JobServiceClient) *UserUseCase {
	return &UserUseCase{
		repo:   userRepo,
		log:    log.NewHelper(logger),
		sms:    smsClient,
		tm:     transaction,
		jobCli: jobClient,
	}
}

// Register is a function that handles user registration.
// It takes a context and a registration request as input and returns a user object and an error.
func (uc *UserUseCase) Register(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	switch req.RegisterType {
	case userv1.RegistrationRequest_SMS:
		return uc.registerBySMS(ctx, req)
	case userv1.RegistrationRequest_PASSWORD:
		return uc.registerByPassword(ctx, req)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid register type")
	}
}

// registerBySMS handles user registration by SMS verification code.
// It takes a context and registration request as input, checks if the phone number is already registered,
// verifies the SMS code, creates a new user in the database, and returns the created user object and any error.
// This is part of the user business logic use case.
func (uc *UserUseCase) registerBySMS(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	// Check if the phone number is already registered
	if _, err := uc.repo.GetUserByPhone(ctx, req.Phone); err == nil {
		return nil, status.Error(codes.AlreadyExists, "phone number already registered")
	}

	// Check SMS verification code
	if checkResult, err := uc.sms.CheckSMSStatus(ctx, &smsv1.QuerySMSStatusRequest{
		MobileNumber: req.Phone,
		SmsCode:      req.SmsCode,
		Scene:        smsv1.SmsScene_USER_REGISTER,
	}); err != nil {
		return nil, err
	} else if !checkResult.Status {
		return nil, errors.New("failed to check SMS verification code")
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

// registerByPassword 根据密码注册用户。
// 接收注册请求的上下文和请求参数,检查用户名和手机号是否已被注册,
// 创建用户对象,在数据库中创建用户,并返回创建的用户对象和错误。
// 这是用户业务逻辑用例的一部分。
func (uc *UserUseCase) registerByPassword(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	// Check if the username is already registered
	if _, err := uc.repo.GetUserByName(ctx, req.Username); err == nil {
		return nil, status.Error(codes.AlreadyExists, "username already registered")
	}

	// Check if the phone number is already registered
	if _, err := uc.repo.GetUserByPhone(ctx, req.Phone); err == nil {
		return nil, status.Error(codes.AlreadyExists, "phone number already registered")
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

// GetUser 根据用户 ID 获取用户信息。
// 从仓库中根据用户 ID 查询用户,如果查询成功,返回用户信息对象,如果失败返回错误。
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
		return errors.New("failed to check SMS verification code")
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

// ManageUserPermission 根据请求更新用户的权限。
//
// 根据请求的模式,它可以添加、删除或更新用户的权限。
// 它通过存储库接口与数据层交互。
// 如果存储库操作失败,将返回错误。
// 成功后,它会重新获取用户信息并返回。
func (uc *UserUseCase) ManageUserPermission(ctx context.Context, req *userv1.ManageUserPermissionRequest) (res *userv1.User, err error) {
	switch req.Mode {
	case userv1.ManageUserPermissionRequest_USER_PERMISSION_ADD:
		err = uc.repo.AddUserPermission(ctx, req.UserId, req.PermissionIds...)
	case userv1.ManageUserPermissionRequest_USER_PERMISSION_DELETE:
		err = uc.repo.DeleteUserPermission(ctx, req.UserId, req.PermissionIds...)
	case userv1.ManageUserPermissionRequest_USER_PERMISSION_UPDATE:
		err = uc.repo.DeleteUserPermission(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		err = uc.repo.AddUserPermission(ctx, req.UserId, req.PermissionIds...)
	default:
		return nil, errors.New("invalid action")
	}
	if err != nil {
		return nil, err
	}
	return uc.GetUser(ctx, &userv1.GetUserRequest{UserId: req.UserId})
}

// LogActivity logs user activity.
// It creates a new user activity entry in the repository.
func (uc *UserUseCase) LogActivity(ctx context.Context, req *userv1.LogEntry) (res *emptypb.Empty, err error) {
	// Create a new user activity based on the request.
	activity := &models.UserActivity{
		UserID:   req.UserId,
		Action:   req.Action,
		Resource: req.Resource,
	}

	// Call the repository to create the user activity.
	err = uc.repo.CreateUserActivity(ctx, activity)
	if err != nil {
		return nil, err
	}

	// Return an empty response.
	res = new(emptypb.Empty)
	return
}

// GetUserActivity 获取用户活动列表。
//
// 该函数从存储库中获取指定用户ID的最近活动,并将其格式化为响应。
// 它接受用户ID,限制和偏移参数来控制返回的活动数量和范围。
//
// 返回的活动列表包含活动操作,资源和时间戳。
// 时间戳被格式化为易读的字符串。
//
// 如果存储库查询失败,将返回错误。
func (uc *UserUseCase) GetUserActivity(ctx context.Context, req *userv1.GetUserActivityRequest) (*userv1.UserActivityListResponse, error) {
	res := &userv1.UserActivityListResponse{
		Activity: make([]*userv1.UserActivity, 0),
	}

	activities, err := uc.repo.GetUserActivity(ctx, int(req.UserId), req.Page, req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user activities: %v", err)
	}

	for _, activity := range activities {
		timestamp := time.Unix(activity.ActionTime, 0).Format("2006-01-02 15:04:05")
		res.Activity = append(res.Activity, &userv1.UserActivity{
			Action:    activity.Action,
			Resource:  activity.Resource,
			Timestamp: timestamp,
		})
	}

	return res, nil
}
