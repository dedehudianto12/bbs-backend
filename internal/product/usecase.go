package product

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound       = errors.New("produk tidak ditemukan")
	ErrSlugDuplicate  = errors.New("slug sudah digunakan")
	ErrNameRequired   = errors.New("nama produk wajib diisi")
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List(ctx context.Context, group, kategori string) ([]Product, error) {
	return u.repo.FindAll(ctx, group, kategori)
}

func (u *Usecase) GetByID(ctx context.Context, id string) (*Product, error) {
	p, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}
	return p, nil
}

func (u *Usecase) GetBySlug(ctx context.Context, slug string) (*Product, error) {
	p, err := u.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, mapError(err)
	}
	return p, nil
}

func (u *Usecase) Create(ctx context.Context, p *Product) error {
	if p.Name == "" {
		return ErrNameRequired
	}
	p.Slug = Slugify(p.Name)
	return u.repo.Create(ctx, p)
}

func (u *Usecase) Update(ctx context.Context, id string, p *Product) (*Product, error) {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}

	if p.Name != "" {
		existing.Name = p.Name
		existing.Slug = Slugify(p.Name)
	}
	if p.Group != "" {
		existing.Group = p.Group
	}
	if p.Kategori != "" {
		existing.Kategori = p.Kategori
	}
	if p.Category != "" {
		existing.Category = p.Category
	}
	existing.Description = p.Description
	existing.Detail = p.Detail
	existing.Image = p.Image
	if p.Specs != nil {
		existing.Specs = p.Specs
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
