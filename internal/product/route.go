package product

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/api/admin/produk", h.List)
	r.Get("/api/admin/produk/{id}", h.GetByID)
	r.Post("/api/admin/produk", h.Create)
	r.Put("/api/admin/produk/{id}", h.Update)
	r.Delete("/api/admin/produk/{id}", h.Delete)
}

func PublicRoutes(r chi.Router, h *Handler) {
	r.Get("/api/produk", h.PublicList)
	r.Get("/api/produk/{slug}", h.GetBySlug)
}
