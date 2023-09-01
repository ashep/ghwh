package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

func Log(next http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs := bytes.NewBuffer([]byte{})
		if err := r.Header.Write(hs); err != nil {
			l.Error("request header read failed", "error", err)
		}

		bb, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			l.Error("request body read failed", "error", err)
			return
		}

		l.Info("request", "header", hs, "body", bb)

		next.ServeHTTP(w, &http.Request{Body: io.NopCloser(bytes.NewReader(hs.Bytes()))})
	})
}
