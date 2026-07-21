package service

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

type serviceRequest struct {
	Name             string   `json:"name"`
	ShortDescription string   `json:"shortDescription"`
	FullDescription  string   `json:"fullDescription"`
	Images           []string `json:"images"`
}

// --- Admin ---

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	services, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, services)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	s, err := h.usecase.GetByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, s)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req serviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	s := &Service{
		Name:             req.Name,
		ShortDescription: req.ShortDescription,
		FullDescription:  req.FullDescription,
		Images:           req.Images,
	}

	if err := h.usecase.Create(r.Context(), s); err != nil {
		code := http.StatusInternalServerError
		if err == ErrNameRequired {
			code = http.StatusBadRequest
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusCreated, s)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req serviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	s := &Service{
		Name:             req.Name,
		ShortDescription: req.ShortDescription,
		FullDescription:  req.FullDescription,
		Images:           req.Images,
	}

	updated, err := h.usecase.Update(r.Context(), chi.URLParam(r, "id"), s)
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
	s, err := h.usecase.GetBySlug(r.Context(), chi.URLParam(r, "slug"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, s)
}
