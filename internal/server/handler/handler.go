package handler

import (
	"io"
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

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.l.Error().Err(err).Msg("request body read failed")

		return
	}

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	if _, err := w.Write(bb); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.l.Error().Err(err).Msg("response write failed")

		return
	}
}
