package article

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/api/admin/artikel", h.List)
	r.Get("/api/admin/artikel/{id}", h.GetByID)
	r.Post("/api/admin/artikel", h.Create)
	r.Put("/api/admin/artikel/{id}", h.Update)
	r.Delete("/api/admin/artikel/{id}", h.Delete)
}

func PublicRoutes(r chi.Router, h *Handler) {
	r.Get("/api/artikel", h.PublicList)
	r.Get("/api/artikel/{slug}", h.GetBySlug)
}
