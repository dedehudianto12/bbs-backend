package industry

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

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

	industries, total, err := h.usecase.ListAdmin(r.Context(), search, sort, page, limit)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.SuccessPaginated(w, http.StatusOK, industries, total, page, limit, sort)
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	ind := &Industry{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	if slugs := r.FormValue("productSlugs"); slugs != "" {
		ind.ProductSlugs = strings.Split(slugs, ",")
		for i := range ind.ProductSlugs {
			ind.ProductSlugs[i] = strings.TrimSpace(ind.ProductSlugs[i])
		}
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "industries", "")
		if err != nil {
			slog.Error("cloudinary upload", "err", err)
			httphelper.Error(w, http.StatusInternalServerError, fmt.Errorf("gagal upload gambar: %w", err))
			return
		}
		ind.Image = &url
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	ind := &Industry{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	if slugs := r.FormValue("productSlugs"); slugs != "" {
		ind.ProductSlugs = strings.Split(slugs, ",")
		for i := range ind.ProductSlugs {
			ind.ProductSlugs[i] = strings.TrimSpace(ind.ProductSlugs[i])
		}
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "industries", chi.URLParam(r, "id"))
		if err != nil {
			slog.Error("cloudinary upload", "err", err)
			httphelper.Error(w, http.StatusInternalServerError, fmt.Errorf("gagal upload gambar: %w", err))
			return
		}
		ind.Image = &url
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

func (h *Handler) PublicList(w http.ResponseWriter, r *http.Request) {
	industries, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, industries)
}

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
