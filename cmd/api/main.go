package main

import (
	"log"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/dedehudianto12/bbs-backend/internal/article"
	"github.com/dedehudianto12/bbs-backend/internal/auth"
	"github.com/dedehudianto12/bbs-backend/internal/category"
	"github.com/dedehudianto12/bbs-backend/internal/gallery"
	"github.com/dedehudianto12/bbs-backend/internal/health"
	"github.com/dedehudianto12/bbs-backend/internal/industry"
	"github.com/dedehudianto12/bbs-backend/internal/product"
	"github.com/dedehudianto12/bbs-backend/internal/service"
	"github.com/dedehudianto12/bbs-backend/internal/shared/database"
	httpserver "github.com/dedehudianto12/bbs-backend/internal/shared/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connected")

	r := httpserver.NewRouter()

	// Auth module
	authRepo := auth.NewRepository(db)
	authUc := auth.NewUsecase(authRepo, cfg.JWT.Secret)
	authH := auth.NewHandler(authUc)
	auth.Routes(r, authH) // login & logout: no auth

	// Health
	healthH := health.NewHandler()
	health.Routes(r, healthH)

	// Public routes
	categoryH := category.NewHandler(category.NewUsecase(category.NewRepository(db)))
	category.PublicRoutes(r, categoryH)

	productH := product.NewHandler(product.NewUsecase(product.NewRepository(db)))
	product.PublicRoutes(r, productH)

	articleH := article.NewHandler(article.NewUsecase(article.NewRepository(db)))
	article.PublicRoutes(r, articleH)

	serviceH := service.NewHandler(service.NewUsecase(service.NewRepository(db)))
	service.PublicRoutes(r, serviceH)

	galleryH := gallery.NewHandler(gallery.NewUsecase(gallery.NewRepository(db)))
	gallery.PublicRoutes(r, galleryH)

	industryH := industry.NewHandler(industry.NewUsecase(industry.NewRepository(db)))
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

	server := httpserver.NewServer(cfg, r)
	log.Printf("Server running on port %s", cfg.Server.Port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
