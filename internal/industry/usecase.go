package industry

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound     = errors.New("industri tidak ditemukan")
	ErrNameRequired = errors.New("nama industri wajib diisi")
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List(ctx context.Context) ([]Industry, error) {
	return u.repo.FindAll(ctx)
}

func (u *Usecase) GetByID(ctx context.Context, id string) (*Industry, error) {
	ind, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}
	return ind, nil
}

func (u *Usecase) GetBySlug(ctx context.Context, slug string) (*Industry, error) {
	ind, err := u.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, mapError(err)
	}
	return ind, nil
}

func (u *Usecase) Create(ctx context.Context, ind *Industry) error {
	if ind.Name == "" {
		return ErrNameRequired
	}
	ind.Slug = Slugify(ind.Name)
	return u.repo.Create(ctx, ind)
}

func (u *Usecase) Update(ctx context.Context, id string, ind *Industry) (*Industry, error) {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}

	if ind.Name != "" {
		existing.Name = ind.Name
		existing.Slug = Slugify(ind.Name)
	}
	if ind.Description != "" {
		existing.Description = ind.Description
	}
	existing.Image = ind.Image // allow clearing
	if ind.ProductSlugs != nil {
		existing.ProductSlugs = ind.ProductSlugs
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
