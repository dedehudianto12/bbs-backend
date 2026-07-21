package gallery

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

type galleryRequest struct {
	Image    string  `json:"image"`
	Caption  string  `json:"caption"`
	Location *string `json:"location"`
}

// --- Admin ---

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	galleries, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, galleries)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	g, err := h.usecase.GetByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, g)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req galleryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	g := &Gallery{
		Image:    req.Image,
		Caption:  req.Caption,
		Location: req.Location,
	}

	if err := h.usecase.Create(r.Context(), g); err != nil {
		code := http.StatusInternalServerError
		if err == ErrImageRequired {
			code = http.StatusBadRequest
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusCreated, g)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req galleryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	g := &Gallery{
		Image:    req.Image,
		Caption:  req.Caption,
		Location: req.Location,
	}

	updated, err := h.usecase.Update(r.Context(), chi.URLParam(r, "id"), g)
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

func (h *Handler) PublicList(w http.ResponseWriter, r *http.Request) {
	galleries, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, galleries)
}

func (h *Handler) PublicGetByID(w http.ResponseWriter, r *http.Request) {
	g, err := h.usecase.GetByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		code := http.StatusInternalServerError
		if err == ErrNotFound {
			code = http.StatusNotFound
		}
		httphelper.Error(w, code, err)
		return
	}
	httphelper.Success(w, http.StatusOK, g)
}
