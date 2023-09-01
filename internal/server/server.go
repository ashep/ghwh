package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ashep/ghwh/internal/server/handler"
	"github.com/ashep/ghwh/internal/server/middleware"
)

type Server struct {
	cfg Config
	srv *http.Server
	l   *slog.Logger
}

func New(cfg Config, l *slog.Logger) *Server {
	hdl := handler.New(l)

	mux := http.NewServeMux()
	mux.Handle("/", middleware.Log(hdl, l))

	srv := &http.Server{
		Addr:        cfg.Address,
		ReadTimeout: time.Second * time.Duration(cfg.ReadTimeout),
		Handler:     mux,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()

		if errF := s.srv.Close(); errF != nil {
			s.l.Error("server close failed")
		}
	}()

	s.l.Info("starting the server", "addr", s.cfg.Address)

	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("ListenAndServe failed: %w", err)
	}

	s.l.Info("server stopped", "addr", s.cfg.Address)

	return nil
}
