package cloudinary

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
)

const maxUploadSize = 10 << 20 // 10 MB

// UploadHandler handles multipart image upload to Cloudinary.
type UploadHandler struct {
	svc *Service
}

func NewUploadHandler(svc *Service) *UploadHandler {
	return &UploadHandler{svc: svc}
}

type uploadResponse struct {
	URL string `json:"url"`
}

// Upload accepts POST with multipart/form-data, field name "file".
// Optional form field "folder" to set Cloudinary folder.
func (h *UploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		httphelper.Error(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	folder := r.FormValue("folder")
	if folder == "" {
		folder = "general"
	}

	url, err := h.svc.Upload(r.Context(), file, folder)
	if err != nil {
		slog.Error("cloudinary upload", "err", err)
		httphelper.Error(w, http.StatusInternalServerError, 
			fmt.Errorf("gagal upload ke server. Pastikan konfigurasi Cloudinary sudah benar"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(httphelper.Response{
		Data:  uploadResponse{URL: url},
		Error: nil,
	})
}
