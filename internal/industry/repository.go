package industry

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Industry, error)
	FindByID(ctx context.Context, id string) (*Industry, error)
	FindBySlug(ctx context.Context, slug string) (*Industry, error)
	Create(ctx context.Context, i *Industry) error
	Update(ctx context.Context, i *Industry) error
	Delete(ctx context.Context, id string) error
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindAll(ctx context.Context) ([]Industry, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, slug, name, description, image, created_at, updated_at FROM industries ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var industries []Industry
	for rows.Next() {
		var i Industry
		if err := rows.Scan(&i.ID, &i.Slug, &i.Name, &i.Description, &i.Image, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, err
		}
		industries = append(industries, i)
	}

	// Load product slugs untuk semua industries
	for idx := range industries {
		slugs, _ := r.getProductSlugs(ctx, industries[idx].ID)
		industries[idx].ProductSlugs = slugs
	}

	return industries, nil
}

func (r *pgxRepo) FindByID(ctx context.Context, id string) (*Industry, error) {
	var i Industry
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, name, description, image, created_at, updated_at FROM industries WHERE id = $1`, id,
	).Scan(&i.ID, &i.Slug, &i.Name, &i.Description, &i.Image, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		return nil, err
	}

	i.ProductSlugs, _ = r.getProductSlugs(ctx, i.ID)
	return &i, nil
}

func (r *pgxRepo) FindBySlug(ctx context.Context, slug string) (*Industry, error) {
	var i Industry
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, name, description, image, created_at, updated_at FROM industries WHERE slug = $1`, slug,
	).Scan(&i.ID, &i.Slug, &i.Name, &i.Description, &i.Image, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		return nil, err
	}

	i.ProductSlugs, _ = r.getProductSlugs(ctx, i.ID)
	return &i, nil
}

func (r *pgxRepo) Create(ctx context.Context, i *Industry) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx,
		`INSERT INTO industries (name, slug, description, image) VALUES ($1,$2,$3,$4) RETURNING id, created_at, updated_at`,
		i.Name, i.Slug, i.Description, i.Image,
	).Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		return err
	}

	if err := r.syncProductSlugs(ctx, tx, i.ID, i.ProductSlugs); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *pgxRepo) Update(ctx context.Context, i *Industry) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx,
		`UPDATE industries SET name=$1, slug=$2, description=$3, image=$4, updated_at=now() WHERE id=$5 RETURNING updated_at`,
		i.Name, i.Slug, i.Description, i.Image, i.ID,
	).Scan(&i.UpdatedAt)
	if err != nil {
		return err
	}

	if err := r.syncProductSlugs(ctx, tx, i.ID, i.ProductSlugs); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *pgxRepo) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM industries WHERE id = $1`, id)
	return err
}

func (r *pgxRepo) getProductSlugs(ctx context.Context, industryID uuid.UUID) ([]string, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT product_slug FROM industry_products WHERE industry_id = $1 ORDER BY product_slug`, industryID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slugs []string
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}
		slugs = append(slugs, s)
	}
	return slugs, nil
}

func (r *pgxRepo) syncProductSlugs(ctx context.Context, tx pgx.Tx, industryID uuid.UUID, slugs []string) error {
	_, err := tx.Exec(ctx, `DELETE FROM industry_products WHERE industry_id = $1`, industryID)
	if err != nil {
		return err
	}

	for _, slug := range slugs {
		_, err := tx.Exec(ctx,
			`INSERT INTO industry_products (industry_id, product_slug) VALUES ($1,$2) ON CONFLICT DO NOTHING`,
			industryID, slug,
		)
		if err != nil {
			return err
		}
	}
	return nil
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
