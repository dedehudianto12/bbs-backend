package category

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
)

type Handler struct {
	usecase *Usecase
}

func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	cats, err := h.usecase.GetAll(r.Context(), group)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, cats)
}

type updateLabelRequest struct {
	Label string `json:"label"`
}

func (h *Handler) UpdateLabel(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	var req updateLabelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	cat, err := h.usecase.UpdateLabel(r.Context(), slug, req.Label)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, cat)
}
