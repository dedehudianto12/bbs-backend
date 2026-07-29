package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Service, error)
	FindAllAdmin(ctx context.Context, search, sort string, page, limit int) ([]Service, int, error)
	FindByID(ctx context.Context, id string) (*Service, error)
	FindBySlug(ctx context.Context, slug string) (*Service, error)
	Create(ctx context.Context, s *Service) error
	Update(ctx context.Context, s *Service) error
	Delete(ctx context.Context, id string) error
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindAll(ctx context.Context) ([]Service, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, slug, name, short_description, full_description, images, created_at, updated_at FROM services ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []Service
	for rows.Next() {
		var s Service
		if err := rows.Scan(&s.ID, &s.Slug, &s.Name, &s.ShortDescription, &s.FullDescription, &s.Images, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return services, nil
}

func (r *pgxRepo) FindAllAdmin(ctx context.Context, search, sort string, page, limit int) ([]Service, int, error) {
	var total int
	countQ := `SELECT count(*) FROM services WHERE 1=1`
	dataQ := `SELECT id, slug, name, short_description, full_description, images, created_at, updated_at FROM services WHERE 1=1`
	args := []any{}
	idx := 1

	if search != "" {
		cond := fmt.Sprintf(` AND (name ILIKE $%d OR slug ILIKE $%d)`, idx, idx)
		countQ += cond
		dataQ += cond
		args = append(args, "%"+search+"%")
		idx++
	}

	if err := r.pool.QueryRow(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	order := "DESC"
	if sort == "asc" {
		order = "ASC"
	}
	dataQ += fmt.Sprintf(` ORDER BY created_at %s LIMIT $%d OFFSET $%d`, order, idx, idx+1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, dataQ, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var services []Service
	for rows.Next() {
		var s Service
		if err := rows.Scan(&s.ID, &s.Slug, &s.Name, &s.ShortDescription, &s.FullDescription, &s.Images, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, 0, err
		}
		services = append(services, s)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return services, total, nil
}

func (r *pgxRepo) FindByID(ctx context.Context, id string) (*Service, error) {
	var s Service
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, name, short_description, full_description, images, created_at, updated_at FROM services WHERE id = $1`, id,
	).Scan(&s.ID, &s.Slug, &s.Name, &s.ShortDescription, &s.FullDescription, &s.Images, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *pgxRepo) FindBySlug(ctx context.Context, slug string) (*Service, error) {
	var s Service
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, name, short_description, full_description, images, created_at, updated_at FROM services WHERE slug = $1`, slug,
	).Scan(&s.ID, &s.Slug, &s.Name, &s.ShortDescription, &s.FullDescription, &s.Images, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *pgxRepo) Create(ctx context.Context, s *Service) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO services (name, slug, short_description, full_description, images) VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at, updated_at`,
		s.Name, s.Slug, s.ShortDescription, s.FullDescription, s.Images,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

func (r *pgxRepo) Update(ctx context.Context, s *Service) error {
	return r.pool.QueryRow(ctx,
		`UPDATE services SET name=$1, slug=$2, short_description=$3, full_description=$4, images=$5, updated_at=now() WHERE id=$6 RETURNING updated_at`,
		s.Name, s.Slug, s.ShortDescription, s.FullDescription, s.Images, s.ID,
	).Scan(&s.UpdatedAt)
}

func (r *pgxRepo) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM services WHERE id = $1`, id)
	return err
}


