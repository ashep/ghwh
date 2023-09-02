package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"

	"github.com/ashep/ghwh/internal/server/handler"
	"github.com/ashep/ghwh/internal/server/middleware"
)

type Server struct {
	cfg Config
	srv *http.Server
	l   zerolog.Logger
}

func New(cfg Config, l zerolog.Logger) *Server {
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
			s.l.Error().Err(errF).Msg("server close failed")
		}
	}()

	s.l.Info().Str("addr", s.cfg.Address).Msg("server is starting")

	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("ListenAndServe failed: %w", err)
	}

	s.l.Info().Str("addr", s.cfg.Address).Msg("server stopped")

	return nil
}
