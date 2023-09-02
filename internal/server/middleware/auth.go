package middleware

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

func Auth(next http.Handler, authToken string, l zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authToken == "" {
			next.ServeHTTP(w, r)
			return
		}

		tok := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		l.Info().Str("tok", tok).Msg("tok")
		if tok != authToken {
			l.Warn().
				Str("remote_addr", r.RemoteAddr).
				Str("request_uri", r.RequestURI).
				Msg("invalid authorization token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
