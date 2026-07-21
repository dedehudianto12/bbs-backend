package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Admin struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Claims struct {
	AdminID uuid.UUID `json:"adminId"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	jwt.RegisteredClaims
}
