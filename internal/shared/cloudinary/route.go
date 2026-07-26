package cloudinary

import "github.com/go-chi/chi/v5"

// Routes registers the image upload endpoint.
func (h *UploadHandler) Routes(r chi.Router) {
	r.Post("/upload/image", h.Upload)
}
