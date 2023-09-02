package handler

import (
	"net/http"

	"github.com/rs/zerolog"
)

type Handler struct {
	l zerolog.Logger
}

func New(l zerolog.Logger) *Handler {
	return &Handler{
		l: l,
	}
}

func (h *Handler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
}
