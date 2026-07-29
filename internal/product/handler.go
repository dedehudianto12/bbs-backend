package product

import (
	"encoding/json"
	"fmt"
	"log/slog"
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
	group := r.URL.Query().Get("group")
	kategori := r.URL.Query().Get("kategori")
	page, limit, search, sort := httphelper.ParsePagination(r)
	products, total, err := h.usecase.ListAdmin(r.Context(), group, kategori, search, sort, page, limit)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.SuccessPaginated(w, http.StatusOK, products, total, page, limit, sort)
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		slog.Error("parse multipart", "err", err)
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	p := &Product{
		Name:        r.FormValue("name"),
		Group:       r.FormValue("group"),
		Kategori:    r.FormValue("kategori"),
		Category:    r.FormValue("category"),
		Description: r.FormValue("description"),
		Detail:      r.FormValue("detail"),
	}

	if specs := r.FormValue("specs"); specs != "" {
		p.Specs = json.RawMessage(specs)
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "products")
		if err != nil {
			slog.Error("cloudinary upload", "err", err)
			httphelper.Error(w, http.StatusInternalServerError, fmt.Errorf("gagal upload gambar: %w", err))
			return
		}
		p.Image = &url
	}

	if err := h.usecase.Create(r.Context(), p); err != nil {
		slog.Error("create product", "err", err)
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	p := &Product{
		Name:        r.FormValue("name"),
		Group:       r.FormValue("group"),
		Kategori:    r.FormValue("kategori"),
		Category:    r.FormValue("category"),
		Description: r.FormValue("description"),
		Detail:      r.FormValue("detail"),
	}

	if specs := r.FormValue("specs"); specs != "" {
		p.Specs = json.RawMessage(specs)
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "products")
		if err != nil {
			slog.Error("cloudinary upload", "err", err)
			httphelper.Error(w, http.StatusInternalServerError, fmt.Errorf("gagal upload gambar: %w", err))
			return
		}
		p.Image = &url
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

func (h *Handler) PublicList(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	kategori := r.URL.Query().Get("kategori")

	products, err := h.usecase.List(r.Context(), group, kategori)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, products)
}

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
