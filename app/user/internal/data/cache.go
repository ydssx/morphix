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

func NewUserRepoCacheDecorator(repo *userRepo, cache cache.Cache) biz.UserRepo {
	return &UserRepoCacheDecorator{repo, cache}
}

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

func (u *UserRepoCacheDecorator) UpdateUser(ctx context.Context, user *models.User) error {
	err := u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("user:%v", user.ID)
	err = u.Delete(key)
	return err
}
