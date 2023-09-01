package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashep/go-cfgloader"

	"github.com/ashep/ghwh/internal/app"
)

func main() {
	time.Local = time.UTC

	var l *slog.Logger
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		l = slog.New(slog.NewTextHandler(os.Stdout, nil))
	} else {
		l = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	cfg := app.Config{}
	if err := cfgloader.LoadFromEnv("ghwh", &cfg); err != nil {
		l.Error("config load failed", "error", err)
		os.Exit(1)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sig
		l.Info("signal received", "signal", s)
		ctxC()
	}()

	if err := app.New(cfg, l).Run(ctx); err != nil {
		l.Error("app run failed", "error", err)
		os.Exit(1)
	}
}
