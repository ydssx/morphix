package biz

import (
	"context"
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

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	return uc.repo.Save(ctx, g)
}
