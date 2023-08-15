package data

import (
	"context"
	"fmt"
	"time"

	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
	"github.com/ydssx/morphix/pkg/cache"
)

type UserRepoCacheDecorator struct {
	*userRepo
	cache.Cache
}

func NewUserRepoCacheDecorator(repo *userRepo, cache cache.Cache) biz.UserRepoWithCache {
	return &UserRepoCacheDecorator{repo, cache}
}

func (u *UserRepoCacheDecorator) ListUser(ctx context.Context) []models.User {
	return u.userRepo.ListUser(ctx)
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
