package gallery

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Gallery, error)
	FindByID(ctx context.Context, id string) (*Gallery, error)
	Create(ctx context.Context, g *Gallery) error
	Update(ctx context.Context, g *Gallery) error
	Delete(ctx context.Context, id string) error
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindAll(ctx context.Context) ([]Gallery, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, image, caption, location, created_at, updated_at FROM galleries ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var galleries []Gallery
	for rows.Next() {
		var g Gallery
		if err := rows.Scan(&g.ID, &g.Image, &g.Caption, &g.Location, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, err
		}
		galleries = append(galleries, g)
	}
	return galleries, nil
}

func (r *pgxRepo) FindByID(ctx context.Context, id string) (*Gallery, error) {
	var g Gallery
	err := r.pool.QueryRow(ctx,
		`SELECT id, image, caption, location, created_at, updated_at FROM galleries WHERE id = $1`, id,
	).Scan(&g.ID, &g.Image, &g.Caption, &g.Location, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *pgxRepo) Create(ctx context.Context, g *Gallery) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO galleries (image, caption, location) VALUES ($1,$2,$3) RETURNING id, created_at, updated_at`,
		g.Image, g.Caption, g.Location,
	).Scan(&g.ID, &g.CreatedAt, &g.UpdatedAt)
}

func (r *pgxRepo) Update(ctx context.Context, g *Gallery) error {
	return r.pool.QueryRow(ctx,
		`UPDATE galleries SET image=$1, caption=$2, location=$3, updated_at=now() WHERE id=$4 RETURNING updated_at`,
		g.Image, g.Caption, g.Location, g.ID,
	).Scan(&g.UpdatedAt)
}

func (r *pgxRepo) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM galleries WHERE id = $1`, id)
	return err
}
