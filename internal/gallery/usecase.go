package gallery

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound      = errors.New("galeri tidak ditemukan")
	ErrImageRequired = errors.New("gambar wajib diisi")
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List(ctx context.Context) ([]Gallery, error) {
	return u.repo.FindAll(ctx)
}

func (u *Usecase) ListAdmin(ctx context.Context, search, sort string, page, limit int) ([]Gallery, int, error) {
	return u.repo.FindAllAdmin(ctx, search, sort, page, limit)
}

func (u *Usecase) GetByID(ctx context.Context, id string) (*Gallery, error) {
	g, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}
	return g, nil
}

func (u *Usecase) Create(ctx context.Context, g *Gallery) error {
	if g.Image == "" {
		return ErrImageRequired
	}
	return u.repo.Create(ctx, g)
}

func (u *Usecase) Update(ctx context.Context, id string, g *Gallery) (*Gallery, error) {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}

	if g.Image != "" {
		existing.Image = g.Image
	}
	if g.Caption != "" {
		existing.Caption = g.Caption
	}
	existing.Location = g.Location // allow clearing

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
