package category

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context, group string) ([]Category, error)
	UpdateLabel(ctx context.Context, slug, label string) error
	Create(ctx context.Context, c *Category) error
	Delete(ctx context.Context, slug string) error
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindAll(ctx context.Context, group string) ([]Category, error) {
	query := `SELECT slug, label, "group", sort_order FROM categories`
	args := []any{}

	if group != "" {
		query += ` WHERE "group" = $1`
		args = append(args, group)
	}
	query += ` ORDER BY sort_order`

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.Slug, &c.Label, &c.Group, &c.SortOrder); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}

func (r *pgxRepo) UpdateLabel(ctx context.Context, slug, label string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE categories SET label = $1 WHERE slug = $2`,
		label, slug,
	)
	return err
}

func (r *pgxRepo) Create(ctx context.Context, c *Category) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO categories (slug, label, "group", sort_order) VALUES ($1,$2,$3,$4)`,
		c.Slug, c.Label, c.Group, c.SortOrder,
	)
	return err
}

func (r *pgxRepo) Delete(ctx context.Context, slug string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM categories WHERE slug = $1`, slug)
	return err
}
