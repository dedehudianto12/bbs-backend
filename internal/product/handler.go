package product

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

type productRequest struct {
	Name        string          `json:"name"`
	Group       string          `json:"group"`
	Kategori    string          `json:"kategori"`
	Category    string          `json:"category"`
	Description string          `json:"description"`
	Detail      string          `json:"detail"`
	Image       *string         `json:"image"`
	Specs       json.RawMessage `json:"specs"`
}

// --- Admin ---

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	kategori := r.URL.Query().Get("kategori")

	products, err := h.usecase.List(r.Context(), group, kategori)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, products)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	p, err := h.usecase.GetByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, p)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req productRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	p := &Product{
		Name:        req.Name,
		Group:       req.Group,
		Kategori:    req.Kategori,
		Category:    req.Category,
		Description: req.Description,
		Detail:      req.Detail,
		Image:       req.Image,
		Specs:       req.Specs,
	}

	if err := h.usecase.Create(r.Context(), p); err != nil {
		code := http.StatusInternalServerError
		if err == ErrNameRequired {
			code = http.StatusBadRequest
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusCreated, p)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req productRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	p := &Product{
		Name:        req.Name,
		Group:       req.Group,
		Kategori:    req.Kategori,
		Category:    req.Category,
		Description: req.Description,
		Detail:      req.Detail,
		Image:       req.Image,
		Specs:       req.Specs,
	}

	updated, err := h.usecase.Update(r.Context(), chi.URLParam(r, "id"), p)
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
	p, err := h.usecase.GetBySlug(r.Context(), chi.URLParam(r, "slug"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, p)
}
