package data

import (
	"context"
	"fmt"
	"time"

	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/cache"
	"github.com/ydssx/morphix/pkg/logger"
	"github.com/ydssx/morphix/pkg/util"
)

type UserRepoCacheDecorator struct {
	*userRepo
	cache.Cache
}

// NewUserRepoCacheDecorator creates a new UserRepoCacheDecorator, which wraps
// a userRepo with caching capabilities using the provided cache.Cache
// implementation.
func NewUserRepoCacheDecorator(repo *userRepo, cache cache.Cache) biz.UserRepo {
	return &UserRepoCacheDecorator{repo, cache}
}

// ListUser retrieves users from cache if available, otherwise from database.
// It caches the retrieved users for 1 hour.
func (u *UserRepoCacheDecorator) ListUser(ctx context.Context, cond *biz.ListUserCond) (data []models.User) {
	key := fmt.Sprintf("user.list:%v", util.CalculateChecksum(cond))
	err := u.Get(key, &data)
	if err == nil {
		return
	}

	data = u.userRepo.ListUser(ctx, cond)

	err = u.Set(key, data, time.Hour)
	if err != nil {
		logger.Errorf(ctx, "缓存用户列表失败：%v", err)
	}

	return
}

// GetUserByID 从缓存中获取用户数据。
// 如果缓存中不存在,则从数据库中获取用户数据,并设置缓存。
// 缓存时间为 1 小时。
func (u *UserRepoCacheDecorator) GetUserByID(ctx context.Context, id uint) (data *models.User, err error) {
	key := fmt.Sprintf("user:%v", id)
	err = u.Get(key, &data)
	if err == nil {
		return
	}

	data, err = u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = u.Set(key, data, time.Hour)
	if err != nil {
		return nil, err
	}

	return
}

// UpdateUser updates the user in the database and invalidates the cache.
func (u *UserRepoCacheDecorator) UpdateUser(ctx context.Context, updatedUser *models.User) error {
	if err := u.userRepo.UpdateUser(ctx, updatedUser); err != nil {
		return err
	}

	key := fmt.Sprintf("user:%v", updatedUser.ID)
	err := u.Delete(key)
	if err != nil {
		logger.Errorf(ctx, "删除缓存失败：%v", err)
	}

	return nil
}
