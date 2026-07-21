package category

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/api/admin/kategori", h.List)
	r.Put("/api/admin/kategori/{slug}", h.UpdateLabel)
}

func PublicRoutes(r chi.Router, h *Handler) {
	r.Get("/api/kategori", h.List)
}
