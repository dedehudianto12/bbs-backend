package industry

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

type industryRequest struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Image        *string  `json:"image"`
	ProductSlugs []string `json:"productSlugs"`
}

// --- Admin ---

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	industries, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, industries)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ind, err := h.usecase.GetByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, ind)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req industryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	ind := &Industry{
		Name:         req.Name,
		Description:  req.Description,
		Image:        req.Image,
		ProductSlugs: req.ProductSlugs,
	}

	if err := h.usecase.Create(r.Context(), ind); err != nil {
		code := http.StatusInternalServerError
		if err == ErrNameRequired {
			code = http.StatusBadRequest
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusCreated, ind)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req industryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	ind := &Industry{
		Name:         req.Name,
		Description:  req.Description,
		Image:        req.Image,
		ProductSlugs: req.ProductSlugs,
	}

	updated, err := h.usecase.Update(r.Context(), chi.URLParam(r, "id"), ind)
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, updated)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.usecase.Delete(r.Context(), chi.URLParam(r, "id")); err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, nil)
}

// --- Public ---

func (h *Handler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	ind, err := h.usecase.GetBySlug(r.Context(), chi.URLParam(r, "slug"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, ind)
}
