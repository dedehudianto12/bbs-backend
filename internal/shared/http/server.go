package http

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dedehudianto12/bbs-backend/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	cfg     *config.Config
	handler http.Handler
	srv     *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		cfg:     cfg,
		handler: handler,
	}
}

func (s *Server) Run() error {
	s.srv = &http.Server{
		Addr:    net.JoinHostPort("", s.cfg.Server.Port),
		Handler: s.handler,
	}

	// Graceful shutdown on SIGINT/SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		slog.Info("shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.srv.Shutdown(ctx); err != nil {
			slog.Error("server forced to shutdown", "err", err)
		}
	}()

	slog.Info("server starting", "port", s.cfg.Server.Port)
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	slog.Info("server stopped")
	return nil
}

func NewRouter(corsOrigins []string) chi.Router {
	r := chi.NewRouter()
	r.Use(SecurityHeaders)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	return r
}
