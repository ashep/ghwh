package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func Log(next http.Handler, l zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs := bytes.NewBuffer([]byte{})
		if err := r.Header.Write(hs); err != nil {
			l.Error().Err(err).Msg("request header read failed")
		}

		bb, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			l.Error().Err(err).Msg("request body read failed")
			return
		}

		rc := r.Clone(r.Context())
		rc.Body = io.NopCloser(bytes.NewReader(bb))

		l.Info().
			Str("method", r.Method).
			Str("header", hs.String()).
			Str("body", string(bb)).
			Msg("request")

		next.ServeHTTP(w, rc)
	})
}
