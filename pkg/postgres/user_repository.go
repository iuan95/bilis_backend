package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/iuan95/bilis_backend/internal/entity"
	"github.com/iuan95/bilis_backend/internal/usecase"
)

type UserRepository struct {
    pool *pgxpool.Pool
}
func NewUserRepository(pool *pgxpool.Pool) usecase.UserRepository {
    return &UserRepository{pool: pool}
}
func (r *UserRepository) GetById(ctx context.Context,id int) (*entity.User, error) {
	user:= &entity.User{}
	query:= "SELECT id, name, email FROM users WHERE id=$1"
    err := r.pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
	return user, nil
}

func (r *UserRepository) Create(ctx context.Context,user *entity.User) error {
	query := "INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id"
    return r.pool.QueryRow(context.Background(), query,user.Name, user.Password, user.Email).Scan(&user.ID)

}
