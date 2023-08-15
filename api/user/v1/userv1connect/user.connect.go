// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/user/v1/user.proto

package userv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ydssx/morphix/api/user/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "userv1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceRegisterProcedure is the fully-qualified name of the UserService's Register RPC.
	UserServiceRegisterProcedure = "/userv1.UserService/Register"
	// UserServiceLoginProcedure is the fully-qualified name of the UserService's Login RPC.
	UserServiceLoginProcedure = "/userv1.UserService/Login"
	// UserServiceLogoutProcedure is the fully-qualified name of the UserService's Logout RPC.
	UserServiceLogoutProcedure = "/userv1.UserService/Logout"
	// UserServiceUpdateProfileProcedure is the fully-qualified name of the UserService's UpdateProfile
	// RPC.
	UserServiceUpdateProfileProcedure = "/userv1.UserService/UpdateProfile"
	// UserServiceResetPasswordProcedure is the fully-qualified name of the UserService's ResetPassword
	// RPC.
	UserServiceResetPasswordProcedure = "/userv1.UserService/ResetPassword"
	// UserServiceAuthenticateProcedure is the fully-qualified name of the UserService's Authenticate
	// RPC.
	UserServiceAuthenticateProcedure = "/userv1.UserService/Authenticate"
	// UserServiceAuthorizeProcedure is the fully-qualified name of the UserService's Authorize RPC.
	UserServiceAuthorizeProcedure = "/userv1.UserService/Authorize"
	// UserServiceGetUserListProcedure is the fully-qualified name of the UserService's GetUserList RPC.
	UserServiceGetUserListProcedure = "/userv1.UserService/GetUserList"
	// UserServiceManageUserPermissionProcedure is the fully-qualified name of the UserService's
	// ManageUserPermission RPC.
	UserServiceManageUserPermissionProcedure = "/userv1.UserService/ManageUserPermission"
	// UserServiceLogActivityProcedure is the fully-qualified name of the UserService's LogActivity RPC.
	UserServiceLogActivityProcedure = "/userv1.UserService/LogActivity"
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/userv1.UserService/GetUser"
)

// UserServiceClient is a client for the userv1.UserService service.
type UserServiceClient interface {
	// 用户注册
	Register(context.Context, *connect_go.Request[v1.RegistrationRequest]) (*connect_go.Response[v1.User], error)
	// 用户登录
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.AuthenticationResponse], error)
	Logout(context.Context, *connect_go.Request[v1.LogoutRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdateProfile(context.Context, *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.User], error)
	ResetPassword(context.Context, *connect_go.Request[v1.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	Authenticate(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.AuthenticationResponse], error)
	Authorize(context.Context, *connect_go.Request[v1.AuthorizationRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetUserList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.UserListResponse], error)
	ManageUserPermission(context.Context, *connect_go.Request[v1.ManageUserPermissionRequest]) (*connect_go.Response[v1.User], error)
	LogActivity(context.Context, *connect_go.Request[v1.LogEntry]) (*connect_go.Response[emptypb.Empty], error)
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error)
}

// NewUserServiceClient constructs a client for the userv1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		register: connect_go.NewClient[v1.RegistrationRequest, v1.User](
			httpClient,
			baseURL+UserServiceRegisterProcedure,
			opts...,
		),
		login: connect_go.NewClient[v1.LoginRequest, v1.AuthenticationResponse](
			httpClient,
			baseURL+UserServiceLoginProcedure,
			opts...,
		),
		logout: connect_go.NewClient[v1.LogoutRequest, emptypb.Empty](
			httpClient,
			baseURL+UserServiceLogoutProcedure,
			opts...,
		),
		updateProfile: connect_go.NewClient[v1.UpdateProfileRequest, v1.User](
			httpClient,
			baseURL+UserServiceUpdateProfileProcedure,
			opts...,
		),
		resetPassword: connect_go.NewClient[v1.ResetPasswordRequest, emptypb.Empty](
			httpClient,
			baseURL+UserServiceResetPasswordProcedure,
			opts...,
		),
		authenticate: connect_go.NewClient[emptypb.Empty, v1.AuthenticationResponse](
			httpClient,
			baseURL+UserServiceAuthenticateProcedure,
			opts...,
		),
		authorize: connect_go.NewClient[v1.AuthorizationRequest, emptypb.Empty](
			httpClient,
			baseURL+UserServiceAuthorizeProcedure,
			opts...,
		),
		getUserList: connect_go.NewClient[emptypb.Empty, v1.UserListResponse](
			httpClient,
			baseURL+UserServiceGetUserListProcedure,
			opts...,
		),
		manageUserPermission: connect_go.NewClient[v1.ManageUserPermissionRequest, v1.User](
			httpClient,
			baseURL+UserServiceManageUserPermissionProcedure,
			opts...,
		),
		logActivity: connect_go.NewClient[v1.LogEntry, emptypb.Empty](
			httpClient,
			baseURL+UserServiceLogActivityProcedure,
			opts...,
		),
		getUser: connect_go.NewClient[v1.GetUserRequest, v1.User](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	register             *connect_go.Client[v1.RegistrationRequest, v1.User]
	login                *connect_go.Client[v1.LoginRequest, v1.AuthenticationResponse]
	logout               *connect_go.Client[v1.LogoutRequest, emptypb.Empty]
	updateProfile        *connect_go.Client[v1.UpdateProfileRequest, v1.User]
	resetPassword        *connect_go.Client[v1.ResetPasswordRequest, emptypb.Empty]
	authenticate         *connect_go.Client[emptypb.Empty, v1.AuthenticationResponse]
	authorize            *connect_go.Client[v1.AuthorizationRequest, emptypb.Empty]
	getUserList          *connect_go.Client[emptypb.Empty, v1.UserListResponse]
	manageUserPermission *connect_go.Client[v1.ManageUserPermissionRequest, v1.User]
	logActivity          *connect_go.Client[v1.LogEntry, emptypb.Empty]
	getUser              *connect_go.Client[v1.GetUserRequest, v1.User]
}

// Register calls userv1.UserService.Register.
func (c *userServiceClient) Register(ctx context.Context, req *connect_go.Request[v1.RegistrationRequest]) (*connect_go.Response[v1.User], error) {
	return c.register.CallUnary(ctx, req)
}

// Login calls userv1.UserService.Login.
func (c *userServiceClient) Login(ctx context.Context, req *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.AuthenticationResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// Logout calls userv1.UserService.Logout.
func (c *userServiceClient) Logout(ctx context.Context, req *connect_go.Request[v1.LogoutRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.logout.CallUnary(ctx, req)
}

// UpdateProfile calls userv1.UserService.UpdateProfile.
func (c *userServiceClient) UpdateProfile(ctx context.Context, req *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.User], error) {
	return c.updateProfile.CallUnary(ctx, req)
}

// ResetPassword calls userv1.UserService.ResetPassword.
func (c *userServiceClient) ResetPassword(ctx context.Context, req *connect_go.Request[v1.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.resetPassword.CallUnary(ctx, req)
}

// Authenticate calls userv1.UserService.Authenticate.
func (c *userServiceClient) Authenticate(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.AuthenticationResponse], error) {
	return c.authenticate.CallUnary(ctx, req)
}

// Authorize calls userv1.UserService.Authorize.
func (c *userServiceClient) Authorize(ctx context.Context, req *connect_go.Request[v1.AuthorizationRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.authorize.CallUnary(ctx, req)
}

// GetUserList calls userv1.UserService.GetUserList.
func (c *userServiceClient) GetUserList(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.UserListResponse], error) {
	return c.getUserList.CallUnary(ctx, req)
}

// ManageUserPermission calls userv1.UserService.ManageUserPermission.
func (c *userServiceClient) ManageUserPermission(ctx context.Context, req *connect_go.Request[v1.ManageUserPermissionRequest]) (*connect_go.Response[v1.User], error) {
	return c.manageUserPermission.CallUnary(ctx, req)
}

// LogActivity calls userv1.UserService.LogActivity.
func (c *userServiceClient) LogActivity(ctx context.Context, req *connect_go.Request[v1.LogEntry]) (*connect_go.Response[emptypb.Empty], error) {
	return c.logActivity.CallUnary(ctx, req)
}

// GetUser calls userv1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error) {
	return c.getUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the userv1.UserService service.
type UserServiceHandler interface {
	// 用户注册
	Register(context.Context, *connect_go.Request[v1.RegistrationRequest]) (*connect_go.Response[v1.User], error)
	// 用户登录
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.AuthenticationResponse], error)
	Logout(context.Context, *connect_go.Request[v1.LogoutRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdateProfile(context.Context, *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.User], error)
	ResetPassword(context.Context, *connect_go.Request[v1.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	Authenticate(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.AuthenticationResponse], error)
	Authorize(context.Context, *connect_go.Request[v1.AuthorizationRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetUserList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.UserListResponse], error)
	ManageUserPermission(context.Context, *connect_go.Request[v1.ManageUserPermissionRequest]) (*connect_go.Response[v1.User], error)
	LogActivity(context.Context, *connect_go.Request[v1.LogEntry]) (*connect_go.Response[emptypb.Empty], error)
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userServiceRegisterHandler := connect_go.NewUnaryHandler(
		UserServiceRegisterProcedure,
		svc.Register,
		opts...,
	)
	userServiceLoginHandler := connect_go.NewUnaryHandler(
		UserServiceLoginProcedure,
		svc.Login,
		opts...,
	)
	userServiceLogoutHandler := connect_go.NewUnaryHandler(
		UserServiceLogoutProcedure,
		svc.Logout,
		opts...,
	)
	userServiceUpdateProfileHandler := connect_go.NewUnaryHandler(
		UserServiceUpdateProfileProcedure,
		svc.UpdateProfile,
		opts...,
	)
	userServiceResetPasswordHandler := connect_go.NewUnaryHandler(
		UserServiceResetPasswordProcedure,
		svc.ResetPassword,
		opts...,
	)
	userServiceAuthenticateHandler := connect_go.NewUnaryHandler(
		UserServiceAuthenticateProcedure,
		svc.Authenticate,
		opts...,
	)
	userServiceAuthorizeHandler := connect_go.NewUnaryHandler(
		UserServiceAuthorizeProcedure,
		svc.Authorize,
		opts...,
	)
	userServiceGetUserListHandler := connect_go.NewUnaryHandler(
		UserServiceGetUserListProcedure,
		svc.GetUserList,
		opts...,
	)
	userServiceManageUserPermissionHandler := connect_go.NewUnaryHandler(
		UserServiceManageUserPermissionProcedure,
		svc.ManageUserPermission,
		opts...,
	)
	userServiceLogActivityHandler := connect_go.NewUnaryHandler(
		UserServiceLogActivityProcedure,
		svc.LogActivity,
		opts...,
	)
	userServiceGetUserHandler := connect_go.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		opts...,
	)
	return "/userv1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceRegisterProcedure:
			userServiceRegisterHandler.ServeHTTP(w, r)
		case UserServiceLoginProcedure:
			userServiceLoginHandler.ServeHTTP(w, r)
		case UserServiceLogoutProcedure:
			userServiceLogoutHandler.ServeHTTP(w, r)
		case UserServiceUpdateProfileProcedure:
			userServiceUpdateProfileHandler.ServeHTTP(w, r)
		case UserServiceResetPasswordProcedure:
			userServiceResetPasswordHandler.ServeHTTP(w, r)
		case UserServiceAuthenticateProcedure:
			userServiceAuthenticateHandler.ServeHTTP(w, r)
		case UserServiceAuthorizeProcedure:
			userServiceAuthorizeHandler.ServeHTTP(w, r)
		case UserServiceGetUserListProcedure:
			userServiceGetUserListHandler.ServeHTTP(w, r)
		case UserServiceManageUserPermissionProcedure:
			userServiceManageUserPermissionHandler.ServeHTTP(w, r)
		case UserServiceLogActivityProcedure:
			userServiceLogActivityHandler.ServeHTTP(w, r)
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) Register(context.Context, *connect_go.Request[v1.RegistrationRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.Register is not implemented"))
}

func (UnimplementedUserServiceHandler) Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.AuthenticationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.Login is not implemented"))
}

func (UnimplementedUserServiceHandler) Logout(context.Context, *connect_go.Request[v1.LogoutRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.Logout is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateProfile(context.Context, *connect_go.Request[v1.UpdateProfileRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.UpdateProfile is not implemented"))
}

func (UnimplementedUserServiceHandler) ResetPassword(context.Context, *connect_go.Request[v1.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.ResetPassword is not implemented"))
}

func (UnimplementedUserServiceHandler) Authenticate(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.AuthenticationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.Authenticate is not implemented"))
}

func (UnimplementedUserServiceHandler) Authorize(context.Context, *connect_go.Request[v1.AuthorizationRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.Authorize is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUserList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.UserListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.GetUserList is not implemented"))
}

func (UnimplementedUserServiceHandler) ManageUserPermission(context.Context, *connect_go.Request[v1.ManageUserPermissionRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.ManageUserPermission is not implemented"))
}

func (UnimplementedUserServiceHandler) LogActivity(context.Context, *connect_go.Request[v1.LogEntry]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.LogActivity is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.User], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("userv1.UserService.GetUser is not implemented"))
}
