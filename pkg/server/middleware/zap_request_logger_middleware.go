package middleware

import (
	"fmt"
	"net/http"
	"skeleton/pkg/logger"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func ZapRequestLoggerMiddleware(
	logger *logger.Logger,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			format := fmt.Sprintf(
				"%s \"%s %s %s\" %d %d %d \"%s\" %s",
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
				r.Proto,
				ww.Status(),
				r.ContentLength,
				ww.BytesWritten(),
				r.UserAgent(),
				time.Since(start),
			)
			logger.Info(format)
		})
	}
}
