package gallery

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Get("/api/admin/galeri", h.List)
	r.Get("/api/admin/galeri/{id}", h.GetByID)
	r.Post("/api/admin/galeri", h.Create)
	r.Put("/api/admin/galeri/{id}", h.Update)
	r.Delete("/api/admin/galeri/{id}", h.Delete)
}

func PublicRoutes(r chi.Router, h *Handler) {
	r.Get("/api/galeri", h.PublicList)
	r.Get("/api/galeri/{id}", h.PublicGetByID)
}
