package http

import (
	"net"
	"net/http"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/go-chi/chi/v5"
)

type Server struct{
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server{
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() error{
	r := chi.NewRouter()

	server := &http.Server{
		Addr: net.JoinHostPort("", s.cfg.Server.Port),
		Handler: r,
	}

	return server.ListenAndServe()
}