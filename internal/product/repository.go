package product

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context, group, kategori string) ([]Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	FindBySlug(ctx context.Context, slug string) (*Product, error)
	Create(ctx context.Context, p *Product) error
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id string) error
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindAll(ctx context.Context, group, kategori string) ([]Product, error) {
	query := `SELECT id, slug, name, "group", kategori, category, description, detail, image, specs, created_at, updated_at FROM products WHERE 1=1`
	args := []any{}
	idx := 1

	if group != "" {
		query += fmt.Sprintf(` AND "group" = $%d`, idx)
		args = append(args, group)
		idx++
	}
	if kategori != "" {
		query += fmt.Sprintf(` AND kategori = $%d`, idx)
		args = append(args, kategori)
		idx++
	}
	query += ` ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Slug, &p.Name, &p.Group, &p.Kategori, &p.Category, &p.Description, &p.Detail, &p.Image, &p.Specs, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *pgxRepo) FindByID(ctx context.Context, id string) (*Product, error) {
	var p Product
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, name, "group", kategori, category, description, detail, image, specs, created_at, updated_at FROM products WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.Slug, &p.Name, &p.Group, &p.Kategori, &p.Category, &p.Description, &p.Detail, &p.Image, &p.Specs, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *pgxRepo) FindBySlug(ctx context.Context, slug string) (*Product, error) {
	var p Product
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, name, "group", kategori, category, description, detail, image, specs, created_at, updated_at FROM products WHERE slug = $1`,
		slug,
	).Scan(&p.ID, &p.Slug, &p.Name, &p.Group, &p.Kategori, &p.Category, &p.Description, &p.Detail, &p.Image, &p.Specs, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *pgxRepo) Create(ctx context.Context, p *Product) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO products (name, slug, "group", kategori, category, description, detail, image, specs) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id, created_at, updated_at`,
		p.Name, p.Slug, p.Group, p.Kategori, p.Category, p.Description, p.Detail, p.Image, p.Specs,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *pgxRepo) Update(ctx context.Context, p *Product) error {
	return r.pool.QueryRow(ctx,
		`UPDATE products SET name=$1, slug=$2, "group"=$3, kategori=$4, category=$5, description=$6, detail=$7, image=$8, specs=$9, updated_at=now() WHERE id=$10 RETURNING updated_at`,
		p.Name, p.Slug, p.Group, p.Kategori, p.Category, p.Description, p.Detail, p.Image, p.Specs, p.ID,
	).Scan(&p.UpdatedAt)
}

func (r *pgxRepo) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM products WHERE id = $1`, id)
	return err
}

func Slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")

	// Hapus karakter selain a-z, 0-9, dash
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
