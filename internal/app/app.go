package app

import (
	"context"
	"log/slog"

	"github.com/ashep/ghwh/internal/server"
)

type App struct {
	srv *server.Server
	l   *slog.Logger
}

func New(cfg Config, l *slog.Logger) *App {
	return &App{
		srv: server.New(cfg.Server, l),
		l:   l,
	}
}

func (a *App) Run(ctx context.Context) error {
	return a.srv.Run(ctx)
}
