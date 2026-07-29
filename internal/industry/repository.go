package industry

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Industry, error)
	FindAllAdmin(ctx context.Context, search, sort string, page, limit int) ([]Industry, int, error)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// ponytail: batch fetch product slugs instead of N+1 per-industry queries
	slugMap, _ := r.batchGetProductSlugs(ctx, industryIDs(industries))
	for idx := range industries {
		industries[idx].ProductSlugs = slugMap[industries[idx].ID]
	}

	return industries, nil
}

func (r *pgxRepo) FindAllAdmin(ctx context.Context, search, sort string, page, limit int) ([]Industry, int, error) {
	var total int
	countQ := `SELECT count(*) FROM industries WHERE 1=1`
	dataQ := `SELECT id, slug, name, description, image, created_at, updated_at FROM industries WHERE 1=1`
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

	var industries []Industry
	for rows.Next() {
		var i Industry
		if err := rows.Scan(&i.ID, &i.Slug, &i.Name, &i.Description, &i.Image, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, 0, err
		}
		industries = append(industries, i)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	// ponytail: batch fetch product slugs instead of N+1 per-industry queries
	slugMap, _ := r.batchGetProductSlugs(ctx, industryIDs(industries))
	for idx := range industries {
		industries[idx].ProductSlugs = slugMap[industries[idx].ID]
	}

	return industries, total, nil
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

func industryIDs(industries []Industry) []uuid.UUID {
	ids := make([]uuid.UUID, len(industries))
	for i, ind := range industries {
		ids[i] = ind.ID
	}
	return ids
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slugs, nil
}

// batchGetProductSlugs fetches product slugs for multiple industry IDs in one query.
func (r *pgxRepo) batchGetProductSlugs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID][]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	rows, err := r.pool.Query(ctx,
		`SELECT industry_id, product_slug FROM industry_products WHERE industry_id = ANY($1) ORDER BY product_slug`,
		ids,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[uuid.UUID][]string)
	for rows.Next() {
		var industryID uuid.UUID
		var slug string
		if err := rows.Scan(&industryID, &slug); err != nil {
			return nil, err
		}
		m[industryID] = append(m[industryID], slug)
	}
	return m, rows.Err()
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


