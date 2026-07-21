package article

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

type articleRequest struct {
	Title       string  `json:"title"`
	Excerpt     string  `json:"excerpt"`
	Content     string  `json:"content"`
	Image       *string `json:"image"`
	Tag         string  `json:"tag"`
	PublishedAt string  `json:"publishedAt"`
	Author      string  `json:"author"`
}

// --- Admin ---

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	articles, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, articles)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	a, err := h.usecase.GetByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, a)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req articleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	a := &Article{
		Title:       req.Title,
		Excerpt:     req.Excerpt,
		Content:     req.Content,
		Image:       req.Image,
		Tag:         req.Tag,
		PublishedAt: req.PublishedAt,
		Author:      req.Author,
	}

	if err := h.usecase.Create(r.Context(), a); err != nil {
		code := http.StatusInternalServerError
		if err == ErrTitleRequired {
			code = http.StatusBadRequest
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusCreated, a)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req articleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	a := &Article{
		Title:       req.Title,
		Excerpt:     req.Excerpt,
		Content:     req.Content,
		Image:       req.Image,
		Tag:         req.Tag,
		PublishedAt: req.PublishedAt,
		Author:      req.Author,
	}

	updated, err := h.usecase.Update(r.Context(), chi.URLParam(r, "id"), a)
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
	a, err := h.usecase.GetBySlug(r.Context(), chi.URLParam(r, "slug"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, a)
}
