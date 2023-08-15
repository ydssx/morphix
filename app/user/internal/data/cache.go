package data

import (
	"context"

	"github.com/ydssx/morphix/app/user/internal/biz"
	"github.com/ydssx/morphix/app/user/internal/models"
)

type UserRepoWithCache struct {
	*userRepo
}

func NewUserRepoWithCache(repo *userRepo) biz.UserRepoWithCache {
	return &UserRepoWithCache{repo}
}

func (u *UserRepoWithCache) ListUser(ctx context.Context) []models.User {
	return u.userRepo.ListUser(ctx)
}
