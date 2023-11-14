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

func NewCasbinEnforcer() *casbin.Enforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/")
		if err != nil {
			panic(err)
		}
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

type RbacManager struct {
	enforcer *casbin.Enforcer
}

func NewRbacManager(dsn string) *RbacManager {
	return &RbacManager{enforcer: NewCasbinEnforcer()}
}

func (r *RbacManager) AddUserRoles(userID int, roles ...string) error {
	_, err := r.enforcer.AddRolesForUser(strconv.Itoa(userID), roles)
	return err
}

func (r *RbacManager) DeleteUserRoles(userID int, roles ...string) error {
	var errs []error

	for _, role := range roles {
		_, err := r.enforcer.DeleteRoleForUser(strconv.Itoa(userID), role)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (r *RbacManager) DeleteRoles(roles ...string) error {
	var errs []error

	for _, role := range roles {
		_, err := r.enforcer.DeleteRole(role)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (r *RbacManager) CheckPermission(userID int, path, meth string) error {
	userRoles, err := r.enforcer.GetRolesForUser(strconv.Itoa(userID))
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

func (r *RbacManager) AddRolePermission(role, path, meth string) error {
	_, err := r.enforcer.AddPolicy(role, path, meth)
	return err
}
