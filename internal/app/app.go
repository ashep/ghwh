package app

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/ashep/ghwh/internal/server"
)

type App struct {
	srv *server.Server
	l   zerolog.Logger
}

func New(cfg Config, l zerolog.Logger) *App {
	return &App{
		srv: server.New(cfg.Server, l),
		l:   l,
	}
}

func (a *App) Run(ctx context.Context) error {
	if err := a.srv.Run(ctx); err != nil {
		return fmt.Errorf("server: %w", err)
	}

	return nil
}
