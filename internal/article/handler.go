package article

import (
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
	page, limit, search, sort := httphelper.ParsePagination(r)
	tag := r.URL.Query().Get("tag")

	articles, total, err := h.usecase.ListAdmin(r.Context(), tag, search, sort, page, limit)
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.SuccessPaginated(w, http.StatusOK, articles, total, page, limit, sort)
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	a := &Article{
		Title:       r.FormValue("title"),
		Excerpt:     r.FormValue("excerpt"),
		Content:     r.FormValue("content"),
		Tag:         r.FormValue("tag"),
		PublishedAt: r.FormValue("publishedAt"),
		Author:      r.FormValue("author"),
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "articles", "")
		if err != nil {
			slog.Error("cloudinary upload", "err", err)
			httphelper.Error(w, http.StatusInternalServerError, fmt.Errorf("gagal upload gambar: %w", err))
			return
		}
		a.Image = &url
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
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	a := &Article{
		Title:       r.FormValue("title"),
		Excerpt:     r.FormValue("excerpt"),
		Content:     r.FormValue("content"),
		Tag:         r.FormValue("tag"),
		PublishedAt: r.FormValue("publishedAt"),
		Author:      r.FormValue("author"),
	}

	file, _, err := r.FormFile("file")
	if err == nil {
		defer file.Close()
		url, err := h.cld.Upload(r.Context(), file, "articles", chi.URLParam(r, "id"))
		if err != nil {
			slog.Error("cloudinary upload", "err", err)
			httphelper.Error(w, http.StatusInternalServerError, fmt.Errorf("gagal upload gambar: %w", err))
			return
		}
		a.Image = &url
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

func (h *Handler) PublicList(w http.ResponseWriter, r *http.Request) {
	articles, err := h.usecase.List(r.Context())
	if err != nil {
		httphelper.Error(w, http.StatusInternalServerError, err)
		return
	}
	httphelper.Success(w, http.StatusOK, articles)
}

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
