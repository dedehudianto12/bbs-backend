package cloudinary

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Service struct {
	cld *cloudinary.Cloudinary
}

func New(cloudName, apiKey, apiSecret string) (*Service, error) {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}
	return &Service{cld: cld}, nil
}

// Upload uploads an image from reader to Cloudinary and returns the secure URL.
// folder is optional (e.g. "articles", "gallery").
// publicID, if set, replaces the existing asset with that ID (idempotent overwrite).
func (s *Service) Upload(ctx context.Context, reader io.Reader, folder string, publicID string) (string, error) {
	params := uploader.UploadParams{Folder: folder}
	if publicID != "" {
		params.PublicID = publicID
		overwrite := true
		params.Overwrite = &overwrite // ponytail: one asset per entity, overwrite in place
	}
	resp, err := s.cld.Upload.Upload(ctx, reader, params)
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
