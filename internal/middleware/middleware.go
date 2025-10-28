package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{w, http.StatusOK}
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			crw := newCustomResponseWriter(w)

			// Пропускаем в следующий handler
			next.ServeHTTP(crw, r)

			logger.Info("request processed",
				zap.String("path", r.URL.Path),
				zap.Int("status", crw.statusCode),
				zap.String("method", r.Method),
			)
		})
	}
}
