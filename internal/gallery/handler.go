package gallery

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/dedehudianto12/bbs-backend/internal/shared/cloudinary"
	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
)

const maxUploadSize = 10 << 20 // 10 MB

type Handler struct {
	usecase *Usecase
	cld     *cloudinary.Service
}

func NewHandler(usecase *Usecase, cld *cloudinary.Service) *Handler {
	return &Handler{usecase: usecase, cld: cld}
}

// --- Admin ---

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	page, limit, search, sort := httphelper.ParsePagination(r)

	galleries, total, err := h.usecase.ListAdmin(r.Context(), search, sort, page, limit)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.SuccessPaginated(w, http.StatusOK, galleries, total, page, limit, sort)
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	g := &Gallery{
		Caption: r.FormValue("caption"),
	}

	if loc := r.FormValue("location"); loc != "" {
		g.Location = &loc
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "galeri")
		if err != nil {
			log.Printf("WARNING: cloudinary upload failed (gallery created without image): %v", err)
		} else {
			g.Image = url
		}
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	g := &Gallery{
		Caption: r.FormValue("caption"),
	}

	if loc := r.FormValue("location"); loc != "" {
		g.Location = &loc
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "galeri")
		if err != nil {
			log.Printf("WARNING: cloudinary upload failed (gallery updated without image): %v", err)
		} else {
			g.Image = url
		}
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
