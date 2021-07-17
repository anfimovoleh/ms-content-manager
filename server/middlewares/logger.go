package middlewares

import (
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/go-chi/chi/middleware"
)

func Logger(logger *zap.Logger, durationThreshold time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			requestID := rand.Int()

			logger = logger.With(
				zap.Int("request_id", requestID),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
			)
			defer func() {
				duration := time.Since(t1)
				lEntry := logger.With(
					zap.String("path", r.URL.Path),
					zap.String("duration", duration.String()),
					zap.Int("status", ww.Status()),
				)
				lEntry.Debug("request finished")

				if duration > durationThreshold {
					lEntry.With(zap.Any("http_request", r)).Warn("slow request")
				}
			}()

			logger.Debug("request started")
			next.ServeHTTP(ww, r)
		})
	}
}
