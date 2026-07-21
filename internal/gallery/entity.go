package gallery

import (
	"time"

	"github.com/google/uuid"
)

type Gallery struct {
	ID        uuid.UUID `json:"id"`
	Image     string    `json:"image"`
	Caption   string    `json:"caption"`
	Location  *string   `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
