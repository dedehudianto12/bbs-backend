package health

import (
	"net/http"

	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	httphelper.Success(w, http.StatusOK, map[string]string{"status": "ok"})
}
