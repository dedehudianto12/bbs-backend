package auth

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, h *Handler) {
	r.Post("/api/admin/login", h.Login)
	r.Post("/api/admin/logout", h.Logout)
}

func AdminRoutes(r chi.Router, h *Handler) {
	r.Get("/api/admin/me", h.Me)
}
