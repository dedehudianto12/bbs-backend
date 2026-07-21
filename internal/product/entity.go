package product

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID       `json:"id"`
	Slug        string          `json:"slug"`
	Name        string          `json:"name"`
	Group       string          `json:"group"`
	Kategori    string          `json:"kategori"`
	Category    string          `json:"category"`
	Description string          `json:"description"`
	Detail      string          `json:"detail"`
	Image       *string         `json:"image"`
	Specs       json.RawMessage `json:"specs"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}
