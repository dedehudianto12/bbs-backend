package industry

import (
	"time"

	"github.com/google/uuid"
)

type Industry struct {
	ID           uuid.UUID `json:"id"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Image        *string   `json:"image"`
	ProductSlugs []string  `json:"productSlugs"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
