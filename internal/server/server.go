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
	c Config
	s *http.Server
	l zerolog.Logger
}

func New(cfg Config, l zerolog.Logger) *Server {
	mux := http.NewServeMux()
	hdl := middleware.Auth(handler.New(l), cfg.AuthToken, l)

	if cfg.Debug {
		l.Warn().Msg("debug mode enabled")
		mux.Handle("/", middleware.Log(hdl, l))
	} else {
		mux.Handle("/", hdl)
	}

	srv := &http.Server{
		Addr:        cfg.Address,
		ReadTimeout: time.Second * time.Duration(cfg.ReadTimeout),
		Handler:     mux,
	}

	return &Server{
		c: cfg,
		s: srv,
		l: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()

		if errF := s.s.Close(); errF != nil {
			s.l.Error().Err(errF).Msg("server close failed")
		}
	}()

	s.l.Info().Str("addr", s.c.Address).Msg("server is starting")

	if s.c.AuthToken == "" {
		s.l.Warn().Msg("empty auth token, should be used only for development purposes")
	}

	if err := s.s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("serve failed: %w", err)
	}

	s.l.Info().Str("addr", s.c.Address).Msg("server stopped")

	return nil
}
