package article

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Article, error)
	FindByID(ctx context.Context, id string) (*Article, error)
	FindBySlug(ctx context.Context, slug string) (*Article, error)
	Create(ctx context.Context, a *Article) error
	Update(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id string) error
}

type pgxRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgxRepo{pool: pool}
}

func (r *pgxRepo) FindAll(ctx context.Context) ([]Article, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, slug, title, excerpt, content, image, tag, published_at, author, created_at, updated_at FROM articles ORDER BY published_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		if err := rows.Scan(&a.ID, &a.Slug, &a.Title, &a.Excerpt, &a.Content, &a.Image, &a.Tag, &a.PublishedAt, &a.Author, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func (r *pgxRepo) FindByID(ctx context.Context, id string) (*Article, error) {
	var a Article
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, title, excerpt, content, image, tag, published_at, author, created_at, updated_at FROM articles WHERE id = $1`, id,
	).Scan(&a.ID, &a.Slug, &a.Title, &a.Excerpt, &a.Content, &a.Image, &a.Tag, &a.PublishedAt, &a.Author, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgxRepo) FindBySlug(ctx context.Context, slug string) (*Article, error) {
	var a Article
	err := r.pool.QueryRow(ctx,
		`SELECT id, slug, title, excerpt, content, image, tag, published_at, author, created_at, updated_at FROM articles WHERE slug = $1`, slug,
	).Scan(&a.ID, &a.Slug, &a.Title, &a.Excerpt, &a.Content, &a.Image, &a.Tag, &a.PublishedAt, &a.Author, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *pgxRepo) Create(ctx context.Context, a *Article) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO articles (title, slug, excerpt, content, image, tag, published_at, author) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id, created_at, updated_at`,
		a.Title, a.Slug, a.Excerpt, a.Content, a.Image, a.Tag, a.PublishedAt, a.Author,
	).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
}

func (r *pgxRepo) Update(ctx context.Context, a *Article) error {
	return r.pool.QueryRow(ctx,
		`UPDATE articles SET title=$1, slug=$2, excerpt=$3, content=$4, image=$5, tag=$6, published_at=$7, author=$8, updated_at=now() WHERE id=$9 RETURNING updated_at`,
		a.Title, a.Slug, a.Excerpt, a.Content, a.Image, a.Tag, a.PublishedAt, a.Author, a.ID,
	).Scan(&a.UpdatedAt)
}

func (r *pgxRepo) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM articles WHERE id = $1`, id)
	return err
}

func Slugify(title string) string {
	s := strings.ToLower(title)
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
