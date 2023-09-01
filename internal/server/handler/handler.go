package handler

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	l *slog.Logger
}

func New(l *slog.Logger) *Handler {
	return &Handler{
		l: l,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
