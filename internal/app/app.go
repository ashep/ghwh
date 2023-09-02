package app

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/ashep/ghwh/internal/server"
)

type App struct {
	s *server.Server
	l zerolog.Logger
}

func New(cfg Config, l zerolog.Logger) *App {
	return &App{
		s: server.New(cfg.Server, l.With().Str("pkg", "server").Logger()),
		l: l,
	}
}

func (a *App) Run(ctx context.Context) error {
	if err := a.s.Run(ctx); err != nil {
		return fmt.Errorf("server: %w", err)
	}

	return nil
}
