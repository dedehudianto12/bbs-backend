package auth

import (
	"context"
	"net/http"
	"strings"

	httphelper "github.com/dedehudianto12/bbs-backend/internal/shared/http"
)

type contextKey string

const adminCtxKey contextKey = "admin"

func Middleware(usecase *Usecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractToken(r)
			if token == "" {
				httphelper.Error(w, http.StatusUnauthorized, ErrInvalidToken)
				return
			}

			admin, err := usecase.Me(token)
			if err != nil {
				httphelper.Error(w, http.StatusUnauthorized, err)
				return
			}

			ctx := context.WithValue(r.Context(), adminCtxKey, admin)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractToken(r *http.Request) string {
	// Cookie first (production)
	cookie, err := r.Cookie("token")
	if err == nil && cookie.Value != "" {
		return cookie.Value
	}

	// Bearer header fallback (development)
	header := r.Header.Get("Authorization")
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}

	return ""
}
