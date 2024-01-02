package data

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *userRepo) GetRolePermission(ctx context.Context, roleID int) ([]models.Permission, error) {
	rolePermissionModel := models.NewRolePermissionModel(r.data.DB(ctx)).SetRoleId(uint(roleID)).List()

	if len(rolePermissionModel) == 0 {
		return nil, errors.New("role permission not found")
	}

	var rolePermissionIDs []int
	for _, rp := range rolePermissionModel {
		rolePermissionIDs = append(rolePermissionIDs, rp.PermissionId)
	}

	return models.NewPermissionModel(r.data.DB(ctx)).SetId(rolePermissionIDs...).List(), nil
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
	userIds, err := models.NewUserRoleModel(r.data.DB(ctx)).SetRoleId((roleID)).ListUserId()
	if err != nil {
		return nil, err
	}
	return models.NewUserModel(r.data.DB(ctx)).SetIds(userIds).ListAll()
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

	userModel := models.NewUserModel(r.data.DB(ctx)).WithContext(ctx)

	if cond.Phone != "" {
		userModel.PhoneNumberLike(cond.Phone)
	}

	if cond.Page == 0 {
		cond.Page = 1
	}

	if cond.Limit == 0 {
		cond.Limit = 10
	}

	users, _, _ := userModel.List(int(cond.Limit), (int(cond.Page)-1)*int(cond.Limit))

	return users
}

func (r *userRepo) AddUserPermission(ctx context.Context, userID int64, permissionID ...int64) error {
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
	return models.NewUserRoleModel(r.data.DB(ctx)).SetUserId(userID).SetRoleId(roleID).Delete()
}

// GetUserRole 根据用户ID获取用户角色
// 从数据库中查询用户角色关系表,提取角色ID
// 根据角色ID查询角色信息
// 返回角色列表
func (r *userRepo) GetUserRole(ctx context.Context, userID int) ([]models.Role, error) {
	userRoles, err := r.getUserRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	roleIDs := extractRoleIDs(userRoles)
	roles, err := r.getRoles(ctx, roleIDs)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *userRepo) getUserRoles(ctx context.Context, userID int) ([]models.UserRole, error) {
	return models.NewUserRoleModel(r.data.DB(ctx)).WithContext(ctx).SetUserId(userID).ListAll()
}

func extractRoleIDs(userRoles []models.UserRole) []int {
	var roleIDs []int
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleId)
	}
	return roleIDs
}

func (r *userRepo) getRoles(ctx context.Context, roleIDs []int) ([]models.Role, error) {
	return models.NewRoleModel(r.data.DB(ctx)).SetIds(roleIDs...).ListAll()
}

func (r *userRepo) DeleteUserPermission(ctx context.Context, userID int64, permissionID ...int64) error {
	return nil
}

func (r *userRepo) ListRole(ctx context.Context, cond *biz.ListRoleCond) []models.Role {
	if cond == nil {
		cond = &biz.ListRoleCond{}
	}

	model := models.NewRoleModel(r.data.DB(ctx)).WithContext(ctx)

	if cond.Name != "" {
		model.NameLike(cond.Name)
	}

	if cond.Page == 0 {
		cond.Page = 1
	}

	if cond.Limit == 0 {
		cond.Limit = 10
	}

	roles, _, _ := model.List(int(cond.Limit), (int(cond.Page)-1)*int(cond.Limit))

	return roles
}

func (r *userRepo) CreateUserActivity(ctx context.Context, userActivity *models.UserActivity) error {
	_, err := r.data.collection.InsertOne(ctx, userActivity)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) GetUserActivity(ctx context.Context, userID int, page, limit int64) ([]models.UserActivity, error) {
	// Create a slice to store the user activities
	var activities []models.UserActivity

	// Set up the options for the query
	findOptions := options.Find().
		SetSkip((page - 1) * limit).
		SetLimit(limit).
		SetSort(bson.D{{"create_at", -1}, {"_id", -1}})

	// Perform the query
	cursor, err := repo.data.collection.Find(ctx, bson.M{"user_id": userID}, findOptions)
	if err != nil {
		return nil, err
	}

	// Decode the results into the activities slice
	if err = cursor.All(ctx, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}
