package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ydssx/morphix/app/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.UserRepo = (*userRepo)(nil)

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
