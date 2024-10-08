package usecase

import (
	"context"

	"github.com/iuan95/bilis_backend/internal/entity"
)

type UserRepository interface{
	GetById(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context,user *entity.User) error
}

type UserService struct{
	Repo UserRepository
}

func NewUserService(repo UserRepository) *UserService{
	return &UserService{Repo: repo}
}

func (u *UserService) Create(ctx context.Context,post *entity.User) error{
	return u.Repo.Create(ctx, post)
}

func (u *UserService) GetById(ctx context.Context,id int) (*entity.User, error){
	return u.Repo.GetById(ctx,id)
}