package category

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("kategori tidak ditemukan")
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetAll(ctx context.Context, group string) ([]Category, error) {
	return u.repo.FindAll(ctx, group)
}

func (u *Usecase) UpdateLabel(ctx context.Context, slug, label string) (*Category, error) {
	err := u.repo.UpdateLabel(ctx, slug, label)
	if err != nil {
		return nil, err
	}

	// Ambil data terbaru — karena hanya update label, return single category
	cats, err := u.repo.FindAll(ctx, "")
	if err != nil {
		return nil, err
	}
	for _, c := range cats {
		if c.Slug == slug {
			return &c, nil
		}
	}
	return nil, ErrNotFound
}
