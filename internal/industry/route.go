package industry

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/api/admin/industri", h.List)
	r.Get("/api/admin/industri/{id}", h.GetByID)
	r.Post("/api/admin/industri", h.Create)
	r.Put("/api/admin/industri/{id}", h.Update)
	r.Delete("/api/admin/industri/{id}", h.Delete)
}

func PublicRoutes(r chi.Router, h *Handler) {
	r.Get("/api/industri", h.PublicList)
	r.Get("/api/industri/{slug}", h.GetBySlug)
}
