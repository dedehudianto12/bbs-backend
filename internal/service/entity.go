package service

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID               uuid.UUID `json:"id"`
	Slug             string    `json:"slug"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"shortDescription"`
	FullDescription  string    `json:"fullDescription"`
	Images           []string  `json:"images"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
