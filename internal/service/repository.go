package service

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Service, error)
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
	return services, nil
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

func Slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
