package service

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/api/admin/jasa", h.List)
	r.Get("/api/admin/jasa/{id}", h.GetByID)
	r.Post("/api/admin/jasa", h.Create)
	r.Put("/api/admin/jasa/{id}", h.Update)
	r.Delete("/api/admin/jasa/{id}", h.Delete)
}

func PublicRoutes(r chi.Router, h *Handler) {
	r.Get("/api/jasa", h.PublicList)
	r.Get("/api/jasa/{slug}", h.GetBySlug)
}
