package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/dedehudianto12/bbs-backend/internal/article"
	"github.com/dedehudianto12/bbs-backend/internal/auth"
	"github.com/dedehudianto12/bbs-backend/internal/category"
	"github.com/dedehudianto12/bbs-backend/internal/gallery"
	"github.com/dedehudianto12/bbs-backend/internal/health"
	"github.com/dedehudianto12/bbs-backend/internal/industry"
	"github.com/dedehudianto12/bbs-backend/internal/product"
	"github.com/dedehudianto12/bbs-backend/internal/service"
	"github.com/dedehudianto12/bbs-backend/internal/shared/cloudinary"
	"github.com/dedehudianto12/bbs-backend/internal/shared/database"
	httpserver "github.com/dedehudianto12/bbs-backend/internal/shared/http"
	"github.com/dedehudianto12/bbs-backend/internal/swagger"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type statsHandler struct {
	db *pgxpool.Pool
}

func (h *statsHandler) Stats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stats := make(map[string]int)

	var count int
	if err := h.db.QueryRow(ctx, `SELECT count(*) FROM products`).Scan(&count); err == nil {
		stats["products"] = count
	}
	if err := h.db.QueryRow(ctx, `SELECT count(*) FROM articles`).Scan(&count); err == nil {
		stats["articles"] = count
	}
	if err := h.db.QueryRow(ctx, `SELECT count(*) FROM services`).Scan(&count); err == nil {
		stats["services"] = count
	}
	if err := h.db.QueryRow(ctx, `SELECT count(*) FROM galleries`).Scan(&count); err == nil {
		stats["galleries"] = count
	}

	httpserver.Success(w, http.StatusOK, stats)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config load", "err", err)
		os.Exit(1)
	}

	db, err := database.New(cfg)
	if err != nil {
		slog.Error("database connect", "err", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("database connected")

	r := httpserver.NewRouter(cfg.Server.CORSOrigins)

	// Auth module
	authRepo := auth.NewRepository(db)
	authUc := auth.NewUsecase(authRepo, cfg.JWT.Secret)
	authH := auth.NewHandler(authUc, cfg.App.Env == "production")
	auth.Routes(r, authH) // logout: no auth

	// Login with rate limiting: 5 req/min per IP
	loginLimiter := httpserver.PerIPRateLimit(5, time.Minute)
	r.Group(func(r chi.Router) {
		r.Use(loginLimiter)
		r.Post("/api/admin/login", authH.Login)
	})

	// Cloudinary
	cldSvc, err := cloudinary.New(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	if err != nil {
		slog.Error("cloudinary init", "err", err)
		os.Exit(1)
	}

	// Health
	healthH := health.NewHandler(db)
	health.Routes(r, healthH)

	// Public routes
	categoryH := category.NewHandler(category.NewUsecase(category.NewRepository(db)))
	category.PublicRoutes(r, categoryH)

	productH := product.NewHandler(product.NewUsecase(product.NewRepository(db)), cldSvc)
	product.PublicRoutes(r, productH)

	articleH := article.NewHandler(article.NewUsecase(article.NewRepository(db)), cldSvc)
	article.PublicRoutes(r, articleH)

	serviceH := service.NewHandler(service.NewUsecase(service.NewRepository(db)))
	service.PublicRoutes(r, serviceH)

	galleryH := gallery.NewHandler(gallery.NewUsecase(gallery.NewRepository(db)), cldSvc)
	gallery.PublicRoutes(r, galleryH)

	industryH := industry.NewHandler(industry.NewUsecase(industry.NewRepository(db)), cldSvc)
	industry.PublicRoutes(r, industryH)

	// Admin routes (protected)
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(authUc))
		auth.AdminRoutes(r, authH)
		category.Routes(r, categoryH)
		product.Routes(r, productH)
		article.Routes(r, articleH)
		service.Routes(r, serviceH)
		gallery.Routes(r, galleryH)
		industry.Routes(r, industryH)
	})

	// Swagger docs
	swaggerHandler := swagger.Handler()
	yamlPath := filepath.Join("docs", "swagger.yaml")
	r.Get("/docs/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, yamlPath)
	})
	r.Handle("/docs/*", http.StripPrefix("/docs", swaggerHandler))

	// Admin stats (protected)
	statsH := &statsHandler{db: db}
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(authUc))
		r.Get("/api/admin/stats", statsH.Stats)
	})

	server := httpserver.NewServer(cfg, r)
	if err := server.Run(); err != nil {
		slog.Error("server run", "err", err)
		os.Exit(1)
	}
}
