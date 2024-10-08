package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/iuan95/bilis_backend/internal/entity"
	"github.com/iuan95/bilis_backend/internal/usecase"
)

type PostRepository struct {
    pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) usecase.PostRepository {
    return &PostRepository{pool: pool}
}

func (r *PostRepository) GetById(ctx context.Context,id int) (*entity.Post, error) {
	post:= &entity.Post{}
	query:= "SELECT id, title, description, date  FROM posts WHERE id=$1"
    err := r.pool.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.Title, &post.Description, &post.Date)
    if err != nil {
        return nil, err
    }
	return post, nil
}

func (r *PostRepository) Create(ctx context.Context,post *entity.Post) error {
	query := "INSERT INTO posts (title, description, date) VALUES ($1, $2, $3) RETURNING id"
    return r.pool.QueryRow(context.Background(), query, post.Title, post.Description, time.Now().UTC()).Scan(&post.ID)

}