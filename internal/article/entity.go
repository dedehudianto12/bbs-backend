package article

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID `json:"id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Excerpt     string    `json:"excerpt"`
	Content     string    `json:"content"`
	Image       *string   `json:"image"`
	Tag         string    `json:"tag"`
	PublishedAt string    `json:"publishedAt"` // YYYY-MM-DD
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
