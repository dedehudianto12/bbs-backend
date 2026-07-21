package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("email atau password salah")
	ErrInvalidToken       = errors.New("token tidak valid")
)

type Usecase struct {
	repo      Repository
	jwtSecret []byte
}

func NewUsecase(repo Repository, jwtSecret string) *Usecase {
	return &Usecase{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (u *Usecase) Login(ctx context.Context, email, password string) (string, *Admin, error) {
	admin, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}


	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := u.generateJWT(admin)
	if err != nil {
		return "", nil, fmt.Errorf("gagal generate token: %w", err)
	}

	return token, admin, nil
}

func (u *Usecase) Me(tokenString string) (*Admin, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return u.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return &Admin{
		ID:    claims.AdminID,
		Email: claims.Email,
		Name:  claims.Name,
	}, nil
}

func (u *Usecase) generateJWT(admin *Admin) (string, error) {
	claims := &Claims{
		AdminID: admin.ID,
		Email:   admin.Email,
		Name:    admin.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(u.jwtSecret)
}
