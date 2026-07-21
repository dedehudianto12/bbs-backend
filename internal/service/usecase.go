package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound     = errors.New("jasa tidak ditemukan")
	ErrNameRequired = errors.New("nama jasa wajib diisi")
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List(ctx context.Context) ([]Service, error) {
	return u.repo.FindAll(ctx)
}

func (u *Usecase) GetByID(ctx context.Context, id string) (*Service, error) {
	s, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}
	return s, nil
}

func (u *Usecase) GetBySlug(ctx context.Context, slug string) (*Service, error) {
	s, err := u.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, mapError(err)
	}
	return s, nil
}

func (u *Usecase) Create(ctx context.Context, s *Service) error {
	if s.Name == "" {
		return ErrNameRequired
	}
	s.Slug = Slugify(s.Name)
	return u.repo.Create(ctx, s)
}

func (u *Usecase) Update(ctx context.Context, id string, s *Service) (*Service, error) {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, mapError(err)
	}

	if s.Name != "" {
		existing.Name = s.Name
		existing.Slug = Slugify(s.Name)
	}
	if s.ShortDescription != "" {
		existing.ShortDescription = s.ShortDescription
	}
	if s.FullDescription != "" {
		existing.FullDescription = s.FullDescription
	}
	if s.Images != nil {
		existing.Images = s.Images
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
