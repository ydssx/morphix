package biz

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	smsv1 "github.com/ydssx/morphix/api/sms/v1"
	userv1 "github.com/ydssx/morphix/api/user/v1"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/interceptors"
	"github.com/ydssx/morphix/pkg/jwt"
	"github.com/ydssx/morphix/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (userId int, err error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	ListUser(ctx context.Context, cond *ListUserCond) []models.User
	GetUsersByRole(ctx context.Context, roleID int) ([]models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByName(ctx context.Context, username string) (*models.User, error)
	GetUserByPhone(ctx context.Context, phoneNumber string) (*models.User, error)
}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
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

type UserUseCase struct {
	repo UserRepoWithCache
	log  *log.Helper
	sms  smsv1.SMSServiceClient
	tm   Transaction
}

func NewUserUseCase(repo UserRepoWithCache, logger log.Logger, sms smsv1.SMSServiceClient, tm Transaction) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger), sms: sms, tm: tm}
}

// 用户注册
func (uc *UserUseCase) Register(ctx context.Context, req *userv1.RegistrationRequest) (*userv1.User, error) {
	checkResult, err := uc.sms.CheckSMSStatus(ctx, &smsv1.QuerySMSStatusRequest{MobileNumber: req.Phone, SmsCode: req.SmsCode, Scene: smsv1.SmsScene_USER_REGISTER})
	if err != nil {
		return nil, err
	}
	if !checkResult.Status {
		return nil, errors.New("校验短信验证码失败")
	}

	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	if !util.IsPhoneNumber(req.Phone) {
		return nil, errors.New("手机号格式不正确")
	}

	// 创建用户对象
	user := &models.User{
		Username: req.Username,
		Password: util.MD5(req.Password),
		Email:    req.Email,
		Phone:    req.Phone,
	}

	userId, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = uint(userId)

	response := &userv1.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return response, nil
}

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

// 用户登录
func (uc *UserUseCase) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.AuthenticationResponse, error) {
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

func (uc *UserUseCase) ResetPassword(ctx context.Context, req *userv1.ResetPasswordRequest) error {
	checkResult, err := uc.sms.CheckSMSStatus(ctx, &smsv1.QuerySMSStatusRequest{MobileNumber: "", SmsCode: req.VerificationCode, Scene: smsv1.SmsScene_USER_RESET_PASSWORD})
	if err != nil {
		return err
	}
	if !checkResult.Status {
		return errors.New("校验短信验证码失败")
	}

	user, err := uc.repo.GetUserByName(ctx, req.Username)
	if err != nil {
		return err
	}

	return uc.repo.UpdateUser(ctx, &models.User{BaseModel: models.BaseModel{ID: user.ID}, Password: util.MD5(req.NewPassword)})
}

func (uc *UserUseCase) UpdateProfile(ctx context.Context, req *userv1.UpdateProfileRequest) (*userv1.User, error) {
	claims, _ := interceptors.AuthFromContext(ctx)

	err := uc.repo.UpdateUser(ctx, &models.User{BaseModel: models.BaseModel{ID: uint(claims.Uid)}, Email: req.Email, Phone: req.Phone})
	if err != nil {
		return nil, err
	}

	return uc.GetUser(ctx, &userv1.GetUserRequest{UserId: claims.Uid})
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

	// TODO:ADD logic here and delete this line.

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

	// TODO:ADD logic here and delete this line.

	return
}
