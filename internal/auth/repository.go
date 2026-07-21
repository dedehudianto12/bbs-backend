package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*Admin, error)
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindByEmail(ctx context.Context, email string) (*Admin, error) {
	var a Admin
	err := r.pool.QueryRow(ctx,
		`SELECT id, email, password_hash, name, created_at FROM admins WHERE email = $1`,
		email,
	).Scan(&a.ID, &a.Email, &a.PasswordHash, &a.Name, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
