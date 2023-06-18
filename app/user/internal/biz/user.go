package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/pkg/logger"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) ListUser(ctx context.Context, g *Greeter) (*Greeter, error) {
	logger.Info(ctx, "用户列表")
	logger.Infof(ctx, "用户列表:%v",123)
	uc.repo.ListAll(ctx)
	return nil, nil
}
