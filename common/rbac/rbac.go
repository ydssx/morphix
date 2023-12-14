package rbac

import (
	"errors"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	enforcer *casbin.Enforcer
	once     sync.Once
)

// NewCasbinEnforcer 创建并返回一个新的 CasbinEnforcer 实例。
func NewCasbinEnforcer() *casbin.Enforcer {
	once.Do(func() {
		// 创建 GORM 的适配器。
		a, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/")
		if err != nil {
			panic(err)
		}
		// 创建 Casbin 的 Enforcer 实例。
		e, err := casbin.NewEnforcer("rbac_model.conf", a)
		if err != nil {
			panic(err)
		}
		e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
		e.LoadPolicy()
		enforcer = e
	})
	return enforcer
}

// RbacManager 是 RBAC 管理器的结构体。
type RbacManager struct {
	enforcer *casbin.Enforcer
}

// NewRbacManager 创建并返回一个新的 RbacManager 实例。
func NewRbacManager(dsn string) *RbacManager {
	return &RbacManager{enforcer: NewCasbinEnforcer()}
}

// AddUserRoles 为用户添加角色。
func (r *RbacManager) AddUserRoles(userID int, roles ...string) error {
	_, err := r.enforcer.AddRolesForUser(strconv.Itoa(userID), roles)
	return err
}

// DeleteUserRoles 删除用户的角色。
func (r *RbacManager) DeleteUserRoles(userID int, roles ...string) error {
	var errs []error

	for _, role := range roles {
		_, err := r.enforcer.DeleteRoleForUser(strconv.Itoa(userID), role)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

// DeleteRoles 删除角色。
func (r *RbacManager) DeleteRoles(roles ...string) error {
	var errs []error

	for _, role := range roles {
		_, err := r.enforcer.DeleteRole(role)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

// CheckPermission checks if the given user has permission to access the specified path and method.
// It first checks if the user has direct permission.
// If not, it retrieves the roles for the user and checks if any role has permission.
// If no permission is found, an error is returned indicating insufficient permissions.
func (r *RbacManager) CheckPermission(userID int, path, meth string) error {
	x := strconv.Itoa(userID)
	ok, _ := r.enforcer.Enforce(x, path, meth)
	if ok {
		return nil
	}

	userRoles, err := r.enforcer.GetRolesForUser(x)
	if err != nil {
		return err
	}

	var errs []error
	for _, role := range userRoles {
		ok, err := r.enforcer.Enforce(role, path, meth)
		if ok {
			return nil
		}
		errs = append(errs, err)
	}

	err = errors.Join(errs...)
	if err != nil {
		return err
	}

	return errors.New("权限不足")
}

// AddRolePermission adds a permission rule to the specified role.
// It takes in the role name, resource path, and access method as strings.
// It calls the enforcer's AddPolicy method to add the permission rule.
// Returns any error from the enforcer.
func (r *RbacManager) AddRolePermission(role, path, meth string) error {
	_, err := r.enforcer.AddPolicy(role, path, meth)
	return err
}

// AddUserPermission adds a permission rule for the given user ID.
// It takes the user ID as an integer, plus the resource path and access method as strings.
// It calls the enforcer's AddPolicy method to add the permission rule for that user.
// Returns any error from the enforcer.
func (r *RbacManager) AddUserPermission(userID int, path, meth string) error {
	_, err := r.enforcer.AddPolicy(strconv.Itoa(userID), path, meth)
	return err
}

// DeleteRolePermission 删除角色的权限。
func (r *RbacManager) DeleteRolePermission(role, path, meth string) error {
	_, err := r.enforcer.RemovePolicy(role, path, meth)
	return err
}

func (r *RbacManager) DeleteUserPermission(userID int, path, meth string) error {
	_, err := r.enforcer.RemovePolicy(strconv.Itoa(userID), path, meth)
	return err
}
