package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// AddRolePermission implements biz.UserRepo.
func (r *userRepo) AddRolePermission(ctx context.Context, roleID int, permissionID int) error {
	return models.NewRolePermissionModel(r.data.DB(ctx)).Create(models.RolePermission{RoleId: int(roleID), PermissionId: int(permissionID)})
}

// DeleteRolePermission implements biz.UserRepo.
func (r *userRepo) DeleteRolePermission(ctx context.Context, roleID int, permissionID int) error {
	return models.NewRolePermissionModel(r.data.DB(ctx)).SetRoleId(uint(roleID)).SetPermissionId(uint(permissionID)).Delete()
}

// GetRolePermission implements biz.UserRepo.
func (*userRepo) GetRolePermission(ctx context.Context, roleID int) ([]models.Permission, error) {
	panic("unimplemented")
}

func NewUserRepo(data *Data, logger log.Logger) *userRepo {
	return &userRepo{data: data, log: log.NewHelper(logger)}
}

// GetUserByPhone implements biz.UserRepo.
func (r *userRepo) GetUserByPhone(ctx context.Context, phoneNumber string) (*models.User, error) {
	user, err := models.NewUserModel(r.data.DB(ctx)).SetPhoneNumber(phoneNumber).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// GetUserByName implements biz.UserRepo.
func (r *userRepo) GetUserByName(ctx context.Context, username string) (*models.User, error) {
	user, err := models.NewUserModel(r.data.DB(ctx)).SetUsername(username).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := models.NewUserModel(r.data.DB(ctx)).SetId(id).FirstOne()
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *models.User) error {
	return models.NewUserModel(r.data.DB(ctx)).SetId(user.ID).Update(user)
}

func (r *userRepo) DeleteUser(ctx context.Context, id uint) error {
	return models.NewUserModel(r.data.DB(ctx)).DeleteById(int(id))
}

func (r *userRepo) GetUsersByRole(ctx context.Context, roleID int) (result []models.User, err error) {
	return
}

// CreateUser implements biz.UserRepo.
func (r *userRepo) CreateUser(ctx context.Context, user *models.User) (userId int, err error) {
	userInfo, err := models.NewUserModel(r.data.DB(ctx)).Create(*user)
	if err != nil {
		return 0, err
	}
	return int(userInfo.ID), nil
}

func (r *userRepo) ListUser(ctx context.Context, cond *biz.ListUserCond) []models.User {
	if cond == nil {
		cond = new(biz.ListUserCond)
	}
	model := models.NewUserModel(r.data.DB(ctx)).WithContext(ctx)
	if cond.Phone != "" {
		model.PhoneNumberLike(cond.Phone)
	}
	if cond.Page == 0 {
		cond.Page = 1
	}
	if cond.Limit == 0 {
		cond.Limit = 10
	}

	users, _, _ := model.List(int(cond.Limit), (int(cond.Page)-1)*int(cond.Limit))

	return users
}

func (r *userRepo) AddUserPermission(ctx context.Context, userID int, permissionID int) error {
	return nil
}

func (r *userRepo) GetUserPermission(ctx context.Context, userID int) ([]models.Permission, error) {
	return nil, nil
}

func (r *userRepo) GetUserPermissionByRole(ctx context.Context, roleID int) ([]models.Permission, error) {
	return nil, nil
}

func (r *userRepo) AddUserRole(ctx context.Context, userID int, roleID int) error {
	return models.NewUserRoleModel(r.data.DB(ctx)).Create(models.UserRole{UserId: userID, RoleId: roleID})
}

func (r *userRepo) DeleteUserRole(ctx context.Context, userID int, roleID int) error {
	return nil
}

// GetUserRole retrieves the roles associated with the given user ID.
// It queries the user_roles table to get the role IDs for the user,
// then queries the roles table to hydrate the role objects.
// Returns a slice of Role structs and any error.
func (r *userRepo) GetUserRole(ctx context.Context, userID int) ([]models.Role, error) {
	userRoles, _, err := models.NewUserRoleModel(r.data.DB(ctx)).SetUserId(userID).List(0, 0)
	if err != nil {
		return nil, err
	}
	var roles []models.Role
	var roleIDs []int
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleId)
	}
	roles, _, err = models.NewRoleModel(r.data.DB(ctx)).SetIds(roleIDs...).List(0, 0)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *userRepo) DeleteUserPermission(ctx context.Context, userID int, permissionID int) error {
	return nil
}
