package middleware

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/massivemadness/articles-server/internal/metrics"
)

func Metrics() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				chiContext := chi.RouteContext(r.Context())

				pattern := chiContext.RoutePattern()
				method := chiContext.RouteMethod
				status := strconv.Itoa(ww.Status())

				metrics.HttpRequestsTotal.WithLabelValues(pattern, method, status).Inc()
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
