package v1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/articles", func(resp http.ResponseWriter, req *http.Request) {
			fmt.Println("Received v1/articles request")
		})
		r.Get("/articles/{id}", func(resp http.ResponseWriter, req *http.Request) {
			fmt.Println("Received v1/articles/{id} request")
		})
	})
	return r
}
