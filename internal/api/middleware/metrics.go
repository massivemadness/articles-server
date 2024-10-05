package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/massivemadness/articles-server/internal/metrics"
)

func Metrics() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				chiContext := chi.RouteContext(r.Context())

				pattern := chiContext.RoutePattern()
				method := chiContext.RouteMethod

				metrics.HttpRequestsTotal.WithLabelValues(pattern, method).Inc()
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
