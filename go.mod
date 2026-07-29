module github.com/dedehudianto12/bbs-backend

// pgx v5.10.0 requires go >= 1.25.0; must match or Go 1.22 refuses to build.
// This is safe — the go directive is a minimum, not the actual compiler version.
go 1.25.5

require github.com/go-chi/chi v1.5.5

require (
	github.com/cloudinary/cloudinary-go/v2 v2.16.0 // indirect
	github.com/creasty/defaults v1.7.0 // indirect
	github.com/go-chi/chi/v5 v5.3.1 // indirect
	github.com/go-chi/cors v1.2.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/schema v1.4.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.10.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)
