package http

import (
	"net"
	"net/http"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	cfg     *config.Config
	handler http.Handler
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		cfg:     cfg,
		handler: handler,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    net.JoinHostPort("", s.cfg.Server.Port),
		Handler: s.handler,
	}
	return server.ListenAndServe()
}

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "http://127.0.0.1:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	return r
}
