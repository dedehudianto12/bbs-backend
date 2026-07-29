package health

import (
	"context"
	"net/http"
	"time"

	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := h.db.Ping(ctx); err != nil {
		httphelper.Error(w, http.StatusServiceUnavailable, err)
		return
	}

	httphelper.Success(w, http.StatusOK, map[string]string{"status": "ok"})
}
