package article

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound     = errors.New("artikel tidak ditemukan")
	ErrTitleRequired = errors.New("judul artikel wajib diisi")
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List(ctx context.Context) ([]Article, error) {
	return u.repo.FindAll(ctx)
}

func (u *Usecase) GetByID(ctx context.Context, id string) (*Article, error) {
	a, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}
	return a, nil
}

func (u *Usecase) GetBySlug(ctx context.Context, slug string) (*Article, error) {
	a, err := u.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, mapError(err)
	}
	return a, nil
}

func (u *Usecase) Create(ctx context.Context, a *Article) error {
	if a.Title == "" {
		return ErrTitleRequired
	}
	a.Slug = Slugify(a.Title)
	return u.repo.Create(ctx, a)
}

func (u *Usecase) Update(ctx context.Context, id string, a *Article) (*Article, error) {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}

	if a.Title != "" {
		existing.Title = a.Title
		existing.Slug = Slugify(a.Title)
	}
	if a.Excerpt != "" {
		existing.Excerpt = a.Excerpt
	}
	if a.Content != "" {
		existing.Content = a.Content
	}
	if a.Image != nil {
		existing.Image = a.Image
	}
	if a.Tag != "" {
		existing.Tag = a.Tag
	}
	if a.PublishedAt != "" {
		existing.PublishedAt = a.PublishedAt
	}
	if a.Author != "" {
		existing.Author = a.Author
	}

	if err := u.repo.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (u *Usecase) Delete(ctx context.Context, id string) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return mapError(err)
	}
	return nil
}

func mapError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	return err
}
