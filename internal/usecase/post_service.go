package usecase

import (
	"context"

	"github.com/iuan95/bilis_backend/internal/entity"
)


type PostRepository interface{
	GetById(id int) (*entity.Post, error)
	Create(post *entity.Post) error
}


type PostService struct{
	Repo PostRepository
}

func NewPostService(repo PostRepository) *PostService{
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context,post *entity.Post) error{
	return s.Repo.Create(post)
}

func (s *PostService) GetPostById(ctx context.Context, id int) (*entity.Post, error){
	return s.Repo.GetById(id)
}